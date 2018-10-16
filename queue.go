package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"gotube/request"
)

type QueueAddDto struct {
	Id string `json:"id"`
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

	url := fmt.Sprintf("https://youtube.com/watch?v=%s", add.Id)

	cmd := exec.Command("/bin/youtube-dl", "-f 171", "-oaudio.webm", url)
	cmd.Dir = filepath.Join(dir, "tunes")

	// e = cmd.Run()
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	log.Printf("result: %s\n", out)

	s := fmt.Sprintf("%s\n", out)
	fmt.Fprintf(w, s)
}
