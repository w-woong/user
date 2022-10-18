package port

//go:generate mockgen -destination=./mocks/mock_user_repo.go -package=mocks -mock_names=UserRepo=MockUserRepo -source=./user_repo.go . UserRepo

import (
	"context"

	"github.com/w-woong/user/entity"
)

type UserRepo interface {
	// ReadUserByID reads user by ID.
	ReadUserByID(ctx context.Context, tx TxController, ID string) (entity.User, error)
	ReadUserByIDNoTx(ctx context.Context, ID string) (entity.User, error)

	// ReadUserByLoginID reads user by loginID.
	ReadUserByLoginID(ctx context.Context, tx TxController, loginID string) (entity.User, error)
	ReadUserByLoginIDNoTx(ctx context.Context, loginID string) (entity.User, error)

	// CreateUser creates a new user.
	CreateUser(ctx context.Context, tx TxController, user entity.User) (int64, error)

	// UpdateUserByID updates user having ID with user.
	// UpdateUserByID(ID string, user entity.User) (int64, error)

	// DeleteUserByID deletes user with ID.
	DeleteUserByID(ctx context.Context, tx TxController, ID string) (int64, error)
}

type PasswordRepo interface {
	UpdateByUserID(ctx context.Context, tx TxController, value string, userID string) (int64, error)

	ReadByUserID(ctx context.Context, tx TxController, userID string) (entity.Password, error)
	ReadByUserIDNoTx(ctx context.Context, userID string) (entity.Password, error)

	DeleteByUserID(ctx context.Context, tx TxController, userID string) (int64, error)
}
