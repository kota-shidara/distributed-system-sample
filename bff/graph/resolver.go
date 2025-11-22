package graph

import (
	postpb "github.com/kota/distributed-system-sample/post-service/proto/post"
	userpb "github.com/kota/distributed-system-sample/user-service/proto/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserClient userpb.UserServiceClient
	PostClient postpb.PostServiceClient
}
