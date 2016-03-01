package webvars

import (
	"testing"
)

//func TestWebVars(t *testing.T) {
//	w, err := GetServerVariables("178.18.196.250")

//	if err != nil {
//		t.Error(err)
//	}

//	if w.WebServer != "Microsoft-IIS/7.5" {
//		t.Error("Web server cannot determine:", w.WebServer)
//	}

//	if w.PoweredBy != "PHP/5.3.13, ASP.NET" {
//		t.Error("X-Powered-by header cannot determine: ", w.PoweredBy)
//	}
//}

//func TestDetectPort(t *testing.T) {

//	hasWeb := detectPorts("178.18.196.250", WEB)

//	if !hasWeb {
//		t.Error("has not web sever")
//	}
//}

func TestDetectClosedPort(t *testing.T) {

	hasWeb := detectPorts("178.18.196.250", SSH, MONGODB, PLESK)

	if hasWeb {
		t.Error("Invalid action")
	}
}
