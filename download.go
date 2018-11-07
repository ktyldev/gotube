package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const _songDir = "tunes"

func DownloadSong(s Song) error {
	dir, err := os.Getwd()

	path := filepath.Join(dir, _songDir)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}

	cmd := exec.Command(
		"/bin/youtube-dl",
		"-f 171", // webm
		fmt.Sprintf("-o%s", s.Filename()),
		s.Id)

	cmd.Dir = path

	_, err = cmd.CombinedOutput()

	return err
}
