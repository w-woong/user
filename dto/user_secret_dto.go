package dto

import (
	"encoding/json"
	"time"
)

type UserSecret struct {
	ID        string     `json:"id,omitempty"`
	UserID    string     `json:"user_id,omitempty"`
	Type      uint       `json:"type,omitempty"`
	Value     string     `json:"value,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (e *UserSecret) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type UserSecrets []UserSecret
