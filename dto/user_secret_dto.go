package dto

import (
	"encoding/json"
	"time"
)

type UserPassword struct {
	ID        string     `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	UserID    string     `json:"user_id,omitempty"`
	Value     string     `json:"value,omitempty"`
}

func (e *UserPassword) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
