package adapter_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/user/adapter"
	"github.com/w-woong/user/entity"
)

func TestCreateUser(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	deleteUser("TEST_ID")
	defer deleteUser("TEST_ID")
	var err error

	txBeginner := txcom.NewGormTxBeginner(gdb)
	userRepo := adapter.NewPgUser(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	birthStr := "2022-10-15T09:10:00+00:00"
	birthDate, _ := time.Parse(time.RFC3339, birthStr)

	userID := "TEST_ID"
	personalID := "TEST_ID"
	passwordID := "TEST_ID"
	emailID := "TEST_ID"

	_, err = userRepo.CreateUser(context.Background(), tx, entity.User{
		ID:        userID,
		LoginID:   "test_login_id",
		LoginType: entity.LoginTypeID,
		CredentialPassword: &entity.CredentialPassword{
			ID:     passwordID,
			UserID: userID,
			Value:  "asdfasdfasdf",
		},
		Personal: &entity.Personal{
			ID:        personalID,
			UserID:    userID,
			BirthDate: &birthDate,
		},
		Emails: entity.Emails{
			entity.Email{
				ID:       emailID,
				UserID:   userID,
				Email:    "wonk@wonk.orgg",
				Priority: 0,
			},
		},
	})
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())
}

// func TestCreateUser2(t *testing.T) {
// 	if !onlinetest {
// 		t.Skip("skipping online tests")
// 	}
// 	var err error

// 	txBeginner := txcom.NewGormTxBeginner(gdb)
// 	userRepo := adapter.NewPgUser(gdb)

// 	tx, err := txBeginner.Begin()
// 	assert.Nil(t, err)
// 	defer tx.Rollback()

// 	birthStr := "2022-10-15T18:10:00+09:00"
// 	birthDate, _ := time.Parse(time.RFC3339, birthStr)

// 	userID := uuid.New().String()
// 	personalID := uuid.New().String()
// 	_, err = userRepo.CreateUser(context.Background(), tx, entity.User{
// 		ID:      userID,
// 		LoginID: "wonk1",
// 		Personal: entity.Personal{
// 			ID:        personalID,
// 			UserID:    userID,
// 			BirthDate: &birthDate,
// 		},
// 	})
// 	assert.Nil(t, err)
// 	assert.Nil(t, tx.Commit())
// }

// func TestCreateUser3(t *testing.T) {
// 	if !onlinetest {
// 		t.Skip("skipping online tests")
// 	}
// 	var err error

// 	txBeginner := txcom.NewGormTxBeginner(gdb)
// 	userRepo := adapter.NewPgUser(gdb)

// 	tx, err := txBeginner.Begin()
// 	assert.Nil(t, err)
// 	defer tx.Rollback()

// 	birthStr := "2022-10-15T05:10:00-04:00"
// 	birthDate, _ := time.Parse(time.RFC3339, birthStr)

// 	userID := uuid.New().String()
// 	passwordID := uuid.New().String()
// 	password := entity.Password{
// 		ID:     passwordID,
// 		UserID: userID,
// 		Value:  "asdf",
// 	}

// 	personalID := uuid.New().String()
// 	_, err = userRepo.CreateUser(context.Background(), tx, entity.User{
// 		ID:       userID,
// 		LoginID:  "wonk3",
// 		Password: password,
// 		Personal: entity.Personal{
// 			ID:        personalID,
// 			UserID:    userID,
// 			BirthDate: &birthDate,
// 		},
// 	})
// 	assert.Nil(t, err)
// 	assert.Nil(t, tx.Commit())
// }

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
	createUser("TEST_ID")
	defer deleteUser("TEST_ID")

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewPgUser(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	user, err := repo.ReadUser(ctx, tx, "TEST_ID")
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())

	assert.EqualValues(t, "TEST_ID", user.ID)
	fmt.Println(user.String())

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

func createUser(id string) error {
	var err error
	deleteUser(id)

	txBeginner := txcom.NewGormTxBeginner(gdb)
	userRepo := adapter.NewPgUser(gdb)

	tx, err := txBeginner.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	birthStr := "2022-10-15T09:10:00+00:00"
	birthDate, _ := time.Parse(time.RFC3339, birthStr)

	userID := id
	personalID := id
	passwordID := id
	emailID := id
	deliveryAddressID := id

	_, err = userRepo.CreateUser(context.Background(), tx, entity.User{
		ID:        userID,
		LoginID:   "test_login_id",
		LoginType: entity.LoginTypeID,
		CredentialPassword: &entity.CredentialPassword{
			ID:     passwordID,
			UserID: userID,
			Value:  "my_password",
		},
		Personal: &entity.Personal{
			ID:        personalID,
			UserID:    userID,
			BirthDate: &birthDate,
		},
		Emails: []entity.Email{
			{
				ID:       emailID,
				UserID:   userID,
				Email:    "test_login_id@test.test",
				Priority: 0,
			},
		},
		DeliveryAddress: &entity.DeliveryAddress{
			ID:              deliveryAddressID,
			UserID:          userID,
			IsDefault:       true,
			ReceiverName:    "Name",
			ReceiverContact: "000-0000-0000",
			PostCode:        "00000",
			Address:         "asdfasdfasdfasdfasdfasdf",
			AddressDetail:   "asdfasdfasdfasdfasdfasdf",
			DeliveryRequest: entity.DeliveryRequest{
				ID:                    "TEST_ID",
				DeliveryAddressID:     deliveryAddressID,
				DeliveryRequestTypeID: "TEST_ID",
				DeliveryRequestType: entity.DeliveryRequestType{
					ID:   "TEST_ID",
					Name: "At door",
				},
				RequestMessage: "message",
			},
		},
	})

	if err != nil {
		return err
	}
	return tx.Commit()
}

func deleteUser(id string) error {
	deleteEmail(id)
	deletePassword(id)
	deletePersonal(id)
	deleteDeliveryAddressByUserID(id)

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	userRepo := adapter.NewPgUser(gdb)

	tx, err := txBeginner.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = userRepo.DeleteUser(ctx, tx, id)

	if err != nil {
		return err
	}
	return tx.Commit()
}

func deletePassword(id string) error {
	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewPasswordPg(gdb)

	tx, err := txBeginner.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = repo.DeleteByUserID(ctx, tx, id)

	if err != nil {
		return err
	}
	return tx.Commit()
}

func deletePersonal(id string) error {
	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewPersonalPg(gdb)

	tx, err := txBeginner.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = repo.Delete(ctx, tx, id)

	if err != nil {
		return err
	}
	return tx.Commit()
}

func deleteEmail(id string) error {
	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewEmailPg(gdb)

	tx, err := txBeginner.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = repo.Delete(ctx, tx, id)

	if err != nil {
		return err
	}
	return tx.Commit()
}
