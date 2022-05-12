package dto

//go:generate easyjson -all -disallow_unknown_fields request_models.go

//easyjson:json
type UserAuthCheckRequest struct {
	Token     string `query:"token" json:"token"`
	ID        int64  `query:"id" json:"id"`
	AuthDate  int64  `query:"auth_date" json:"auth_date"`
	FirstName string `query:"first_name" json:"first_name"`
	LastName  string `query:"last_name" json:"last_name"`
	Username  string `query:"username" json:"username"`
	Avatar    string `query:"photo_url" json:"photo_url"`
	Hash      string `query:"hash" json:"hash"`
}

func (req *UserAuthCheckRequest) ToUserAuthUsecase() *UserAuthUsecase {
	return &UserAuthUsecase{
		ID:        req.ID,
		AuthDate:  req.AuthDate,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Avatar:    req.Avatar,
		Hash:      req.Hash,
	}
}

//easyjson:json
type UserAuthRequest struct {
	Token string `query:"token" json:"token"`
}
