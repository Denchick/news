package models

// Feed ...
type Feed struct {
	ID uint `pg:"id"`
	URL string `pg:"url"`
}