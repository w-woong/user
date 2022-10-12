package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/w-woong/user/pkg/common"
	"github.com/w-woong/user/pkg/common/mapper"
	"github.com/w-woong/user/pkg/dto"
	"github.com/w-woong/user/pkg/entity"
	"github.com/w-woong/user/pkg/port"
)

func init() {
	mapper.StoreMapper(&dto.User{}, &entity.User{})
	mapper.StoreMapper(&entity.User{}, &dto.User{})
}

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

func (u *User) RegisterUser(ctx context.Context, userDto dto.User) (dto.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.defaultTimeout)
	defer cancel()

	tx, err := u.txBeginner.Begin()
	if err != nil {
		return dto.NilUser, err
	}
	defer tx.Rollback()

	if err := u.takenLoginID(ctx, tx, userDto.LoginID); err != nil {
		return dto.NilUser, err
	}

	user := entity.User{}
	if err := mapper.Map(&userDto, &user); err != nil {
		return dto.NilUser, err
	}

	err = user.PrepareToRegister()
	if err != nil {
		return dto.NilUser, err
	}

	userToCreate := dto.User{}
	if err := mapper.Map(&user, &userToCreate); err != nil {
		return dto.NilUser, err
	}

	rowsAffected, err := u.userRepo.CreateUser(ctx, tx, userToCreate)
	if err != nil {
		return dto.NilUser, err
	}
	if rowsAffected != 1 {
		return dto.NilUser, common.ErrCreateUser
	}

	if err = tx.Commit(); err != nil {
		return dto.NilUser, err
	}

	return userToCreate, nil
}

func (u *User) FindUserByID(ID string) (dto.User, error) {
	user, err := u.userRepo.ReadUserByID(ID)
	if err != nil {
		return dto.NilUser, err
	}

	return user, nil
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
