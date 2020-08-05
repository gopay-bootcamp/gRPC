# gRPC
sample gRPC for book example

# Requirements

1. go SDK 1.14.6
2. google.golang.org/grpc
3. protoc for Go
4. protobuff plugin for Go

# Run Instructions
1. Clone the repository in `$GOPATH` of your system.
2. Set pwd to the project directory by `cd /go/src/gRPC`
3. Run the server by `go run example.com/server/server.go`
4. In a new terminal window run the client by `go run example.com/client/client.go`

# Development and Play-around Instructions

You can change services.proto file to modify services. Remove services.pb.go before moving ahead.

1. Modify services.go file as per requirements.
2. Generate the services.pb.go file by running `protoc --go_out=plugins=grpc:. example.com/services/services.proto`
3. Follow the run instructions as stated above to run the gRPC with modified services.
