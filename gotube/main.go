package main

import (
	"log"
	"net/http"
)

func main() {
	InitConfig()

	conf := GetConfig()

	log.Printf(
		"starting gotube v%s on port %s\n",
		conf.Version,
		conf.Port)

	log.Printf(
		"using youtube-dl at %s\n",
		conf.YoutubeDl)

	if conf.GoogleApiKey != "" {
		log.Println("found api key - fast search enabled")
	}

	log.Fatal(http.ListenAndServe(conf.Port, NewRouter()))
}
