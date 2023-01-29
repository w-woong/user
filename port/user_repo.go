package port

//go:generate mockgen -destination=./mocks/mock_user_repo.go -package=mocks -mock_names=UserRepo=MockUserRepo -source=./user_repo.go . UserRepo

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/user/entity"
)

type UserRepo interface {
	// CreateUser creates a new user.
	CreateUser(ctx context.Context, tx common.TxController, user entity.User) (int64, error)

	// ReadUser reads user by ID.
	ReadUser(ctx context.Context, tx common.TxController, id string) (entity.User, error)
	ReadUserNoTx(ctx context.Context, id string) (entity.User, error)

	// ReadByLoginID reads user by loginID.
	ReadByLoginID(ctx context.Context, tx common.TxController, loginID string) (entity.User, error)
	ReadByLoginIDNoTx(ctx context.Context, loginID string) (entity.User, error)

	// UpdateUserByID updates user having ID with user.
	// UpdateUserByID(ID string, user entity.User) (int64, error)

	// DeleteUser deletes user with ID.
	DeleteUser(ctx context.Context, tx common.TxController, id string) (int64, error)
}

type PasswordRepo interface {
	ReadByUserID(ctx context.Context, tx common.TxController, userID string) (entity.CredentialPassword, error)
	ReadByUserIDNoTx(ctx context.Context, userID string) (entity.CredentialPassword, error)

	UpdateByUserID(ctx context.Context, tx common.TxController, value string, userID string) (int64, error)

	DeleteByUserID(ctx context.Context, tx common.TxController, userID string) (int64, error)
}

type PersonalRepo interface {
	Create(ctx context.Context, tx common.TxController, o entity.Personal) (int64, error)
	Read(ctx context.Context, tx common.TxController, id string) (entity.Personal, error)
	ReadNoTx(ctx context.Context, id string) (entity.Personal, error)
	Update(ctx context.Context, tx common.TxController, o entity.Personal) (int64, error)
	Delete(ctx context.Context, tx common.TxController, id string) (int64, error)
}

type EmailRepo interface {
	Create(ctx context.Context, tx common.TxController, o entity.Email) (int64, error)
	Read(ctx context.Context, tx common.TxController, id string) (entity.Email, error)
	ReadNoTx(ctx context.Context, id string) (entity.Email, error)
	Update(ctx context.Context, tx common.TxController, o entity.Email) (int64, error)
	Delete(ctx context.Context, tx common.TxController, id string) (int64, error)
}
