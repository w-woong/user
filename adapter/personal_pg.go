package adapter

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/user/entity"
	"gorm.io/gorm"
)

type personalPg struct {
	db *gorm.DB
}

func NewPersonalPg(db *gorm.DB) *personalPg {
	return &personalPg{
		db: db,
	}
}

func (a *personalPg) Create(ctx context.Context, tx common.TxController, o entity.Personal) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Create(&o)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}
func (a *personalPg) Read(ctx context.Context, tx common.TxController, id string) (entity.Personal, error) {
	return a.read(ctx, tx.(*txcom.GormTxController).Tx, id)
}
func (a *personalPg) ReadNoTx(ctx context.Context, id string) (entity.Personal, error) {
	return a.read(ctx, a.db, id)
}

func (a *personalPg) read(ctx context.Context, db *gorm.DB, id string) (entity.Personal, error) {
	out := entity.Personal{}
	res := db.WithContext(ctx).
		Where("id = ?", id).
		Limit(1).Find(&out)

	if res.Error != nil {
		logger.Error(res.Error.Error())
		return entity.NilPersonal, txcom.ConvertErr(res.Error)
	}
	if res.RowsAffected == 0 {
		logger.Error(common.ErrRecordNotFound.Error())
		return entity.NilPersonal, common.ErrRecordNotFound
	}

	return out, nil
}
func (a *personalPg) Update(ctx context.Context, tx common.TxController, o entity.Personal) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Save(&o)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}
func (a *personalPg) Delete(ctx context.Context, tx common.TxController, id string) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Delete(&entity.Personal{ID: id})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}
