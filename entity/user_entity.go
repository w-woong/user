package entity

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
)

var NilUser = User{}

// User entity.
type User struct {
	ID          string     `gorm:"primaryKey;type:string;size:64;comment:id"`
	LoginID     string     `gorm:"unique;not null;type:string;size:256;index:idx_users_1;comment:login id"`
	FirstName   string     `gorm:"not null;type:string;size:256;comment:first name"`
	LastName    string     `gorm:"not null;type:string;size:256;comment:last name"`
	BirthDate   time.Time  `gorm:"comment:yyyymmdd"`
	Gender      string     `gorm:"type:string;size:1;comment:M or F"`
	Nationality string     `gorm:"type:string;size:3;comment:Nationality(ISO 3166-1)"`
	CreatedAt   *time.Time `gorm:"<-:create"`
	UpdatedAt   *time.Time `gorm:"<-:update"`
	DeletedAt   *time.Time `gorm:"<-:update"`

	UserEmails UserEmails `gorm:"foreignKey:UserID;references:ID"`
}

// IsNill returns true if underlying ID is empty.
func (e User) IsNil() bool {
	return e.ID == ""
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
	// if _, err := time.Parse("20060102", e.BirthDate); err != nil {
	// 	return err
	// }

	return nil
}
