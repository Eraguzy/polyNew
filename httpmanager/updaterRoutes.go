package httpmanager

import (
	"net/http"
)

func UpdaterRoutes(mux *http.ServeMux, db DBHandler) *http.ServeMux {
	basePath := "/updater"
	mux.HandleFunc(basePath+"/insert-all-sports", db.HandlerInsertSports)
	mux.HandleFunc(basePath+"/set-tracked-sport", db.HandlerSetTrackedSport)

	return mux
}
