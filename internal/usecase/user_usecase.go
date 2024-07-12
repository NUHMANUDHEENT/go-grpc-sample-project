// internal/usecase/user_usecase.go
package usecase

import "github.com/nuhmanudheent/go-microservices/internal/entity"

type UserUseCase interface {
	GetUser(id string) (*entity.User, error)
	
}

type userUseCase struct{}

func NewUserUseCase() UserUseCase {
	return &userUseCase{}
}

func (u *userUseCase) GetUser(id string) (*entity.User, error) {
	// Business logic to get a user
	return &entity.User{ID: id, Name: "John Doe", Email: "john@example.com"}, nil
}
