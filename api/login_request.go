package api

type LoginRequest struct {
	PhoneNumber string `validate:"required" json:"phone_number"`
	Password    string `validate:"required" json:"password"`
}
