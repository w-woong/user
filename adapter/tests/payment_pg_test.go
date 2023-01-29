package adapter_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/user/adapter"
	"github.com/w-woong/user/entity"
)

func Test_paymentTypePg_Create(t *testing.T) {

	if !onlinetest {
		t.Skip("skipping online tests")
	}
	deletePaymentType("TEST_ID")
	defer deletePaymentType("TEST_ID")

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewPaymentTypePg(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	res, err := repo.Create(ctx, tx, entity.PaymentType{
		ID:   "TEST_ID",
		Name: "CREDIT CARD",
	})
	assert.Equal(t, int64(1), res)
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())
}

func Test_paymentTypePg_Read(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	deletePaymentType("TEST_ID")
	createPaymentType("TEST_ID")
	defer deletePaymentType("TEST_ID")

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewPaymentTypePg(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	res, err := repo.Read(ctx, tx, "TEST_ID")
	assert.Equal(t, "TEST_ID", res.ID)
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())

	res, err = repo.ReadNoTx(ctx, "TEST_ID")
	assert.Equal(t, "TEST_ID", res.ID)
	assert.Nil(t, err)
}

func Test_paymentTypePg_Update(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	deletePaymentType("TEST_ID")
	createPaymentType("TEST_ID")
	defer deletePaymentType("TEST_ID")

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewPaymentTypePg(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	tim := time.Now().Add(time.Minute * 3600)
	res, err := repo.Update(ctx, tx, entity.PaymentType{
		ID:        "TEST_ID",
		CreatedAt: &tim,
		UpdatedAt: &tim,
		Name:      "BANK DEPOSIT",
	})
	assert.Equal(t, int64(1), res)
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())
}

func Test_paymentTypePg_Delete(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	deletePaymentType("TEST_ID")
	createPaymentType("TEST_ID")

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewPaymentTypePg(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	res2, err := repo.Delete(ctx, tx, "TEST_ID")
	assert.Equal(t, int64(1), res2)
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())
}

// PaymentMethod

func Test_paymentMethodPg_Create(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	createUser("TEST_ID")
	deletePaymentMethod("TEST_ID")
	deletePaymentType("TEST_ID")
	createPaymentType("TEST_ID")
	defer func() {
		deletePaymentMethod("TEST_ID")
		deletePaymentType("TEST_ID")
	}()

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewPaymentMethodPg(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	res, err := repo.Create(ctx, tx, entity.PaymentMethod{
		ID:            "TEST_ID",
		UserID:        "TEST_ID",
		PaymentTypeID: "TEST_ID",
		PaymentType: entity.PaymentType{
			ID:   "TEST_ID",
			Name: "CREDIT CARD",
		},
		Identity: "1234-12**-****-1234",
		Option:   "0",
	})
	assert.Equal(t, int64(1), res)
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())
}

func Test_paymentMethodPg_Read(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	createUser("TEST_ID")
	deletePaymentMethod("TEST_ID")
	deletePaymentType("TEST_ID")
	createPaymentType("TEST_ID")
	createPaymentMethod("TEST_ID")
	defer func() {
		deletePaymentMethod("TEST_ID")
		deletePaymentType("TEST_ID")
	}()

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewPaymentMethodPg(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	res, err := repo.Read(ctx, tx, "TEST_ID")
	assert.Equal(t, "TEST_ID", res.ID)
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())

	res, err = repo.ReadNoTx(ctx, "TEST_ID")
	assert.Equal(t, "TEST_ID", res.ID)
	assert.Nil(t, err)

}

func Test_paymentMethodPg_Update(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	createUser("TEST_ID")
	deletePaymentMethod("TEST_ID")
	deletePaymentType("TEST_ID")
	createPaymentType("TEST_ID")
	deletePaymentType("TEST_ID2")
	createPaymentType("TEST_ID2")
	createPaymentMethod("TEST_ID")
	defer func() {
		deletePaymentMethod("TEST_ID")
		deletePaymentType("TEST_ID")
		deletePaymentType("TEST_ID2")
	}()

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewPaymentMethodPg(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	tim := time.Now().Add(time.Minute * 3600)
	res, err := repo.Update(ctx, tx, entity.PaymentMethod{
		ID:            "TEST_ID",
		CreatedAt:     &tim,
		UpdatedAt:     &tim,
		UserID:        "TEST_ID",
		PaymentTypeID: "TEST_ID2",
		Identity:      "abcd",
		Option:        "2",
	})
	assert.Equal(t, int64(1), res)
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())
}

func Test_paymentMethodPg_Delete(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	createUser("TEST_ID")
	deletePaymentMethod("TEST_ID")
	deletePaymentType("TEST_ID")
	createPaymentType("TEST_ID")
	createPaymentMethod("TEST_ID")
	defer func() {
		deletePaymentMethod("TEST_ID")
		deletePaymentType("TEST_ID")
	}()

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewPaymentMethodPg(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	res, err := repo.Delete(ctx, tx, "TEST_ID")
	assert.Equal(t, int64(1), res)
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())
}

func deletePaymentType(id string) error {

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewPaymentTypePg(gdb)

	tx, err := txBeginner.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	repo.Delete(ctx, tx, id)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func createPaymentType(id string) error {

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewPaymentTypePg(gdb)

	tx, err := txBeginner.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = repo.Create(ctx, tx, entity.PaymentType{
		ID:   id,
		Name: "CREDIT CARD",
	})
	if err != nil {
		return err
	}
	return tx.Commit()
}

func deletePaymentMethod(id string) error {

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewPaymentMethodPg(gdb)

	tx, err := txBeginner.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	repo.Delete(ctx, tx, id)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func createPaymentMethod(id string) error {

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewPaymentMethodPg(gdb)

	tx, err := txBeginner.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = repo.Create(ctx, tx, entity.PaymentMethod{
		ID:            id,
		UserID:        id,
		PaymentTypeID: id,
		Identity:      "123412******1234",
		Option:        "0",
	})
	if err != nil {
		return err
	}
	return tx.Commit()
}
