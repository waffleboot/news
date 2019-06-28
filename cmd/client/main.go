package main

import (
	"net/http"
	"time"

	"github.com/waffleboot/news/client"
	"github.com/waffleboot/news/messaging"
)

func main() {
	http.Handle("/", client.GetMuxRouter(messaging.NewService(messaging.WithTimeout(10*time.Second))))
	http.ListenAndServe(":8000", nil)
}
