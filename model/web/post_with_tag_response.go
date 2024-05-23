package web

import (
	"time"
)

type PostWithTagResponse struct {
	Id          int           `json:"id"`
	Title       string        `json:"title"`
	Content     string        `json:"content"`
	Status      string        `json:"status"`
	PublishDate time.Time     `json:"publish_date"`
	Tags        []TagResponse `json:"tags"`
}
