package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/w-woong/common"
	"github.com/w-woong/user/dto"
	"github.com/w-woong/user/entity"
	"github.com/w-woong/user/port"
	"github.com/wonksing/structmapper"
)

func init() {
	structmapper.StoreMapper(&dto.User{}, &entity.User{})
	structmapper.StoreMapper(&entity.User{}, &dto.User{})
}

type User struct {
	txBeginner     port.TxBeginner
	userRepo       port.UserRepo
	defaultTimeout time.Duration
}

func NewUser(txBeginner port.TxBeginner, userRepo port.UserRepo, defaultTimeout time.Duration) *User {
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
	if err := structmapper.Map(&userDto, &user); err != nil {
		return dto.NilUser, err
	}

	if err = user.PrepareToRegister(); err != nil {
		return dto.NilUser, err
	}

	rowsAffected, err := u.userRepo.CreateUser(ctx, tx, user)
	if err != nil {
		return dto.NilUser, err
	}
	if rowsAffected != 1 {
		return dto.NilUser, common.ErrCreateUser
	}

	if err = tx.Commit(); err != nil {
		return dto.NilUser, err
	}

	res := dto.User{}
	if err = structmapper.Map(&user, &res); err != nil {
		return dto.NilUser, err
	}

	return res, nil
}

func (u *User) FindUserByID(ID string) (dto.User, error) {
	user, err := u.userRepo.ReadUserByID(ID)
	if err != nil {
		return dto.NilUser, err
	}

	res := dto.User{}
	if err = structmapper.Map(&user, &res); err != nil {
		return dto.NilUser, err
	}

	return res, nil
}

// takenLoginID checks if loginID is already taken.
// Returns nil if loginID is available.
func (u *User) takenLoginID(ctx context.Context, tx port.TxController, loginID string) error {
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
	conv := entity.User{}
	if err := structmapper.Map(&user, &conv); err != nil {
		return err
	}
	_, err := u.userRepo.UpdateUserByID(ID, conv)
	return err
}

func (u *User) RemoveUser(ID string) error {
	_, err := u.userRepo.DeleteUserByID(ID)
	return err
}
