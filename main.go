package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	limit := 500

	for offset := 0; ; offset += 500 {
		url := fmt.Sprintf(
			"https://gamma-api.polymarket.com/events?active=true&closed=false&limit=%d&offset=%d",
			limit,
			offset,
		)

		fmt.Println("fetch offset =", offset)

		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			panic(err)
		}

		fmt.Println(string(body))

		if string(body) == "[]" {
			fmt.Println("done")
			break
		}
	}
}
