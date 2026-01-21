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

	io.WriteString(w, "Server is up. To test arguments, try to pass 'first' and 'second' through the URL\n")
	io.WriteString(w, fmt.Sprintf("first(%t)=%s, second(%t)=%s\n",
		hasFirst, first,
		hasSecond, second,
	))
}
func HandlerGetHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "/hello received\n")
}
