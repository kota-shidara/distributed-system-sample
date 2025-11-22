package postgres

import (
	"context"
	"time"

	"github.com/kota/distributed-system-sample/post-service/domain"
	"gorm.io/gorm"
)

type Post struct {
	ID        string `gorm:"primaryKey"`
	Title     string
	Content   string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (p *Post) toDomain() *domain.Post {
	return &domain.Post{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		UserID:    p.UserID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func fromDomainPost(p *domain.Post) *Post {
	return &Post{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		UserID:    p.UserID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) domain.PostRepository {
	return &postRepository{db}
}

func (r *postRepository) Create(ctx context.Context, p *domain.Post) error {
	postModel := fromDomainPost(p)
	return r.db.WithContext(ctx).Create(postModel).Error
}

func (r *postRepository) GetByID(ctx context.Context, id string) (*domain.Post, error) {
	var postModel Post
	if err := r.db.WithContext(ctx).First(&postModel, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return postModel.toDomain(), nil
}

func (r *postRepository) List(ctx context.Context) ([]*domain.Post, error) {
	var postModels []Post
	if err := r.db.WithContext(ctx).Find(&postModels).Error; err != nil {
		return nil, err
	}
	var posts []*domain.Post
	for _, p := range postModels {
		posts = append(posts, p.toDomain())
	}
	return posts, nil
}

func (r *postRepository) ListByUserID(ctx context.Context, userID string) ([]*domain.Post, error) {
	var postModels []Post
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&postModels).Error; err != nil {
		return nil, err
	}
	var posts []*domain.Post
	for _, p := range postModels {
		posts = append(posts, p.toDomain())
	}
	return posts, nil
}
