package dto

import "time"

type Personal struct {
	ID        string     `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	UserID string `json:"user_id"`

	FirstName   string    `json:"first_name,omitempty"`
	LastName    string    `json:"last_name,omitempty"`
	BirthYear   int       `json:"birth_year,omitempty"`
	BirthMonth  int       `json:"birth_month,omitempty"`
	BirthDay    int       `json:"birth_day,omitempty"`
	BirthDate   time.Time `json:"birth_date,omitempty"`
	Gender      string    `json:"gender,omitempty"`
	Nationality string    `json:"nationality,omitempty"`
}
