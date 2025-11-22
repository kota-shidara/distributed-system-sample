package domain

import (
	"context"
	"time"
)

type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id string) (*User, error)
	List(ctx context.Context) ([]*User, error)
}

type UserUsecase interface {
	Create(ctx context.Context, name, email string) (*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	List(ctx context.Context) ([]*User, error)
}
