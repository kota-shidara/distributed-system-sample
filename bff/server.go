package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kota/distributed-system-sample/bff/graph"
	postpb "github.com/kota/distributed-system-sample/post-service/proto/post"
	userpb "github.com/kota/distributed-system-sample/user-service/proto/user"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Connect to User Service
	userServiceURL := os.Getenv("USER_SERVICE_URL")
	if userServiceURL == "" {
		userServiceURL = "localhost:50051"
	}
	userConn, err := grpc.NewClient(userServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to user service: %v", err)
	}
	defer userConn.Close()

	// Connect to Post Service
	postServiceURL := os.Getenv("POST_SERVICE_URL")
	if postServiceURL == "" {
		postServiceURL = "localhost:50052"
	}
	postConn, err := grpc.NewClient(postServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to post service: %v", err)
	}
	defer postConn.Close()

	userClient := userpb.NewUserServiceClient(userConn)
	postClient := postpb.NewPostServiceClient(postConn)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UserClient: userClient,
		PostClient: postClient,
	}}))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
