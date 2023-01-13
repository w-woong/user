package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

var (
	NilAddress = Address{}
)

type Address struct {
	ID        string     `gorm:"primaryKey;type:string;size:64;comment:id" json:"id"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"`
	UpdatedAt *time.Time `gorm:"<-" json:"updated_at"`

	UserID string `gorm:"type:string;size:64;comment:user id" json:"user_id"`

	ZoneCode    string `gorm:"column:zone_code;type:string;size:64;comment:우편번호" json:"zone_code"`
	Address     string `gorm:"column:address;type:string;size:2048;comment:주소" json:"address"`
	AddressEng  string `gorm:"column:address_eng;type:string;size:2048;comment:주소(영문)" json:"address_eng"`
	AddressType string `gorm:"column:address_type;type:string;size:8;comment:주소유형(도로명/지번, R/J)" json:"address_type"`
}

func (e *Address) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e Address) IsNil() bool {
	return e.ID == "" && e.UserID == ""
}

func (e Address) CreateID() string {
	return uuid.New().String()
}

func (e *Address) CreateSetID() {
	e.ID = e.CreateID()
}

func (e *Address) ReferTo(userID string) {
	e.UserID = userID
}

type Addresses []Address

func (e *Addresses) CreateSetID() {
	for i := range *e {
		(*e)[i].ID = (*e)[i].CreateID()
	}
}

func (e *Addresses) ReferTo(userID string) {
	for i := range *e {
		(*e)[i].ReferTo(userID)
	}
}
