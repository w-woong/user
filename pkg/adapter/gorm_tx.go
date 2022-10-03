package adapter

import (
	"github.com/w-woong/user/pkg/common"
	"gorm.io/gorm"
)

type GormTxBeginner struct {
	db *gorm.DB
}

func NewGormTxBeginner(db *gorm.DB) *GormTxBeginner {
	return &GormTxBeginner{
		db: db,
	}
}

func (a *GormTxBeginner) Begin() (common.TxController, error) {
	tx := a.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return NewGormTxController(tx), nil
}

type GormTxController struct {
	Tx *gorm.DB
}

func NewGormTxController(tx *gorm.DB) *GormTxController {
	return &GormTxController{
		Tx: tx,
	}
}

func (a *GormTxController) Commit() error {
	return a.Tx.Commit().Error
}

func (a *GormTxController) Rollback() error {
	return a.Tx.Rollback().Error
}
