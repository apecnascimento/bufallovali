package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Cidade struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Nome      string    `json:"nome" db:"nome"`
	Estado    string    `json:"estado" db:"estado"`
	Cep       string    `json:"cep" db:"cep"`
}

// String is not required by pop and may be deleted
func (c Cidade) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Cidades is not required by pop and may be deleted
type Cidades []Cidade

// String is not required by pop and may be deleted
func (c Cidades) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Cidade) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.Nome, Name: "Nome"},
		&validators.StringIsPresent{Field: c.Estado, Name: "Estado"},
		&validators.StringIsPresent{Field: c.Cep, Name: "Cep"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Cidade) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Cidade) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
