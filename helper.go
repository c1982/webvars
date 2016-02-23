package webvars

import (
	"io/ioutil"
	"net/http"
)

func getWebVars(url string) (content string, headers http.Header, statusCode int, err error) {

	statusCode = http.StatusCreated

	r, err := http.Get(url)

	if err != nil {
		return "", nil, statusCode, err
	}

	defer r.Body.Close()

	bt, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return "", nil, statusCode, err
	}

	statusCode = r.StatusCode
	headers = r.Header

	return string(bt), headers, statusCode, err
}
