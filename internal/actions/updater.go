package actions

import (
	"context"
	"encoding/json"
	"io"
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

	io.WriteString(io.Discard, "Inserted tag "+slug+"\n")
	return nil
}

func CompareEvents(ctx context.Context, db storage.DBManager) error {
	// might add parameters e.g. eventid, tagids, etc later
	// get all tracked tags from db
	tableTags, err := db.GetTrackedTags(ctx)
	if err != nil {
		return err
	}
	seriesIDs := []int{}
	for _, tag := range tableTags {
		seriesIDs = append(seriesIDs, tag.TagID)
	}

	// get events from polymarket
	for _, serieID := range seriesIDs {
		args := []polymarket.GetArgs{
			{Param: polymarket.URLParamActive, Value: "true"},
			{Param: polymarket.URLParamClosed, Value: "false"},
			{Param: polymarket.URLParamLimit, Value: "500"},
			{Param: polymarket.URLParamSeriesID, Value: strconv.Itoa(serieID)},
			// {Param: polymarket.URLParamOffset, Value: fmt.Sprintf("%d", offset)},
		}
		body, err := polymarket.SendGetRequest(
			polymarket.URLGammaAPI,
			polymarket.PathEvents,
			args,
		)
		if err != nil {
			return err
		}

		events := []polymarket.Event{}
		err = json.Unmarshal(body, &events)
		if err != nil {
			return err
		}
	}

	// compare to events in db

	// update db accordingly
	// add event if missing
	// if present in db but not in polymarket, remove it

	// notify telegram

	return nil
}
