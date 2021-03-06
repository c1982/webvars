package webvars

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

//getWebVars web request info
func getWebVars(url string) (content string, headers http.Header, statusCode int, err error) {

	statusCode = 0

	r, err := http.Head(url)

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

//getWebHeaders get all headers from web server
func getWebHeaders(url string) (headers http.Header, err error) {
	_, headers, _, err = getWebVars(url)
	return headers, err
}

//detectPorts check all ports on host
func detectPorts(ipaddr string, ports ...int) (result bool) {
	result = false
	c := make(chan bool)

	for _, port := range ports {
		go isPortOpen(c, ipaddr, port)
	}

	for _, _ = range ports {
		result = <-c
		if result {
			break
		}
	}

	return result
}

//checkPort check port from host ip address
func isPortOpen(c chan bool, ipaddr string, port int) {

	result := false

	connection, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ipaddr, port), time.Duration(2)*time.Second)

	if err == nil {
		result = true
	}

	if connection != nil {
		connection.Close()
	}

	c <- result
}

//getServerHeader just retrieve server header on web
func getHeader(ipaddr string, headerKey string) string {
	serverHeader := ""
	hostUrl := fmt.Sprintf("http://%s", ipaddr)

	headers, err := getWebHeaders(hostUrl)

	if err == nil {
		if len(headers) > 0 {
			serverHeader = headers.Get(headerKey)
		}
	}

	return serverHeader
}
