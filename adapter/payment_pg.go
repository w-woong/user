package adapter

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/user/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type paymentTypePg struct {
	db *gorm.DB
}

func NewPaymentTypePg(db *gorm.DB) *paymentTypePg {
	return &paymentTypePg{
		db: db,
	}
}

func (a *paymentTypePg) Create(ctx context.Context, tx common.TxController, o entity.PaymentType) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Create(&o)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}
func (a *paymentTypePg) Read(ctx context.Context, tx common.TxController, id string) (entity.PaymentType, error) {
	return a.read(ctx, tx.(*txcom.GormTxController).Tx, id)
}
func (a *paymentTypePg) ReadNoTx(ctx context.Context, id string) (entity.PaymentType, error) {
	return a.read(ctx, a.db, id)
}

func (a *paymentTypePg) read(ctx context.Context, db *gorm.DB, id string) (entity.PaymentType, error) {
	out := entity.PaymentType{}
	res := db.WithContext(ctx).
		Where("id = ?", id).
		Limit(1).Find(&out)

	if res.Error != nil {
		logger.Error(res.Error.Error())
		return entity.NilPaymentType, txcom.ConvertErr(res.Error)
	}
	if res.RowsAffected == 0 {
		logger.Error(common.ErrRecordNotFound.Error())
		return entity.NilPaymentType, common.ErrRecordNotFound
	}

	return out, nil
}
func (a *paymentTypePg) Update(ctx context.Context, tx common.TxController, o entity.PaymentType) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Save(&o)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}
func (a *paymentTypePg) Delete(ctx context.Context, tx common.TxController, id string) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Delete(&entity.PaymentType{ID: id})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}

// PaymentMethod
type paymentMethodPg struct {
	db *gorm.DB
}

func NewPaymentMethodPg(db *gorm.DB) *paymentMethodPg {
	return &paymentMethodPg{
		db: db,
	}
}

func (a *paymentMethodPg) Create(ctx context.Context, tx common.TxController, o entity.PaymentMethod) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Create(&o)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}
func (a *paymentMethodPg) Read(ctx context.Context, tx common.TxController, id string) (entity.PaymentMethod, error) {
	return a.read(ctx, tx.(*txcom.GormTxController).Tx, id)
}
func (a *paymentMethodPg) ReadNoTx(ctx context.Context, id string) (entity.PaymentMethod, error) {
	return a.read(ctx, a.db, id)
}

func (a *paymentMethodPg) read(ctx context.Context, db *gorm.DB, id string) (entity.PaymentMethod, error) {
	out := entity.PaymentMethod{}
	res := db.WithContext(ctx).
		Preload(clause.Associations).
		Where("id = ?", id).
		Limit(1).Find(&out)

	if res.Error != nil {
		logger.Error(res.Error.Error())
		return entity.NilPaymentMethod, txcom.ConvertErr(res.Error)
	}
	if res.RowsAffected == 0 {
		logger.Error(common.ErrRecordNotFound.Error())
		return entity.NilPaymentMethod, common.ErrRecordNotFound
	}

	return out, nil
}
func (a *paymentMethodPg) Update(ctx context.Context, tx common.TxController, o entity.PaymentMethod) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Save(&o)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}
func (a *paymentMethodPg) Delete(ctx context.Context, tx common.TxController, id string) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Delete(&entity.PaymentMethod{ID: id})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}
