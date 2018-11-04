package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"gotube/request"
)

const _songDir = "tunes"

var _queue []Song

type QueueClearAction struct {
	Index int `json:"index"`
}

func QueueAdd(w http.ResponseWriter, r *http.Request) {
	var add Song

	err := request.Read(r, &add)
	if err != nil {
		panic(err)
	}

	err = _downloadSong(add)
	if err != nil {
		panic(err)
	}

	log.Printf("downloaded %s\n", add.Filename())

	_enqueue(add)
	log.Printf("added %s to queue\n", add.Filename())

	fmt.Fprintln(w, "ok")
}

func QueueGet(w http.ResponseWriter, r *http.Request) {
	out, err := json.Marshal(_queue)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "%s\n", out)
}

func QueueNext(w http.ResponseWriter, r *http.Request) {
	if len(_queue) != 0 {
		_discardTop()
	}

	fmt.Fprintln(w, "ok")
}

func QueueClear(w http.ResponseWriter, r *http.Request) {
	var clearAction QueueClearAction

	err := request.Read(r, &clearAction)
	if err != nil {
		panic(err)
	}

	index := clearAction.Index
	if index >= len(_queue) {
		msg := "index out of range"
		print(msg)
		fmt.Fprintln(w, msg)
		w.WriteHeader(400) // bad request
		return
	}

	// clear the whole queue
	if index == -1 {
		_queue = make([]Song, 0)
		fmt.Fprintln(w, "queue cleared")
		return
	}
}

func QueueGetCurrentFilename() string {
	return _queue[0].Filename()
}

func _enqueue(s Song) {
	_queue = append(_queue, s)
}

// not quite dequeue since we're not returning the result
func _discardTop() {
	_queue = _queue[1:]
}

func _downloadSong(s Song) error {
	cmd := exec.Command(
		"/bin/youtube-dl",
		"-f 171", // save as web,
		fmt.Sprintf("-o%s", s.Filename()),
		s.Id)

	dir, err := os.Getwd()
	cmd.Dir = filepath.Join(dir, _songDir)

	_, err = cmd.CombinedOutput()

	return err
}
