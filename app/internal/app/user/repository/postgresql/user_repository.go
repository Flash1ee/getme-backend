package repository_postgresql

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	"getme-backend/internal/app"
	"getme-backend/internal/app/user/entities"
	postgresql_utilits "getme-backend/internal/pkg/utilits/postgresql"
)

type UserRepository struct {
	store *sqlx.DB
}

func NewUserRepository(store *sqlx.DB) *UserRepository {
	return &UserRepository{
		store: store,
	}
}

const querySimpleCreate = `
	INSERT INTO users (tg_id, first_name, last_name, nickname, avatar) 
	VALUES (?, ?, ?, ?, ?);
	`

// Create Errors:
// 		app.GeneralError with Error
// 			CreateError
func (r *UserRepository) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	query := r.store.Rebind(querySimpleCreate)

	err := r.store.QueryRowxContext(ctx, query, user.TelegramID, user.FirstName, user.LastName, user.Nickname, user.Avatar).Scan(&user.ID)
	if err != nil {
		return nil, app.GeneralError{
			Err:         CreateError,
			ExternalErr: err,
		}
	}

	return user, nil
}

const queryCreateUpdateUser = `
	INSERT INTO users (tg_id, first_name, last_name, nickname, avatar) 
VALUES (?, ?, ?, ?, ?) ON CONFLICT (tg_id) DO UPDATE
	SET
		first_name = excluded.first_name, 
		last_name = excluded.last_name,
		nickname = excluded.nickname,
		avatar = excluded.avatar,
		updated_at = now()
	RETURNING id;
	`

// CreateWithUpdate Errors:
// 		app.GeneralError with Error
// 			CreateError
func (r *UserRepository) CreateWithUpdate(ctx context.Context, user *entities.User) (*entities.User, error) {
	query := r.store.Rebind(queryCreateUpdateUser)

	err := r.store.QueryRowxContext(ctx, query, user.TelegramID, user.FirstName, user.LastName, user.Nickname, user.Avatar).Scan(&user.ID)
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
	from users WHERE tg_id = ?
	`

// GetUserByTelegramID Errors:
// 		app.GeneralError with Error
// 			GetError
func (r *UserRepository) GetUserByTelegramID(ctx context.Context, tgID int64) (*entities.User, error) {
	query := r.store.Rebind(queryGetUserByTelegramID)

	user := &entities.User{}
	if err := r.store.GetContext(ctx, user, query, tgID); err != nil {
		if err != sql.ErrNoRows {
			return nil, app.GeneralError{
				Err:         GetError,
				ExternalErr: err,
			}
		}
	}
	return user, nil
}

const queryCheckExistsUser = `
	SELECT tg_id from users where tg_id=?
`

// CheckExists Errors:
// 		app.GeneralError with Errors:
// 			postgresql_utilits.DefaultErrDB
func (repo *UserRepository) CheckExists(tgID int64) (bool, error) {
	user := &entities.User{}
	query := repo.store.Rebind(queryCheckExistsUser)

	if err := repo.store.QueryRow(query, tgID).Scan(&user.TelegramID); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, postgresql_utilits.NewDBError(err)
	}

	return true, nil
}

const queryFindBySimpleLogin = `SELECT id, login, encrypted_password, user_id from users_simple_auth where login=?`

// FindByLoginSimple Errors:
// 		postgresql_utilits.NotFound
// 		app.GeneralError with Errors:
// 			postgresql_utilits.DefaultErrDB
func (repo *UserRepository) FindByLoginSimple(login string) (*entities.UserSimpleAuth, error) {
	user := &entities.UserSimpleAuth{}

	userID := sql.NullInt64{}
	query := repo.store.Rebind(queryFindBySimpleLogin)

	if err := repo.store.QueryRow(query, login).
		Scan(&user.ID, &user.Login, &user.EncryptedPassword, &userID); err != nil {
		if err == sql.ErrNoRows {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)

	}
	if userID.Int64 != 0 {
		user.UserID = userID.Int64
	}

	return user, nil
}

const createSimpleQuery = `INSERT INTO users_simple_auth (login, encrypted_password) VALUES (?, ?) RETURNING id`

// CreateSimple Errors:
// 		LoginAlreadyExist
// 		NicknameAlreadyExist
// 		app.GeneralError with Errors
// 			repository.DefaultErrDB
func (repo *UserRepository) CreateSimple(u *entities.UserSimpleAuth) (*entities.UserSimpleAuth, error) {
	query := repo.store.Rebind(createSimpleQuery)

	if err := repo.store.QueryRow(query, u.Login, u.EncryptedPassword).Scan(&u.ID); err != nil {
		if _, ok := err.(*pq.Error); ok {
			return nil, parsePQError(err.(*pq.Error))
		}
		return nil, postgresql_utilits.NewDBError(err)
	}

	return u, nil
}
