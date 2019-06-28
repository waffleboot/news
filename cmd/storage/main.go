package main

import (
	"os"
	"os/signal"

	"github.com/go-redis/redis"
)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	msgClient := redis.NewClient(&redis.Options{
		Addr:     "messaging:6379",
		Password: "",
		DB:       0,
	})
	defer msgClient.Close()

	dbClient := redis.NewClient(&redis.Options{
		Addr:     "database:6379",
		Password: "",
		DB:       0,
	})
	defer dbClient.Close()

	pubsubCR := msgClient.PSubscribe("create-request-*")
	pubsubFR := msgClient.PSubscribe("find-request-*")
	defer pubsubCR.Close()
	defer pubsubFR.Close()

	createChannel := pubsubCR.Channel()
	findChannel := pubsubFR.Channel()
	for {
		select {

		case msg := <-createChannel:
			id := msg.Channel[len("create-request-"):]
			dbClient.Set(id, msg.Payload, 0)

			msgClient.Publish("create-reply-"+id, id)

		case msg := <-findChannel:
			id := msg.Channel[len("find-request-"):]

			val, err := dbClient.Get(id).Result()
			if err != nil {
				msgClient.Publish("find-reply-"+id, "")
			} else {
				msgClient.Publish("find-reply-"+id, val)
			}

		case <-c:
			return
		}
	}

}
