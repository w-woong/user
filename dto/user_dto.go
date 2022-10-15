package dto

import (
	"time"
)

type User struct {
	ID          string     `json:"id,omitempty"`
	LoginID     string     `json:"login_id,omitempty"`
	FirstName   string     `json:"first_name,omitempty"`
	LastName    string     `json:"last_name,omitempty"`
	BirthYear   int        `json:"birth_year,omitempty"`
	BirthMonth  int        `json:"birth_month,omitempty"`
	BirthDay    int        `json:"birth_day,omitempty"`
	BirthDate   time.Time  `json:"birth_date,omitempty"`
	Gender      string     `json:"gender,omitempty"`
	Nationality string     `json:"nationality,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`

	UserEmails  []UserEmail  `json:"user_emails,omitempty"`
	UserSecrets []UserSecret `json:"user_secrets,omitempty"`
}

type Users []User

var NilUser = User{}

type UserEmail struct {
	ID        string     `json:"id,omitempty"`
	UserID    string     `json:"user_id,omitempty"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
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
