package repositories

import (
	"unbound/db"
	"unbound/models"
)

func FindBusByID(id string) (models.KYBResponse, error) {
	supabase := db.GetSupabase()

	var business models.KYBResponse

	_, err := supabase.From("cust_business").
		Select("*", "", false).
		Eq("user_id", id).
		Single().
		ExecuteTo(&business)

	if err != nil {
		return models.KYBResponse{}, err
	}

	return business, nil
}

func CreateCustomerBusiness(kybData *models.KYBSubmission, userID string) (models.KYBResponse, error) {
	supabase := db.GetSupabase()

	newBuss := map[string]interface{}{
		"streetline_one":  kybData.StreetLineOne,
		"streetline_two":  kybData.StreetLineOne,
		"address_city":     kybData.AddressCity,
		"address_country":  kybData.AddressCountry,
		"postal_code":      kybData.PostalCode,
		"document_type":    kybData.DocumentType,
		"document_value":   kybData.DocumentNumber,
		"document_country": kybData.DocumentCountry,
		"user_id": userID,
	}

	var createdBusiness models.KYBResponse

	_, err := supabase.From("business").
		Insert(newBuss, false, "", "", "").
		Single().
		ExecuteTo(&createdBusiness)

	if err != nil {
		return models.KYBResponse{}, err
	}

	return createdBusiness, nil
}

func CreateCustomerIndividual(kycData *models.KYCSubmission, userID string) (models.KYCResponse, error) {
	supabase := db.GetSupabase()

	newBuss := map[string]interface{}{
		"streetline_one":  kycData.StreetLineOne,
		"streetline_two":  kycData.StreetLineOne,
		"address_city":     kycData.AddressCity,
		"address_country":  kycData.AddressCountry,
		"postal_code":      kycData.PostalCode,
		"document_type":    kycData.DocumentType,
		"document_value":   kycData.DocumentNumber,
		"document_country": kycData.DocumentCountry,
		"user_id": userID,
	}

	var createdIndividual models.KYCResponse

	_, err := supabase.From("business").
		Insert(newBuss, false, "", "", "").
		Single().
		ExecuteTo(&createdIndividual)

	if err != nil {
		return models.KYCResponse{}, err
	}

	return createdIndividual, nil
}
