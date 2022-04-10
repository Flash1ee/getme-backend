package postgresql

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"

	"getme-backend/internal/app/user/entities"
	"getme-backend/internal/pkg/utilits/postgresql"
)

const (
	updateUser = `
					UPDATE users SET 
					    fullname = COALESCE(NULLIF(TRIM($1), ''), fullname),
					    about = COALESCE(NULLIF(TRIM($2), ''), about),
					    email = COALESCE(NULLIF(TRIM($3), ''), email) 
					WHERE nickname = $4
					RETURNING nickname, fullname, about, email`

	getQuery = "SELECT nickname, fullname, about, email FROM users WHERE nickname = $1"

	createQuery = `    
						WITH sel AS (
						    SELECT nickname, fullname, about, email
							FROM users
							WHERE nickname = $1 OR email = $4
						), ins as (
							INSERT INTO users (nickname, fullname, about, email)
								SELECT $1, $2, $3, $4
								WHERE not exists (select 1 from sel)
							RETURNING nickname, fullname, about, email
						)
						SELECT nickname, fullname, about, email, 0
						FROM ins
						UNION ALL
						SELECT nickname, fullname, about, email, 1
						FROM sel
					`
)

type UserRepository struct {
	store *sqlx.DB
}

func NewUserRepository(store *sqlx.DB) *UserRepository {
	return &UserRepository{
		store: store,
	}
}

func (r *UserRepository) Create(us *entities.User) ([]entities.User, error) {
	rows, err := r.store.Queryx(createQuery,
		us.Nickname,
		us.Fullname,
		us.About,
		us.Email,
	)
	if err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}

	isCorrect := 0
	var res []entities.User
	for rows.Next() {
		if err = rows.Scan(
			&us.Nickname,
			&us.Fullname,
			&us.About,
			&us.Email,
			&isCorrect,
		); err != nil {
			_ = rows.Close()
			return nil, postgresql_utilits.NewDBError(err)
		}
		res = append(res, *us)
	}

	if err = rows.Err(); err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}

	if isCorrect == 1 {
		return res, postgresql_utilits.Conflict
	}
	return res, nil
}

func (r *UserRepository) Get(nickname string) (*entities.User, error) {
	res := &entities.User{Nickname: nickname}
	if err := r.store.QueryRowx(getQuery, res.Nickname).
		Scan(
			&res.Nickname,
			&res.Fullname,
			&res.About,
			&res.Email,
		); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}
	return res, nil
}

func (r *UserRepository) Update(us *entities.User) (*entities.User, error) {
	if err := r.store.QueryRowx(updateUser,
		us.Fullname,
		us.About,
		us.Email,
		us.Nickname,
	).
		Scan(
			&us.Nickname,
			&us.Fullname,
			&us.About,
			&us.Email,
		); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NotFound
		}
		return nil, parsePQError(err.(*pq.Error))
	}
	return us, nil
}
