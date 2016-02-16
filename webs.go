package webvars

import (
	"net/http"
)

type Webvars struct {
	Name    string
	IPAddr  string
	Server  string
	Os      string
	Panel   string
	Headers http.Header
}
