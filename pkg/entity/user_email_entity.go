package entity

import "time"

type UserEmail struct {
	ID        string
	UserID    string
	Email     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
