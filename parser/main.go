package main

import (
	"github.com/denchick/news/database"
	"log"
)

func parseArticles() []database.Article {
	articlesMap := make(map[string]database.Article)
	for url, rule := range getRules() {
		log.Printf("Start parsing %s\n", url)
		for _, article := range parseSite(url, rule) {
			if len(article.Link) > 0 {
				articlesMap[article.Link] = article
			}
		}
	}
	for _, url := range getFeeds() {
		log.Printf("Start parsing %s\n", url)
		for _, article := range parseFeed(url) {
			if len(article.Link) > 0 {
				articlesMap[article.Link] = article
			}
		}
	}
	return toArray(articlesMap)
}

func toArray(articlesMap map[string]database.Article) []database.Article{
	result := make([]database.Article, len(articlesMap))
	i := 0
	for _, article := range articlesMap {
		result[i] = article
		i++
	}
	return result
}

func main() {
	db, err := database.Connect()
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&database.Article{})
	database.CreateOrUpdateArticles(db, parseArticles())
}