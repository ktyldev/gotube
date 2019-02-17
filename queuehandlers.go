package main

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"
)

type QueueAddAction struct {
	Id    string `json:"id"`
	Index int    `json:"index"`
}

type QueueAction struct {
	Id    string `json:"id"`
	Index int    `json:"index"`
}

type QueueMoveAction struct {
	Old int `json:"oldIndex"`
	New int `json:"newIndex"`
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

	err = GetQueue().Add(s, add.Index)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	log.Printf("added song at position %d: %s\n", add.Index, s.Title)

	w.WriteHeader(http.StatusOK)
}

func QueueMove(w http.ResponseWriter, r *http.Request) {
	var move QueueMoveAction

	err := ReadJsonRequest(r, &move)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	err = GetQueue().Move(move.Old, move.New)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	log.Printf(
		"moved %s: %d -> %d\n",
		GetQueue().Songs[move.New],
		move.Old,
		move.New)

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
	var a QueueAction

	ReadJsonRequest(r, &a)
	log.Println(a.Index)

	/*
		parts := strings.Split(r.RequestURI, "/")

		index, err := strconv.Atoi(parts[len(parts)-1])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			return
		}
	*/

	err := GetQueue().Remove(a.Index)
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
