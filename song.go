package main

import (
	"fmt"
	"strings"
)

type Song struct {
	Title string `json:"title"` // video title
	Id    string `json:"id"`    // youtube video id
}

func (s Song) Filename() string {
	f := strings.ToLower(s.Title)
	f = strings.Replace(f, " ", "_", -1)
	f = fmt.Sprintf("%s.webm", f)

	return f
}
