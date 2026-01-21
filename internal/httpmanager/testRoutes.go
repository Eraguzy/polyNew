package httpmanager

import "net/http"

func TestRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/", HandlerGetRoot)
	mux.HandleFunc("/hello", HandlerGetHello)

	return mux
}
