package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	if !Config.Exists() {
		Config.Create()
	}

	port := Config.Read(CFG_PORT)
	log.Printf(
		"starting gotube v%s on port %s\n",
		Config.Version(),
		port)

	log.Printf(
		"using youtube-dl at %s\n",
		Config.YoutubeDl())
	log.Printf(
		"using du at %s\n",
		Config.Du())

	log.Println(Cache.DiskUsage())

	apiKey := Config.Read(CFG_G_API_KEY)
	if apiKey != "" {
		log.Println("found api key - fast search enabled")
	}

	port = fmt.Sprintf(":%s", port)
	log.Fatal(http.ListenAndServe(port, NewRouter()))
}
