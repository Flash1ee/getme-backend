package dto

import "time"

//go:generate easyjson -all -disallow_unknown_fields response_models.go

//easyjson:json
type UserResponse struct {
	ID           int64     `json:"id"`
	FirstName    string    `json:"first_name,omitempty"`
	LastName     string    `json:"last_name,omitempty"`
	Nickname     string    `json:"nickname"`
	About        string    `json:"about,omitempty"`
	Avatar       string    `json:"avatar,omitempty"`
	IsSearchable bool      `json:"is_mentor"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

//easyjson:json
type UsersResponse struct {
	Users []UserResponse `json:"users"`
}

//easyjson:json
type UserWithSkillResponse struct {
	UserResponse
	Skill []string `json:"skill"`
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
			Nickname:     user.Nickname,
			About:        user.About,
			Avatar:       user.Avatar,
			IsSearchable: user.IsSearchable,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
		},
		Skill: user.Skill,
	}

}
func ToUserResponse(user *UserUsecase) UserResponse {
	return UserResponse{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Nickname:     user.Nickname,
		About:        user.About,
		Avatar:       user.Avatar,
		IsSearchable: user.IsSearchable,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
