package adapter

import (
	"github.com/w-woong/user/pkg/dto"
)

// NopUser is a no-op user adapter that impelments port.UserRepo interface.
type NopUser struct{}

// NewNopUser returns new pointer to NopUser
func NewNopUser() *NopUser {
	return &NopUser{}
}

// ReadUserByID returns fixed dto.User
func (a *NopUser) ReadUserByID(ID string) (dto.User, error) {
	return dto.User{
		ID:        ID,
		LoginID:   "wonk@wonk.orgg",
		FirstName: "wonk",
		LastName:  "sun",
	}, nil
}
