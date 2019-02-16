package main

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"
	"strconv"
	"strings"
)

type QueueAddAction struct {
	Id    string `json:"id"`
	Index int    `json:"index"`
}

type QueueClearAction struct {
	Index int `json:"index"`
}

// /queue/add
// POST
// params: id
func QueueAdd(w http.ResponseWriter, r *http.Request) {
	var add QueueAddAction

	err := ReadJsonRequest(r, &add)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}

	// TODO: check if cache already contains the song
	s, err := DownloadSong(add.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	log.Printf("downloaded %s\n", s.Filename())

	err = GetQueue().Add(s, add.Index)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	log.Printf("added %s to queue\n", s.Title)

	w.WriteHeader(http.StatusOK)
}

func QueueGet(w http.ResponseWriter, r *http.Request) {
	out, err := json.Marshal(GetQueue().Songs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err)
		return
	}

	fmt.Fprintf(w, "%s\n", out)
}

// /queue/remove
func QueueRemove(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.RequestURI, "/")

	index, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	err = GetQueue().Remove(index)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func QueueClear(w http.ResponseWriter, r *http.Request) {
	GetQueue().Clear()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "")
}
