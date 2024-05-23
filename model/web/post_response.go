package web

import "time"

type PostResponse struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Status      string    `json:"status"`
	PublishDate time.Time `json:"publish_date"`
}
