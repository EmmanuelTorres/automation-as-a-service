package dto

type Designer struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CountryID string `json:"country_id"`
}
