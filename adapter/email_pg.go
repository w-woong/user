package adapter

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/user/entity"
	"gorm.io/gorm"
)

type emailPg struct {
	db *gorm.DB
}

func NewEmailPg(db *gorm.DB) *emailPg {
	return &emailPg{
		db: db,
	}
}

func (a *emailPg) Create(ctx context.Context, tx common.TxController, o entity.Email) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Create(&o)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}
func (a *emailPg) Read(ctx context.Context, tx common.TxController, id string) (entity.Email, error) {
	return a.read(ctx, tx.(*txcom.GormTxController).Tx, id)
}
func (a *emailPg) ReadNoTx(ctx context.Context, id string) (entity.Email, error) {
	return a.read(ctx, a.db, id)
}

func (a *emailPg) read(ctx context.Context, db *gorm.DB, id string) (entity.Email, error) {
	out := entity.Email{}
	res := db.WithContext(ctx).
		Where("id = ?", id).
		Limit(1).Find(&out)

	if res.Error != nil {
		logger.Error(res.Error.Error())
		return entity.NilEmail, txcom.ConvertErr(res.Error)
	}
	if res.RowsAffected == 0 {
		logger.Error(common.ErrRecordNotFound.Error())
		return entity.NilEmail, common.ErrRecordNotFound
	}

	return out, nil
}
func (a *emailPg) Update(ctx context.Context, tx common.TxController, o entity.Email) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Save(&o)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}
func (a *emailPg) Delete(ctx context.Context, tx common.TxController, id string) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Delete(&entity.Email{ID: id})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}
