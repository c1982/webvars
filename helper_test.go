package webvars

import (
	"net/http"
	"testing"
)
import _ "fmt"

func TestGetWebContent(t *testing.T) {

	content, h, scode, err := getWebVars("http://www.maestropanel.com")

	if err != nil {
		t.Errorf(err.Error())
	}

	if content == "" {
		t.Error("Content is null")
	}

	if len(h) == 0 {
		t.Error("Header is null")
	}

	if scode != http.StatusOK {
		t.Errorf("Status code is %s", scode)
	}

	//	for k := range h {
	//		fmt.Println(k, ":", h.Get(k))
	//	}
}

func TestGetWebContentNotWorkingWebSite(t *testing.T) {
	content, h, scode, err := getWebVars("http://www.za123za123.com")

	if err == nil {
		t.Error("Error object cannto be nil")
	}

	t.Log("Status Code:", scode)

	if content != "" {
		t.Error("Content must be empty")
	}

	if len(h) > 0 {
		t.Error("Header list is not empty")
	}

	if scode == http.StatusOK {
		t.Errorf("Status code is %s", scode)
	}
}
