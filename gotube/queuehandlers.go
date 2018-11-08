package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type QueueClearAction struct {
	Index int `json:"index"`
}

func QueueAdd(w http.ResponseWriter, r *http.Request) {
	var s Song

	err := ReadJsonRequest(r, &s)
	err = DownloadSong(s)
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

func QueueGetTop(w http.ResponseWriter, r *http.Request) {
	song, err := GetQueue().Top()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, song.Id)
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

func QueueNext(w http.ResponseWriter, r *http.Request) {
	err := GetQueue().Next()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func QueueClear(w http.ResponseWriter, r *http.Request) {
	var clearAction QueueClearAction

	err := ReadJsonRequest(r, &clearAction)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(w, err)
		return
	}

	q := GetQueue()
	index := clearAction.Index
	if index >= len(q.Songs) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(w, "index out of range")
		return
	}

	if index == -1 {
		q.Clear()
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusNotImplemented)
}
