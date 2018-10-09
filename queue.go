package main

import (
	"log"
	"net/http"

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
}
