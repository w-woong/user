package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/w-woong/user/pkg/common"
	"github.com/w-woong/user/pkg/dto"
	"github.com/w-woong/user/pkg/entity"
	"github.com/w-woong/user/pkg/port"
)

type User struct {
	txBeginner     common.TxBeginner
	userRepo       port.UserRepo
	defaultTimeout time.Duration
}

func NewUser(txBeginner common.TxBeginner, userRepo port.UserRepo, defaultTimeout time.Duration) *User {
	return &User{
		txBeginner:     txBeginner,
		userRepo:       userRepo,
		defaultTimeout: defaultTimeout,
	}
}

func (u *User) FindUserByID(ID string) (dto.User, error) {
	user, err := u.userRepo.ReadUserByID(ID)
	if err != nil {
		return dto.NilUser, err
	}

	return user, nil
}

func (u *User) RegisterUser(ctx context.Context, userDto dto.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.defaultTimeout)
	defer cancel()

	tx, err := u.txBeginner.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := u.takenLoginID(ctx, tx, userDto.LoginID); err != nil {
		return err
	}

	user := entity.User{}
	common.ScanStruct(&userDto, &user)

	err = user.PrepareToRegister()
	if err != nil {
		return err
	}

	userToCreate := dto.User{}
	common.ScanStruct(&user, &userToCreate)

	rowsAffected, err := u.userRepo.CreateUser(ctx, tx, userToCreate)
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return common.ErrCreateUser
	}

	return tx.Commit()
}

func (u *User) takenLoginID(ctx context.Context, tx common.TxController, loginID string) error {
	foundUser, err := u.userRepo.ReadUserByLoginID(ctx, tx, loginID)
	if err != nil {
		if !errors.Is(err, common.ErrRecordNotFound) {
			return err
		}
	}
	if !foundUser.IsNil() {
		return common.ErrLoginIDAlreadyExists
	}

	return nil
}

func (u *User) ModifyUser(ID string, user dto.User) error {
	_, err := u.userRepo.UpdateUserByID(ID, user)
	return err
}

func (u *User) RemoveUser(ID string) error {
	_, err := u.userRepo.DeleteUserByID(ID)
	return err
}
