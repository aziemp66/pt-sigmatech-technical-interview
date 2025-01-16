package customer_model

// CreateCustomerRequest represents the request to create a new customer
type CreateCustomerRequest struct {
	NIK         string  `json:"nik"`
	FullName    string  `json:"full_name"`
	LegalName   string  `json:"legal_name"`
	BirthPlace  string  `json:"birth_place"`
	BirthDate   string  `json:"birth_date"` // in RFC3339 format
	Salary      float64 `json:"salary"`
	IDCardPhoto string  `json:"id_card_photo"`
	SelfiePhoto string  `json:"selfie_photo"`
}

// UpdateCustomerRequest represents the request to update an existing customer
type UpdateCustomerRequest struct {
	FullName    string  `json:"full_name"`
	LegalName   string  `json:"legal_name"`
	BirthPlace  string  `json:"birth_place"`
	Salary      float64 `json:"salary"`
	IDCardPhoto string  `json:"id_card_photo"`
	SelfiePhoto string  `json:"selfie_photo"`
}
