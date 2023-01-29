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

func Test_deliveryRequestTypePg_Create(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	deleteDeliveryRequestType("TEST_ID")
	defer deleteDeliveryRequestType("TEST_ID")

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewDeliveryRequestTypePg(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	res, err := repo.Create(ctx, tx, entity.DeliveryRequestType{
		ID:   "TEST_ID",
		Name: "문앞",
	})
	assert.Equal(t, int64(1), res)
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())
}

func Test_deliveryRequestTypePg_Read(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	createDeliveryRequestType("TEST_ID")
	defer deleteDeliveryRequestType("TEST_ID")

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewDeliveryRequestTypePg(gdb)

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

func Test_deliveryRequestTypePg_Update(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	createDeliveryRequestType("TEST_ID")
	defer deleteDeliveryRequestType("TEST_ID")

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewDeliveryRequestTypePg(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	tim := time.Now().Add(time.Minute * 3600)
	res, err := repo.Update(ctx, tx, entity.DeliveryRequestType{
		ID:        "TEST_ID",
		CreatedAt: &tim,
		UpdatedAt: &tim,
		Name:      "New name",
	})
	assert.Equal(t, int64(1), res)
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())
}

func Test_deliveryRequestTypePg_Delete(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	createDeliveryRequestType("TEST_ID")

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewDeliveryRequestTypePg(gdb)

	tx2, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx2.Rollback()

	res2, err := repo.Delete(ctx, tx2, "TEST_ID")
	assert.Equal(t, int64(1), res2)
	assert.Nil(t, err)
	assert.Nil(t, tx2.Commit())
}

// DeliveryAddress

func Test_deliveryAddressPg_Create(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	createUser("TEST_ID")
	deleteDeliveryAddress("TEST_ID")
	defer func() {
		deleteDeliveryAddress("TEST_ID")
		deleteUser("TEST_ID")
	}()

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewDeliveryAddressPg(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	res, err := repo.Create(ctx, tx, entity.DeliveryAddress{
		ID:              "TEST_ID",
		UserID:          "TEST_ID",
		IsDefault:       true,
		ReceiverName:    "receiver",
		ReceiverContact: "0000000000",
		PostCode:        "00000",
		Address:         "asdf asdf asdf",
		AddressDetail:   "afwe awef",
	})
	assert.Equal(t, int64(1), res)
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())
}

func Test_deliveryAddressPg_Read(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	createUser("TEST_ID")
	createDeliveryAddress("TEST_ID")
	defer func() {
		deleteDeliveryAddress("TEST_ID")
		deleteUser("TEST_ID")
	}()

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewDeliveryAddressPg(gdb)

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

func Test_deliveryAddressPg_Update(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	createUser("TEST_ID")
	createDeliveryAddress("TEST_ID")
	createDeliveryRequestType("TEST_ID")
	defer func() {
		deleteDeliveryAddress("TEST_ID")
		deleteDeliveryRequestType("TEST_ID")
		deleteUser("TEST_ID")
	}()

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewDeliveryAddressPg(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	tim := time.Now().Add(time.Minute * 3600)
	res, err := repo.Update(ctx, tx, entity.DeliveryAddress{
		ID:              "TEST_ID",
		CreatedAt:       &tim,
		UpdatedAt:       &tim,
		UserID:          "TEST_ID",
		IsDefault:       true,
		ReceiverName:    "receiver",
		ReceiverContact: "0000000000",
		PostCode:        "00000",
		Address:         "asdf asdf asdf",
		AddressDetail:   "afwe awef",
		DeliveryRequest: entity.DeliveryRequest{
			ID:                    "TEST_ID",
			DeliveryAddressID:     "TEST_ID",
			DeliveryRequestTypeID: "TEST_ID",
			RequestMessage:        "message modified",
		},
	})
	assert.Equal(t, int64(1), res)
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())
}

func Test_deliveryAddressPg_Delete(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	createUser("TEST_ID")
	createDeliveryAddress("TEST_ID")
	deleteDeliveryRequest("TEST_ID")
	defer func() {
		deleteUser("TEST_ID")
	}()

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewDeliveryAddressPg(gdb)

	tx2, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx2.Rollback()

	res2, err := repo.Delete(ctx, tx2, "TEST_ID")
	assert.Equal(t, int64(1), res2)
	assert.Nil(t, err)
	assert.Nil(t, tx2.Commit())
}

// DeliveryRequest

func Test_deliveryRequestPg_Create(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	createDeliveryRequestType("TEST_ID")
	createDeliveryAddress("TEST_ID")
	deleteDeliveryRequest("TEST_ID")
	defer func() {
		deleteDeliveryAddress("TEST_ID")
		deleteDeliveryRequestType("TEST_ID")
	}()
	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewDeliveryRequestPg(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	res, err := repo.Create(ctx, tx, entity.DeliveryRequest{
		ID:                    "TEST_ID",
		DeliveryAddressID:     "TEST_ID",
		DeliveryRequestTypeID: "TEST_ID",
		RequestMessage:        "hello",
	})
	assert.Equal(t, int64(1), res)
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())
}

func Test_deliveryRequestPg_Read(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	createDeliveryAddress("TEST_ID")
	createDeliveryRequest("TEST_ID")
	defer func() {
		deleteDeliveryAddress("TEST_ID")
	}()

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewDeliveryRequestPg(gdb)

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

func Test_deliveryRequestPg_Update(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	createDeliveryAddress("TEST_ID")
	createDeliveryRequest("TEST_ID")
	defer func() {
		deleteDeliveryAddress("TEST_ID")
	}()

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewDeliveryRequestPg(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	tim := time.Now().Add(time.Minute * 3600)
	res, err := repo.Update(ctx, tx, entity.DeliveryRequest{
		ID:                    "TEST_ID",
		CreatedAt:             &tim,
		UpdatedAt:             &tim,
		DeliveryAddressID:     "TEST_ID",
		DeliveryRequestTypeID: "TEST_ID",
		DeliveryRequestType: entity.DeliveryRequestType{
			ID:   "TEST_ID",
			Name: "modified from request",
		},
		RequestMessage: "hello updated",
	})
	assert.Equal(t, int64(1), res)
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())
}

func Test_deliveryRequestPg_Delete(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	createDeliveryAddress("TEST_ID")
	createDeliveryRequest("TEST_ID")
	defer func() {
		deleteDeliveryAddress("TEST_ID")
	}()

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewDeliveryRequestPg(gdb)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	res, err := repo.Delete(ctx, tx, "TEST_ID")
	assert.Equal(t, int64(1), res)
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())
}

func createDeliveryRequestType(id string) error {
	deleteDeliveryRequestType(id)

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewDeliveryRequestTypePg(gdb)

	tx, err := txBeginner.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = repo.Create(ctx, tx, entity.DeliveryRequestType{
		ID:   id,
		Name: "At door",
	})
	if err != nil {
		return err
	}
	return tx.Commit()
}

func deleteDeliveryRequestType(id string) error {
	deleteDeliveryRequestByRequestTypeID(id)

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewDeliveryRequestTypePg(gdb)

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

func createDeliveryAddress(id string) error {
	deleteDeliveryAddress(id)
	createUser(id)

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewDeliveryAddressPg(gdb)

	tx, err := txBeginner.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = repo.Create(ctx, tx, entity.DeliveryAddress{
		ID:              id,
		UserID:          id,
		IsDefault:       true,
		ReceiverName:    "Receiver",
		ReceiverContact: "0000000000",
		PostCode:        "00000",
		Address:         "asdf asdf asdf",
		AddressDetail:   "afwe awef",
	})
	if err != nil {
		return err
	}
	return tx.Commit()
}

func deleteDeliveryAddress(id string) error {
	deleteDeliveryRequestByAddressID(id)

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewDeliveryAddressPg(gdb)

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
func deleteDeliveryAddressByUserID(userID string) error {

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewDeliveryAddressPg(gdb)

	list, _ := repo.ReadByUserIDNoTx(ctx, userID)
	for _, e := range list {
		deleteDeliveryAddress(e.ID)
	}
	deleteDeliveryRequestByAddressID(userID)

	tx, err := txBeginner.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = repo.DeleteByUserID(ctx, tx, userID)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func createDeliveryRequest(id string) error {
	deleteDeliveryRequest(id)
	createDeliveryRequestType(id)
	createDeliveryAddress(id)

	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewDeliveryRequestPg(gdb)

	tx, err := txBeginner.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = repo.Create(ctx, tx, entity.DeliveryRequest{
		ID:                    id,
		DeliveryAddressID:     id,
		DeliveryRequestTypeID: id,
		RequestMessage:        "Hello",
	})
	if err != nil {
		return err
	}
	return tx.Commit()
}

func deleteDeliveryRequest(id string) error {
	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewDeliveryRequestPg(gdb)

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
func deleteDeliveryRequestByRequestTypeID(id string) error {
	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewDeliveryRequestPg(gdb)

	tx, err := txBeginner.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = repo.DeleteByRequestTypeID(ctx, tx, id)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func deleteDeliveryRequestByAddressID(id string) error {
	var err error
	ctx := context.Background()

	txBeginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewDeliveryRequestPg(gdb)

	tx, err := txBeginner.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = repo.DeleteByAddressID(ctx, tx, id)
	if err != nil {
		return err
	}
	return tx.Commit()
}
