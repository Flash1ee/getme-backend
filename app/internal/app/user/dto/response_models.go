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
type UserWithSkillsResponse struct {
	UserResponse
	Skills []string `json:"skills"`
}

//easyjson:json
type UsersWithSkillResponse struct {
	Users []UserWithSkillsResponse `json:"users"`
}

func ToUsersWithSkillResponse(users []UserWithSkillsUsecase) UsersWithSkillResponse {
	res := UsersWithSkillResponse{
		Users: []UserWithSkillsResponse{},
	}
	for _, val := range users {
		res.Users = append(res.Users, ToUserWithSkillsResponse(&val))
	}
	return res
}

func ToUserWithSkillsResponse(user *UserWithSkillsUsecase) UserWithSkillsResponse {
	return UserWithSkillsResponse{
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
func ToUserResponse(user UserUsecase) UserResponse {
	return UserResponse{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		About:        user.About,
		Avatar:       user.Avatar,
		IsSearchable: user.IsSearchable,
	}
}
func ToUsersResponse(user []UserUsecase) UsersResponse {
	res := &UsersResponse{
		Users: make([]UserResponse, 0),
	}
	for _, val := range user {
		res.Users = append(res.Users, ToUserResponse(val))
	}
	return *res
}
