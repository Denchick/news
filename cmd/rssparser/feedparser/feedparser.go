package feedparser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/denchick/news/internal/logger"
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
func (fp *FeedParser) Parse(feed *models.DBFeed) ([]*models.DBArticle, error) {
	fp.logger.Logger.Debug().Msgf("Start %s parsing...", feed.Name)
	parsedFeed, err := fp.parser.ParseURL(feed.FeedURL)
	if err != nil {
		return nil, errors.Wrap(err, "feedparser.ParseFeed")
	}

	fp.logger.Logger.Debug().Msg("Got feed, start items parsing...")
	var articles []*models.DBArticle
	for _, item := range parsedFeed.Items {
		article := &models.DBArticle{
			Title:       strings.TrimSpace(item.Title),
			Description: fp.getItemDescription(item),
			URL:         strings.TrimSpace(item.Link),
			FeedID:      feed.ID,
			PublishedAt:   *item.PublishedParsed,
		}
		if len(article.Description) > 0 {
			fp.logger.Logger.Debug().Msgf("OK: %s with desc: '%s'", article.URL, article.Description)
			articles = append(articles, article)
		}
	}
	fp.logger.Logger.Debug().Msgf("Parsed %d articles from %s", len(articles), feed.FeedURL)
	return articles, nil
}

func (fp *FeedParser) getItemDescription(item *gofeed.Item) string {
	var description string
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
