package models

import "time"

// NewsSubcategory...
type NewsSubcategory struct {
	Name    string        `json:"subcategoryName"`
	Columns []*NewsColumn `json:"columns"`
}

// NewsColumn...
type NewsColumn struct {
	FeedName string        `json:"feedName"`
	URL      string        `json:"feedUrl"`
	Articles []*WebArticle `json:"articles"`
}

// WebArticle...
type WebArticle struct {
	URL         string    `json:"url"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	PublishedAt time.Time `json:"publishedAt"`
}

// ToWebArticle converts DBArticle to WebArticle
func (article *DBArticle) ToWebArticle() *WebArticle {
	return &WebArticle{
		URL:         article.URL,
		Title:       article.Title,
		Description: article.Description,
		CreatedAt:   article.CreatedAt,
		PublishedAt: article.PublishedAt,
	}
}
