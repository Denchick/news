package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/denchick/news/database"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func parseSite(address string, rule Rule) (result []database.Article) {
	for _, link := range findAllArticleLinks(address) {
		html := loadHtml(link)
		if html == nil || !isArticlePage(rule, html) {
			log.Printf("Not parsed %s", link)
			continue
		}
		log.Printf("Parsed %s", link)
		article := createArticle(html.Find(rule.Title).Text(), html.Find(rule.Description).Text(), link)
		result = append(result, article)
	}
	return
}

func findAllArticleLinks(address string) []string {
	doc := loadHtml(address)
	if doc == nil {
		return make([]string, 0)
	}
	parsedAddress, _ := url.Parse(address)
	links := make([]string, 0)
	doc.Find("a[href^='" + address + "']").Each(func(i int, s *goquery.Selection) {
		href, ok := parseHref(parsedAddress.Host, s)
		if ok {
			links = append(links, href)
		}
	})

	doc.Find("a[href^='/']").Each(func(i int, s *goquery.Selection) {
		href, ok := parseHref(parsedAddress.Host, s)
		if ok {
			links = append(links, href)
		}
	})
	return links
}

func parseHref(host string, s *goquery.Selection) (string, bool) {
	href, exists := s.Attr("href")
	if !exists {
		return "", false
	}
	parsed, err := url.Parse(href)
	if err != nil || parsed.Path == "/" {
		return "", false
	}
	parsedScheme := parsed.Scheme
	if parsedScheme == "" {
		parsedScheme = "https"
	}
	return strings.Join([]string{parsedScheme, "://", host, parsed.Path}, ""), true
}

func loadHtml(address string) *goquery.Document {
	res, err := http.Get(address)
	if err != nil {
		log.Fatalf("%s: %e", address, err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("%s: status code error: %s", address, res.Status)
		return nil
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalf("%s: %e", address, err)
	}
	return doc
}

func isArticlePage(rule Rule, document *goquery.Document) bool {
	return len(document.Find(rule.Exists).Nodes) != 0
}
