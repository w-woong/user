package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserEmail struct {
	ID        string
	UserID    string
	Email     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (e *UserEmail) GenerateAndSetID() {
	e.ID = e.generateID()
}

func (e UserEmail) generateID() string {
	return uuid.New().String()
}
