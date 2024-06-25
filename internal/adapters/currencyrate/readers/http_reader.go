package readers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func ReadHTTP(URL string) (*map[string]interface{}, error) {
	parsedURL, err := url.Parse(URL)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(parsedURL.String())
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
