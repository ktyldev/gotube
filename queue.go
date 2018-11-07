package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const _songDir = "tunes"

var _queue Queue = Queue{make([]Song, 0)}

type Queue struct {
	Songs []Song
}

func GetQueue() *Queue {
	return &_queue
}

func (q *Queue) Top() (Song, error) {
	if q.IsEmpty() {
		return Song{}, errors.New("queue is empty")
	}

	return q.Songs[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.Songs) == 0
}

func (q *Queue) Add(s Song) {
	q.Songs = append(q.Songs, s)
	log.Println(q.Songs)
}

func (q *Queue) Next() error {
	if q.IsEmpty() {
		return errors.New("queue is empty")
	}

	q.Songs = q.Songs[1:]
	log.Println("bump")

	return nil
}

func (q *Queue) Clear() {
	q.Songs = make([]Song, 0)
	log.Println("clear")
}

func (q *Queue) GetById(id string) (Song, error) {
	var result Song

	for _, s := range q.Songs {
		if s.Id == id {
			result = s
			break
		}
	}

	if result.Id == "" {
		msg := fmt.Sprintf("no song matching %s\n", id)
		return Song{}, errors.New(msg)
	}

	return result, nil
}

type QueueClearAction struct {
	Index int `json:"index"`
}

func QueueAdd(w http.ResponseWriter, r *http.Request) {
	var add Song

	err := ReadJsonRequest(r, &add)
	if err != nil {
		panic(err)
	}

	err = _downloadSong(add)
	if err != nil {
		panic(err)
	}

	log.Printf("downloaded %s\n", add.Filename())

	_queue.Add(add)
	log.Printf("added %s to queue\n", add.Filename())
	log.Println(_queue.Songs)

	fmt.Fprintln(w, "ok")
}

func QueueGetTop(w http.ResponseWriter, r *http.Request) {
	song, err := _queue.Top()
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	fmt.Fprintln(w, song.Id)
}

func QueueGet(w http.ResponseWriter, r *http.Request) {
	out, err := json.Marshal(_queue.Songs)
	if err != nil {
		panic(err)
	}

	log.Println(_queue.Songs)

	fmt.Fprintf(w, "%s\n", out)
}

func QueueNext(w http.ResponseWriter, r *http.Request) {
	err := _queue.Next()
	if err != nil {
		fmt.Fprintln(w, err)
	}

	w.WriteHeader(http.StatusOK)
}

func QueueClear(w http.ResponseWriter, r *http.Request) {
	var clearAction QueueClearAction

	err := ReadJsonRequest(r, &clearAction)
	if err != nil {
		panic(err)
	}

	index := clearAction.Index
	if index >= len(_queue.Songs) {
		msg := "index out of range"
		print(msg)
		fmt.Fprintln(w, msg)
		w.WriteHeader(400) // bad request
		return
	}

	// clear the whole queue
	if index == -1 {
		_queue.Clear()
		fmt.Fprintln(w, "queue cleared")
		return
	}
}

func _downloadSong(s Song) error {
	dir, err := os.Getwd()

	path := filepath.Join(dir, _songDir)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}

	cmd := exec.Command(
		"/bin/youtube-dl",
		"-f 171", // save as web,
		fmt.Sprintf("-o%s", s.Filename()),
		s.Id)

	cmd.Dir = path

	_, err = cmd.CombinedOutput()

	return err
}
