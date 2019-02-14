package main

import (
	"fmt"
)

type Song struct {
	Title       string `json:"title"`       // video title
	Id          string `json:"id"`          // youtube video id
    Thumbnail   string `json:"thumbnail"`   // highest-res available thumbnail
    Duration    string `json:"duration"`    // video duration (in youtube's weird format)
}

func (s Song) Filename() string {
	return fmt.Sprintf("%s.webm", s.Id)
}
