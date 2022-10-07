package port

import (
	"context"

	"github.com/w-woong/user/pkg/dto"
)

type UserUsc interface {
	// RegisterUser registers a new user
	RegisterUser(ctx context.Context, input dto.User) (dto.User, error)

	// FindUserByID finds user with ID
	FindUserByID(ID string) (dto.User, error)

	// ModifyUser modifies user information with input
	ModifyUser(ID string, input dto.User) error

	// RemoveUser removes user with ID
	RemoveUser(ID string) error
}
