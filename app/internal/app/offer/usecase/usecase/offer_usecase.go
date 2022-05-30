package offer_usecase

import (
	"github.com/pkg/errors"

	"getme-backend/internal/app"
	"getme-backend/internal/app/offer/dto"
	offer_repository "getme-backend/internal/app/offer/repository"
	offer_usecase "getme-backend/internal/app/offer/usecase"
	plan_dto "getme-backend/internal/app/plans/dto"
	"getme-backend/internal/app/plans/entities"
	plan_repository "getme-backend/internal/app/plans/repository"
	entities2 "getme-backend/internal/app/skill/entities"
	skill_repository "getme-backend/internal/app/skill/repository"
	user_dto "getme-backend/internal/app/user/dto"
	user_repository "getme-backend/internal/app/user/repository"
	"getme-backend/internal/pkg/usecase"
	postgresql_utilits "getme-backend/internal/pkg/utilits/postgresql"
)

type OfferUsecase struct {
	usecase.BaseUsecase
	offerRepository offer_repository.Repository
	userRepository  user_repository.Repository
	skillRepository skill_repository.Repository
	planRepository  plan_repository.Repository
}

func NewOfferUsecase(repoOffer offer_repository.Repository, repoUser user_repository.Repository, repoSkill skill_repository.Repository, repoPlan plan_repository.Repository) *OfferUsecase {
	return &OfferUsecase{
		offerRepository: repoOffer,
		userRepository:  repoUser,
		skillRepository: repoSkill,
		planRepository:  repoPlan,
	}
}

//	Create with Errors:
// 	skill_usecase.SkillNotExists
//	offer_usecase.LogicError
// 	offer_usecase.MentorNotExist
// 	offer_usecase.AlreadyExists
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (u *OfferUsecase) Create(data *dto.OfferUsecaseDTO) (int64, error) {
	if data == nil {
		return app.InvalidInt, offer_usecase.NilDataArg
	}
	if data.MentorID == data.MenteeID {
		return app.InvalidInt, offer_usecase.LogicError
	}
	if _, err := u.userRepository.FindMentorByID(data.MentorID); err != nil {
		if errors.Is(err, postgresql_utilits.NotFound) {
			return app.InvalidInt, offer_usecase.MentorNotExist
		}
	}
	if err := u.skillRepository.CheckExists(data.SkillName); err != postgresql_utilits.Exists {
		//return app.InvalidInt, skill_usecase.SkillNotExists
	}
	if err := u.offerRepository.CheckExists(data.MenteeID, data.MentorID); err != postgresql_utilits.NotFound {
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
func (u *OfferUsecase) Get(mentorID int64) ([]user_dto.UserWithOfferIDUsecase, error) {
	if _, err := u.userRepository.FindMentorByID(mentorID); err != nil {
		if errors.Is(err, postgresql_utilits.NotFound) {
			return nil, offer_usecase.NotMentor
		}
	}
	res, err := u.userRepository.GetMenteeByMentorWithOfferID(mentorID)
	if err != nil {
		return nil, err
	}
	return user_dto.ToUserWithOfferIDUsecases(res), nil
}

//Accept with Errors:
//		InvalidOfferID
//		postgresql_utilits.NotFound
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (u *OfferUsecase) Accept(userID int64, data *dto.OfferAcceptUsecaseDTO) (*plan_dto.PlansCreateUsecaseDTO, error) {
	res, err := u.offerRepository.GetByID(data.OfferID)
	if err != nil {
		return nil, err
	}
	if res.MentorID != userID {
		return nil, offer_usecase.InvalidOfferID
	}
	createdPlan, err := u.planRepository.Create(data.OfferID, entities2.GetSkills(data.Skills), entities.Plan{
		Name:     data.Title,
		About:    data.Description,
		MenteeID: res.MenteeID,
		MentorID: res.MentorID,
	})
	if err != nil {
		return nil, err
	}

	return plan_dto.EntityToPlanCreateUsecaseDTO(createdPlan), nil

}

//Delete with Errors:
//		InvalidOfferID
//		postgresql_utilits.NotFound
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (u *OfferUsecase) Delete(userID int64, offerID int64) error {
	res, err := u.offerRepository.GetByID(offerID)
	if err != nil {
		return err
	}
	if res.MentorID != userID {
		return offer_usecase.InvalidOfferID
	}
	if res.Status == false {
		return postgresql_utilits.NotFound
	}

	if err := u.offerRepository.Delete(offerID); err != nil {
		return err
	}
	return nil
}
