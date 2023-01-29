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
	createUser("TEST_ID")
	defer deleteUser("TEST_ID")

	var err error

	txBeginner := txcom.NewGormTxBeginner(gdb)
	userRepo := adapter.NewPasswordPg(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	res, err := userRepo.UpdateByUserID(context.Background(), tx, "a", "TEST_ID")
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())
	assert.Equal(t, int64(1), res)
}

func TestReadByUserID(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	createUser("TEST_ID")
	defer deleteUser("TEST_ID")

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewPasswordPg(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	password, err := repo.ReadByUserID(context.Background(), tx, "TEST_ID")
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())

	assert.Equal(t, "TEST_ID", password.ID)
}

func TestReadByUserIDNoTx(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	createUser("TEST_ID")
	defer deleteUser("TEST_ID")

	repo := adapter.NewPasswordPg(gdb)

	password, err := repo.ReadByUserIDNoTx(context.Background(), "TEST_ID")
	assert.Nil(t, err)
	assert.Equal(t, "TEST_ID", password.ID)
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
