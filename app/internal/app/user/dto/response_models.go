package dto

//go:generate easyjson -all -disallow_unknown_fields response_models.go

//easyjson:json
type UserResponse struct {
	ID           int64  `json:"id"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	About        string `json:"about,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
	IsSearchable bool   `json:"is_mentor"`
}

//easyjson:json
type UsersResponse struct {
	Users []UserResponse `json:"users"`
}

//easyjson:json
type UserWithSkillResponse struct {
	UserResponse
	Skills []string `json:"skills"`
}

//easyjson:json
type UsersWithSkillResponse struct {
	Users []UserWithSkillResponse `json:"users"`
}

func ToUsersWithSkillResponse(users []UserWithSkillsUsecase) UsersWithSkillResponse {
	res := UsersWithSkillResponse{}
	for _, val := range users {
		res.Users = append(res.Users, ToUserWithSkillResponse(&val))
	}
	return res
}

func ToUserWithSkillResponse(user *UserWithSkillsUsecase) UserWithSkillResponse {
	return UserWithSkillResponse{
		UserResponse: UserResponse{
			ID:           user.ID,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			About:        user.About,
			Avatar:       user.Avatar,
			IsSearchable: user.IsSearchable,
		},
		Skills: user.Skills,
	}

}
func ToUserResponse(user *UserUsecase) UserResponse {
	return UserResponse{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		About:        user.About,
		Avatar:       user.Avatar,
		IsSearchable: user.IsSearchable,
	}
}
