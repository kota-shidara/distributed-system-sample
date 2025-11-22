package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kota/distributed-system-sample/post-service/domain"
)

type postUsecase struct {
	postRepo domain.PostRepository
}

func NewPostUsecase(postRepo domain.PostRepository) domain.PostUsecase {
	return &postUsecase{
		postRepo: postRepo,
	}
}

func (u *postUsecase) Create(ctx context.Context, title, content, userID string) (*domain.Post, error) {
	post := &domain.Post{
		ID:        uuid.New().String(),
		Title:     title,
		Content:   content,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := u.postRepo.Create(ctx, post); err != nil {
		return nil, err
	}

	return post, nil
}

func (u *postUsecase) GetByID(ctx context.Context, id string) (*domain.Post, error) {
	return u.postRepo.GetByID(ctx, id)
}

func (u *postUsecase) List(ctx context.Context) ([]*domain.Post, error) {
	return u.postRepo.List(ctx)
}

func (u *postUsecase) ListByUserID(ctx context.Context, userID string) ([]*domain.Post, error) {
	return u.postRepo.ListByUserID(ctx, userID)
}
