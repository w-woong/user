package adapter_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w-woong/user/adapter"
)

func TestUpdateValueByUserID(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	var err error

	txBeginner := adapter.NewGormTxBeginner(gdb)
	userRepo := adapter.NewPgPassword(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	userID := "faad3cfb-a23e-4f17-a580-b7e3bcf8de43"

	_, err = userRepo.UpdateValueByUserID(context.Background(), tx, "a", userID)
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())
}

func TestReadByUserID(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	txBeginner := adapter.NewGormTxBeginner(gdb)
	userRepo := adapter.NewPgPassword(gdb)

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

	userRepo := adapter.NewPgPassword(gdb)

	userID := "faad3cfb-a23e-4f17-a580-b7e3bcf8de43"

	password, err := userRepo.ReadByUserIDNoTx(context.Background(), userID)
	assert.Nil(t, err)

	fmt.Println(password)
}
