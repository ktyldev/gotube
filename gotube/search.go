package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"net/http"
)

type SearchDto struct {
	Query string `json:"query"`
}

func Search(w http.ResponseWriter, r *http.Request) {
	var search SearchDto
	var results []Song

	err := ReadJsonRequest(r, &search)
	query := strings.Replace(search.Query, " ", "+", -1)

	if GetConfig().GoogleApiKey == "" {
		// key not set, use slow search
		results, err = YtdlSearch(query)
	} else {
		// fast search! :D
		results, err = GSearch(query)
	}

	jsonResult, err := json.Marshal(results)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	s := fmt.Sprintf("%s\n", jsonResult)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, s)
}
