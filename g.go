package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

func GService() *youtube.Service {
	key := Config.Read(CFG_G_API_KEY)
	client := &http.Client{
		Transport: &transport.APIKey{Key: key},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("error creating youtube client: %v\n", err)
	}

	return service
}

func GGetVideoTitle(id string) (string, error) {
	call := GService().
		Videos.
		List("snippet")

	if id != "" {
		call = call.Id(id)
	}

	response, err := call.Do()
	if err != nil {
		panic(err)
	}

	l := len(response.Items)
	switch l {
	case 0:
		msg := fmt.Sprintf("found nothing with id: %s\n", id)
		return "", errors.New(msg)
	case 1: // this is ok
		break
	default:
		return "", errors.New("?!?!?!")
	}

	title := response.Items[0].Snippet.Title
	return title, nil
}

// set snippet as a blank interface
func GThumbnail(details *youtube.ThumbnailDetails) string {
    if details.Maxres != nil {
        return details.Maxres.Url
    }
    if details.Standard != nil {
        return details.Standard.Url
    }
    // this isn't a bug, "standard" is higher resolution than "high". go figure.
    if details.High != nil {
        return details.High.Url
    }
    if details.Medium != nil {
        return details.Medium.Url
    }
    return details.Default.Url
}

// https://developers.google.com/youtube/v3/docs/search/list#examples
func GSearch(query string, resultCount int64) ([]Song, error) {
	var results []Song

	call := GService().
		Search.
		List("id,snippet").
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
        thumbnail := GThumbnail(item.Snippet.Thumbnails)

		song := Song{
			title,
			id,
			thumbnail,
			"",
		}

		results = append(results, song)
	}

	return results, nil
}

// https://developers.google.com/youtube/v3/docs/videos/list
func GDetails(id string) (Song, error) {
    var details Song

    call := GService().
        Videos.
        List("snippet,contentDetails")

    if id == "" {
        return details, errors.New("cannot get the details of a song with blank id.")
    }

    call = call.Id(id)
    response, err := call.Do()
    if err != nil {
        return details, err
    }

    if len(response.Items) < 1 {
        return details, errors.New("song with that id is not available.")
    }

    title   := response.Items[0].Snippet.Title
    thumbnail := GThumbnail(response.Items[0].Snippet.Thumbnails)
    duration := response.Items[0].ContentDetails.Duration

    details = Song{
        title,
        id,
        thumbnail,
        duration,
    }

    return details, nil
}
