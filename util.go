package main

import (
	"fmt"
	"net/http"
	"time"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%d\n", time.Now().UnixNano())
}

func GetVersion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	s := fmt.Sprintf("gotube/%s", Config.Version())

	fmt.Fprintln(w, s)
}
