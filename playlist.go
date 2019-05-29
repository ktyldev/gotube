package main

import (
	//"fmt"
	//"encoding/json"
	"net/http"

	"log"
)

type PlaylistIds struct {
	Ids []string `json:"ids"`
}

// /playlist/load
// POST
// takes a list of IDs, clears the current queue and loads videos
// from IDs into new queue
func PlaylistLoad(w http.ResponseWriter, r *http.Request) {
	var ids []string

	err := ReadStringArray(r, &ids)
	if err != nil {
		panic(err)
	}

	for _, id := range ids {
		log.Println(id)
	}

	w.WriteHeader(http.StatusOK)
}

// /playlist/load/{youtubePlaylistId}
// POST
// loads a youtube playlist from a playlist ID. get the songs from the
// playlist and load them into the queue.
func PlaylistLoadYoutube(w http.ResponseWriter, r *http.Request) {

}
