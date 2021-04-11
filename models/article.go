package models

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10"
)

type DBArticle struct {
	tableName   struct{}  `pg:"articles"`
	ID          uint      `pg:"id"`
	URL         string    `pg:"url"`
	Title       string    `pg:"title"`
	Description string    `pg:"description"`
	CreatedAt   time.Time `pg:"created_at"`
	PublishedAt time.Time `pg:"published_at"`
	FeedID      uint      `pg:"feed_id"`
}

var _ pg.BeforeInsertHook = (*DBArticle)(nil)

func (article *DBArticle) BeforeInsert(ctx context.Context) (context.Context, error) {
	article.CreatedAt = time.Now()
	return ctx, nil
}


