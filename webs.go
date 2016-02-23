package webvars

import (
	"net/http"
)

//Webvars web sitesi degiskenlerini tutar.
type Webvars struct {
	Name        string
	IPAddr      string
	IsWordPress bool
	XPoweredBy  string
	Server      string
	headers     http.Header
}

type ServerVars struct {
	IPAddr       string
	Linux        bool
	Windows      bool
	Plesk        bool
	Cpanel       bool
	DirectAdmin  bool
	WebSitePanel bool
	MaestroPanel bool
	EnableGzip   bool
	ReverseDns   string
	TTL          string
}
