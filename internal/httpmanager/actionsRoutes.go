package httpmanager

import (
	"net/http"
)

func UpdaterRoutes(mux *http.ServeMux, db DBHandler, basicUser string, basicPass string) *http.ServeMux {
	basePath := "/actions"
	//tags
	mux.HandleFunc(basePath+"/fetch-store-tag", BasicAuth(db.HandlerFetchStoreTag, basicUser, basicPass))
	mux.HandleFunc(basePath+"/set-tracked-tag", BasicAuth(db.HandlerSetTrackedTag, basicUser, basicPass))
	mux.HandleFunc(basePath+"/notify-tracked-tags", BasicAuth(db.HandlerNotifyTrackedTags, basicUser, basicPass))

	mux.HandleFunc(basePath+"/compare-events", BasicAuth(db.HandlerCompareEvents, basicUser, basicPass))

	return mux
}
