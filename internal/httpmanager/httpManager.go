package httpmanager

import (
	"net/http"

	"github.com/Eraguzy/PolyNew/internal/storage"
)

type DBHandler struct {
	Db storage.DBManager
}

func BasicAuth(
	handler func(w http.ResponseWriter, r *http.Request),
	basicUser string,
	basicPass string,
) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok || u != basicUser || p != basicPass {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		handler(w, r)
	}
}
