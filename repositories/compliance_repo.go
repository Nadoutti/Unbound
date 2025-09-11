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
