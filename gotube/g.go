package main

import (
	"log"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var (
	_maxResults int64 = 5
)

// https://developers.google.com/youtube/v3/docs/search/list#examples
func GSearch(query string) ([]Song, error) {
	var results []Song

	client := &http.Client{
		Transport: &transport.APIKey{Key: GetConfig().GoogleApiKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("error creating youtube client: %v\n", err)
	}

	call := service.Search.List("id,snippet").
		Q(query).
		MaxResults(_maxResults).
		Type("video")

	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	for _, item := range response.Items {
		title := item.Snippet.Title
		id := item.Id.VideoId

		song := Song{
			title,
			id,
		}

		results = append(results, song)
	}

	return results, nil
}
