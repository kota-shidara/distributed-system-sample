module github.com/kota/distributed-system-sample/bff

go 1.24.1

replace github.com/kota/distributed-system-sample/user-service => ../user-service

replace github.com/kota/distributed-system-sample/post-service => ../post-service

require (
	github.com/99designs/gqlgen v0.17.83
	github.com/kota/distributed-system-sample/user-service v0.0.0-00010101000000-000000000000
	github.com/kota/distributed-system-sample/post-service v0.0.0-00010101000000-000000000000
	github.com/rs/cors v1.11.1
	github.com/vektah/gqlparser/v2 v2.5.31
	google.golang.org/grpc v1.77.0
)

require (
	github.com/agnivade/levenshtein v1.2.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	golang.org/x/net v0.47.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/text v0.31.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251111163417-95abcf5c77ba // indirect
	google.golang.org/protobuf v1.36.10 // indirect
)
