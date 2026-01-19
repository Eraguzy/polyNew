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
	if !r.URL.Query().Has("TagID") {
		io.WriteString(w, "'TagID' param is required\n")
		return
	}
	tracked, err := strconv.ParseBool(r.URL.Query().Get("tracked"))
	if err != nil {
		io.WriteString(w, fmt.Sprintf("error while parsing 'tracked' param: %v\n", err))
		return
	}
	TagID, err := strconv.Atoi(r.URL.Query().Get("TagID"))
	if err != nil {
		io.WriteString(w, fmt.Sprintf("error while parsing 'TagID' param: %v\n", err))
		return
	}

	err = db.Db.SetTrackedTag(r.Context(), TagID, tracked)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("error while setting tag's tracked value : %v\n", err))
		return
	}
}

func (db DBHandler) HandlerCompareEvents(w http.ResponseWriter, r *http.Request) {
	err := actions.CompareEvents(r.Context(), db.Db)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("error while comparing events: %v\n", err))
		return
	}
}

func (db DBHandler) HandlerFetchStoreTag(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("slug") {
		io.WriteString(w, "'tracked' slug is required\n")
		return
	}
	err := actions.FetchAndStoreTags(r.Context(), db.Db, r.URL.Query().Get("slug"))
	if err != nil {
		io.WriteString(w, fmt.Sprintf("error during fetching/storing tag: %v\n", err))
		return
	}
}
