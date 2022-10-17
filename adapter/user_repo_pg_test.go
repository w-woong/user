package adapter_test

import (
	"context"
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

	userID := uuid.New().String()
	personalID := uuid.New().String()
	passwordID := uuid.New().String()

	emails := make(entity.Emails, 0)
	emails = append(emails, entity.Email{
		ID:       uuid.New().String(),
		UserID:   userID,
		Email:    "wonk@wonk.orgg",
		Priority: 0,
	})
	emails = append(emails, entity.Email{
		ID:       uuid.New().String(),
		UserID:   userID,
		Email:    "monk@wonk.orgg",
		Priority: 1,
	})
	_, err = userRepo.CreateUser(context.Background(), tx, entity.User{
		ID:        userID,
		LoginID:   "wonk1",
		LoginType: "id",
		Password: entity.Password{
			ID:     passwordID,
			UserID: userID,
			Value:  "asdfasdfasdf",
		},
		Personal: entity.Personal{
			ID:        personalID,
			UserID:    userID,
			BirthDate: birthDate,
		},
		Emails: emails,
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

	userID := uuid.New().String()
	personalID := uuid.New().String()
	_, err = userRepo.CreateUser(context.Background(), tx, entity.User{
		ID:      userID,
		LoginID: "wonk1",
		Personal: entity.Personal{
			ID:        personalID,
			UserID:    userID,
			BirthDate: birthDate,
		},
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
	passwordID := uuid.New().String()
	password := entity.Password{
		ID:     passwordID,
		UserID: userID,
		Value:  "asdf",
	}

	personalID := uuid.New().String()
	_, err = userRepo.CreateUser(context.Background(), tx, entity.User{
		ID:       userID,
		LoginID:  "wonk3",
		Password: password,
		Personal: entity.Personal{
			ID:        personalID,
			UserID:    userID,
			BirthDate: birthDate,
		},
	})
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())
}

func TestUpdateUser(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	var err error

	userID := "faad3cfb-a23e-4f17-a580-b7e3bcf8de43"
	personalID := "32d31258-ddd2-419c-9511-c1f1bb23687f"
	passwordID := "85b8fe01-d979-4e96-84f1-9f48179ce0fd"
	emails := make(entity.Emails, 0)
	emails = append(emails, entity.Email{
		ID:       "8397895b-0e96-428b-8d35-87a74fc7d0e7",
		UserID:   userID,
		Email:    "wonk@wonk.orgg",
		Priority: 0,
	})
	emails = append(emails, entity.Email{
		ID:       "3266dac3-02ea-49b9-aa63-ce50f9b10266",
		UserID:   userID,
		Email:    "monk@wonk.orgg",
		Priority: 1,
	})
	birthStr := "2011-10-15T09:10:00+00:00"
	birthDate, _ := time.Parse(time.RFC3339, birthStr)

	userRepo := adapter.NewPgUser(gdb)
	_, err = userRepo.UpdateUserByID("faad3cfb-a23e-4f17-a580-b7e3bcf8de43", entity.User{
		ID:        "faad3cfb-a23e-4f17-a580-b7e3bcf8de43",
		LoginID:   "wonksing",
		LoginType: "id",
		Password: entity.Password{
			ID:     passwordID,
			UserID: userID,
			Value:  "a",
		},
		Personal: entity.Personal{
			ID:        personalID,
			UserID:    userID,
			BirthDate: birthDate,
		},
		Emails: emails,
	})
	assert.Nil(t, err)
}

func TestUpdateUserBirthDate(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	// var err error
	// userRepo := adapter.NewPgUser(gdb)

	// birthStr := "2022-10-15T09:10:00+00:00"
	// birthDate, _ := time.Parse(time.RFC3339, birthStr)

	// _, err = userRepo.UpdateUserByID("85bf6aeb-459c-445a-be1e-0b67b8c100ef", entity.User{
	// 	ID:        "e557bccf-7665-46db-a1b6-8e418fed01b3",
	// 	BirthDate: birthDate,
	// })
	// assert.Nil(t, err)
}

func TestReadByID(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	// var err error
	// userRepo := adapter.NewPgUser(gdb)

	// id := "cdb497b8-3698-41b8-bd4c-605e0e0a0446"
	// user, err := userRepo.ReadUserByID(id)
	// assert.Nil(t, err)
	// fmt.Println(user.String())

	// // location, err := time.LoadLocation("America/New_York")
	// location, err := time.LoadLocation("Asia/Seoul")
	// assert.Nil(t, err)
	// fmt.Println(user.BirthDate.In(location))
	// assert.Nil(t, err)
}

func TestReadByID2(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	// var err error
	// userRepo := adapter.NewPgUser(gdb)

	// id := "cdb497b8-3698-41b8-bd4c-605e0e0a0446"
	// user, err := userRepo.ReadUserByID(id)
	// assert.Nil(t, err)
	// fmt.Println(user.String())

	// location, err := time.LoadLocation("America/New_York")
	// // location, err := time.LoadLocation("Asia/Seoul")
	// assert.Nil(t, err)
	// fmt.Println(user.BirthDate.In(location))
	// assert.Nil(t, err)
}
