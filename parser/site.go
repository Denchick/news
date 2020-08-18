package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"server/articles"
	"strings"
)

func parseSite(url string, rule Rule) (result []articles.Article) {
	for _, link := range findAllArticleLinks(url) {
		html := loadHtml(link)
		if !isArticlePage(rule, html) {
			continue
		}
		article := articles.Article{
			Title: strings.TrimSpace(html.Find(rule.Title).Text()),
			Description: strings.TrimSpace(html.Find(rule.Description).Text()),
			Link: strings.TrimSpace(link),
		}
		result = append(result, article)
	}
	return
}

func findAllArticleLinks(url string) []string {
	doc := loadHtml(url)
	links := make([]string, 0)
	doc.Find("a[href^='" + url + "']").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			links = append(links, href)
		}
	})
	return links
}

func loadHtml(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func isArticlePage(rule Rule, document *goquery.Document) bool {
	return len(document.Find(rule.Exists).Nodes) != 0
}

