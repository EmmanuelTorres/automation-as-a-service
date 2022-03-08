package dto

type Brand struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CountryID int64  `json:"country_id"`
}
