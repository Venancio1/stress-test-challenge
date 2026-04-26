package client

import (
	"net/http"
)

func RequestDefault(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return resp.Status, nil
}
