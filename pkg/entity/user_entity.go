package entity

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
)

// User entity.
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

	UserEmails UserEmails
}

// PrepareToRegister prepares user entity, e, before registering a new user.
// It validates underlying fields.
// It generates and set a new ID.
// It sets references to child entities.
func (e *User) PrepareToRegister() error {

	err := e.validate()
	if err != nil {
		return err
	}

	e.GenerateAndSetID()

	if err = e.UserEmails.PrepareToRegister(e.ID); err != nil {
		return err
	}

	return nil
}

func (e *User) GenerateAndSetID() {
	e.ID = e.generateID()
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

func (e *User) prepareUserEmailsToRegister() {
	for i := range e.UserEmails {
		e.UserEmails[i].GenerateAndSetID()
		e.UserEmails[i].RefersUserIDTo(e.ID)
	}
}
