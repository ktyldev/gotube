package main

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"
	"strconv"
	"strings"
)

type QueueClearAction struct {
	Index int `json:"index"`
}

// /queue/add
// POST
// params: id
func QueueAdd(w http.ResponseWriter, r *http.Request) {
	id, err := ReadStringRequest(r)

	s, err := DownloadSong(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	log.Printf("downloaded %s\n", s.Filename())

	GetQueue().Add(s)
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

// /queue/remove/{index}
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
