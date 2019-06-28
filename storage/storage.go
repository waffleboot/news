package storage

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
)
import "github.com/waffleboot/news/client"

type StorageService struct {
	msgClient *redis.Client
}

func NewStorageService() *StorageService {
	msgClient := redis.NewClient(&redis.Options{
		Addr:     "messaging:6379",
		Password: "",
		DB:       0,
	})
	return &StorageService{msgClient: msgClient}
}

func (f StorageService) CreateNews(newsobj client.News) (client.NewsId, error) {
	newsid := uuid.New()
	newsobj.Id = newsid
	obj, err := json.Marshal(newsobj)
	if err != nil {
		return uuid.UUID{}, err
	}
	ch1name := fmt.Sprintf("create-request-%v", newsid)
	ch2name := fmt.Sprintf("create-reply-%v", newsid)

	pubsub := f.msgClient.Subscribe(ch2name)
	defer pubsub.Close()

	ch2 := pubsub.Channel()

	f.msgClient.Publish(ch1name, string(obj))

	fmt.Printf("send %v to %v\n", string(obj), newsid)

	select {
	case <-ch2:
		fmt.Println("received response")
		return newsid, nil
	}
}

func (f StorageService) FindNewsById(newsid client.NewsId) (client.News, error) {

	ch1name := fmt.Sprintf("find-request-%v", newsid)
	ch2name := fmt.Sprintf("find-reply-%v", newsid)

	pubsub := f.msgClient.Subscribe(ch2name)
	defer pubsub.Close()

	ch2 := pubsub.Channel()

	f.msgClient.Publish(ch1name, "find")

	select {
	case msg := <-ch2:
		var newsobj client.News
		err := json.Unmarshal([]byte(msg.Payload), &newsobj)
		if err != nil {
			return client.News{}, err
		}
		return newsobj, nil
	}
}
