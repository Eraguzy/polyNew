package httpmanager

import (
	"net/http"
	"os"

	"github.com/Eraguzy/PolyNew/internal/storage"
)

type DBHandler struct {
	Db storage.DBManager
}

func BasicAuth(
	handler func(w http.ResponseWriter, r *http.Request),
) func(w http.ResponseWriter, r *http.Request) {

	var user = os.Getenv("BASICAUTH_USER")
	var pass = os.Getenv("BASICAUTH_PASS")

	return func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok || u != user || p != pass {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		handler(w, r)
	}
}
