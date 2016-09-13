package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func download(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error loading '%s' got status %d (%s)", url, resp.StatusCode, resp.Status)
	}

	defer resp.Body.Close()
	return html.Parse(resp.Body)
}
