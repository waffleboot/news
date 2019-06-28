package client

import (
	"github.com/google/uuid"
)

type NewsId = uuid.UUID

type News struct {
	Id    NewsId `json:"id,omitempty"`
	Date  string `json:"date,omitempty"`
	Title string `json:"title"`
}

func parseNewsId(s string) (NewsId, error) {
	return uuid.Parse(s)
}
