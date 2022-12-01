package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Personal struct {
	ID        string     `gorm:"primaryKey;type:string;size:64;comment:id" json:"id"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"`
	UpdatedAt *time.Time `gorm:"<-" json:"updated_at"`

	UserID string `gorm:"type:string;size:64;comment:user id" json:"user_id"`

	FirstName   string     `gorm:"not null;type:string;size:256;comment:first name" json:"first_name"`
	MiddleName  string     `gorm:"not null;type:string;size:256;comment:middle name" json:"middle_name"`
	LastName    string     `gorm:"not null;type:string;size:256;comment:last name" json:"last_name"`
	BirthYear   int        `gorm:"column:birth_year" json:"birth_year"`
	BirthMonth  int        `gorm:"column:birth_month" json:"birth_month"`
	BirthDay    int        `gorm:"column:birth_day" json:"birth_day"`
	BirthDate   *time.Time `gorm:"comment:yyyymmdd" json:"birth_date"`
	Gender      string     `gorm:"type:string;size:1;comment:M or F or empty/null" json:"gender"`
	Nationality string     `gorm:"type:string;size:3;comment:Nationality(ISO 3166-1)" json:"nationality"`
}

func (e *Personal) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// IsNill returns true if underlying ID is empty.
func (e Personal) IsNil() bool {
	return e.ID == ""
}

func (e *Personal) PrepareToRegister(userID string) error {

	e.GenerateAndSetID()
	e.RefersUserIDTo(userID)

	return nil
}

func (e *Personal) GenerateAndSetID() {
	e.ID = e.generateID()
}

func (e Personal) generateID() string {
	return uuid.New().String()
}

func (e *Personal) RefersUserIDTo(userID string) {
	e.UserID = userID
}

func (e *Personal) CombineBirthdate() {
	if e.BirthYear == 0 || e.BirthMonth == 0 || e.BirthDay == 0 {
		e.BirthDate = nil
		return
	}
	d := time.Date(e.BirthYear, time.Month(e.BirthMonth), e.BirthDay, 0, 0, 0, 0, time.Local)
	e.BirthDate = &d
}
