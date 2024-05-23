package domain

import "time"

type Post struct {
	ID          int
	Title       string
	Content     string
	Status      *Status //many to one
	PublishDate time.Time
	Tags        []*Tag //many to many
}
