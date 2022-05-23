package dto

//go:generate easyjson -all -disallow_unknown_fields response_models.go

//easyjson:json
type AuthResponse struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
}

func ToUserResponseFromUsecase(usr *AuthUsecase) *AuthResponse {
	return &AuthResponse{
		Nickname: usr.Username,
		ID:       usr.ID,
	}
}

//easyjson:json
type IDResponse struct {
	ID int64 `json:"id"`
}
