package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"gotube/request"
)

type QueueAddDto struct {
	Url string `json:"url"`
}

func QueueAdd(w http.ResponseWriter, r *http.Request) {
	var add QueueAddDto

	err := request.Read(r, &add)
	if err != nil {
		panic(err)
	}

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("/bin/youtube-dl", "-f 171", add.Url)
	cmd.Dir = filepath.Join(dir, "tunes")

	// e = cmd.Run()
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	log.Printf("result: %s\n", out)
}
