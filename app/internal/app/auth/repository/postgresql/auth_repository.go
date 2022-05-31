package repository_postgresql

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	"getme-backend/internal/app/auth/entities"
	postgresql_utilits "getme-backend/internal/pkg/utilits/postgresql"
)

type AuthRepository struct {
	store *sqlx.DB
}

func NewAuthRepository(store *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		store: store,
	}
}

const queryFindBySimpleLogin = `SELECT id, login, encrypted_password, user_id from users_simple_auth where login=?;`

// FindByLoginSimple Errors:
// 		postgresql_utilits.NotFound
// 		app.GeneralError with Errors:
// 			postgresql_utilits.DefaultErrDB
func (repo *AuthRepository) FindByLoginSimple(login string) (*entities.SimpleAuth, error) {
	user := &entities.SimpleAuth{}

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

const queryCheckExistsUser = `
	SELECT tg_id, created_at, last_auth, user_id from telegram_auth where tg_id=?
`

// FindByTelegramID Errors:
//		postgresql_utilits.NotFound
// 		app.GeneralError with Errors:
// 			postgresql_utilits.DefaultErrDB
func (repo *AuthRepository) FindByTelegramID(tgID int64) (*entities.TelegramAuth, error) {
	user := &entities.TelegramAuth{}
	query := repo.store.Rebind(queryCheckExistsUser)

	if err := repo.store.QueryRow(query, tgID).Scan(&user.TelegramID, &user.CreatedAt, &user.LastAuth, &user.UserID); err != nil {
		if err == sql.ErrNoRows {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}

	return user, nil
}

const queryCreateTelegramRecord = `
	INSERT INTO telegram_auth (tg_id, user_id)
	VALUES (?, ?);
	`

// CreateTelegramAuthRecord Errors:
// 		app.GeneralError with Errors:
// 			postgresql_utilits.DefaultErrDB
func (repo *AuthRepository) CreateTelegramAuthRecord(auth *entities.TelegramAuth) (*entities.TelegramAuth, error) {
	user := &entities.TelegramAuth{}
	query := repo.store.Rebind(queryCreateTelegramRecord)

	if _, err := repo.store.Exec(query, auth.TelegramID, auth.UserID); err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}

	return user, nil
}

const createSimpleQuery = `INSERT INTO users_simple_auth (login, encrypted_password, user_id) VALUES (?, ?, ?) RETURNING user_id`

// CreateSimple Errors:
// 		auth_repository.EmailAlreadyExist
// 		auth_repository.NicknameAlreadyExist
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (repo *AuthRepository) CreateSimple(u *entities.SimpleAuth) (*entities.SimpleAuth, error) {
	query := repo.store.Rebind(createSimpleQuery)

	if err := repo.store.QueryRow(query, u.Login, u.EncryptedPassword, u.UserID).Scan(&u.ID); err != nil {
		if _, ok := err.(*pq.Error); ok {
			return nil, parsePQError(err.(*pq.Error))
		}
		return nil, postgresql_utilits.NewDBError(err)
	}

	return u, nil
}

const updateAuthTimeQuery = `UPDATE telegram_auth SET last_auth = now() where tg_id = ?;`

func (repo *AuthRepository) UpdateTelegramAuthTime(tgID int64) error {
	query := repo.store.Rebind(updateAuthTimeQuery)
	_, err := repo.store.Exec(query, tgID)
	if err != nil {
		return postgresql_utilits.NewDBError(err)
	}
	return nil
}
