package customer_model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	CustomerID  uuid.UUID `json:"customer_id" db:"customer_id"`     // Primary Key
	NIK         string    `json:"nik" db:"nik"`                     // Nomor KTP
	FullName    string    `json:"full_name" db:"full_name"`         // Nama lengkap
	LegalName   string    `json:"legal_name" db:"legal_name"`       // Nama sesuai KTP
	BirthPlace  string    `json:"birth_place" db:"birth_place"`     // Tempat lahir
	BirthDate   time.Time `json:"birth_date" db:"birth_date"`       // Tanggal lahir
	Salary      float64   `json:"salary" db:"salary"`               // Gaji customer
	IDCardPhoto string    `json:"id_card_photo" db:"id_card_photo"` // Foto KTP
	SelfiePhoto string    `json:"selfie_photo" db:"selfie_photo"`   // Foto selfie customer
}

func (c Customer) MarshalJSON() ([]byte, error) {
	type Alias Customer
	return json.Marshal(&struct {
		CustomerID string `json:"customer_id"`
		BirthDate  string `json:"birth_date"`
		*Alias
	}{
		CustomerID: c.CustomerID.String(),
		BirthDate:  c.BirthDate.Format(time.RFC3339),
		Alias:      (*Alias)(&c),
	})
}

func (c *Customer) UnmarshalJSON(data []byte) error {
	type Alias Customer
	aux := &struct {
		CustomerID string `json:"customer_id"`
		BirthDate  string `json:"birth_date"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Convert string to uuid.UUID
	customerID, err := uuid.Parse(aux.CustomerID)
	if err != nil {
		return err
	}
	c.CustomerID = customerID

	// Convert string to time.Time
	birthDate, err := time.Parse(time.RFC3339, aux.BirthDate)
	if err != nil {
		return err
	}
	c.BirthDate = birthDate

	return nil
}
