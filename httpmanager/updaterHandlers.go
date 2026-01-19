package httpmanager

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Eraguzy/PolyNew/internal/actions"
)

func (db DBHandler) HandlerInsertSports(w http.ResponseWriter, r *http.Request) {
	err := actions.InsertAllSports(r.Context(), db.Db)
	if err != nil {
		log.Fatalf("error while inserting all sports: %v\n", err)
	}
}

func (db DBHandler) HandlerSetTrackedSport(w http.ResponseWriter, r *http.Request) {
	// check and parse GET params
	if !r.URL.Query().Has("tracked") {
		log.Fatalf("'tracked' param is required\n")
	}
	if !r.URL.Query().Has("sportID") {
		log.Fatalf("'sportID' param is required\n")
	}
	tracked, err := strconv.ParseBool(r.URL.Query().Get("tracked"))
	if err != nil {
		log.Fatalf("error while parsing 'tracked' param: %v\n", err)
	}
	sportID, err := strconv.Atoi(r.URL.Query().Get("tracked"))
	if err != nil {
		log.Fatalf("error while parsing 'sportID' param: %v\n", err)
	}

	err = db.Db.SetTrackedSport(r.Context(), sportID, tracked)
	if err != nil {
		log.Fatalf("error while inserting all sports: %v\n", err)
	}
}
