package client

import (
	"net/http"
)

const application_json string = "application/json"

func isApplicationJson(s string) bool {
	return len(s) >= len(application_json) && s[:len(application_json)] == application_json
}

func hasApplicationJsonContentType(h http.Header) bool {
	for _, ct := range h["Content-Type"] {
		if isApplicationJson(ct) {
			return true
		}
	}
	return false
}
