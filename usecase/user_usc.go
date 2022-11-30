package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/w-woong/common"
	"github.com/w-woong/user/conv"
	"github.com/w-woong/user/dto"
	"github.com/w-woong/user/entity"
	"github.com/w-woong/user/port"
)

type User struct {
	txBeginner     common.TxBeginner
	userRepo       port.UserRepo
	pwRepo         port.PasswordRepo
	defaultTimeout time.Duration
}

func NewUser(txBeginner common.TxBeginner,
	userRepo port.UserRepo, pwRepo port.PasswordRepo, defaultTimeout time.Duration) *User {
	return &User{
		txBeginner:     txBeginner,
		userRepo:       userRepo,
		pwRepo:         pwRepo,
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

	user, err := conv.ToUserEntity(&userDto)
	if err != nil {
		return dto.NilUser, err
	}

	if err := u.takenLoginID(ctx, tx, user.LoginID); err != nil {
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

	return conv.ToUserDto(&user)
}

func (u *User) RegisterGoogleUser(ctx context.Context, userDto dto.User) (dto.User, error) {
	tx, err := u.txBeginner.Begin()
	if err != nil {
		return dto.NilUser, err
	}
	defer tx.Rollback()

	user, err := conv.ToUserEntity(&userDto)
	if err != nil {
		return dto.NilUser, err
	}
	if err := user.GenerateGoogleLoginID(); err != nil {
		return dto.NilUser, err
	}

	if err := u.takenLoginID(ctx, tx, userDto.LoginID); err != nil {
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

	return conv.ToUserDto(&user)
}

func (u *User) FindUser(ctx context.Context, id string) (dto.User, error) {
	user, err := u.userRepo.ReadUserNoTx(ctx, id)
	if err != nil {
		return dto.NilUser, err
	}

	return conv.ToUserDto(&user)
}

// takenLoginID checks if loginID is already taken.
// Returns nil if loginID is available.
func (u *User) takenLoginID(ctx context.Context, tx common.TxController, loginID string) error {
	foundUser, err := u.userRepo.ReadByLoginID(ctx, tx, loginID)
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

func (u *User) RemoveUser(ctx context.Context, ID string) error {
	tx, err := u.txBeginner.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = u.userRepo.DeleteUser(ctx, tx, ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (u *User) LoginWithPassword(ctx context.Context, loginID, password string) error {
	tx, err := u.txBeginner.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	user, err := u.userRepo.ReadByLoginID(ctx, tx, loginID)
	if err != nil {
		return err
	}

	var auth port.Authenticator
	switch user.LoginType {
	case entity.IDLoginType:
		fallthrough
	case entity.EmailLoginType:
		auth = &PasswordAuthenticator{
			UserID:       user.ID,
			Password:     password,
			PasswordRepo: u.pwRepo,
			Tx:           tx,
		}
	default:
		return errors.New("unsupported login_type, " + string(user.LoginType))
	}

	if err = auth.Authenticate(ctx); err != nil {
		return err
	}

	fmt.Println(user)
	return nil
}

type PasswordAuthenticator struct {
	UserID       string
	Password     string
	PasswordRepo port.PasswordRepo
	Tx           common.TxController
}

func (u *PasswordAuthenticator) Authenticate(ctx context.Context) error {
	pw, err := u.PasswordRepo.ReadByUserID(ctx, u.Tx, u.UserID)
	if err != nil {
		return err
	}
	if u.Password != pw.Value {
		return errors.New("incorrect user id and password")
	}
	return nil
}
