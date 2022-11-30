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
)

var NilPassword = Password{}

type Password struct {
	ID        string     `gorm:"primaryKey;type:string;size:64;comment:id" json:"id"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"`
	UpdatedAt *time.Time `gorm:"<-" json:"updated_at"`

	UserID string `gorm:"unique;type:string;size:64;comment:user id" json:"user_id"`
	Value  string `gorm:"type:string;size:2048;comment:secret value" json:"value"`
}

func (e *Password) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// IsNill returns true if underlying ID is empty.
func (e Password) IsNil() bool {
	return e.ID == ""
}

func (e *Password) PrepareToRegister(userID string) error {

	e.GenerateAndSetID()
	e.RefersUserIDTo(userID)

	return nil
}

func (e *Password) GenerateAndSetID() {
	e.ID = e.generateID()
}

func (e Password) generateID() string {
	return uuid.New().String()
}

func (e *Password) RefersUserIDTo(userID string) {
	e.UserID = userID
}
