package repository_postgresql

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

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

//const queryCreateUpdateUser = `
//	INSERT INTO users (tg_id, first_name, last_name, nickname, avatar)
//VALUES (?, ?, ?, ?, ?) ON CONFLICT (tg_id) DO UPDATE
//	SET
//		first_name = excluded.first_name,
//		last_name = excluded.last_name,
//		nickname = excluded.nickname,
//		avatar = excluded.avatar,
//		updated_at = now()
//	RETURNING id;
//	`

//// CreateWithUpdate Errors:
//// 		app.GeneralError with Error
//// 			CreateError
//func (r *UserRepository) CreateWithUpdate(ctx context.Context, user *entities.User) (*entities.User, error) {
//	query := r.store.Rebind(queryCreateUpdateUser)
//
//	err := r.store.QueryRowxContext(ctx, query, user.TelegramID, user.FirstName, user.LastName, user.Nickname, user.Avatar).Scan(&user.ID)
//	if err != nil {
//		return nil, app.GeneralError{
//			Err:         CreateError,
//			ExternalErr: err,
//		}
//	}
//
//	return user, nil
//}

const queryGetUserByNickname = `
SELECT id, first_name, last_name, nickname, about, avatar, 
       is_searchable, created_at, updated_at from users where nickname = ?;`

func (repo *UserRepository) FindByNickname(nickname string) (*entities.User, error) {
	query := repo.store.Rebind(queryGetUserByNickname)

	user := &entities.User{}
	if err := repo.store.QueryRow(query, nickname).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Nickname,
		&user.About, &user.Avatar, &user.IsSearchable, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, postgresql_utilits.NotFound
		}
		if _, ok := err.(*pq.Error); ok {
			return nil, parsePQError(err.(*pq.Error))
		}
		return nil, postgresql_utilits.NewDBError(err)
	}

	return user, nil
}

const createBaseUserQuery = `INSERT INTO users (nickname) VALUES (?) RETURNING id;`

// CreateBaseUser Errors:
// 		auth_repository.EmailAlreadyExist
// 		auth_repository.NicknameAlreadyExist
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (repo *UserRepository) CreateBaseUser(nickname string) (int64, error) {
	query := repo.store.Rebind(createBaseUserQuery)
	ID := int64(-1)
	if err := repo.store.QueryRow(query, nickname).Scan(&ID); err != nil {
		if _, ok := err.(*pq.Error); ok {
			return ID, parsePQError(err.(*pq.Error))
		}
		return ID, postgresql_utilits.NewDBError(err)
	}

	return ID, nil
}

const createFilledUserQuery = `INSERT INTO users (first_name, last_name, nickname, avatar)
VALUES (?, ?, ?, ?) RETURNING id;`

func (repo *UserRepository) CreateFilledUser(data *entities.User) (int64, error) {
	query := repo.store.Rebind(createFilledUserQuery)
	ID := int64(-1)
	if err := repo.store.QueryRow(query, data.FirstName, data.FirstName, data.LastName, data.Nickname, data.Avatar).
		Scan(&ID); err != nil {
		if _, ok := err.(*pq.Error); ok {
			return ID, parsePQError(err.(*pq.Error))
		}
		return ID, postgresql_utilits.NewDBError(err)
	}

	return ID, nil

}

const queryFindByID = `
SELECT * from users where id = ?;`

//	FindByID with Errors:
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (repo *UserRepository) FindByID(id int64) (*entities.User, error) {
	query := repo.store.Rebind(queryFindByID)
	user := &entities.User{}

	err := repo.store.Get(user, query, id)
	if err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}
	return user, nil
}
