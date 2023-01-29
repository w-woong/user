package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/w-woong/common"
	commondto "github.com/w-woong/common/dto"
	"github.com/w-woong/user/conv"
	"github.com/w-woong/user/entity"
	"github.com/w-woong/user/port"
)

type User struct {
	txBeginner common.TxBeginner
	userRepo   port.UserRepo
	pwRepo     port.PasswordRepo
}

func NewUser(txBeginner common.TxBeginner,
	userRepo port.UserRepo, pwRepo port.PasswordRepo) *User {
	return &User{
		txBeginner: txBeginner,
		userRepo:   userRepo,
		pwRepo:     pwRepo,
	}
}

func (u *User) RegisterUser(ctx context.Context, userDto commondto.User) (commondto.User, error) {
	// ctx, cancel := context.WithTimeout(ctx, u.defaultTimeout)
	// defer cancel()

	tx, err := u.txBeginner.Begin()
	if err != nil {
		return commondto.NilUser, err
	}
	defer tx.Rollback()

	user, err := conv.ToUserEntity(&userDto)
	if err != nil {
		return commondto.NilUser, err
	}

	user.LoginID, err = user.LoginSource.LoginID(user.LoginID)
	if err != nil {
		return commondto.NilUser, err
	}

	if err := u.takenLoginID(ctx, tx, user.LoginID); err != nil {
		if errors.Is(err, common.ErrLoginIDAlreadyExists) && user.LoginType == entity.LoginTypeToken {
			return u.ModifyUser(ctx, userDto)
		}
		return commondto.NilUser, err
	}

	if err = user.PrepareToRegister(); err != nil {
		return commondto.NilUser, err
	}

	rowsAffected, err := u.userRepo.CreateUser(ctx, tx, user)
	if err != nil {
		return commondto.NilUser, err
	}
	if rowsAffected != 1 {
		return commondto.NilUser, common.ErrCreateUser
	}

	if err = tx.Commit(); err != nil {
		return commondto.NilUser, err
	}

	return conv.ToUserDto(&user)
}

func (u *User) FindUser(ctx context.Context, id string) (commondto.User, error) {
	user, err := u.userRepo.ReadUserNoTx(ctx, id)
	if err != nil {
		return commondto.NilUser, err
	}

	return conv.ToUserDto(&user)
}

func (u *User) FindByLoginID(ctx context.Context, loginSource string, loginID string) (commondto.User, error) {

	loginSourceEntity := entity.LoginSource(loginSource)
	loginIDWithSource, err := loginSourceEntity.LoginID(loginID)
	if err != nil {
		return commondto.NilUser, err
	}

	user, err := u.userRepo.ReadByLoginIDNoTx(ctx, loginIDWithSource)
	if err != nil {
		return commondto.NilUser, err
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

func (u *User) ModifyUser(ctx context.Context, userNew commondto.User) (commondto.User, error) {
	tx, err := u.txBeginner.Begin()
	if err != nil {
		return commondto.NilUser, err
	}
	defer tx.Rollback()

	loginID, err := entity.LoginSource(userNew.LoginSource).LoginID(userNew.LoginID)
	if err != nil {
		return commondto.NilUser, err
	}

	oldUser, err := u.userRepo.ReadByLoginID(ctx, tx, loginID)
	if err != nil {
		return commondto.NilUser, err
	}

	oldUser.CredentialPassword.Value = userNew.CredentialPassword.Value
	oldUser.CredentialToken.Value = userNew.CredentialToken.Value
	oldUser.Personal.FirstName = userNew.Personal.FirstName
	oldUser.Personal.LastName = userNew.Personal.LastName

	_, err = u.pwRepo.UpdateByUserID(ctx, tx, oldUser.CredentialPassword.Value, oldUser.ID)
	if err != nil {
		return commondto.NilUser, err
	}

	oldUserDto, err := conv.ToUserDto(&oldUser)
	if err != nil {
		return commondto.NilUser, err
	}

	return oldUserDto, tx.Commit()
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
	case entity.LoginTypeID:
		fallthrough
	case entity.LoginTypeEmail:
		auth = &PasswordAuthenticator{
			UserID:       user.ID,
			UserPassword: password,
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
	UserPassword string
	PasswordRepo port.PasswordRepo
	Tx           common.TxController
}

func (u *PasswordAuthenticator) Authenticate(ctx context.Context) error {
	pw, err := u.PasswordRepo.ReadByUserID(ctx, u.Tx, u.UserID)
	if err != nil {
		return err
	}
	if u.UserPassword != pw.Value {
		return errors.New("incorrect user id and password")
	}
	return nil
}
