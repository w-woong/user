package adapter_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/user/adapter"
)

func TestUpdateByUserID(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	var err error

	txBeginner := txcom.NewGormTxBeginner(gdb)
	userRepo := adapter.NewPasswordPg(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	userID := "faad3cfb-a23e-4f17-a580-b7e3bcf8de43"

	_, err = userRepo.UpdateByUserID(context.Background(), tx, "a", userID)
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())
}

func TestReadByUserID(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	txBeginner := txcom.NewGormTxBeginner(gdb)
	userRepo := adapter.NewPasswordPg(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	userID := "faad3cfb-a23e-4f17-a580-b7e3bcf8de43"

	password, err := userRepo.ReadByUserID(context.Background(), tx, userID)
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())

	fmt.Println(password)
}

func TestReadByUserIDNoTx(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	userRepo := adapter.NewPasswordPg(gdb)

	userID := "faad3cfb-a23e-4f17-a580-b7e3bcf8de43"

	password, err := userRepo.ReadByUserIDNoTx(context.Background(), userID)
	assert.Nil(t, err)

	fmt.Println(password)
}

func TestDeleteByUserID(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	txBeginner := txcom.NewGormTxBeginner(gdb)
	userRepo := adapter.NewPasswordPg(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	userID := "c85f631f-f22c-4be5-bfab-e17d5d37f484"

	rowsAffected, err := userRepo.DeleteByUserID(context.Background(), tx, userID)
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())

	fmt.Println(rowsAffected)
}
