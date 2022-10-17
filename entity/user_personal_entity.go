package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type UserPersonal struct {
	ID        string     `gorm:"primaryKey;type:string;size:64;comment:id"`
	CreatedAt *time.Time `gorm:"<-:create"`
	UpdatedAt *time.Time `gorm:"<-:update"`
	DeletedAt *time.Time `gorm:"<-:update"`

	UserID string `gorm:"type:string;size:64;comment:user id"`

	FirstName   string    `gorm:"not null;type:string;size:256;comment:first name"`
	LastName    string    `gorm:"not null;type:string;size:256;comment:last name"`
	BirthYear   int       `gorm:"column:birth_year"`
	BirthMonth  int       `gorm:"column:birth_month"`
	BirthDay    int       `gorm:"column:birth_day"`
	BirthDate   time.Time `gorm:"comment:yyyymmdd"`
	Gender      string    `gorm:"type:string;size:1;comment:M or F or empty/null"`
	Nationality string    `gorm:"type:string;size:3;comment:Nationality(ISO 3166-1)"`
}

func (e *UserPersonal) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// IsNill returns true if underlying ID is empty.
func (e UserPersonal) IsNil() bool {
	return e.ID == ""
}

func (e *UserPersonal) PrepareToRegister(userID string) error {

	e.GenerateAndSetID()
	e.RefersUserIDTo(userID)

	return nil
}

func (e *UserPersonal) GenerateAndSetID() {
	e.ID = e.generateID()
}

func (e UserPersonal) generateID() string {
	return uuid.New().String()
}

func (e *UserPersonal) RefersUserIDTo(userID string) {
	e.UserID = userID
}

func (e *UserPersonal) CombineBirthdate() {
	e.BirthDate = time.Date(e.BirthYear, time.Month(e.BirthMonth), e.BirthDay, 0, 0, 0, 0, time.Local)
}
