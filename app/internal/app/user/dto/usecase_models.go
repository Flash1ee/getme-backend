package dto

type UserAuthUsecase struct {
	ID        int64  `query:"id" json:"id"`
	AuthDate  int64  `query:"auth_date" json:"auth_date"`
	FirstName string `query:"first_name" json:"first_name"`
	LastName  string `query:"last_name" json:"last_name"`
	Username  string `query:"username" json:"username"`
	Avatar    string `query:"photo_url" json:"photo_url"`
	Hash      string `query:"hash" json:"hash"`
}
