package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// api/queue/{id}
func GetStream(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.RequestURI, "/")
	id := parts[len(parts)-1]

	song, err := GetQueue().GetById(id)

	filename := song.Filename()

	cwd, err := os.Getwd()

	dir := Config.Read(CFG_SONG_DIR)
	path := filepath.Join(cwd, dir, filename)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	audio, err := os.Open(path)
	defer audio.Close()

	http.ServeContent(w, r, filename, time.Now(), audio)
}
