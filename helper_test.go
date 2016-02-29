package webvars

import (
	_ "net/http"
	"testing"
)
import _ "fmt"

//func TestGetWebContent(t *testing.T) {

//	content, h, scode, err := getWebVars("http://www.maestropanel.com")

//	if err != nil {
//		t.Errorf(err.Error())
//	}

//	if content == "" {
//		t.Error("Content is null")
//	}

//	if len(h) == 0 {
//		t.Error("Header is null")
//	}

//	if scode != http.StatusOK {
//		t.Errorf("Status code is %s", scode)
//	}

//	//	for k := range h {
//	//		fmt.Println(k, ":", h.Get(k))
//	//	}
//}

//func TestGetWebContentNotWorkingWebSite(t *testing.T) {
//	content, h, scode, err := getWebVars("http:/mocalhost")

//	if err == nil {
//		t.Error("Error object cannto be nil")
//	}

//	t.Log("Status Code:", scode)

//	if content != "" {
//		t.Error("Content must be empty")
//	}

//	if len(h) > 0 {
//		t.Error("Header list is not empty")
//	}

//	if scode == http.StatusOK {
//		t.Errorf("Status code is ", scode)
//	}
//}

func TestOpenedPorts(t *testing.T) {

	ipaddr := "178.18.196.250"

	web := checkPort(ipaddr, 80)

	if !web {
		t.Error("Web not opened")
	}

	rdp := checkPort(ipaddr, 3389)

	if rdp {
		t.Error("RDP not opened")
	}

	mysql := checkPort(ipaddr, 3306)

	if mysql {
		t.Error("MySQL not opened")
	}
}
