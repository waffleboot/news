package client

type NewsId uint64

type News struct {
	Id    NewsId `json:"id,omitempty"`
	Date  string `json:"date,omitempty"`
	Title string `json:"title"`
}
