package main

import (
	"log"
	"net/http"
	"os/exec"

	"gotube/request"
)

type QueueAddDto struct {
	Url string `json:"url"`
}

func QueueAdd(w http.ResponseWriter, r *http.Request) {
	var add QueueAddDto

	e := request.Read(r, &add)
	if e != nil {
		panic(e)
	}

	log.Println(add.Url)

	cmd := exec.Command("/bin/youtube-dl", "-f 171", add.Url)
	cmd.Dir = "/home/mono/test"

	e = cmd.Run()
	if e != nil {
		panic(e)
	}
}
