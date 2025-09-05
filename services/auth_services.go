package services

import (
	"errors"
	"unbound/auth_data_processing/middleware"
	"unbound/repositories"
)

func LoginUser(email, password string) (map[string]interface{}, error) {

	// encontrar usuario por email
	currentUser, err := repositories.FindByEmail(email)
	if err != nil {
		return map[string]interface{}{"error": "User not found"}, err

	}

	// ver se da match a senha com a senha hasheada

	isValidPass := middleware.CheckPasswordHash(password, currentUser.Password)

	if isValidPass != nil {
		return map[string]interface{}{"error": "Invalid password"}, errors.New("invalid password")
	}

	// gerar token JWT

	token, err := middleware.CreateJWT(currentUser.ID, currentUser.Email)

	if err != nil {
		return map[string]interface{}{"error": "Could not create token"}, err
	}

	return map[string]interface{}{"token": token, "user": currentUser}, nil
}
