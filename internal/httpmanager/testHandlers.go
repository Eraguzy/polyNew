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
func HandlerKeepAlive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("keepalive triggered.")
	io.WriteString(w, "i'm alive for free :)\n")
}
