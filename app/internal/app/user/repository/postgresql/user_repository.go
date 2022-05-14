package repository_postgresql

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	"getme-backend/internal/app"
	"getme-backend/internal/app/user/entities"
)

type UserRepository struct {
	store *sqlx.DB
}

func NewUserRepository(store *sqlx.DB) *UserRepository {
	return &UserRepository{
		store: store,
	}
}

const queryCreateUser = `
	INSERT INTO users (tg_id, first_name, last_name, nickname, avatar) 
VALUES ($1, $2, $3, $4, $5) ON CONFLICT (tg_id) DO UPDATE
	SET
		first_name = excluded.first_name, 
		last_name = excluded.last_name,
		nickname = excluded.nickname,
		avatar = excluded.avatar,
		updated_at = now()
	RETURNING id;
	`

// Create Errors:
// 		app.GeneralError with Error
// 			CreateError
func (r *UserRepository) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	err := r.store.QueryRowxContext(ctx, queryCreateUser, user.TelegramID, user.FirstName, user.LastName, user.Nickname, user.Avatar).Scan(&user.ID)
	if err != nil {
		return nil, app.GeneralError{
			Err:         CreateError,
			ExternalErr: err,
		}
	}

	return user, nil
}

var queryGetUserByTelegramID = `
	SELECT id, tg_id, first_name, last_name, nickname, about, avatar, is_searchable, created_at, updated_at
	from users WHERE tg_id = $1
	`

// GetUserByTelegramID Errors:
// 		app.GeneralError with Error
// 			GetError
func (r *UserRepository) GetUserByTelegramID(ctx context.Context, tgID int64) (*entities.User, error) {
	user := &entities.User{}
	if err := r.store.GetContext(ctx, user, queryGetUserByTelegramID, tgID); err != nil {
		if err != sql.ErrNoRows {
			return nil, app.GeneralError{
				Err:         GetError,
				ExternalErr: err,
			}
		}
	}
	return user, nil
}
