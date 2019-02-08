package main

import (
	"fmt"
	"strings"
)

var (
	_badChars = " /-" // characters in song titles to replace
	_goodChar = "_"   // character to replace them with
)

type Song struct {
	Title string `json:"title"` // video title
	Id    string `json:"id"`    // youtube video id
}

func (s Song) Filename() string {

	f := strings.ToLower(s.Title)
	// replace bad chars
	for _, c := range _badChars {
		f = strings.Replace(f, string(c), _goodChar, -1)
	}

	for i := 1; i < len(f); i++ {
		test := strings.Repeat(_goodChar, i)
		f = strings.Replace(f, test, _goodChar, -1)
	}

	f = fmt.Sprintf("%s_%s.webm", s.Id, f)

	return f
}
