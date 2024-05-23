package web

import (
	"time"
)

type PostCreateRequest struct {
	Title       string    `validate:"required, max=20, min=1" json:"title"`
	Content     string    `validate:"required" json:"content"`
	Status      string    `validate:"required" json:"status"`
	PublishDate time.Time `validate:"required" json:"publish_date"`
}
