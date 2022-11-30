package adapter

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/user/entity"
	"gorm.io/gorm"
)

type passwordPg struct {
	db *gorm.DB
}

func NewPasswordPg(db *gorm.DB) *passwordPg {
	return &passwordPg{
		db: db,
	}
}

func (a *passwordPg) ReadByUserID(ctx context.Context, tx common.TxController, userID string) (entity.Password, error) {
	return a.readByUserID(ctx, tx.(*txcom.GormTxController).Tx, userID)
}

func (a *passwordPg) ReadByUserIDNoTx(ctx context.Context, userID string) (entity.Password, error) {
	return a.readByUserID(ctx, a.db, userID)
}
func (a *passwordPg) UpdateByUserID(ctx context.Context, tx common.TxController, value string, userID string) (int64, error) {
	// res := a.db.Save(&user)
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Model(&entity.Password{}).
		Where("user_id = ?", userID).
		Updates(entity.Password{Value: value})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}

func (a *passwordPg) DeleteByUserID(ctx context.Context, tx common.TxController, userID string) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Where("user_id = ?", userID).
		Delete(&entity.Password{})
	if res.Error != nil {
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}

func (a *passwordPg) readByUserID(ctx context.Context, db *gorm.DB, userID string) (entity.Password, error) {
	password := entity.Password{}
	res := db.WithContext(ctx).
		Where("user_id = ?", userID).
		First(&password)

	if res.Error != nil {
		logger.Error(res.Error.Error())
		return entity.NilPassword, txcom.ConvertErr(res.Error)
	}

	return password, nil
}
