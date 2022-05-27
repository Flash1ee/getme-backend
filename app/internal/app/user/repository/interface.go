package user_repository

import (
	skill_entities "getme-backend/internal/app/skill/entities"
	"getme-backend/internal/app/user/entities"
)

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type Repository interface {
	//	FindByID with Errors:
	//		postgresql_utilits.NotFound
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	FindByID(id int64) (*[]entities.UserWithSkill, error)

	FindMentorByID(id int64) (*[]entities.UserWithSkill, error)

	//	FindByNickname with Errors:
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	FindByNickname(nickname string) (*entities.User, error)
	// CreateBaseUser Errors:
	// 		user_repository.EmailAlreadyExist
	// 		user_repository.NicknameAlreadyExist
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	CreateBaseUser(nickname string) (int64, error)
	CreateFilledUser(data *entities.User) (int64, error)
	// UpdateUser Errors:
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	UpdateUser(user *entities.User) (*entities.User, error)
	//	GetUsersBySkills with Errors:
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	GetUsersBySkills(data []skill_entities.Skill) ([]entities.UserWithSkill, error)
	//	GetMenteeByMentor with Errors:
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	GetMenteeByMentor(mentorID int64) ([]entities.User, error)
}
