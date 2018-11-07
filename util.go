package main

import (
	"fmt"
	"net/http"
	"time"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%d\n", time.Now().UnixNano())
}
