package adapter

import (
	"context"

	"github.com/w-woong/user/dto"
)

type UserSvcNop struct {
}

func NewUserSvcNop() *UserSvcNop {
	return &UserSvcNop{}
}
func (UserSvcNop) RegisterGoogleUser(ctx context.Context, user dto.User) (dto.User, error) {
	return dto.NilUser, nil
}
