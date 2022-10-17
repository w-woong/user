package port

//go:generate mockgen -destination=./mocks/mock_user_usc.go -package=mocks -mock_names=UserUsc=MockUserUsc -source=./user_usc.go . UserUsc

import (
	"context"

	"github.com/w-woong/user/dto"
)

type UserUsc interface {
	// RegisterUser registers a new user
	RegisterUser(ctx context.Context, input dto.User) (dto.User, error)

	// FindUserByID finds user with ID
	FindUserByID(ctx context.Context, ID string) (dto.User, error)

	// ModifyUser modifies user information with input
	// ModifyUser(ID string, input dto.User) error

	// RemoveUser removes user with ID
	RemoveUser(ctx context.Context, ID string) error

	LoginWithPassword(ctx context.Context, loginID, password string) error
}
