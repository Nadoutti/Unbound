package repositories

import (
	"unbound/models"
	"unbound/db"
)

func FindUserbyID(id string) (models.User, error) {
	supabase := db.GetSupabase()

	var user models.User

	_, err := supabase.From("users").
		Select("*", "", false).
		Eq("id", id).
		Single().
		ExecuteTo(&user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
