package adapter

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/user/dto"
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

	res := tx.(*GormTxController).Tx.WithContext(ctx).Create(&user)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}

func (a *PgUser) ReadUserByID(ctx context.Context, tx common.TxController, ID string) (entity.User, error) {
	return a.readByID(ctx, tx.(*GormTxController).Tx, ID)
}
func (a *PgUser) ReadUserByIDNoTx(ctx context.Context, ID string) (entity.User, error) {
	return a.readByID(ctx, a.db, ID)
}

func (a *PgUser) readByID(ctx context.Context, db *gorm.DB, ID string) (entity.User, error) {
	user := entity.User{}
	res := db.WithContext(ctx).Where("ID = ?", ID).
		First(&user)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return entity.NilUser, ConvertErr(res.Error)
	}

	return user, nil
}

func (a *PgUser) ReadUserByLoginID(ctx context.Context, tx common.TxController, loginID string) (entity.User, error) {
	return a.readByLoginID(ctx, tx.(*GormTxController).Tx, loginID)
}
func (a *PgUser) ReadUserByLoginIDNoTx(ctx context.Context, loginID string) (entity.User, error) {
	return a.readByLoginID(ctx, a.db, loginID)
}

func (a *PgUser) readByLoginID(ctx context.Context, db *gorm.DB, loginID string) (entity.User, error) {
	user := entity.User{}
	res := db.WithContext(ctx).Where("login_id = ?", loginID).
		First(&user)

	if res.Error != nil {
		logger.Error(res.Error.Error())
		return entity.NilUser, ConvertErr(res.Error)
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

func (a *PgUser) DeleteUserByID(ctx context.Context, tx common.TxController, ID string) (int64, error) {
	res := tx.(*GormTxController).Tx.WithContext(ctx).Delete(&dto.User{ID: ID})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}
