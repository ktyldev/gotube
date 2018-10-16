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

var _queue []Song

func QueueAdd(w http.ResponseWriter, r *http.Request) {
	var add Song

	err := request.Read(r, &add)
	if err != nil {
		panic(err)
	}

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("/bin/youtube-dl", "-f 171", "-o%(title)s.webm", add.Id)
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
