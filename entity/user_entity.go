package entity

import (
	"encoding/json"
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
)

var NilUser = User{}

type LoginType string

const (
	IDLoginType    LoginType = "id"
	EmailLoginType LoginType = "email"
)

// User entity.
type User struct {
	ID        string     `gorm:"primaryKey;type:string;size:64;comment:id"`
	CreatedAt *time.Time `gorm:"<-:create"`
	UpdatedAt *time.Time `gorm:"<-:update"`
	DeletedAt *time.Time `gorm:"<-:update"`

	LoginID   string    `gorm:"unique;not null;type:string;size:4096;comment:login id"`
	LoginType LoginType `gorm:"not null;type:string;size:32;comment:login type"`

	Password Password `gorm:"foreignKey:UserID;references:ID"`
	Personal Personal `gorm:"foreignKey:UserID;references:ID"`
	Emails   Emails   `gorm:"foreignKey:UserID;references:ID"`
}

func (e *User) String() string {
	b, _ := json.Marshal(e)
	return string(b)
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
	e.GenerateAndSetID()
	err := e.Validate()
	if err != nil {
		return err
	}

	if err = e.Password.PrepareToRegister(e.ID); err != nil {
		return err
	}

	if err = e.Personal.PrepareToRegister(e.ID); err != nil {
		return err
	}

	if err = e.Emails.PrepareToRegister(e.ID); err != nil {
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

func (e User) Validate() error {
	// login_id
	if err := e.validateLoginID(); err != nil {
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
