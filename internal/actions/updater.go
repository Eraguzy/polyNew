package actions

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/Eraguzy/PolyNew/internal/storage"
	"github.com/Eraguzy/PolyNew/messager/polymarket"
)

// if found in polymarket's api, insert/update in db
func FetchAndStoreTags(ctx context.Context, db storage.DBManager, slug string) error {
	body, err := polymarket.SendGetRequest(
		polymarket.URLGammaAPI,
		polymarket.PathTagsSlug(slug),
		nil,
	)
	if err != nil {
		return err
	}

	tag := polymarket.Tag{}
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

		args := []polymarket.GetArgs{
			{Param: polymarket.URLParamActive, Value: "true"},
			{Param: polymarket.URLParamClosed, Value: "false"},
			{Param: polymarket.URLParamLimit, Value: "500"},
			{Param: polymarket.URLParamTagID, Value: strconv.Itoa(tag.TagID)},
			// {Param: polymarket.URLParamOffset, Value: fmt.Sprintf("%d", offset)}, // required if api response is too long
		}
		body, err := polymarket.SendGetRequest(
			polymarket.URLGammaAPI,
			polymarket.PathEvents,
			args,
		)
		if err != nil {
			return nil, nil, err
		}

		events := []polymarket.Event{}
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

	// notify telegram (TODO)
	return inserted, deleted, nil
}
