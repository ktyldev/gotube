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

func _enqueue(s Song) {
	_queue = append(_queue, s)
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
