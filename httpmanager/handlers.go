package httpmanager

import (
	"fmt"
	"io"
	"net/http"
)

func HandlerGetRoot(w http.ResponseWriter, r *http.Request) {
	hasFirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")
	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")

	fmt.Printf("got / request. first(%t)=%s, second(%t)=%s\n",
		hasFirst, first,
		hasSecond, second)
	io.WriteString(w, "Server is up\n")
}
func HandlerGetHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "/hello received\n")
}
