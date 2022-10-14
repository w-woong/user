package entity

import (
	"errors"
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

func (e *UserEmail) RefersUserIDTo(userID string) {
	e.UserID = userID
}

type UserEmails []UserEmail

func (e *UserEmails) PrepareToRegister(userID string) error {
	if userID == "" {
		return errors.New("userID was not provided")
	}

	if e == nil || len(*e) == 0 {
		return errors.New("a user must have at least 1 email address")
	}

	for i := range *e {
		(*e)[i].GenerateAndSetID()
		(*e)[i].RefersUserIDTo(userID)
	}
	return nil
}
