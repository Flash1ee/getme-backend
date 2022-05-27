package offer_postgresql

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"getme-backend/internal/app/offer/entities"
	postgresql_utilits "getme-backend/internal/pkg/utilits/postgresql"
)

type OfferRepository struct {
	store *sqlx.DB
}

func NewOfferRepository(store *sqlx.DB) *OfferRepository {
	return &OfferRepository{
		store: store,
	}
}

const queryCreateOffer = `
INSERT INTO offers(skill_name, mentor_id, mentee_id) VALUES (?, ?, ?) RETURNING id;`

//	Create with Errors:
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (r *OfferRepository) Create(data *entities.Offer) (int64, error) {
	id := int64(-1)
	query := r.store.Rebind(queryCreateOffer)

	if err := r.store.QueryRowx(query, data).Scan(&id); err != nil {
		return id, postgresql_utilits.NewDBError(err)
	}
	return id, nil
}
func (r *OfferRepository) Accept(id int64) error {
	return nil
}

const checkExistsOffer = "SELECT count(*) from getme_db.public.offers where mentee_id = ? and mentor_id = ?;"

// CheckExists Errors:
//		postgresql_utilits.Exists
//		postgresql_utilits.NotFound
// 		app.GeneralError with Errors:
// 			postgresql_utilits.DefaultErrDB
func (repo *OfferRepository) CheckExists(menteeID, mentorID int64) error {
	cnt := int64(0)
	if err := repo.store.Get(&cnt, checkExistsOffer, menteeID, mentorID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return postgresql_utilits.NotFound
		}
		return postgresql_utilits.NewDBError(errors.Wrap(err, fmt.Sprintf("OfferRepository - CheckExists(%v, %v)", menteeID, mentorID)))
	}

	if cnt != 0 {
		return postgresql_utilits.Exists
	}
	return postgresql_utilits.NotFound
}
