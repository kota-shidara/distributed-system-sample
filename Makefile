.PHONY: proto

proto: proto-user proto-post

proto-user:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		user-service/proto/user/user.proto

proto-post:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		post-service/proto/post/post.proto
