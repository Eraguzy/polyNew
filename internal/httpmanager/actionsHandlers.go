package httpmanager

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Eraguzy/PolyNew/internal/actions"
)

func (db DBHandler) HandlerSetTrackedTag(w http.ResponseWriter, r *http.Request) {
	// check and parse GET params
	if !r.URL.Query().Has("tracked") {
		io.WriteString(w, "'tracked' param is required\n")
		return
	}
	if !r.URL.Query().Has("tagID") {
		io.WriteString(w, "'tagID' param is required\n")
		return
	}
	tracked, err := strconv.ParseBool(r.URL.Query().Get("tracked"))
	if err != nil {
		io.WriteString(w, fmt.Sprintf("error while parsing 'tracked' param: %v\n", err))
		return
	}
	tagID, err := strconv.Atoi(r.URL.Query().Get("tagID"))
	if err != nil {
		io.WriteString(w, fmt.Sprintf("error while parsing 'tagID' param: %v\n", err))
		return
	}

	err = actions.SetTrackedTag(r.Context(), db.Db, tagID, tracked)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("error while setting tag's tracked value : %v\n", err))
		return
	}
}

func (db DBHandler) HandlerCompareEvents(w http.ResponseWriter, r *http.Request) {
	inserted, deleted, err := actions.CompareEvents(r.Context(), db.Db)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("error while comparing events: %v\n", err))
		return
	}
	io.WriteString(w, fmt.Sprintf("Events inserted: %v\n", inserted))
	io.WriteString(w, fmt.Sprintf("Events deleted: %v\n", deleted))
}

func (db DBHandler) HandlerFetchStoreTag(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("slug") {
		io.WriteString(w, "'slug' param is required\n")
		return
	}
	err := actions.FetchAndStoreTags(r.Context(), db.Db, r.URL.Query().Get("slug"))
	if err != nil {
		io.WriteString(w, fmt.Sprintf("error during fetching/storing tag: %v\n", err))
		return
	}
	io.WriteString(
		w,
		fmt.Sprintf("Tag %s fetched and stored successfully\n",
			r.URL.Query().Get("slug"),
		),
	)
}
