package adapter_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/w-woong/user/adapter"
	"github.com/w-woong/user/entity"
)

func TestCreateUser(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	var err error

	txBeginner := adapter.NewGormTxBeginner(gdb)
	userRepo := adapter.NewPgUser(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	birthStr := "2022-10-15T09:10:00+00:00"
	birthDate, _ := time.Parse(time.RFC3339, birthStr)

	_, err = userRepo.CreateUser(context.Background(), tx, entity.User{
		ID:        uuid.New().String(),
		LoginID:   "wonk1",
		FirstName: "wonk",
		LastName:  "sun",
		BirthDate: birthDate,
	})
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())
}

func TestCreateUser2(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	var err error

	txBeginner := adapter.NewGormTxBeginner(gdb)
	userRepo := adapter.NewPgUser(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	birthStr := "2022-10-15T18:10:00+09:00"
	birthDate, _ := time.Parse(time.RFC3339, birthStr)

	_, err = userRepo.CreateUser(context.Background(), tx, entity.User{
		ID:        uuid.New().String(),
		LoginID:   "wonk2",
		FirstName: "wonk",
		LastName:  "sun",
		BirthDate: birthDate,
	})
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())
}

func TestCreateUser3(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	var err error

	txBeginner := adapter.NewGormTxBeginner(gdb)
	userRepo := adapter.NewPgUser(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	birthStr := "2022-10-15T05:10:00-04:00"
	birthDate, _ := time.Parse(time.RFC3339, birthStr)

	userID := uuid.New().String()
	userSecretID := uuid.New().String()
	userSecret := entity.UserSecret{
		ID:     userSecretID,
		UserID: userID,
		Type:   entity.LoginPassword,
		Value:  "asdf",
	}
	userSecrets := make(entity.UserSecrets, 0)
	userSecrets = append(userSecrets, userSecret)

	_, err = userRepo.CreateUser(context.Background(), tx, entity.User{
		ID:          userID,
		LoginID:     "wonk3",
		FirstName:   "wonk",
		LastName:    "sun",
		BirthYear:   2022,
		BirthMonth:  10,
		BirthDay:    15,
		BirthDate:   birthDate,
		Gender:      "M",
		Nationality: "KOR",
		UserSecrets: userSecrets,
	})
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())
}

func TestUpdateUser(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	var err error
	userRepo := adapter.NewPgUser(gdb)
	_, err = userRepo.UpdateUserByID("85bf6aeb-459c-445a-be1e-0b67b8c100ef", entity.User{
		ID:      "85bf6aeb-459c-445a-be1e-0b67b8c100ef",
		LoginID: "wonksing",
	})
	assert.Nil(t, err)
}

func TestUpdateUserBirthDate(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	var err error
	userRepo := adapter.NewPgUser(gdb)

	birthStr := "2022-10-15T09:10:00+00:00"
	birthDate, _ := time.Parse(time.RFC3339, birthStr)

	_, err = userRepo.UpdateUserByID("85bf6aeb-459c-445a-be1e-0b67b8c100ef", entity.User{
		ID:        "e557bccf-7665-46db-a1b6-8e418fed01b3",
		BirthDate: birthDate,
	})
	assert.Nil(t, err)
}

func TestReadByID(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	var err error
	userRepo := adapter.NewPgUser(gdb)

	id := "cdb497b8-3698-41b8-bd4c-605e0e0a0446"
	user, err := userRepo.ReadUserByID(id)
	assert.Nil(t, err)
	fmt.Println(user.String())

	// location, err := time.LoadLocation("America/New_York")
	location, err := time.LoadLocation("Asia/Seoul")
	assert.Nil(t, err)
	fmt.Println(user.BirthDate.In(location))
	assert.Nil(t, err)
}

func TestReadByID2(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	var err error
	userRepo := adapter.NewPgUser(gdb)

	id := "cdb497b8-3698-41b8-bd4c-605e0e0a0446"
	user, err := userRepo.ReadUserByID(id)
	assert.Nil(t, err)
	fmt.Println(user.String())

	location, err := time.LoadLocation("America/New_York")
	// location, err := time.LoadLocation("Asia/Seoul")
	assert.Nil(t, err)
	fmt.Println(user.BirthDate.In(location))
	assert.Nil(t, err)
}
