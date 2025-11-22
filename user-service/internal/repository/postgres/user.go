package postgres

import (
	"context"
	"time"

	"github.com/kota/distributed-system-sample/user-service/domain"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) toDomain() *domain.User {
	return &domain.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func fromDomainUser(u *domain.User) *User {
	return &User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, u *domain.User) error {
	userModel := fromDomainUser(u)
	return r.db.WithContext(ctx).Create(userModel).Error
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	var userModel User
	if err := r.db.WithContext(ctx).First(&userModel, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return userModel.toDomain(), nil
}

func (r *userRepository) List(ctx context.Context) ([]*domain.User, error) {
	var userModels []User
	if err := r.db.WithContext(ctx).Find(&userModels).Error; err != nil {
		return nil, err
	}
	var users []*domain.User
	for _, u := range userModels {
		users = append(users, u.toDomain())
	}
	return users, nil
}
