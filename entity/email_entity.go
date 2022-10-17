package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Email struct {
	ID        string     `gorm:"primaryKey;type:string;size:64;comment:id"`
	CreatedAt *time.Time `gorm:"<-:create"`
	UpdatedAt *time.Time `gorm:"<-:update"`
	DeletedAt *time.Time `gorm:"<-:update"`
	UserID    string     `gorm:"type:string;size:64;comment:user id"`
	Email     string     `gorm:"type:string;size:256;comment:email"`
	Priority  uint8      `gorm:"type:int;comment:email priority starting from 0"`
}

func (e *Email) GenerateAndSetID() {
	e.ID = e.generateID()
}

func (e Email) generateID() string {
	return uuid.New().String()
}

func (e *Email) RefersUserIDTo(userID string) {
	e.UserID = userID
}

type Emails []Email

func (e *Emails) PrepareToRegister(userID string) error {
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
