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
}

type Users []User

var NilUser = User{}
