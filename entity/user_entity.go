package entity

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

var NilUser = User{}

type LoginType string

const (
	LoginTypeID    LoginType = "id"
	LoginTypeEmail LoginType = "email"
	LoginTypeToken LoginType = "token"
)

type LoginSource string

var (
	LoginSourceWoong  LoginSource = "woong"
	LoginSourceGoogle LoginSource = "google"
	LoginSourceKakao  LoginSource = "kakao"
)

func (e LoginSource) LoginID(id string) (string, error) {
	switch e {
	case LoginSourceWoong:
		return id, nil
	case LoginSourceGoogle:
		if strings.HasPrefix(id, string(e)+"_") {
			return id, nil
		}
		return string(e) + "_" + id, nil
	case LoginSourceKakao:
		if strings.HasPrefix(id, string(e)+"_") {
			return id, nil
		}
		return string(e) + "_" + id, nil
	default:
		return "", errors.New("invalid login source")
	}
}

// User entity.
type User struct {
	ID        string     `gorm:"primaryKey;type:string;size:64;comment:id" json:"id"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"`
	UpdatedAt *time.Time `gorm:"<-" json:"updated_at"`

	LoginID     string      `gorm:"unique;not null;type:string;size:2048;comment:login id" json:"login_id"`
	LoginType   LoginType   `gorm:"not null;type:string;size:32;comment:login type" json:"login_type"`
	LoginSource LoginSource `gorm:"not null;type:string;size:32;comment:login source" json:"login_source"`

	CredentialPassword *CredentialPassword `gorm:"foreignKey:UserID;references:ID" json:"password,omitempty"`
	CredentialToken    *CredentialToken    `gorm:"foreignKey:UserID;references:ID" json:"token,omitempty"`
	Personal           *Personal           `gorm:"foreignKey:UserID;references:ID" json:"personal,omitempty"`
	Emails             Emails              `gorm:"foreignKey:UserID;references:ID" json:"emails,omitempty"`
	DeliveryAddress    *DeliveryAddress    `gorm:"foreignKey:UserID;references:ID" json:"delivery_address,omitempty"`
	PaymentMethod      *PaymentMethod      `gorm:"foreignKey:UserID;references:ID" json:"payment_method,omitempty"`
}

func (e *User) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// IsNill returns true if underlying ID is empty.
func (e User) IsNil() bool {
	return e.ID == "" && e.LoginID == ""
}

// PrepareToRegister prepares user entity, e, before registering a new user.
// It validates underlying fields.
// It generates and set a new ID.
// It sets references to child entities.
func (e *User) PrepareToRegister() error {
	e.CreateSetID()
	err := e.Validate()
	if err != nil {
		return err
	}

	if e.CredentialPassword != nil && e.CredentialPassword.Value != "" {
		if err = e.CredentialPassword.PrepareToRegister(e.ID); err != nil {
			return err
		}
	} else {
		e.CredentialPassword = nil
	}
	if e.CredentialToken != nil && e.CredentialToken.Value != "" {
		if err = e.CredentialToken.PrepareToRegister(e.ID); err != nil {
			return err
		}
	} else {
		e.CredentialToken = nil
	}

	if err = e.Personal.PrepareToRegister(e.ID); err != nil {
		return err
	}

	if err = e.Emails.PrepareToRegister(e.ID); err != nil {
		return err
	}

	e.DeliveryAddress = nil
	e.PaymentMethod = nil

	return nil
}

func (e *User) CreateSetID() {
	e.ID = e.CreateID()
}

func (e User) CreateID() string {
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
