package main

import "net/http"

func TestRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/", GetRoot)
	mux.HandleFunc("/hello", GetHello)

	return mux
}
