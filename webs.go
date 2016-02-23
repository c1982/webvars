package webvars

import (
	"net/http"
)

//Webvars web sitesi değişkenlerini tutar.
type Webvars struct {
	Name    string
	IPAddr  string
	Server  string
	Os      string
	Panel   string
	Headers http.Header
}
