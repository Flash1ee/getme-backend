package offer_usecase

import (
	"getme-backend/internal/app/offer/dto"
	plan_dto "getme-backend/internal/app/plans/dto"
	dto2 "getme-backend/internal/app/user/dto"
)

type Usecase interface {
	//	Create with Errors:
	// 	skill_usecase.SkillNotExists
	// 	offer_usecase.MentorNotExist
	//	offer_usecase.LogicError
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	Create(data *dto.OfferUsecaseDTO) (int64, error)
	//	Get with Errors:
	// 	offer_usecase.NotMentor
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	Get(mentorID int64) ([]dto2.UserWithOfferIDUsecase, error)
	//Accept with Errors:
	//		AlreadyAccepted
	//		InvalidOfferID
	//		postgresql_utilits.NotFound
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	Accept(userID int64, data *dto.OfferAcceptUsecaseDTO) (*plan_dto.PlansCreateUsecaseDTO, error)
	//Delete with Errors:
	//		InvalidOfferID
	//		postgresql_utilits.NotFound
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	Delete(userID int64, offerID int64) error
}
