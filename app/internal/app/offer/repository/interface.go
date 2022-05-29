package offer_repository

import (
	"getme-backend/internal/app/offer/entities"
)

type Repository interface {
	//	Create with Errors:
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	Create(data *entities.Offer) (int64, error)
	// CheckExists Errors:
	//		postgresql_utilits.Exists
	//		postgresql_utilits.NotFound
	// 		app.GeneralError with Errors:
	// 			postgresql_utilits.DefaultErrDB
	CheckExists(menteeID, mentorID int64) error
	// GetByID Errors:
	//		postgresql_utilits.NotFound
	// 		app.GeneralError with Errors:
	// 			postgresql_utilits.DefaultErrDB
	GetByID(id int64) (*entities.Offer, error)
	// Delete Errors:
	// 		app.GeneralError with Errors:
	// 			postgresql_utilits.DefaultErrDB
	Delete(id int64) error
}
