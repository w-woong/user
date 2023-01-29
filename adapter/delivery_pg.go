package adapter

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/user/entity"
	"github.com/w-woong/user/port"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ port.DeliveryAddressRepo = (*deliveryAddressPg)(nil)
var _ port.DeliveryRequestRepo = (*deliveryRequestPg)(nil)
var _ port.DeliveryRequestTypeRepo = (*deliveryRequestTypePg)(nil)

type deliveryRequestTypePg struct {
	db *gorm.DB
}

func NewDeliveryRequestTypePg(db *gorm.DB) *deliveryRequestTypePg {
	return &deliveryRequestTypePg{
		db: db,
	}
}

func (a *deliveryRequestTypePg) Create(ctx context.Context, tx common.TxController, o entity.DeliveryRequestType) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Create(&o)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}
func (a *deliveryRequestTypePg) Read(ctx context.Context, tx common.TxController, id string) (entity.DeliveryRequestType, error) {
	return a.read(ctx, tx.(*txcom.GormTxController).Tx, id)
}
func (a *deliveryRequestTypePg) ReadNoTx(ctx context.Context, id string) (entity.DeliveryRequestType, error) {
	return a.read(ctx, a.db, id)
}

func (a *deliveryRequestTypePg) read(ctx context.Context, db *gorm.DB, id string) (entity.DeliveryRequestType, error) {
	out := entity.DeliveryRequestType{}
	res := db.WithContext(ctx).
		Where("id = ?", id).
		Limit(1).Find(&out)

	if res.Error != nil {
		logger.Error(res.Error.Error())
		return entity.NilDeliveryRequestType, txcom.ConvertErr(res.Error)
	}
	if res.RowsAffected == 0 {
		logger.Error(common.ErrRecordNotFound.Error())
		return entity.NilDeliveryRequestType, common.ErrRecordNotFound
	}

	return out, nil
}
func (a *deliveryRequestTypePg) Update(ctx context.Context, tx common.TxController, o entity.DeliveryRequestType) (int64, error) {

	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Save(&o)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}
func (a *deliveryRequestTypePg) Delete(ctx context.Context, tx common.TxController, id string) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Delete(&entity.DeliveryRequestType{ID: id})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}

type deliveryRequestPg struct {
	db *gorm.DB
}

func NewDeliveryRequestPg(db *gorm.DB) *deliveryRequestPg {
	return &deliveryRequestPg{
		db: db,
	}
}

func (a *deliveryRequestPg) Create(ctx context.Context, tx common.TxController, o entity.DeliveryRequest) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Create(&o)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}
func (a *deliveryRequestPg) Read(ctx context.Context, tx common.TxController, id string) (entity.DeliveryRequest, error) {
	return a.read(ctx, tx.(*txcom.GormTxController).Tx, id)
}
func (a *deliveryRequestPg) ReadNoTx(ctx context.Context, id string) (entity.DeliveryRequest, error) {
	return a.read(ctx, a.db, id)
}

func (a *deliveryRequestPg) read(ctx context.Context, db *gorm.DB, id string) (entity.DeliveryRequest, error) {
	out := entity.DeliveryRequest{}
	res := db.WithContext(ctx).
		Preload(clause.Associations).
		Where("id = ?", id).
		Limit(1).Find(&out)

	if res.Error != nil {
		logger.Error(res.Error.Error())
		return entity.NilDeliveryRequest, txcom.ConvertErr(res.Error)
	}
	if res.RowsAffected == 0 {
		logger.Error(common.ErrRecordNotFound.Error())
		return entity.NilDeliveryRequest, common.ErrRecordNotFound
	}

	return out, nil
}
func (a *deliveryRequestPg) Update(ctx context.Context, tx common.TxController, o entity.DeliveryRequest) (int64, error) {

	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Save(&o)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}
func (a *deliveryRequestPg) Delete(ctx context.Context, tx common.TxController, id string) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Delete(&entity.DeliveryRequest{ID: id})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}

func (a *deliveryRequestPg) DeleteByRequestTypeID(ctx context.Context, tx common.TxController, deliveryRequestTypeID string) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Where("delivery_request_type_id = ?", deliveryRequestTypeID).
		Delete(&entity.DeliveryRequest{})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}

func (a *deliveryRequestPg) DeleteByAddressID(ctx context.Context, tx common.TxController, deliveryAddressID string) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Where("delivery_address_id = ?", deliveryAddressID).
		Delete(&entity.DeliveryRequest{})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}

type deliveryAddressPg struct {
	db *gorm.DB
}

func NewDeliveryAddressPg(db *gorm.DB) *deliveryAddressPg {
	return &deliveryAddressPg{
		db: db,
	}
}

func (a *deliveryAddressPg) Create(ctx context.Context, tx common.TxController, o entity.DeliveryAddress) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Create(&o)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}
func (a *deliveryAddressPg) Read(ctx context.Context, tx common.TxController, id string) (entity.DeliveryAddress, error) {
	return a.read(ctx, tx.(*txcom.GormTxController).Tx, id)
}
func (a *deliveryAddressPg) ReadNoTx(ctx context.Context, id string) (entity.DeliveryAddress, error) {
	return a.read(ctx, a.db, id)
}
func (a *deliveryAddressPg) ReadByUserID(ctx context.Context, tx common.TxController, userID string) (entity.DeliveryAddresses, error) {
	return a.readByUserID(ctx, tx.(*txcom.GormTxController).Tx, userID)
}
func (a *deliveryAddressPg) ReadByUserIDNoTx(ctx context.Context, userID string) (entity.DeliveryAddresses, error) {
	return a.readByUserID(ctx, a.db, userID)
}

func (a *deliveryAddressPg) read(ctx context.Context, db *gorm.DB, id string) (entity.DeliveryAddress, error) {
	out := entity.DeliveryAddress{}
	res := db.WithContext(ctx).
		Preload(clause.Associations).
		Where("id = ?", id).
		Limit(1).Find(&out)

	if res.Error != nil {
		logger.Error(res.Error.Error())
		return entity.NilDeliveryAddress, txcom.ConvertErr(res.Error)
	}
	if res.RowsAffected == 0 {
		logger.Error(common.ErrRecordNotFound.Error())
		return entity.NilDeliveryAddress, common.ErrRecordNotFound
	}

	return out, nil
}

func (a *deliveryAddressPg) readByUserID(ctx context.Context, db *gorm.DB, userID string) (entity.DeliveryAddresses, error) {
	out := entity.DeliveryAddresses{}
	res := db.WithContext(ctx).
		Preload(clause.Associations).
		Where("user_id = ?", userID).
		Find(&out)

	if res.Error != nil {
		logger.Error(res.Error.Error())
		return nil, txcom.ConvertErr(res.Error)
	}

	return out, nil
}

func (a *deliveryAddressPg) Update(ctx context.Context, tx common.TxController, o entity.DeliveryAddress) (int64, error) {

	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Save(&o)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}
func (a *deliveryAddressPg) Delete(ctx context.Context, tx common.TxController, id string) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Delete(&entity.DeliveryAddress{ID: id})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}

func (a *deliveryAddressPg) DeleteByUserID(ctx context.Context, tx common.TxController, userID string) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Where("user_id = ?", userID).
		Delete(&entity.DeliveryAddress{})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}
