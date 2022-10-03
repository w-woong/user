package usecase

import (
	"fmt"

	"github.com/w-woong/user/pkg/common"
	"github.com/w-woong/user/pkg/core/model"
	"github.com/w-woong/user/pkg/core/port"
	"github.com/w-woong/user/pkg/dto"
)

type User struct {
	txBeginner common.TxBeginner
	userRepo   port.UserRepo
}

func NewUser(txBeginner common.TxBeginner, userRepo port.UserRepo) *User {
	return &User{
		txBeginner: txBeginner,
		userRepo:   userRepo,
	}
}

func (u *User) FindUserByID(ID string) (dto.User, error) {
	dto, err := u.userRepo.ReadUserByID(ID)
	if err != nil {
		return dto, err
	}

	var user model.User
	common.Scan(&dto, &user)
	fmt.Println(user)

	return dto, nil
}

func (u *User) RegisterUser(input dto.User) error {
	user := model.User{}
	common.Scan(&input, &user)

	user.CreateAndSetID()
	common.Scan(&user, &input)

	tx, err := u.txBeginner.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = u.userRepo.CreateUser(tx, input)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (u *User) ModifyUser(ID string, input dto.User) error {
	_, err := u.userRepo.UpdateUserByID(ID, input)
	return err
}

func (u *User) RemoveUser(ID string) error {
	_, err := u.userRepo.DeleteUserByID(ID)
	return err
}
