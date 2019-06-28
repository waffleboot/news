package client

type News struct {
	Id    string `json:"id,omitempty"`
	Date  string `json:"date,omitempty"`
	Title string `json:"title"`
}
