package grpc

import (
	"context"

	"github.com/kota/distributed-system-sample/user-service/domain"
	pb "github.com/kota/distributed-system-sample/user-service/proto/user"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	usecase domain.UserUsecase
}

func NewUserHandler(u domain.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: u,
	}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	user, err := h.usecase.Create(ctx, req.Name, req.Email)
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	user, err := h.usecase.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (h *UserHandler) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users, err := h.usecase.List(ctx)
	if err != nil {
		return nil, err
	}

	var pbUsers []*pb.User
	for _, u := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		})
	}
	return &pb.ListUsersResponse{Users: pbUsers}, nil
}
