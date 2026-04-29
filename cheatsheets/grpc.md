---
title: Protobuf / gRPC
icon: fa-arrows-left-right
primary: "#244C5A"
lang: protobuf
---

## fa-file-code Proto3 Messages

```protobuf
syntax = "proto3";

package myapp.v1;

option go_package = "myapp/v1";

message User {
  int64 id = 1;
  string name = 2;
  string email = 3;
}
```

## fa-list Scalar Types

```protobuf
message Types {
  double price = 1;
  float rate = 2;
  int32 count = 3;
  int64 id = 4;
  uint32 uid = 5;
  uint64 big_id = 6;
  sint32 neg = 7;
  sint64 neg64 = 8;
  fixed32 hash = 9;
  fixed64 hash64 = 10;
  sfixed32 sneg = 11;
  sfixed64 sneg64 = 12;
  bool active = 13;
  string label = 14;
  bytes data = 15;
}
```

## fa-tags Enums

```protobuf
enum Status {
  STATUS_UNSPECIFIED = 0;
  STATUS_ACTIVE = 1;
  STATUS_INACTIVE = 2;
  STATUS_SUSPENDED = 3;
}

message Account {
  int64 id = 1;
  Status status = 2;
}
```

## fa-object-group Nested Messages

```protobuf
message Order {
  int64 id = 1;

  message Item {
    string product = 1;
    int32 quantity = 2;
    double price = 3;
  }

  repeated Item items = 2;
  double total = 3;
}
```

## fa-repeat Repeated & Optional

```protobuf
message Playlist {
  string name = 1;
  repeated string tags = 2;

  message Track {
    string title = 1;
    string artist = 2;
    int32 duration_sec = 3;
  }

  repeated Track tracks = 3;
}
```

## fa-map Maps

```protobuf
message Config {
  map<string, string> headers = 1;
  map<string, int32> limits = 2;
}
```

## fa-code-branch Oneof

```protobuf
message Payment {
  int64 amount = 1;

  oneof method {
    string credit_card = 2;
    string paypal_id = 3;
    string bank_account = 4;
  }
}
```

## fa-network-wired Service Definition

```protobuf
service UserService {
  rpc GetUser(GetUserRequest) returns (GetUserResponse);
  rpc ListUsers(ListUsersRequest) returns (ListUsersResponse);
  rpc CreateUser(CreateUserRequest) returns (CreateUserResponse);
}

message GetUserRequest {
  int64 id = 1;
}

message GetUserResponse {
  User user = 1;
}
```

## fa-terminal protoc Commands

```bash
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  api/v1/*.proto

protoc --python_out=. --grpc_python_out=. api/v1/*.proto

protoc --js_out=import_style=commonjs,. \
  --grpc-web_out=import_style=commonjs,mode=grpcwebtext,. \
  api/v1/*.proto
```

## fa-server Go gRPC Server

```go
type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user := &pb.User{Id: req.Id, Name: "Alice"}
	return &pb.GetUserResponse{User: user}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	s.Serve(lis)
}
```

## fa-laptop-code Go gRPC Client

```go
func main() {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	resp, _ := client.GetUser(context.Background(), &pb.GetUserRequest{Id: 1})
	fmt.Println(resp.User.Name)
}
```

## fa-arrow-down Streaming (Server)

```protobuf
rpc ListUsers(ListUsersRequest) returns (stream User);
```

```go
func (s *server) ListUsers(req *pb.ListUsersRequest, stream pb.UserService_ListUsersServer) error {
	users := []*pb.User{
		{Id: 1, Name: "Alice"},
		{Id: 2, Name: "Bob"},
	}
	for _, u := range users {
		if err := stream.Send(u); err != nil {
			return err
		}
	}
	return nil
}
```

## fa-arrow-up Streaming (Client)

```protobuf
rpc UploadUsers(stream User) returns (UploadSummary);
```

```go
func (s *server) UploadUsers(stream pb.UserService_UploadUsersServer) error {
	var count int32
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.UploadSummary{Count: count})
		}
		if err != nil {
			return err
		}
		count++
	}
}
```

## fa-arrows-left-right-to-line Streaming (Bidirectional)

```protobuf
rpc Chat(stream ChatMessage) returns (stream ChatMessage);
```

```go
func (s *server) Chat(stream pb.UserService_ChatServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		stream.Send(&pb.ChatMessage{Text: "echo: " + msg.Text})
	}
}
```

## fa-clock Deadline & Cancellation

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

resp, err := client.GetUser(ctx, &pb.GetUserRequest{Id: 1})
if err != nil {
	status, _ := status.FromError(err)
	fmt.Println(status.Code(), status.Message())
}
```

## fa-shield Interceptors/Middleware

```go
func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	log.Printf("%s %v", info.FullMethod, time.Since(start))
	return resp, err
}

func main() {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryInterceptor),
	)
}
```

## fa-triangle-exclamation Error Codes

```go
import "google.golang.org/grpc/codes"
import "google.golang.org/grpc/status"

return nil, status.Error(codes.NotFound, "user not found")
return nil, status.Error(codes.InvalidArgument, "id is required")
return nil, status.Error(codes.PermissionDenied, "access denied")
return nil, status.Error(codes.Internal, "database error")
return nil, status.Error(codes.Unavailable, "service down")
return nil, status.Error(codes.DeadlineExceeded, "request timed out")
return nil, status.Error(codes.AlreadyExists, "email taken")
```

## fa-magnifying-glass Reflection & grpcurl

```go
import "google.golang.org/grpc/reflection"

func main() {
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	reflection.Register(s)
	s.Serve(lis)
}
```

```bash
grpcurl -plaintext localhost:50051 list
grpcurl -plaintext localhost:50051 list myapp.v1.UserService
grpcurl -plaintext localhost:50051 describe myapp.v1.GetUserRequest
grpcurl -plaintext -d '{"id": 1}' localhost:50051 myapp.v1.UserService/GetUser
```
