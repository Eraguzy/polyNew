package httpmanager

import (
	"net/http"
)

func UpdaterRoutes(mux *http.ServeMux, db DBHandler) *http.ServeMux {
	basePath := "/actions"
	//tags
	mux.HandleFunc(basePath+"/fetch-store-tag", BasicAuth(db.HandlerFetchStoreTag))
	mux.HandleFunc(basePath+"/set-tracked-tag", BasicAuth(db.HandlerSetTrackedTag))
	mux.HandleFunc(basePath+"/notify-tracked-tags", BasicAuth(db.HandlerNotifyTrackedTags))

	mux.HandleFunc(basePath+"/compare-events", BasicAuth(db.HandlerCompareEvents))

	return mux
}
