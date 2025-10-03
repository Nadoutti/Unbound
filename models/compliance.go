package models

// KYBSubmission representa os dados necessrios para a submissao KYB

type KYBSubmission struct {
	StreetLineOne   string `json:"street_line_one"`
	StreetAdressTwo string `json:"street_line_two"`
	AddressCity     string `json:"address_city"`
	AddressCountry  string `json:"address_country"`
	PostalCode      string `json:"postal_code"`
	DocumentType    string `json:"document_type"`
	DocumentNumber  string `json:"document_value"`
	DocumentCountry string `json:"document_country"`
}

type KYBResponse struct {
	Status string `json:"status"`
}

type KYCSubmission struct {
	StreetLineOne   string `json:"street_line_one"`
	StreetAdressTwo string `json:"street_line_two"`
	AddressCity     string `json:"address_city"`
	AddressCountry  string `json:"address_country"`
	PostalCode      string `json:"postal_code"`
	DocumentType    string `json:"document_type"`
	DocumentNumber  string `json:"document_value"`
	DocumentCountry string `json:"document_country"`
}

type KYCResponse struct {
	Status string `json:"status"`
}
