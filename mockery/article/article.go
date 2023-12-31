package article

import (
	"context"
	"mockery/author"
	"time"
)

// Article is representing the Article data struct
type Article struct {
	ID        int64         `json:"id"`
	Title     string        `json:"title" validate:"required"`
	Content   string        `json:"content" validate:"required"`
	Author    author.Author `json:"author"`
	UpdatedAt time.Time     `json:"updated_at"`
	CreatedAt time.Time     `json:"created_at"`
}

// ArticleRepository represent the article's repository contract
type ArticleRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Article, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (Article, error)
	GetByTitle(ctx context.Context, title string) (Article, error)
	Update(ctx context.Context, ar *Article) error
	Store(ctx context.Context, a *Article) error
	Delete(ctx context.Context, id int64) error
}

// AuthorRepository represent the author's repository contract
type AuthorRepository interface {
	GetByID(ctx context.Context, id int64) (Article, error)
}
