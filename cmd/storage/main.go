package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/go-redis/redis"
)

func main() {

	storage := make(map[string]string)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	msgClient := redis.NewClient(&redis.Options{
		Addr:     "messaging:6379",
		Password: "",
		DB:       0,
	})
	defer msgClient.Close()

	pubsub1 := msgClient.PSubscribe("create-request-*")
	pubsub2 := msgClient.PSubscribe("find-request-*")
	defer pubsub1.Close()
	defer pubsub2.Close()

	ch1 := pubsub1.Channel()
	ch2 := pubsub2.Channel()
	for {
		select {
		case msg := <-ch1:
			id := msg.Channel[len("create-request-"):]
			storage[id] = msg.Payload
			msgClient.Publish("create-reply-"+id, id)
			fmt.Println(msg)
		case msg := <-ch2:
			id := msg.Channel[len("find-request-"):]
			r := storage[id]
			msgClient.Publish("find-reply-"+id, r)
			fmt.Println(msg)
		case <-c:
			return
		}
	}

}
