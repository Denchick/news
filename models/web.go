package models

import "time"

type NewsSubcategory struct {
	Name    string        `json:"subcategoryName"`
	Columns []*NewsColumn `json:"columns"`
}

type NewsColumn struct {
	FeedName string        `json:"feedName"`
	URL      string        `json:"feedUrl"`
	Articles []*WebArticle `json:"articles"`
}

type WebArticle struct {
	URL         string    `json:"url"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	PublishedAt time.Time `json:"publishedAt"`
}

func (article *DBArticle) ToWebArticle() *WebArticle {
	return &WebArticle{
		URL:         article.URL,
		Title:       article.Title,
		Description: article.Description,
		CreatedAt:   article.CreatedAt,
		PublishedAt: article.PublishedAt,
	}
}
