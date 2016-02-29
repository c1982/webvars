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

//getWebHeaders get all headers from web server
func getWebHeaders(url string) (headers http.Header, err error) {
	_, headers, _, err = getWebVars(url)
	return headers, err
}

//isLinux function is determines the server operating system.
func isLinux(ipaddr string) (linux bool, err error) {

	linux := checkPort(ipaddr, 22)
	if linux {
		return linux
	}

	web := checkPort(ipaddr, 80)

	if web {
		name := getServerHeader(ipaddr)

	}

	return linux, err
}

//isWindows function is determines the server operating system.
func isWindows(ipaddr string) (windows bool, err error) {
	return windows, err
}

//checkPort check port from host ip address
func checkPort(ipaddr string, port int) bool {

	fmt.Println(fmt.Sprintf("%s:%d", ipaddr, port))
	connection, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ipaddr, port), time.Duration(3)*time.Second)

	if err != nil {
		return false
	}

	defer connection.Close()

	return true
}

//getServerHeader just retrieve server header on web
func getServerHeader(ipaddr string) string {
	serverHeader := ""
	hostUrl := fmt.Sprintf("http://%s", ipaddr)

	headers, err := getWebHeaders(hostUrl)

	if err == nil {
		if len(headers) > 0 {
			serverHeader = headers.Get("Server")
		}
	}

	return serverHeader
}
