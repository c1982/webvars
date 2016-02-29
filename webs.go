package webvars

import (
	"net/http"
)

var webServers = make(map[string]string)

func init{
	webServers["apache"] = "linux"
	webServers["nginx"] = "linux"
	webServers["LiteSpeed"] = "linux"	
	webServers["Lighttpd"] = "linux"	
	webServers["Microsoft"] = "windows"	
	webServers["IIS"] = "windows"
	webServers["Microsoft-IIS/8.5"] = "windows"
	webServers["Microsoft-IIS/8.0"] = "windows"
	webServers["Microsoft-IIS/7.5"] = "windows"
	webServers["Microsoft-IIS/7.0"] = "windows"
	webServers["Microsoft-IIS/6.0"] = "windows"
	
	
}

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
