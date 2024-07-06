// internal/interfaces/grpc/user_handler.go
package grpc

import (
	"context"

	pb "github.com/nuhmanudheent/go-microservices/api/grpc/protos"
	"github.com/nuhmanudheent/go-microservices/internal/usecase"
)

type UserHandler struct {
	useCase usecase.UserUseCase
	pb.UnimplementedUserServiceServer
}

func NewUserHandler(u usecase.UserUseCase) *UserHandler {
	return &UserHandler{useCase: u}
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := h.useCase.GetUser(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{Id: user.ID, Name: user.Name, Email: user.Email}, nil
}
