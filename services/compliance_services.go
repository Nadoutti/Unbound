package services

import (
	"errors"
	"unbound/models"
	"unbound/repositories"
)

func SubmitKYB(kybData *models.KYBSubmission, userID string) (map[string]interface{}, error) {
	// verificando se o business ja existe

	_, err := repositories.FindBusByID(userID)
	if err != nil {
		return map[string]interface{}{"error": "Business not found"}, errors.New("business not found")
	}

	// criando o business

	newKYB, err := repositories.CreateCustomerBusiness(kybData)

	if err != nil {
		return map[string]interface{}{"error": "Nao foi possivel criar a business"}, err
	}

	return map[string]interface{}{"created_cust_business": &newKYB}, nil

}
