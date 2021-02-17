package models

import "time"

// Article is an article model
type Article struct {
	Link        string
	Title       string
	Description string
	CreatedAt   time.Time
}
