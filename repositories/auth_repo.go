package repositories

import (
	"unbound/db"
	"unbound/models"
)

func FindByEmail(email string) (models.User, error) {
	supabase := db.GetSupabase()

	var user models.User

	_, err := supabase.From("users").
		Select("*", "", false).
		Eq("email", email).
		Single().
		ExecuteTo(&user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
