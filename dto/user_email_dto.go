package dto

import "time"

type UserEmail struct {
	ID        string     `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	UserID    string     `json:"user_id,omitempty"`
	Email     string     `json:"email"`
	Priority  uint8      `json:"priority"`
}

var NilUserEmail = UserEmail{}

func (m UserEmail) IsNil() bool {
	return m.ID == ""
}
