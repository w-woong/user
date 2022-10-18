package dto

import (
	"encoding/json"
	"time"
)

type User struct {
	ID        string     `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	LoginID   string `json:"login_id,omitempty"`
	LoginType string `json:"login_type,omitempty"`

	Password Pasword  `json:"password"`
	Personal Personal `json:"personal"`
	Emails   []Email  `json:"emails,omitempty"`
}

func (d *User) String() string {
	b, _ := json.Marshal(d)
	return string(b)
}

var NilUser = User{}

type Users []User

// type UserMobile struct {
// 	ID       string    `json:"id"`
// 	UserID   string    `json:"user_id"`
// 	Email    string    `json:"email"`
// 	Created  time.Time `json:"created"`
// 	Modified time.Time `json:"modified"`
// }

// type UserAddress struct {
// 	ID         string    `json:"id"`
// 	UserID     string    `json:"user_id"`
// 	CountryID  string    `json:"country_id"`
// 	PostalCode string    `json:"postal_code"`
// 	State      string    `json:"state"`
// 	City       string    `json:"city"`
// 	Addr1      string    `json:"addr1"`
// 	Addr2      string    `json:"addr2"`
// 	Created    time.Time `json:"created"`
// 	Modified   time.Time `json:"modified"`
// }
