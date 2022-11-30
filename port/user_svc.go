package port

import (
	"context"

	"github.com/w-woong/user/dto"
)

type UserSvc interface {
	RegisterGoogleUser(ctx context.Context, user dto.User) (dto.User, error)
}
