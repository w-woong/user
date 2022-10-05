package port

import (
	"context"

	"github.com/w-woong/user/pkg/dto"
)

type UserUsc interface {
	FindUserByID(ID string) (dto.User, error)

	RegisterUser(ctx context.Context, input dto.User) error

	ModifyUser(ID string, input dto.User) error

	RemoveUser(ID string) error
}
