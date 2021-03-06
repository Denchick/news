package models

// DBFeed is the database model for feed
type DBFeed struct {
	tableName struct{} `pg:"feeds"`
	ID        uint     `pg:"id"`
	URL       string   `pg:"url"`
	Name      string   `pg:"name"`
	FeedURL   string   `pg:"feed_url"`
}

// DBFeedCategory is the database model for feed category
type DBFeedCategory struct {
	tableName     struct{} `pg:"feeds_categories"`
	ID            uint     `pg:"id"`
	FeedID        uint     `pg:"feed_id"`
	CategoryID    uint     `pg:"category_id"`
	SubcategoryID uint     `pg:"subcategory_id"`
}
