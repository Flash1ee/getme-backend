package offer_repository

import (
	"getme-backend/internal/app/offer/entities"
)

type Repository interface {
	//	Create with Errors:
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	Create(data *entities.Offer) (int64, error)
	Accept(id int64) error
	// CheckExists Errors:
	//		postgresql_utilits.Exists
	//		postgresql_utilits.NotFound
	// 		app.GeneralError with Errors:
	// 			postgresql_utilits.DefaultErrDB
	CheckExists(menteeID, mentorID int64) error
}
