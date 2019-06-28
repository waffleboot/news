package main

import (
	"net/http"

	"github.com/waffleboot/news/client"
)

type LiteApiStorage struct {
	news map[client.NewsId]client.News
}

func (apiStorage *LiteApiStorage) CreateNews(newsobj client.News) (client.NewsId, error) {
	id := client.NewsId(len(apiStorage.news) + 1)
	newsobj.Id = id
	newsobj.Date = "now"
	apiStorage.news[id] = newsobj
	return id, nil
}

func (apiStorage *LiteApiStorage) FindNewsById(newsId client.NewsId) (client.News, error) {
	if newsobj, ok := apiStorage.news[newsId]; ok {
		return newsobj, nil
	}
	return client.News{}, client.ApiStorageNotFound
}

func main() {
	http.Handle("/", client.GetRouter(&LiteApiStorage{news: make(map[client.NewsId]client.News)}))
	http.ListenAndServe(":8000", nil)
}
