package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"gotube/request"
)

type SearchDto struct {
	Query string `json:"query"`
}

// number of results to get from the query; since they come back one by one,
// increasing this number will slow down searches
const _results = 5

func Search(w http.ResponseWriter, r *http.Request) {
	var search SearchDto

	err := request.Read(r, &search)
	if err != nil {
		panic(err)
	}

	// replace any spaces in the query with '+'. should probably do some
	// actual validation as well
	query := strings.Replace(search.Query, " ", "+", -1)
	searchStr := fmt.Sprintf("ytsearch%d:%s", _results, query)

	log.Println(searchStr)

	// TODO: get the whole json and return the urls as well as the titles
	cmd := exec.Command("/bin/youtube-dl", "--get-title", searchStr)

	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	s := fmt.Sprintf("%s\n", out)
	fmt.Fprintf(w, s)
}
