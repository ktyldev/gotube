package main

import (
	"fmt"
)

type Song struct {
	Title string `json:"title"` // video title
	Id    string `json:"id"`    // youtube video id
}

func (s Song) Filename() string {
	return fmt.Sprintf("%s.webm", s.Id)
}
