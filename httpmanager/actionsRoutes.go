package httpmanager

import (
	"net/http"
)

func UpdaterRoutes(mux *http.ServeMux, db DBHandler) *http.ServeMux {
	basePath := "/actions"
	mux.HandleFunc(basePath+"/fetch-and-store-tag", db.HandlerFetchStoreTag)
	mux.HandleFunc(basePath+"/set-tracked-tag", db.HandlerSetTrackedTag)
	mux.HandleFunc(basePath+"/compare-events", db.HandlerCompareEvents)

	return mux
}
