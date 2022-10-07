package port

import (
	"context"

	"github.com/w-woong/user/pkg/common"
	"github.com/w-woong/user/pkg/dto"
)

type UserRepo interface {
	// ReadUserByID reads user by ID.
	ReadUserByID(ID string) (dto.User, error)

	// ReadUserByLoginID reads user by loginID.
	ReadUserByLoginID(ctx context.Context, tx common.TxController, loginID string) (dto.User, error)

	// CreateUser creates a new user.
	CreateUser(ctx context.Context, tx common.TxController, user dto.User) (int64, error)

	// UpdateUserByID updates user having ID with user.
	UpdateUserByID(ID string, user dto.User) (int64, error)

	// DeleteUserByID deletes user with ID.
	DeleteUserByID(ID string) (int64, error)
}
