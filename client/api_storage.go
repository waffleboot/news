package client

import (
	"errors"
)

var ApiStorageTimeout = errors.New("timeout")
var ApiStorageNotFound = errors.New("not found")

type ApiStorage interface {
	CreateNews(restobj News) (string, error)
	FindNewsById(newsid string) (News, error)
}
