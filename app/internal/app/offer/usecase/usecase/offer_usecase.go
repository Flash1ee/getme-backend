package offer_usecase

import (
	"github.com/pkg/errors"

	"getme-backend/internal/app"
	"getme-backend/internal/app/offer/dto"
	offer_repository "getme-backend/internal/app/offer/repository"
	offer_usecase "getme-backend/internal/app/offer/usecase"
	skill_repository "getme-backend/internal/app/skill/repository"
	skill_usecase "getme-backend/internal/app/skill/usecase"
	dto2 "getme-backend/internal/app/user/dto"
	user_repository "getme-backend/internal/app/user/repository"
	"getme-backend/internal/pkg/usecase"
	postgresql_utilits "getme-backend/internal/pkg/utilits/postgresql"
)

type OfferUsecase struct {
	usecase.BaseUsecase
	offerRepository offer_repository.Repository
	userRepository  user_repository.Repository
	skillRepository skill_repository.Repository
}

func NewOfferUsecase(repoOffer offer_repository.Repository, repoUser user_repository.Repository, repoSkill skill_repository.Repository) *OfferUsecase {
	return &OfferUsecase{
		offerRepository: repoOffer,
		userRepository:  repoUser,
		skillRepository: repoSkill,
	}
}

//	Create with Errors:
// 	skill_usecase.SkillNotExists
// 	offer_usecase.MentorNotExist
// 	offer_usecase.AlreadyExists
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (u *OfferUsecase) Create(data *dto.OfferUsecaseDTO) (int64, error) {
	if data == nil {
		return app.InvalidInt, offer_usecase.NilDataArg
	}
	if _, err := u.userRepository.FindMentorByID(data.MentorID); err != nil {
		if errors.Is(err, postgresql_utilits.NotFound) {
			return app.InvalidInt, offer_usecase.MentorNotExist
		}
	}
	if err := u.skillRepository.CheckExists(data.SkillName); err != postgresql_utilits.Exists {
		return app.InvalidInt, skill_usecase.SkillNotExists
	}
	if err := u.offerRepository.CheckExists(data.MenteeID, data.MentorID); err != nil {
		return app.InvalidInt, offer_usecase.AlreadyExists
	}
	res, err := u.offerRepository.Create(data.ToEntity())
	if err != nil {
		return app.InvalidInt, err
	}

	return res, err
}

//	Get with Errors:
// 	offer_usecase.NotMentor
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (u *OfferUsecase) Get(mentorID int64) ([]dto2.UserUsecase, error) {
	if _, err := u.userRepository.FindMentorByID(mentorID); err != nil {
		if errors.Is(err, postgresql_utilits.NotFound) {
			return nil, offer_usecase.NotMentor
		}
	}
	res, err := u.userRepository.GetMenteeByMentor(mentorID)
	if err != nil {
		return nil, err
	}
	return dto2.ToUserUsecases(res), nil
}
