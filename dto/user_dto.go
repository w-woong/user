package dto

import (
	"time"
)

type User struct {
	ID          string     `gorm:"primaryKey;type:varchar(64);comment:id" json:"id,omitempty"`
	LoginID     string     `gorm:"unique;not null;type:varchar(256);index:idx_users_1;comment:login id" json:"login_id,omitempty"`
	FirstName   string     `gorm:"not null;type:varchar(256);comment:first name" json:"first_name,omitempty"`
	LastName    string     `gorm:"not null;type:varchar(256);comment:last name" json:"last_name,omitempty"`
	BirthDate   string     `gorm:"type:varchar(8);comment:yyyymmdd" json:"birth_date,omitempty"`
	Gender      string     `gorm:"type:varchar(1);comment:M or F" json:"gender,omitempty"`
	Nationality string     `gorm:"type:varchar(3);comment:Nationality(ISO 3166-1)" json:"nationality,omitempty"`
	CreatedAt   *time.Time `gorm:"<-:create" json:"created_at,omitempty"`
	UpdatedAt   *time.Time `gorm:"<-:update" json:"updated_at,omitempty"`
	DeletedAt   *time.Time `gorm:"<-:update" json:"deleted_at,omitempty"`

	UserEmails  []UserEmail `gorm:"foreignKey:UserID;references:ID"`
	UserSecrets []UserSecret
}

type Users []User

var NilUser = User{}

func (d User) IsNil() bool {
	return d.LoginID == ""
}

type UserEmail struct {
	ID        string     `gorm:"privaryKey;type:varchar(64)" json:"id,omitempty"`
	UserID    string     `gorm:"type:varchar(64)" json:"user_id,omitempty"`
	Email     string     `gorm:"unique;not null;type:varchar(256);index:idx_useremails_1" json:"email"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at,omitempty"`
	DeletedAt *time.Time `gorm:"<-:update" json:"deleted_at,omitempty"`
}

var NilUserEmail = UserEmail{}

func (m UserEmail) IsNil() bool {
	return m.ID == ""
}

type UserMobile struct {
	ID       string    `json:"id"`
	UserID   string    `json:"user_id"`
	Email    string    `json:"email"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

type UserAddress struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	CountryID  string    `json:"country_id"`
	PostalCode string    `json:"postal_code"`
	State      string    `json:"state"`
	City       string    `json:"city"`
	Addr1      string    `json:"addr1"`
	Addr2      string    `json:"addr2"`
	Created    time.Time `json:"created"`
	Modified   time.Time `json:"modified"`
}

type UserSecret struct {
	ID       string    `json:"id"`
	UserID   string    `json:"user_id"`
	Type     string    `json:"type"`
	Value    string    `json:"value"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}
