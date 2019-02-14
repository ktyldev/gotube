package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func YtdlSearch(query string, resultCount int64) ([]Song, error) {
	if strings.ContainsAny(query, " ") {
		return nil, errors.New("query contains invalid characters")
	}

	query = fmt.Sprintf(
		"ytsearch%d:%s",
		resultCount,
		query)

	cmd := exec.Command(
		Config.YoutubeDl(),
		"--dump-json",
		query)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return makeResults(out)
}

func DownloadSong(id string) (Song, error) {
	dir, err := os.Getwd()
	songDir := Config.Read(CFG_SONG_DIR)

    s, err := GDetails(id)
	if err != nil {
		return Song{}, err
	}

	path := filepath.Join(dir, songDir)

	if _, err = os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}

	cmd := exec.Command(
		Config.YoutubeDl(),
		"-f 171", // webm
		fmt.Sprintf("-o%s", s.Filename()),
		"--",
		id)

	cmd.Dir = path

	e, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(fmt.Sprintf("%s\n", e))
		return Song{}, err
	}

	return s, err
}

func makeResults(jsonDump []byte) ([]Song, error) {
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

	return searchResults, nil
}
