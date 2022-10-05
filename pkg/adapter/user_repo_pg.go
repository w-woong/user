package adapter

import (
	"context"

	"github.com/w-woong/user/pkg/common"
	"github.com/w-woong/user/pkg/dto"
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

func (a *PgUser) ReadUserByID(ID string) (dto.User, error) {
	user := dto.User{}
	res := a.db.Where("ID = ?", ID).First(&user)
	if res.Error != nil {
		return dto.NilUser, ConvertErr(res.Error)
	}

	return user, nil
}

func (a *PgUser) ReadUserByLoginID(ctx context.Context, tx common.TxController, loginID string) (dto.User, error) {
	user := dto.User{}
	res := tx.(*GormTxController).Tx.WithContext(ctx).
		Where("login_id = ?", loginID).
		First(&user)

	if res.Error != nil {
		return dto.NilUser, ConvertErr(res.Error)
	}

	return user, nil
}

func (a *PgUser) CreateUser(ctx context.Context, tx common.TxController, user dto.User) (int64, error) {

	res := tx.(*GormTxController).Tx.WithContext(ctx).Create(&user)
	if res.Error != nil {
		return 0, ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}

func (a *PgUser) UpdateUserByID(ID string, user dto.User) (int64, error) {
	// res := a.db.Save(&user)
	res := a.db.Model(&dto.User{ID: ID}).Updates(&user)
	if res.Error != nil {
		return 0, ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}

func (a *PgUser) DeleteUserByID(ID string) (int64, error) {
	res := a.db.Delete(&dto.User{ID: ID})
	if res.Error != nil {
		return 0, ConvertErr(res.Error)
	}
	tx := a.db.Begin()
	tx.Commit()
	tx.Rollback()
	return res.RowsAffected, nil
}
