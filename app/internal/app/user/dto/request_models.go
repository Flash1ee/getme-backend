package dto

//go:generate easyjson -all -disallow_unknown_fields request_models.go

//easyjson:json
type RequestUserUpdate struct {
	FirstName string `json:"first_name,omitempty" validate:"alpha,min=3"`
	LastName  string `json:"last_name,omitempty" validate:"alpha,min=3"`
	About     string `json:"about,omitempty" validate:"min=10,max=100"`
}

func (req *RequestUserUpdate) ToUserUsecase() *UserUsecase {
	return &UserUsecase{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		About:     req.About,
	}
}
