package main

import (
	"net/http"

	"github.com/waffleboot/news/client"
	"github.com/waffleboot/news/messaging"
)

func main() {
	http.Handle("/", client.GetMuxRouter(messaging.NewService()))
	http.ListenAndServe(":8000", nil)
}
