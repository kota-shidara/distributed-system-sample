package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kota/distributed-system-sample/user-service/domain"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) Create(ctx context.Context, name, email string) (*domain.User, error) {
	user := &domain.User{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) GetByID(ctx context.Context, id string) (*domain.User, error) {
	return u.userRepo.GetByID(ctx, id)
}

func (u *userUsecase) List(ctx context.Context) ([]*domain.User, error) {
	return u.userRepo.List(ctx)
}
