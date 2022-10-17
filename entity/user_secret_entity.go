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

type UserPassword struct {
	ID        string     `gorm:"primaryKey;type:string;size:64;comment:id"`
	CreatedAt *time.Time `gorm:"<-:create"`
	UpdatedAt *time.Time `gorm:"<-:update"`
	DeletedAt *time.Time `gorm:"<-:update"`
	UserID    string     `gorm:"unique;type:string;size:64;comment:user id"`
	Value     string     `gorm:"type:string;size:2048;comment:secret value"`
}

func (e *UserPassword) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// IsNill returns true if underlying ID is empty.
func (e UserPassword) IsNil() bool {
	return e.ID == ""
}

func (e *UserPassword) PrepareToRegister(userID string) error {

	e.GenerateAndSetID()
	e.RefersUserIDTo(userID)

	return nil
}

func (e *UserPassword) GenerateAndSetID() {
	e.ID = e.generateID()
}

func (e UserPassword) generateID() string {
	return uuid.New().String()
}

func (e *UserPassword) RefersUserIDTo(userID string) {
	e.UserID = userID
}
