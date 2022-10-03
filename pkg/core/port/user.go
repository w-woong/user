package port

import (
	"github.com/w-woong/user/pkg/common"
	"github.com/w-woong/user/pkg/dto"
)

type UserRepo interface {
	ReadUserByID(ID string) (dto.User, error)

	CreateUser(tx common.TxController, user dto.User) (int64, error)

	UpdateUserByID(ID string, user dto.User) (int64, error)

	DeleteUserByID(ID string) (int64, error)
}

type UserUsc interface {
	FindUserByID(ID string) (dto.User, error)

	RegisterUser(input dto.User) error

	ModifyUser(ID string, input dto.User) error

	RemoveUser(ID string) error
}
