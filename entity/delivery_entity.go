package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

var (
	NilDeliveryAddress     = DeliveryAddress{}
	NilDeliveryRequest     = DeliveryRequest{}
	NilDeliveryRequestType = DeliveryRequestType{}
)

// DeliveryAddress
// ->User
// <-DeliveryRequest
type DeliveryAddress struct {
	ID        string     `gorm:"primaryKey;type:string;size:64;comment:id" json:"id"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"`
	UpdatedAt *time.Time `gorm:"<-" json:"updated_at"`

	UserID string `gorm:"type:string;size:64;comment:user id" json:"user_id"`

	IsDefault       bool   `gorm:"column:is_default;type:bool;comment:기본배송지" json:"is_default"`
	ReceiverName    string `gorm:"column:receiver_name;type:string;size:1024;comment:받는사람" json:"receiver_name"`
	ReceiverContact string `gorm:"column:receiver_contact;type:string;size:1024;comment:받는사람연락처" json:"receiver_contact"`
	PostCode        string `gorm:"column:post_code;type:string;size:64;comment:우편번호" json:"post_code"`
	Address         string `gorm:"column:address;type:string;size:2048;comment:주소" json:"address"`
	AddressDetail   string `gorm:"column:address_detail;type:string;size:2048;comment:주소상세" json:"address_detail"`

	DeliveryRequest DeliveryRequest `gorm:"foreignKey:DeliveryAddressID;references:ID" json:"delivery_request"`
}

func (e *DeliveryAddress) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e DeliveryAddress) IsNil() bool {
	return e.ID == "" && e.UserID == ""
}

func (e DeliveryAddress) CreateID() string {
	return uuid.New().String()
}

func (e *DeliveryAddress) CreateSetID() {
	e.ID = e.CreateID()
}

func (e *DeliveryAddress) ReferTo(userID string) {
	e.UserID = userID
}

type DeliveryAddresses []DeliveryAddress

func (e *DeliveryAddresses) CreateSetID() {
	for i := range *e {
		(*e)[i].ID = (*e)[i].CreateID()
	}
}

func (e *DeliveryAddresses) ReferTo(userID string) {
	for i := range *e {
		(*e)[i].ReferTo(userID)
	}
}

// DeliveryRequestType
// <-DeliveryRequest
type DeliveryRequestType struct {
	ID        string     `gorm:"primaryKey;type:string;size:64;comment:id" json:"id"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"`
	UpdatedAt *time.Time `gorm:"<-" json:"updated_at"`

	Name string `gorm:"column:name;type:string;size:512" json:"name"`
}

func (e *DeliveryRequestType) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e DeliveryRequestType) IsNil() bool {
	return e.ID == ""
}

// DeliveryRequest
// ->DeliveryAddress
// ->DeliveryRequestType
type DeliveryRequest struct {
	ID        string     `gorm:"primaryKey;type:string;size:64;comment:id" json:"id"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"`
	UpdatedAt *time.Time `gorm:"<-" json:"updated_at"`

	DeliveryAddressID string `gorm:"column:delivery_address_id;type:string;size:64;comment:delivery_address_id;not null" json:"delivery_address_id"`

	DeliveryRequestTypeID string              `gorm:"column:delivery_request_type_id;type:string;size:64;not null" json:"delivery_request_type_id"`
	DeliveryRequestType   DeliveryRequestType `gorm:"foreignKey:DeliveryRequestTypeID;references:ID" json:"delivery_request_type"`
	RequestMessage        string              `gorm:"column:request_message;type:string;size:2048;comment:요청메시지" json:"request_message"`
}

func (e *DeliveryRequest) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e DeliveryRequest) IsNil() bool {
	return e.ID == "" && e.DeliveryAddressID == ""
}

func (e DeliveryRequest) CreateID() string {
	return uuid.New().String()
}

func (e *DeliveryRequest) CreateSetID() {
	e.ID = e.CreateID()
}

func (e *DeliveryRequest) ReferTo(deliveryAddressID string) {
	e.DeliveryAddressID = deliveryAddressID
}
