package adapter

import (
	commondto "github.com/w-woong/common/dto"
)

// NopUser is a no-op user adapter that impelments port.UserRepo interface.
type NopUser struct{}

// NewNopUser returns new pointer to NopUser
func NewNopUser() *NopUser {
	return &NopUser{}
}

// ReadUserByID returns fixed dto.User
func (a *NopUser) ReadUserByID(ID string) (commondto.User, error) {
	return commondto.User{
		ID:      ID,
		LoginID: "wonk@wonk.orgg",
	}, nil
}
