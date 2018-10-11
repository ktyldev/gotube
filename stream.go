package main

import (
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func GetStream(w http.ResponseWriter, r *http.Request) {
	filename := "audio.webm"
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
