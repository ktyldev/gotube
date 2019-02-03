package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var versionPath = "version.txt"

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%d\n", time.Now().UnixNano())
}

func GetVersion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	s := fmt.Sprintf("gotube/%s", Version())

	fmt.Fprintln(w, s)
}

func Version() string {
	b, err := ioutil.ReadFile(versionPath)
	if err != nil {
		panic(err)
	}

	return strings.TrimSuffix(string(b), "\n")
}
