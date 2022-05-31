package dto

//go:generate easyjson -all -disallow_unknown_fields request_models.go

//easyjson:json
type RequestUserUpdate struct {
	FirstName string   `json:"first_name,omitempty" validate:"alpha,min=3"`
	LastName  string   `json:"last_name,omitempty" validate:"alpha,min=3"`
	About     string   `json:"about,omitempty" validate:"min=10,max=100"`
	TgTag     string   `json:"tg_tag,omitempty" validate:"min=0,max=100"`
	Skills    []string `json:"skills,omitempty" validate:"max=100"`
}

//easyjson:json
type RequestUpdateStatus struct {
	IsMentor bool `query:"mentor" validate:"required"`
}

func (req *RequestUserUpdate) ToUserUsecase() *UserUsecase {
	return &UserUsecase{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		About:     req.About,
	}
}

func (req *RequestUserUpdate) ToUserWithSkillsUsecase() *UserWithSkillsUsecase {
	return &UserWithSkillsUsecase{
		UserUsecase: UserUsecase{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			About:     req.About,
			TgTag:     req.TgTag,
		},
		Skills: req.Skills,
	}
}
func (req *RequestUpdateStatus) ToStatusUpdateUsecase() *UserStatusUsecase {
	return &UserStatusUsecase{
		IsMentor: req.IsMentor,
	}

}
