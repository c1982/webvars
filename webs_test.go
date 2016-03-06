package webvars

import (
	"testing"
)

func TestWebVars(t *testing.T) {
	w, err := GetServerVariables("178.18.196.250")

	if err != nil {
		t.Error(err)
	}

	if w.IsWindows {
		t.Log("Yes it's Windows OS")
	} else {
		t.Error("OS cannot be determine")
	}
}

//func TestDetectPort(t *testing.T) {

//	hasWeb := detectPorts("178.18.196.250", WEB)

//	if !hasWeb {
//		t.Error("has not web sever")
//	}
//}

//func TestDetectClosedPort(t *testing.T) {

//	hasWeb := detectPorts("178.18.196.250", SSH, MONGODB, PLESK)

//	if hasWeb {
//		t.Error("Invalid action")
//	}
//}
