package entity

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          string
	LoginID     string
	FirstName   string
	LastName    string
	BirthDate   string
	Gender      string
	Nationality string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time

	UserEmails []UserEmail
}

func (e *User) PrepareToRegister() error {

	err := e.validate()
	if err != nil {
		return err
	}

	e.ID = e.generateID()

	return nil
}

func (e User) generateID() string {
	return uuid.New().String()
}

func (e User) validate() error {
	// login_id
	if err := e.validateLoginID(); err != nil {
		return err
	}
	if err := e.validateBirthDate(); err != nil {
		return err
	}
	return nil
}

func (e User) validateLoginID() error {
	if ok, _ := regexp.MatchString("[a-zA-Z0-9]{6,}", e.LoginID); !ok {
		return errors.New("login_id is not valid")
	}
	return nil
}

func (e User) validateBirthDate() error {
	// if ok, _ := regexp.MatchString(`\d{4}\d{2}\d{2}`, m.BirthDate); !ok {
	// 	return errors.New("birth_date is not valid")
	// }
	if _, err := time.Parse("20060102", e.BirthDate); err != nil {
		return err
	}

	return nil
}
