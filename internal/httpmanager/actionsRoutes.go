package httpmanager

import (
	"net/http"
)

func UpdaterRoutes(mux *http.ServeMux, db DBHandler) *http.ServeMux {
	basePath := "/actions"
	//tags
	mux.HandleFunc(basePath+"/fetch-store-tag", db.HandlerFetchStoreTag)
	mux.HandleFunc(basePath+"/set-tracked-tag", db.HandlerSetTrackedTag)
	mux.HandleFunc(basePath+"/notify-tracked-tags", db.HandlerNotifyTrackedTags)

	mux.HandleFunc(basePath+"/compare-events", db.HandlerCompareEvents)

	return mux
}
