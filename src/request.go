package src

import (
	"io"
	"net/http"
)

func getRequest(method, path string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, baseUrl+path, body)
	assert1(err)

	if path != ApiUpload {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	} else {
	}

	return req
}
