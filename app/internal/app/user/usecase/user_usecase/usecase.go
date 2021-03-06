package user_usecase

import (
	"getme-backend/internal/app/user/dto"
	"getme-backend/internal/app/user/entities"
	"getme-backend/internal/app/user/repository"
	user_usecase "getme-backend/internal/app/user/usecase"
	"getme-backend/internal/pkg/usecase"
	postgresql_utilits "getme-backend/internal/pkg/utilits/postgresql"
)

type UserUsecase struct {
	usecase.BaseUsecase
	userRepository user_repository.Repository
	authChecker    authChecker
}

func NewUserUsecase(repo user_repository.Repository, authCheck authChecker) *UserUsecase {
	return &UserUsecase{
		userRepository: repo,
		authChecker:    authCheck,
	}
}

func (u *UserUsecase) FindByNickname(nickname string) (*entities_user.User, error) {
	res, err := u.userRepository.FindByNickname(nickname)
	if err == postgresql_utilits.NotFound {
		return nil, user_usecase.UserNotFound
	}
	return res, err
}

func (u *UserUsecase) CreateBaseUser(nickname string) (int64, error) {
	return u.userRepository.CreateBaseUser(nickname)
}

func (u *UserUsecase) CreateFilledUser(data *dto.UserUsecase) (int64, error) {
	entityUser := data.ToUserEntity()
	us, err := u.userRepository.CreateFilledUser(entityUser)

	return us, err
}
func (u *UserUsecase) FindByID(id int64) (*dto.UserWithSkillsUsecase, error) {
	user, err := u.userRepository.FindByIDWithSkill(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return &dto.UserWithSkillsUsecase{}, nil
	}
	res := filterUsersData(*user)

	return dto.ToUserWithSkillUsecase(res), nil
}

func (u *UserUsecase) UpdateUser(user *dto.UserWithSkillsUsecase) (*dto.UserWithSkillsUsecase, error) {
	userDTO := user.ToUserWithSkillEntity()
	res, err := u.userRepository.UpdateUser(userDTO)
	if err != nil {
		return nil, err
	}
	return dto.ToUserWithSkillsUsecase(res), nil
}

//GetMentorStatus with Errors:
//		postgresql_utilits.NotFound
//		app.GeneralError with Errors
//			postgresql_utilits.DefaultErrDB
func (u *UserUsecase) GetMentorStatus(id int64) (*dto.UserStatusUsecase, error) {
	user, err := u.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	res := &dto.UserStatusUsecase{}
	switch user.IsSearchable {
	case true:
		res.IsMentor = true
	default:
		res.IsMentor = false
	}

	return res, nil
}

//	UpdateMentorStatus with Errors:
//		postgresql_utilits.NotFound
//		app.GeneralError with Errors
//			postgresql_utilits.DefaultErrDB
func (u *UserUsecase) UpdateMentorStatus(data *dto.UserStatusUsecase) (*dto.UserStatusUsecase, error) {
	_, err := u.userRepository.FindByID(data.UserID)
	if err != nil {
		return nil, err
	}
	res, err := u.userRepository.UpdateMentorStatus(data.UserID)
	//res, err := u.userRepository.SetMentorStatus(data.UserID, data.IsMentor)

	return &dto.UserStatusUsecase{
		UserID:   data.UserID,
		IsMentor: res,
	}, nil
}
