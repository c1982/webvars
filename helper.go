package webvars

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getWebVars(url string) (content string, headers http.Header, statusCode int, err error) {

	statusCode = 0

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

func getWebHeaders(url string) (headers http.Header, err error) {
	_, headers, status, err := getWebVars(url)
	return headers, err
}

//isLinux function is determines the server operating system.
func isLinux(ipaddr string) (bool, error) {

	yeahLinux := false
	hostUrl := fmt.Sprintf("http://%s", ipaddr)
	//Check Server HTTP Header
	headers, err := getWebHeaders(hostUrl)

	if err == nil {
		if len(headers) > 0 {
			serverHeader := headers.Get("Server")

		}
	}

	//Check ICPM TTL

	return yeahLinux, err
}

//isWindows function is determines the server operating system.
func isWindows(ipaddr string) bool {

	//Check Server HTTP Header
	//Check ICPM TTL

	return true
}
