package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type Show struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Number    int       `json:"number" db:"number"`
	Date      string    `json:"date" db:"date"`
	GuestID   uuid.UUID `json:"guest_id" db:"guest_id"`
	Topic     string    `json:"topic" db:"topic"`
}

// String is not required by pop and may be deleted
func (s Show) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Shows is not required by pop and may be deleted
type Shows []Show

// String is not required by pop and may be deleted
func (s Shows) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Show) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: s.Number, Name: "Number"},
		&validators.StringIsPresent{Field: s.Date, Name: "Date"},
		&validators.StringIsPresent{Field: s.Topic, Name: "Topic"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *Show) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *Show) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
