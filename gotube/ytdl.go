package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const _binPath = "/bin/youtube-dl"

type SearchDto struct {
	Query string `json:"query"`
}

const _results = 5

func Search(w http.ResponseWriter, r *http.Request) {
	var search SearchDto

	err := ReadJsonRequest(r, &search)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	query := strings.Replace(search.Query, " ", "+", -1)
	searchStr := fmt.Sprintf("ytsearch%d:%s", _results, query)

	cmd := exec.Command(_binPath, "--dump-json", searchStr)

	out, err := cmd.CombinedOutput()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	jsonResult, err := _makeSearchResults(out)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	s := fmt.Sprintf("%s\n", jsonResult)
	fmt.Fprintf(w, s)
}

func DownloadSong(s Song) error {
	dir, err := os.Getwd()

	path := filepath.Join(dir, GetConfig().SongDir)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}

	cmd := exec.Command(
		_binPath,
		"-f 171", // webm
		fmt.Sprintf("-o%s", s.Filename()),
		s.Id)

	cmd.Dir = path

	_, err = cmd.CombinedOutput()

	return err
}

func _makeSearchResults(jsonDump []byte) ([]byte, error) {
	var searchResults []Song
	var err error

	first := 0
	for i, v := range jsonDump {
		if v != '\n' {
			continue
		}

		var song Song
		err = json.Unmarshal(jsonDump[first:i], &song)
		if err != nil {
			panic(err)
		}

		searchResults = append(searchResults, song)

		first = i + 1
	}

	return json.Marshal(searchResults)
}
