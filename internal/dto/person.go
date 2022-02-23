package dto

type Person struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
}
