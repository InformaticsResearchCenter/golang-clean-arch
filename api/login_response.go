package api

import "iteung-api/model"

type LoginResponse struct {
	UserId       string `json:"user_id"`
	UserName     string `json:"user_name"`
	PhoneNumber  string `json:"phone_number"`
	EmailAddress string `json:"email_address"`
}

func ToLoginResponse(userInfo model.UserInfo) LoginResponse {
	return LoginResponse{
		UserId:       userInfo.UserId,
		UserName:     userInfo.UserName,
		PhoneNumber:  userInfo.PhoneNumber,
		EmailAddress: userInfo.EmailAddress,
	}
}
