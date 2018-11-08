package main

import (
	"log"
	"net/http"
)

func main() {
	InitConfig()

	port := GetConfig().Port

	log.Printf("starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, NewRouter()))
}
