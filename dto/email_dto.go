package dto

import "time"

type Email struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Priority uint8  `json:"priority"`
}

var NilEmail = Email{}

func (m Email) IsNil() bool {
	return m.ID == ""
}

type Emails []Email
