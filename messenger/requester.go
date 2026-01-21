package messenger

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func SendGetRequest(urlBase URL, route URLPath, args []GetArgs) ([]byte, error) {
	u, err := url.Parse(string(urlBase) + string(route))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	for _, arg := range args {
		q.Set(string(arg.Param), arg.Value)
	}
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func SendPostRequest(urlBase URL, route URLPath, args []GetArgs) ([]byte, error) {
	endpoint := string(urlBase) + string(route)

	form := url.Values{}
	for _, arg := range args {
		form.Set(string(arg.Param), arg.Value)
	}

	resp, err := http.PostForm(endpoint, form)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
