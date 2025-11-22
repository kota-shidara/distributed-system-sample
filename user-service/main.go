package main

import (
	"log"
	"net"

	"github.com/kota/distributed-system-sample/user-service/db"
	delivery "github.com/kota/distributed-system-sample/user-service/internal/delivery/grpc"
	"github.com/kota/distributed-system-sample/user-service/internal/repository/postgres"
	"github.com/kota/distributed-system-sample/user-service/internal/usecase"
	userpb "github.com/kota/distributed-system-sample/user-service/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db.Init()
	db.DB.AutoMigrate(&postgres.User{})

	// Repositories
	userRepo := postgres.NewUserRepository(db.DB)

	// Usecases
	userUsecase := usecase.NewUserUsecase(userRepo)

	// Handlers
	userHandler := delivery.NewUserHandler(userUsecase)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, userHandler)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Printf("user-service listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
