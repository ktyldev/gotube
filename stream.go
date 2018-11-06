package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// api/queue/{id}
func GetStreamId(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.RequestURI, "/")
	id := parts[len(parts)-1]

	song := QueueGetSongById(id)

	_getStream(w, r, song.Filename())
}

func GetStream(w http.ResponseWriter, r *http.Request) {
	_getStream(w, r, QueueGetCurrentFilename())
}

func _getStream(w http.ResponseWriter, r *http.Request, filename string) {
	cwd, err := os.Getwd()
	if err != nil {
		// TODO: gracefully handle the case that no audio is available
		panic(err)
	}
	path := filepath.Join(cwd, "tunes", filename)

	audio, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer audio.Close()

	http.ServeContent(w, r, filename, time.Now(), audio)
}
