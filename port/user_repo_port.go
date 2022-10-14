package port

import (
	"context"

	"github.com/w-woong/user/entity"
)

type UserRepo interface {
	// ReadUserByID reads user by ID.
	ReadUserByID(ID string) (entity.User, error)

	// ReadUserByLoginID reads user by loginID.
	ReadUserByLoginID(ctx context.Context, tx TxController, loginID string) (entity.User, error)

	// CreateUser creates a new user.
	CreateUser(ctx context.Context, tx TxController, user entity.User) (int64, error)

	// UpdateUserByID updates user having ID with user.
	UpdateUserByID(ID string, user entity.User) (int64, error)

	// DeleteUserByID deletes user with ID.
	DeleteUserByID(ID string) (int64, error)
}
