package domain

import (
	"context"
	"time"
)

type Post struct {
	ID        string
	Title     string
	Content   string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PostRepository interface {
	Create(ctx context.Context, post *Post) error
	GetByID(ctx context.Context, id string) (*Post, error)
	List(ctx context.Context) ([]*Post, error)
	ListByUserID(ctx context.Context, userID string) ([]*Post, error)
}

type PostUsecase interface {
	Create(ctx context.Context, title, content, userID string) (*Post, error)
	GetByID(ctx context.Context, id string) (*Post, error)
	List(ctx context.Context) ([]*Post, error)
	ListByUserID(ctx context.Context, userID string) ([]*Post, error)
}
