package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Limit incoming request size
const _requestSizeLimit int64 = 1048576

func ReadStringArray(r *http.Request, arr *[]string) error {
	b, e := ioutil.ReadAll(io.LimitReader(r.Body, _requestSizeLimit))
	if e != nil {
		return e
	}

	e = json.Unmarshal(b, arr)

	return e
}

func ReadJsonRequest(r *http.Request, obj interface{}) error {
	b, e := ioutil.ReadAll(io.LimitReader(r.Body, _requestSizeLimit))
	if e != nil {
		return e
	}

	e = json.Unmarshal(b, &obj)

	return e
}

func ReadStringRequest(r *http.Request) (string, error) {
	b, err := ioutil.ReadAll(io.LimitReader(r.Body, _requestSizeLimit))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", b), nil
}
