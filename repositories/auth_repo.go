package repositories

import (
	"time"
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

func CreateUser(email, password, nome, phone string) (models.User, error) {
	supabase := db.GetSupabase()

	newUser := map[string]interface{}{
		"email":         email,
		"password_hash": password,
		"nome":          nome,
		"phone":         phone,
		"user_type":     "individual",
		"created_at":    time.Now(),
	}

	var createdUser models.User

	_, err := supabase.From("users").
		Insert(newUser, false, "", "", "").
		Single().
		ExecuteTo(&createdUser)

	if err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}
