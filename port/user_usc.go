package port

//go:generate mockgen -destination=./mocks/mock_user_usc.go -package=mocks -mock_names=UserUsc=MockUserUsc -source=./user_usc.go . UserUsc

import (
	"context"

	"github.com/w-woong/user/dto"
)

type UserUsc interface {
	// RegisterUser registers a new user
	RegisterUser(ctx context.Context, input dto.User) (dto.User, error)

	// FindUser finds user with ID
	FindUser(ctx context.Context, id string) (dto.User, error)
	FindByLoginID(ctx context.Context, loginSource string, loginID string) (dto.User, error)
	// FindByGoogleLoginID(ctx context.Context, loginID string) (dto.User, error)

	// ModifyUserPassword(ctx context.Context)

	// ModifyUser modifies user information with input
	ModifyUser(ctx context.Context, userNew dto.User) (dto.User, error)

	// RemoveUser removes user with ID
	RemoveUser(ctx context.Context, id string) error

	LoginWithPassword(ctx context.Context, loginID, password string) error
}

type Authenticator interface {
	Authenticate(ctx context.Context) error
}
