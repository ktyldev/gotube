package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type Song struct {
	Title     string `json:"title"`     // video title
	Id        string `json:"id"`        // youtube video id
	Thumbnail string `json:"thumbnail"` // highest-res available thumbnail
	Duration  string `json:"duration"`  // video duration (in youtube's weird format)
}

func (s Song) Filename() string {
	return fmt.Sprintf("%s.ogg", s.Id)
}

func (s *Song) Path() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	songDir := Config.Read(CFG_SONG_DIR)

	return filepath.Join(cwd, songDir, s.Filename())
}
