package main

import (
	"net/http"

	"github.com/waffleboot/news/client"
	"github.com/waffleboot/news/storage"
)

func main() {
	sf := storage.NewStorageService()
	http.Handle("/", client.GetRouter(sf))
	http.ListenAndServe(":8000", nil)
}
