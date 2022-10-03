package model

import (
	"time"

	"github.com/google/uuid"
)

var NilUser = User{}

type User struct {
	ID          string
	LoginID     string
	FirstName   string
	LastName    string
	BirthDate   string
	Gender      string
	Nationality string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time

	UserEmails  []UserEmail
	UserSecrets []UserSecret
}

func (m User) IsNil() bool {
	return m.ID == ""
}

func (m *User) CreateAndSetID() {
	m.ID = uuid.New().String()
}

var NilUserEmail = UserEmail{}

type UserEmail struct {
	ID       string    `json:"id"`
	UserID   string    `json:"user_id"`
	Email    string    `json:"email"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

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
