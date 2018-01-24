package main

import (
	"net/http"
	"golang.org/x/net/html"
)

type Link struct {
	url   string
	text  string
	depth int
}

type HttpError struct {
	original string
}

func LinkReader(resp *http.Response, depth int) []Link {
	page := html.NewTokenizer(resp.Body)
	links := []Link{}
	var start *html.Token
	var text string

	for {
		_ = page.Next()
		token := page.Token()

		if token.Type {
			
		}
	}
}
