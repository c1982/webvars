package webvars

import "testing"
import "fmt"

func TestGetWebContent(t *testing.T) {

	content, h, err := getWebVars("http://www.maestropanel.com")

	if err != nil {
		t.Errorf(err.Error())
	}

	if content == "" {
		t.Error("Content is null")
	}

	if len(h) == 0 {
		t.Error("Header is null")
	}

	for k := range h {
		fmt.Println(k, ":", h.Get(k))
	}
}
