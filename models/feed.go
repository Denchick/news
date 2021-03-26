package models

// Feed ...
type Feed struct {
	ID  uint   `pg:"id"`
	URL string `pg:"url"`
}

type FeedCategory struct {
	tableName  struct{} `pg:"feeds_categories"`
	ID         uint     `pg:"id"`
	FeedID     uint     `pg:"feed_id"`
	CategoryID uint     `pg:"category_id"`
}
