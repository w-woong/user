package adapter

import (
	"context"

	"github.com/w-woong/common/logger"
	"github.com/w-woong/user/dto"
	"github.com/w-woong/user/entity"
	"github.com/w-woong/user/port"
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

func (a *PgUser) ReadUserByID(ID string) (entity.User, error) {
	user := entity.User{}
	res := a.db.Where("ID = ?", ID).First(&user)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return entity.NilUser, ConvertErr(res.Error)
	}

	return user, nil
}

func (a *PgUser) ReadUserByLoginID(ctx context.Context, tx port.TxController, loginID string) (entity.User, error) {
	user := entity.User{}
	res := tx.(*GormTxController).Tx.WithContext(ctx).
		Where("login_id = ?", loginID).
		First(&user)

	if res.Error != nil {
		logger.Error(res.Error.Error())
		return entity.NilUser, ConvertErr(res.Error)
	}

	return user, nil
}

func (a *PgUser) CreateUser(ctx context.Context, tx port.TxController, user entity.User) (int64, error) {

	res := tx.(*GormTxController).Tx.WithContext(ctx).Create(&user)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}

func (a *PgUser) UpdateUserByID(ID string, user entity.User) (int64, error) {
	// res := a.db.Save(&user)
	res := a.db.Model(&dto.User{ID: ID}).Updates(&user)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}

func (a *PgUser) DeleteUserByID(ID string) (int64, error) {
	res := a.db.Delete(&dto.User{ID: ID})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}
