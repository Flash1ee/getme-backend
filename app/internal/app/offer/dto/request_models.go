package dto

//go:generate easyjson -all -disallow_unknown_fields request_models.go

//easyjson:json
type RequestCreateOffer struct {
	//SkillName string `json:"skill_name" validate:"required,min=2,alphanumunicode"`
	SkillName string `json:"skill_name"`
	MentorID  int64  `json:"mentor_id" validate:"required,min=0"`
}
