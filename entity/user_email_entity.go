package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type UserEmail struct {
	ID        string     `gorm:"primaryKey;type:string;size:64;comment:id"`
	UserID    string     `gorm:"type:string;size:64;comment:user id"`
	Email     string     `gorm:"type:string;size:256;comment:email"`
	CreatedAt *time.Time `gorm:"<-:create"`
	UpdatedAt *time.Time `gorm:"<-:update"`
	DeletedAt *time.Time `gorm:"<-:update"`
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
