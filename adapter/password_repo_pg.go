package adapter

import (
	"context"

	"github.com/w-woong/common/logger"
	"github.com/w-woong/user/entity"
	"github.com/w-woong/user/port"
	"gorm.io/gorm"
)

type PgPassword struct {
	db *gorm.DB
}

func NewPgPassword(db *gorm.DB) *PgPassword {
	return &PgPassword{
		db: db,
	}
}
func (a *PgPassword) UpdateByUserID(ctx context.Context, tx port.TxController, value string, userID string) (int64, error) {
	// res := a.db.Save(&user)
	res := tx.(*GormTxController).Tx.WithContext(ctx).
		Model(&entity.Password{}).
		Where("user_id = ?", userID).
		Updates(entity.Password{Value: value})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}

func (a *PgPassword) ReadByUserID(ctx context.Context, tx port.TxController, userID string) (entity.Password, error) {
	return a.readByUserID(ctx, tx.(*GormTxController).Tx, userID)
}

func (a *PgPassword) ReadByUserIDNoTx(ctx context.Context, userID string) (entity.Password, error) {
	return a.readByUserID(ctx, a.db, userID)
}

func (a *PgPassword) readByUserID(ctx context.Context, db *gorm.DB, userID string) (entity.Password, error) {
	password := entity.Password{}
	res := db.WithContext(ctx).
		Where("user_id = ?", userID).
		First(&password)

	if res.Error != nil {
		logger.Error(res.Error.Error())
		return entity.NilPassword, ConvertErr(res.Error)
	}

	return password, nil
}

func (a *PgPassword) DeleteByUserID(ctx context.Context, tx port.TxController, userID string) (int64, error) {
	res := tx.(*GormTxController).Tx.WithContext(ctx).
		Where("user_id = ?", userID).
		Delete(&entity.Password{})
	if res.Error != nil {
		return 0, ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}
