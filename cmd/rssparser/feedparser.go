package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/denchick/news/logger"
	"github.com/denchick/news/models"
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
)

// FeedParser ...
type FeedParser struct {
	logger *logger.Logger
	parser *gofeed.Parser
}

// NewFeedParser create new FeedParser
func NewFeedParser(logger *logger.Logger) *FeedParser {
	return &FeedParser{logger, gofeed.NewParser()}
}

// Parse parses feed by URL
func (fp *FeedParser) Parse(feedURL string) ([]*models.Article, error) {
	fp.logger.Logger.Debug().Msgf("Start %s parsing...", feedURL)
	feed, err := fp.parser.ParseURL(feedURL)
	if err != nil {
		return nil, errors.Wrap(err, "feedparser.ParseFeed")
	}

	fp.logger.Logger.Debug().Msg("Got feed, start items parsing...")
	var articles []*models.Article
	for _, item := range feed.Items {
		article := &models.Article{
			Title:       strings.TrimSpace(item.Title),
			Description: fp.getItemDescription(item),
			Link:        strings.TrimSpace(item.Link),
		}
		fp.logger.Logger.Debug().Msgf("OK: %s", article.Link)
		articles = append(articles, article)
	}
	return articles, nil
}

func (fp *FeedParser) getItemDescription(item *gofeed.Item) string {
	var description string;
	if len(item.Description) > 0 {
		description = item.Description
	} else {
		description = item.Content
	}

	newDescription, err := fp.cleanText(description, 500)
	if err != nil {
		fp.logger.Info().Err(err)
		return ""
	}
	return newDescription
}

func (fp *FeedParser) cleanText(text string, limit int) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(text))
	if err != nil {
		return "", errors.Wrap(err, "feedparser.cleanText")
	}

	text = strings.TrimSpace(doc.Text())
	text = strings.ReplaceAll(text, "\n", " ")
	
	runes := []rune(text)
	if len(runes) >= limit {
		return string(runes[:limit]) + "...", nil
	}
	return text, nil
}