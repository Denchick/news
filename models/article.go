package models

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10"
)

type DBArticle struct {
	tableName   struct{}  `pg:"articles"`
	ID          uint      `pg:"id"`
	Link        string    `pg:"link"`
	Title       string    `pg:"title"`
	Description string    `pg:"description"`
	CreatedAt   time.Time `pg:"created_at"`
	FeedID      uint      `pg:"feed_id"`
}

var _ pg.BeforeInsertHook = (*DBArticle)(nil)

func (article *DBArticle) BeforeInsert(ctx context.Context) (context.Context, error) {
	article.CreatedAt = time.Now()
	return ctx, nil
}

type ArticleGroups struct {
	FeedName string       `json:"feedName"`
	Link     string       `json:"link"`
	Articles []*DBArticle `json:"articles"`
}
