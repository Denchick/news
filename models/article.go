package models

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10"
)

// Article is an article model
type Article struct {
	Link        string `pg:"link"`
	Title       string `pg:"title"`
	Description string `pg:"description"`
	CreatedAt   time.Time `pg:"created_at"`
}

var _ pg.BeforeInsertHook = (*Article)(nil)

// BeforeInsert ...
func (article *Article) BeforeInsert(ctx context.Context) (context.Context, error) {
	article.CreatedAt = time.Now()
    return ctx, nil
}