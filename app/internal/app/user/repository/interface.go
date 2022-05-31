package user_repository

import (
	skill_entities "getme-backend/internal/app/skill/entities"
	"getme-backend/internal/app/user/entities"
)

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type Repository interface {
	//	FindByIDWithSkill with Errors:
	//		postgresql_utilits.NotFound
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	FindByIDWithSkill(id int64) (*[]entities_user.UserWithSkill, error)
	FindMentorByID(id int64) (*[]entities_user.UserWithSkill, error)
	//	FindByNickname with Errors:
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	FindByNickname(nickname string) (*entities_user.User, error)
	//	FindByID with Errors:
	//		postgresql_utilits.NotFound
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	FindByID(id int64) (*entities_user.User, error)
	// CreateBaseUser Errors:
	// 		user_repository.EmailAlreadyExist
	// 		user_repository.NicknameAlreadyExist
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	CreateBaseUser(nickname string) (int64, error)
	CreateFilledUser(data *entities_user.User) (int64, error)
	// UpdateUser Errors:
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	UpdateUser(user *entities_user.User) (*entities_user.User, error)
	//	GetUsersBySkills with Errors:
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	GetUsersBySkills(data []skill_entities.Skill) ([]entities_user.UserWithSkill, error)
	//	GetMenteeByMentorWithOfferID with Errors:
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	GetMenteeByMentorWithOfferID(mentorID int64) ([]entities_user.UserWithOfferID, error)
	//UpdateMentorStatus with Errors:
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	UpdateMentorStatus(mentorID int64) (bool, error)
	//SetMentorStatus with Errors:
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	SetMentorStatus(mentorID int64, status bool) (bool, error)
}
