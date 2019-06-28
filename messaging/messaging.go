package messaging

//go:generate protoc -I=. --go_out=. news.proto

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/go-redis/redis"
	"github.com/google/uuid"

	"github.com/waffleboot/news/client"
)

type Opts func(s *Service)

type Service struct {
	redisdb *redis.Client
	timeout time.Duration
}

func NewService(opts ...Opts) Service {
	redisdb := redis.NewClient(&redis.Options{
		Addr:     "messaging:6379",
		Password: "",
		DB:       0,
	})
	s := Service{redisdb: redisdb}
	for _, opt := range opts {
		opt(&s)
	}
	return s
}

func WithTimeout(timeout time.Duration) Opts {
	return func(s *Service) {
		s.timeout = timeout
	}
}

func (s Service) CreateNews(restobj client.News) (string, error) {

	newsid := uuid.New().String()

	protoobj := &News{}
	protoobj.Id = newsid
	protoobj.Date = time.Now().Format(time.UnixDate)
	protoobj.Title = restobj.Title

	protoout, err := proto.Marshal(protoobj)
	if err != nil {
		return "", err
	}

	pubsub := s.redisdb.Subscribe(fmt.Sprintf("create-reply-%v", newsid))
	replyChannel := pubsub.Channel()
	defer pubsub.Close()

	s.redisdb.Publish(fmt.Sprintf("create-request-%v", newsid), string(protoout))

	select {
	case <-time.After(s.timeout):
		return "", client.ApiStorageTimeout
	case <-replyChannel:
		return newsid, nil
	}
}

func (s Service) FindNewsById(newsid string) (client.News, error) {

	pubsub := s.redisdb.Subscribe(fmt.Sprintf("find-reply-%v", newsid))
	replyChannel := pubsub.Channel()
	defer pubsub.Close()

	s.redisdb.Publish(fmt.Sprintf("find-request-%v", newsid), ".")

	select {
	case <-time.After(s.timeout):
		return client.News{}, client.ApiStorageTimeout
	case msg := <-replyChannel:
		if msg.Payload == "" {
			return client.News{}, client.ApiStorageNotFound
		}
		protoobj := &News{}
		err := proto.Unmarshal([]byte(msg.Payload), protoobj)
		if err != nil {
			return client.News{}, err
		}
		var restobj client.News
		restobj.Id = protoobj.Id
		restobj.Date = protoobj.Date
		restobj.Title = protoobj.Title
		return restobj, nil
	}
}
