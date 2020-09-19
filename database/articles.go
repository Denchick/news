package database

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	Link        string `gorm:"column:id"`
	Title       string `gorm:"index"`
	Description string
	CreatedAt   time.Time
}

func (a *Article) BeforeCreate(tx *gorm.DB) error {
	a.CreatedAt = time.Now()
	return nil
}

func ReadAllArticles(db *gorm.DB) []Article {
	var articles []Article
	db.Order("created_at").Find(&articles)
	return articles
}

func ReadLastArticles(db *gorm.DB) []Article {
	var articles []Article
	db.Order("created_at").Limit(20).Find(&articles)
	return articles
}

func CreateOrUpdateArticles(db *gorm.DB, articles []Article) {
	addedLinks := readLinks(db)
	added, notAdded := splitArticles(articles, addedLinks)
	db.Create(&notAdded)
	db.Save(&added)
}

func FindArticleByTitle(db *gorm.DB, query string) (articles []Article) {
	pattern := "%" + query + "%"
	db.Where("title LIKE ?", pattern).Find(&articles)
	return
}

func readLinks(db *gorm.DB) []string {
	var links []string
	db.Select("id").Find(&links)
	return links
}

func splitArticles(articles []Article, addedLinks []string) (notAdded []Article, added []Article) {
	for _, article := range articles {
		if contains(addedLinks, article.Link) {
			added = append(added, article)
		} else {
			notAdded = append(notAdded, article)
		}
	}
	return
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
