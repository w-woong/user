package entity

import (
	"encoding/json"
	"time"
)

type SecretType uint

const (
	LoginPassword SecretType = iota
	CI
)

type UserSecret struct {
	ID        string     `gorm:"primaryKey;type:string;size:64;comment:id"`
	UserID    string     `gorm:"uniqueIndex:idx_user_secrets_1;type:string;size:64;comment:user id"`
	Type      SecretType `gorm:"uniqueIndex:idx_user_secrets_1;type:uint;comment:secret type"`
	Value     string     `gorm:"type:string;size:2048;comment:secret value"`
	CreatedAt *time.Time `gorm:"<-:create"`
	UpdatedAt *time.Time `gorm:"<-:update"`
	DeletedAt *time.Time `gorm:"<-:update"`
}

func (e *UserSecret) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type UserSecrets []UserSecret
