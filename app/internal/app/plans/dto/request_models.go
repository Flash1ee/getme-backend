package dto

//go:generate easyjson -all -disallow_unknown_fields request_models.go

const (
	MentorAlias = "mentor"
	MenteeAlias = "mentee"
)

//easyjson:json
type RequestPlan struct {
	Role string `query:"role" validate:"oneof=mentor mentee"`
}
