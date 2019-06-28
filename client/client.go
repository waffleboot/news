package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMuxRouter(apiStorage ApiStorage) *mux.Router {
	c := cli{ApiStorage: apiStorage}
	r := mux.NewRouter()
	r.HandleFunc("/news", c.createNewsHandler).Methods("POST")
	r.HandleFunc("/news/{id}", c.findNewsByIdHandler).Methods("GET")
	return r
}

type cli struct {
	ApiStorage
}

func (c cli) createNewsHandler(w http.ResponseWriter, r *http.Request) {
	if !hasApplicationJsonContentType(r.Header) {
		w.Header()["X-Reason"] = []string{"unsupported or absent content-type"}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if r.Body == nil {
		w.Header()["X-Reason"] = []string{"empty request body"}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var restobj News
	if err := json.NewDecoder(r.Body).Decode(&restobj); err != nil {
		w.Header()["X-Reason"] = []string{"error on parsing request body"}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := c.ApiStorage.CreateNews(restobj)
	if err != nil {
		if err == ApiStorageTimeout {
			w.Header()["X-Reason"] = []string{"request timeout"}
			w.WriteHeader(http.StatusRequestTimeout)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header()["Location"] = []string{fmt.Sprintf("/news/%v", id)}
	w.WriteHeader(http.StatusCreated)
}

func (c cli) findNewsByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var ok bool
	var id string
	if id, ok = vars["id"]; !ok || id == "" {
		w.Header()["X-Reason"] = []string{"absent id"}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	restobj, err := c.ApiStorage.FindNewsById(id)
	if err != nil {
		if err == ApiStorageNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		} else if err == ApiStorageTimeout {
			w.WriteHeader(http.StatusRequestTimeout)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(restobj)
	if err != nil {
		w.Header()["X-Reason"] = []string{"error on encoding response body"}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header()["Content-Type"] = []string{application_json}
}
