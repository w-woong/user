package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

var (
	NilPaymentType   = PaymentType{}
	NilPaymentMethod = PaymentMethod{}
)

type PaymentType struct {
	ID        string     `gorm:"primaryKey;type:string;size:64" json:"id"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"`
	UpdatedAt *time.Time `gorm:"<-" json:"updated_at"`

	Name string `gorm:"column:name;type:string;size:1024" json:"name"`
}

func (e *PaymentType) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e PaymentType) IsNil() bool {
	return e.ID == ""
}

type PaymentMethod struct {
	ID        string     `gorm:"primaryKey;type:string;size:64" json:"id"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"`
	UpdatedAt *time.Time `gorm:"<-" json:"updated_at"`

	UserID string `gorm:"type:string;size:64;comment:user id" json:"user_id"`

	PaymentTypeID string      `gorm:"column:payment_type_id;type:string;size:64" json:"payment_type_id"`
	PaymentType   PaymentType `gorm:"foreignKey:PaymentTypeID;references:ID" json:"payment_type"`
	Identity      string      `gorm:"column:identity;type:string;size:2048" json:"identity"`
	Option        string      `gorm:"column:option;type:string;size:64" json:"option"`
}

func (e *PaymentMethod) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e PaymentMethod) CreateID() string {
	return uuid.New().String()
}

func (e *PaymentMethod) CreateSetID() {
	e.ID = e.CreateID()
}

func (e PaymentMethod) IsNil() bool {
	return e.ID == "" && e.UserID == ""
}
