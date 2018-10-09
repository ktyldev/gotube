package request

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// Limit incoming request size
const _requestSizeLimit int64 = 1048576

func Read(r *http.Request, obj interface{}) error {
	b, e := ioutil.ReadAll(io.LimitReader(r.Body, _requestSizeLimit))
	if e != nil {
		return e
	}

	e = json.Unmarshal(b, &obj)

	return e
}
