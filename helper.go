package webvars

import (
	"io/ioutil"
	"net/http"
)

func getWebVars(url string) (content string, headers http.Header, err error) {

	r, err := http.Get(url)

	if err != nil {
		return "", nil, err
	}

	defer r.Body.Close()

	bt, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return "", nil, err
	}

	headers = r.Header

	return string(bt), headers, err
}
