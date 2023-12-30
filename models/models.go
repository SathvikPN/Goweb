package models

import "time"

// Post schema of post database table
type Post struct {
	ID        int64     `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Body      string    `json:"body,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
