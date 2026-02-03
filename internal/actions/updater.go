package actions

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/Eraguzy/PolyNew/internal/storage"
	"github.com/Eraguzy/PolyNew/messenger"
)

// if found in polymarket's api, insert/update in db
func FetchAndStoreTags(ctx context.Context, db storage.DBManager, slug string) error {
	body, err := messenger.SendGetRequest(
		messenger.URLGammaAPI,
		messenger.PathTagsSlug(slug),
		nil,
	)
	if err != nil {
		return err
	}

	tag := messenger.Tag{}
	err = json.Unmarshal(body, &tag)
	if err != nil {
		return err
	}
	tabletag, err := tag.ToTableTag()
	if err != nil {
		return err
	}
	err = db.InsertTag(ctx, tabletag)
	if err != nil {
		return err
	}

	err = NotifyCurrentTags(ctx, db)
	if err != nil {
		return err
	}
	return nil
}

func SetTrackedTag(ctx context.Context, db storage.DBManager, tagID int, tracked bool) error {
	err := db.SetTrackedTag(ctx, tagID, tracked)
	if err != nil {
		return err
	}

	err = NotifyCurrentTags(ctx, db)
	if err != nil {
		return err
	}
	return nil
}

// compares events from polymarket API to those in our db. Inserts new events, deletes closed/nonexistent events
func CompareEvents(ctx context.Context, db storage.DBManager) (inserted []string, deleted []string, err error) {
	// might add parameters e.g. eventid, tagids, etc later
	// get all tracked tags from db
	tableTags, err := db.GetTrackedTags(ctx)
	if err != nil {
		return nil, nil, err
	}

	// get events from polymarket
	apiEvents := []storage.TableEvent{}
	for _, tag := range tableTags {
		if !tag.Tracked {
			continue
		}

		args := []messenger.GetArgs{
			{Param: messenger.URLParamActive, Value: "true"},
			{Param: messenger.URLParamClosed, Value: "false"},
			{Param: messenger.URLParamLimit, Value: "500"},
			{Param: messenger.URLParamTagID, Value: strconv.Itoa(tag.TagID)},
			// {Param: messenger.URLParamOffset, Value: fmt.Sprintf("%d", offset)}, // required if api response is too long
		}
		body, err := messenger.SendGetRequest(
			messenger.URLGammaAPI,
			messenger.PathEvents,
			args,
		)
		if err != nil {
			return nil, nil, err
		}

		events := []messenger.Event{}
		err = json.Unmarshal(body, &events)
		if err != nil {
			return nil, nil, err
		}
		for _, event := range events {
			tableEvent, err := event.ToTableEvent()
			if err != nil {
				return nil, nil, err
			}
			apiEvents = append(apiEvents, tableEvent)
		}
	}

	// compare to events in db
	// insert if missing
	eventsInDb, err := db.GetTableEvents(ctx, nil, nil)
	if err != nil {
		return nil, nil, err
	}
	notInDb := []storage.TableEvent{}
	for _, apiEvent := range apiEvents {
		found := false
		for _, dbEvent := range eventsInDb {
			if apiEvent.EventID == dbEvent.EventID {
				found = true
				break
			}
		}
		if !found && apiEvent.Title != nil {
			inserted = append(inserted, *apiEvent.Title)
			notInDb = append(notInDb, apiEvent)
		}
	}
	err = db.BulkInsertEvents(ctx, notInDb)
	if err != nil {
		return nil, nil, err
	}

	// remove if not present in polymarket anymore
	for _, dbEvent := range eventsInDb {
		found := false
		for _, apiEvent := range apiEvents {
			if apiEvent.EventID == dbEvent.EventID {
				found = true
				break
			}
		}
		if !found && dbEvent.Title != nil {
			deleted = append(deleted, *dbEvent.Title)
			err = db.DeleteEvent(ctx, dbEvent.EventID)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	if len(inserted) == 1 {
		link := "[" + inserted[0] + "](https://polymarket.com/search?_q=" + url.QueryEscape(inserted[0]) + ")"
		_, err = notifyTelegramChannel("New event found : "+link, true)
		if err != nil {
			return nil, nil, err
		}
	} else if len(inserted) > 1 {
		message := "New events found : \n"
		for _, title := range inserted {
			link := "[" + title + "](https://polymarket.com/search?_q=" + url.QueryEscape(title) + ")"

			if len(message+"- "+link+"\n") > 3000 { // avoid 400 error from telegram api
				_, err = notifyTelegramChannel(message, true)
				if err != nil {
					return nil, nil, err
				}
				message = ""
			}
			message += "- " + link + "\n"
		}
		_, err = notifyTelegramChannel(message, true)
		if err != nil {
			return nil, nil, err
		}
	}

	return inserted, deleted, nil
}

func NotifyCurrentTags(ctx context.Context, db storage.DBManager) error {
	tags, err := db.GetTrackedTags(ctx)
	if err != nil {
		return fmt.Errorf("error while getting tag by ID : %v", err)
	}
	output := "List updated. Now tracking : \n"
	for _, tag := range tags {
		output += fmt.Sprintf("(%d) %s\n", tag.TagID, tag.Name)
	}

	_, err = notifyTelegramChannel(output, false)
	if err != nil {
		return fmt.Errorf("error while notifying telegram channel : %v", err)
	}
	return nil
}

func notifyTelegramChannel(message string, disableNotif bool) ([]byte, error) {
	args := []messenger.GetArgs{
		{Param: messenger.URLDisableNotif, Value: strconv.FormatBool(disableNotif)},
		{Param: messenger.URLParamText, Value: message},
		{Param: messenger.URLParseMode, Value: "Markdown"},
		{Param: messenger.URLParamLinkPreviewOptions, Value: `{"is_disabled":true}`},
		{Param: messenger.URLParamChatID, Value: os.Getenv("TELEGRAM_CHANNEL_TAG")},
	}
	response, err := messenger.SendPostRequest(
		messenger.URLTelegramBotAPI,
		messenger.TelegramBotPathBuilder(
			os.Getenv("TELEGRAM_BOT_TOKEN"),
			messenger.PathSendMessage,
		),
		args,
	)
	return response, err
}

// goroutine to poll polymarket's api
func Polling(ctx context.Context, db storage.DBManager) {
	ticker := time.NewTicker(15 * time.Second)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				_, _, err := CompareEvents(ctx, db)
				if err != nil {
					fmt.Printf("error during polling: %v\n", err)
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}
