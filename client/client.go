package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetRouter(apiStorage ApiStorage) *mux.Router {
	c := cli{ApiStorage: apiStorage}
	r := mux.NewRouter()
	r.HandleFunc("/news", c.createNewsHandler).Methods("POST")
	r.HandleFunc("/news/{id}", c.getNewsByIdHandler).Methods("GET")
	return r
}

type cli struct {
	ApiStorage
}

func (c cli) createNewsHandler(w http.ResponseWriter, r *http.Request) {
	if !hasApplicationJsonContentType(r.Header) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unsupported or absent content type"))
		return
	}
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var newsobj News
	if err := json.NewDecoder(r.Body).Decode(&newsobj); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := c.ApiStorage.CreateNews(newsobj)
	if err != nil {
		if err == ApiStorageTimeout {
			w.WriteHeader(http.StatusRequestTimeout)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header()["Location"] = []string{fmt.Sprintf("/news/%v", id)}
	w.WriteHeader(http.StatusCreated)
}

func (c cli) getNewsByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var ok bool
	var id string
	if id, ok = vars["id"]; !ok || id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newsid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newsobj, err := c.ApiStorage.FindNewsById(NewsId(newsid))
	if err != nil {
		if err == ApiStorageNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(newsobj)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header()["Content-Type"] = []string{application_json}
}
