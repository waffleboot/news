package client

import (
	"net/http"
	"testing"
)

func TestValidApplicationJson(t *testing.T) {
	h := make(http.Header)
	for _, ct := range []string{"application/json", "application/json; charset=UTF-8"} {
		h["Content-Type"] = []string{ct}
		if hasApplicationJsonContentType(h) == false {
			t.Fatal(ct)
		}
	}
}

func TestInvalidApplicationJson(t *testing.T) {
	h := make(http.Header)
	for _, ct := range []string{"application/jso_"} {
		h["Content-Type"] = []string{ct}
		if hasApplicationJsonContentType(h) {
			t.Fatal(ct)
		}
	}
}
