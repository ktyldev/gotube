package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const _results = 5

func YtdlSearch(query string) ([]Song, error) {
	if strings.ContainsAny(query, " ") {
		return nil, errors.New("query contains invalid characters")
	}

	query = fmt.Sprintf(
		"ytsearch%d:%s",
		_results,
		query)

	cmd := exec.Command(
		GetConfig().YoutubeDl,
		"--dump-json",
		query)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return makeResults(out)
}

func DownloadSong(s Song) error {
	dir, err := os.Getwd()
	config := GetConfig()

	path := filepath.Join(dir, config.SongDir)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}

	cmd := exec.Command(
		config.YoutubeDl,
		"-f 171", // webm
		fmt.Sprintf("-o%s", s.Filename()),
		s.Id)

	cmd.Dir = path

	_, err = cmd.CombinedOutput()

	return err
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
