package main

import (
	"encoding/json"
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

	cmd := exec.Command("/bin/youtube-dl", "--dump-json", searchStr)

	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	jsonResult, err := makeSearchResults(out)

	s := fmt.Sprintf("%s\n", jsonResult)
	fmt.Fprintf(w, s)
}

func makeSearchResults(jsonDump []byte) ([]byte, error) {
	// data comes back as a json dump on each line, not an array
	// {"id":"...", ...} # no commas! D:
	// {"id":"...", ...}
	// this means we need to split the output on newlines and deal
	// with each entry individually

	var searchResults []Song
	var err error

	first := 0
	for i, v := range jsonDump {
		// find newlines in byte array
		if v != '\n' {
			continue
		}

		// turn bytes into search result
		var song Song
		err = json.Unmarshal(jsonDump[first:i], &song)
		if err != nil {
			break
		}

		searchResults = append(searchResults, song)

		// increment first index so as not to include newline
		first = i + 1
	}

	return json.Marshal(searchResults)
}
