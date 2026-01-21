package httpmanager

import "net/http"

func TestRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/", HandlerGetRoot)
	mux.HandleFunc("/keepalive", HandlerKeepAlive)

	return mux
}
