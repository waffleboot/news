package client

import (
	"errors"
)

var ApiStorageTimeout = errors.New("timeout")
var ApiStorageNotFound = errors.New("not found")

type ApiStorage interface {
	CreateNews(newsobj News) (NewsId, error)
	FindNewsById(newsid NewsId) (News, error)
}
