package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type SecretType uint

const (
	LoginPassword SecretType = iota
	CI
	Token
)

var (
	NilCredentialPassword = CredentialPassword{}
	NilCredentialToken    = CredentialToken{}
)

type CredentialPassword struct {
	ID        string     `gorm:"primaryKey;type:string;size:64;comment:id" json:"id"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"`
	UpdatedAt *time.Time `gorm:"<-" json:"updated_at"`

	UserID string `gorm:"unique;type:string;size:64;comment:user id" json:"user_id"`
	Value  string `gorm:"type:string;size:2048;comment:secret value" json:"value"`
}

func (e *CredentialPassword) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// IsNill returns true if underlying ID is empty.
func (e CredentialPassword) IsNil() bool {
	return e.ID == ""
}

func (e *CredentialPassword) PrepareToRegister(userID string) error {

	e.CreateSetID()
	e.ReferTo(userID)

	return nil
}

func (e *CredentialPassword) CreateSetID() {
	e.ID = e.CreateID()
}

func (e CredentialPassword) CreateID() string {
	return uuid.New().String()
}

func (e *CredentialPassword) ReferTo(userID string) {
	e.UserID = userID
}

type CredentialToken struct {
	ID        string     `gorm:"primaryKey;type:string;size:64;comment:id" json:"id"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"`
	UpdatedAt *time.Time `gorm:"<-" json:"updated_at"`

	UserID string `gorm:"unique;type:string;size:64;comment:user id" json:"user_id"`
	Value  string `gorm:"type:string;size:2048;comment:secret value" json:"value"`
}

func (e *CredentialToken) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// IsNill returns true if underlying ID is empty.
func (e CredentialToken) IsNil() bool {
	return e.ID == ""
}

func (e *CredentialToken) PrepareToRegister(userID string) error {

	e.CreateSetID()
	e.ReferTo(userID)

	return nil
}

func (e *CredentialToken) CreateSetID() {
	e.ID = e.CreateID()
}

func (e CredentialToken) CreateID() string {
	return uuid.New().String()
}

func (e *CredentialToken) ReferTo(userID string) {
	e.UserID = userID
}
