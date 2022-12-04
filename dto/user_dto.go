package dto

import (
	"encoding/json"
	"time"
)

var NilUser = User{}

type User struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	LoginID     string `json:"login_id"`
	LoginType   string `json:"login_type"`
	LoginSource string `json:"login_source"`

	Password Password `json:"password"`
	Personal Personal `json:"personal"`
	Emails   Emails   `json:"emails"`
}

func (d *User) String() string {
	b, _ := json.Marshal(d)
	return string(b)
}

func (d User) IsNil() bool {
	return d.ID == "" && d.LoginID == ""
}

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
