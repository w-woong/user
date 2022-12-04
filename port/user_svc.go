package port

import (
	"context"

	"github.com/w-woong/user/dto"
)

type UserSvc interface {
	RegisterUser(ctx context.Context, user dto.User) (dto.User, error)
}
