package main

import (
	"log"
	"net"

	"github.com/kota/distributed-system-sample/post-service/db"
	delivery "github.com/kota/distributed-system-sample/post-service/internal/delivery/grpc"
	"github.com/kota/distributed-system-sample/post-service/internal/repository/postgres"
	"github.com/kota/distributed-system-sample/post-service/internal/usecase"
	postpb "github.com/kota/distributed-system-sample/post-service/proto/post"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db.Init()
	db.DB.AutoMigrate(&postgres.Post{})

	// Repositories
	postRepo := postgres.NewPostRepository(db.DB)

	// Usecases
	postUsecase := usecase.NewPostUsecase(postRepo)

	// Handlers
	postHandler := delivery.NewPostHandler(postUsecase)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	postpb.RegisterPostServiceServer(s, postHandler)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Printf("post-service listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
