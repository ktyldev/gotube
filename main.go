package main

import (
	"log"
	"net/http"
)

func main() {
	InitConfig()

	config := GetConfig()

	log.Printf(
		"starting gotube v%s on port %s\n",
		config.Version,
		config.Port)

	log.Printf(
		"using youtube-dl at %s\n",
		config.YoutubeDl)
	log.Printf(
		"using du at %s\n",
		config.Du)

	if config.GoogleApiKey != "" {
		log.Println("found api key - fast search enabled")
	}

	log.Fatal(http.ListenAndServe(config.Port, NewRouter()))
}
