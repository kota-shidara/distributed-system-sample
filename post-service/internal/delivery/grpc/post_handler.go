package grpc

import (
	"context"

	"github.com/kota/distributed-system-sample/post-service/domain"
	pb "github.com/kota/distributed-system-sample/post-service/proto/post"
)

type PostHandler struct {
	pb.UnimplementedPostServiceServer
	usecase domain.PostUsecase
}

func NewPostHandler(u domain.PostUsecase) *PostHandler {
	return &PostHandler{
		usecase: u,
	}
}

func (h *PostHandler) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.Post, error) {
	post, err := h.usecase.Create(ctx, req.Title, req.Content, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.Post{
		Id:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		UserId:  post.UserID,
	}, nil
}

func (h *PostHandler) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {
	post, err := h.usecase.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Post{
		Id:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		UserId:  post.UserID,
	}, nil
}

func (h *PostHandler) ListPosts(ctx context.Context, req *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	posts, err := h.usecase.List(ctx)
	if err != nil {
		return nil, err
	}

	var pbPosts []*pb.Post
	for _, p := range posts {
		pbPosts = append(pbPosts, &pb.Post{
			Id:      p.ID,
			Title:   p.Title,
			Content: p.Content,
			UserId:  p.UserID,
		})
	}
	return &pb.ListPostsResponse{Posts: pbPosts}, nil
}

func (h *PostHandler) ListPostsByUser(ctx context.Context, req *pb.ListPostsByUserRequest) (*pb.ListPostsResponse, error) {
	posts, err := h.usecase.ListByUserID(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	var pbPosts []*pb.Post
	for _, p := range posts {
		pbPosts = append(pbPosts, &pb.Post{
			Id:      p.ID,
			Title:   p.Title,
			Content: p.Content,
			UserId:  p.UserID,
		})
	}
	return &pb.ListPostsResponse{Posts: pbPosts}, nil
}
