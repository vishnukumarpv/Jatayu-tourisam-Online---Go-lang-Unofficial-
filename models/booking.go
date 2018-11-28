package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Booking struct {
	ID        uuid.UUID `json:"id" db:"id"`
	BookedTo  time.Time `json:"booked_to" db:"booked_to"`
	Amount    float32   `json:"amount" db:"amount"`
	Adults    int       `json:"adults" db:"adults"`
	Children  int       `json:"children" db:"children"`
	Infants   int       `json:"infants" db:"infants"`
	Package   string    `json:"package" db:"package"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (b Booking) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Bookings is not required by pop and may be deleted
type Bookings []Booking

// String is not required by pop and may be deleted
func (b Bookings) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (b *Booking) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.IntIsPresent{Field: b.Adults, Name: "Adults"},
		// &validators.StringIsPresent{Field: b.BookedTo, Name: "Booked to"},
	), err
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (b *Booking) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (b *Booking) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
