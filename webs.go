package webvars

import (
	"net/http"
)

//Webvars web sitesi degiskenlerini tutar.
type Webvars struct {
	Name    string
	IPAddr  string
	Server  string
	Os      string
	Panel   string
	Headers http.Header
}
