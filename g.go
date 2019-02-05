package main

import (
	"log"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

// https://developers.google.com/youtube/v3/docs/search/list#examples
func GSearch(query string, resultCount int64) ([]Song, error) {
	var results []Song

	key := Config.Read(CFG_G_API_KEY)

	client := &http.Client{
		Transport: &transport.APIKey{Key: key},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("error creating youtube client: %v\n", err)
	}

	call := service.Search.List("id,snippet").
		Q(query).
		MaxResults(resultCount).
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
