// internal/infrastructure/grpc/grpc_server.go
package grpc

import (
	"log"
	"net"

	grpc "google.golang.org/grpc"

	pb "github.com/nuhmanudheent/go-microservices/api/grpc/protos" // Import the generated protobuf code
	localGrpc "github.com/nuhmanudheent/go-microservices/internal/interfaces/grpc"
	"github.com/nuhmanudheent/go-microservices/internal/usecase"
)

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	userUseCase := usecase.NewUserUseCase()
	userHandler := localGrpc.NewUserHandler(userUseCase)
	pb.RegisterUserServiceServer(s, userHandler)

	log.Printf("gRPC server listening on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
