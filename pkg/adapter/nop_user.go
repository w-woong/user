package adapter

import (
	"github.com/w-woong/user/pkg/dto"
)

type NopUser struct {
}

func NewNopUser() *NopUser {
	return &NopUser{}
}

func (a *NopUser) ReadUserByID(ID string) (dto.User, error) {
	return dto.User{
		ID:        ID,
		LoginID:   "wonk@wonk.orgg",
		FirstName: "wonk",
		LastName:  "sun",
	}, nil
}
