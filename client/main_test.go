package client

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateNews(t *testing.T) {
	req, err := http.NewRequest("POST", "/news", ioutil.NopCloser(strings.NewReader(`{"title":"Hello, world"}`)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header["Content-Type"] = []string{application_json}
	rr := httptest.NewRecorder()
	router := GetMuxRouter(apiStorageMock{})
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Fatalf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
	if rr.Header()["Location"] == nil {
		t.Fatal("handler returned no location")
	}
	for _, location := range rr.Header()["Location"] {
		if location != "/news/123456" {
			t.Fatalf("handler returned wrong location: got %v want %v", location, "/news/123456")
		}
	}
}

func TestFindNewsById(t *testing.T) {
	req, err := http.NewRequest("GET", "/news/123456", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := GetMuxRouter(apiStorageMock{})
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	if !hasApplicationJsonContentType(rr.Header()) {
		t.Fatal("handler returned invalid content type")
	}
	wantbody := `{"id":"123456","date":"now","title":"Hello, world"}
`
	if rr.Body.String() != wantbody {
		t.Fatalf("handler returned wrong body: got %v want %v", rr.Body.String(), wantbody)
	}
}

type apiStorageMock struct{}

func (apiStorageMock) CreateNews(news News) (string, error) {
	return "123456", nil
}

func (apiStorageMock) FindNewsById(newsid string) (News, error) {
	return News{Title: "Hello, world", Id: "123456", Date: "now"}, nil
}
