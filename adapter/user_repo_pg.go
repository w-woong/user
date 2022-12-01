package adapter

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/user/entity"
	"gorm.io/gorm"
)

type PgUser struct {
	db *gorm.DB
}

func NewPgUser(db *gorm.DB) *PgUser {
	return &PgUser{
		db: db,
	}
}

func (a *PgUser) CreateUser(ctx context.Context, tx common.TxController, user entity.User) (int64, error) {

	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).Create(&user)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}

func (a *PgUser) ReadUser(ctx context.Context, tx common.TxController, id string) (entity.User, error) {
	return a.readUser(ctx, tx.(*txcom.GormTxController).Tx, id)
}
func (a *PgUser) ReadUserNoTx(ctx context.Context, id string) (entity.User, error) {
	return a.readUser(ctx, a.db, id)
}

func (a *PgUser) readUser(ctx context.Context, db *gorm.DB, id string) (entity.User, error) {
	user := entity.User{}
	res := db.WithContext(ctx).Where("id = ?", id).
		First(&user)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return entity.NilUser, txcom.ConvertErr(res.Error)
	}

	return user, nil
}

func (a *PgUser) ReadByLoginID(ctx context.Context, tx common.TxController, loginID string) (entity.User, error) {
	return a.readByLoginID(ctx, tx.(*txcom.GormTxController).Tx, loginID)
}
func (a *PgUser) ReadByLoginIDNoTx(ctx context.Context, loginID string) (entity.User, error) {
	return a.readByLoginID(ctx, a.db, loginID)
}

func (a *PgUser) readByLoginID(ctx context.Context, db *gorm.DB, loginID string) (entity.User, error) {
	user := entity.User{}
	res := db.WithContext(ctx).Where("login_id = ?", loginID).
		First(&user)

	if res.Error != nil {
		logger.Error(res.Error.Error())
		return entity.NilUser, txcom.ConvertErr(res.Error)
	}

	return user, nil
}

// func (a *PgUser) UpdateUserByID(ID string, user entity.User) (int64, error) {
// 	// res := a.db.Save(&user)
// 	res := a.db.Model(&entity.User{}).
// 		Where("id = ?", user.ID).
// 		Updates(&user)
// 	if res.Error != nil {
// 		logger.Error(res.Error.Error())
// 		return 0, ConvertErr(res.Error)
// 	}

// 	return res.RowsAffected, nil
// }

func (a *PgUser) DeleteUser(ctx context.Context, tx common.TxController, id string) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).Delete(&entity.User{ID: id})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}
