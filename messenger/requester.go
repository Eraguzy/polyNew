package messenger

import (
	"fmt"
	"io"
	"net/http"
)

// logic for fetching polymarket data

func SendGetRequest(url URL, route URLPath, args []GetArgs) ([]byte, error) {
	urlString := fmt.Sprintf("%s%s", string(url), string(route))
	if len(args) > 0 {
		urlString += "?"
		for i, arg := range args {
			if i > 0 {
				urlString += "&"
			}
			urlString += fmt.Sprintf("%s=%s", arg.Param, arg.Value)
		}
	}
	fmt.Println(urlString)
	resp, err := http.Get(urlString)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
