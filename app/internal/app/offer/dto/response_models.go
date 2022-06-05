package dto

import "getme-backend/internal/app/user/dto"

//go:generate easyjson -all -disallow_unknown_fields response_models.go

//easyjson:json
type ResponseOffer struct {
	ID int64 `json:"offer_id"`
}

//easyjson:json
type RespondOffersWithUser struct {
	ID        int64  `json:"offer_id"`
	SkillName string `json:"skill_name"`
	dto.UserResponse
}
