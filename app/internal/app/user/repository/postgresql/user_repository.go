package repository_postgresql

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	skill_entities "getme-backend/internal/app/skill/entities"
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
	if err := repo.store.QueryRow(query, data.FirstName.String, data.LastName.String, data.Nickname, data.Avatar.String).
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
//		postgresql_utilits.NotFound
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (repo *UserRepository) FindByID(id int64) (*entities.User, error) {
	query := repo.store.Rebind(queryFindByID)
	user := &entities.User{}

	err := repo.store.Get(user, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}
	return user, nil
}

const queryUpdateUser = `update users set
    first_name = COALESCE(NULLIF(TRIM(?), ''), first_name),
    last_name = COALESCE(NULLIF(TRIM(?), ''), last_name),
    about = COALESCE(NULLIF(TRIM(?), ''), about)
WHERE id = ? returning first_name, last_name, nickname, about, avatar, is_searchable;`

func (repo *UserRepository) UpdateUser(user *entities.User) (*entities.User, error) {
	query := repo.store.Rebind(queryUpdateUser)
	userFromDB := &entities.User{}

	err := repo.store.QueryRowx(query, user.FirstName.String, user.LastName.String, user.About.String, user.ID).
		Scan(&userFromDB.FirstName, &userFromDB.LastName, &userFromDB.Nickname, &userFromDB.About, &userFromDB.Avatar, &userFromDB.IsSearchable)
	if err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}

	return user, nil
}

const queryGetUsersBySkillsFirst = `SELECT users.id, first_name, last_name, nickname, about, avatar, is_searchable, skill_name from users
JOIN users_skills as us on us.user_id = users.id where us.skill_name IN (?)`
const queryGetUsersBySkillsSecond = ` and skill_name = (
select skill_name from users_skills where user_id = us.user_id and skill_name IN (?) order by skill_name limit 1
    );`

//const queryGetUsersBySkillsFirst = `SELECT first_name, last_name, nickname, about, avatar, is_searchable, skill_name from users
//JOIN users_skills as us on us.user_id = users.id where us.skill_name IN (?)`
//const queryGetUsersBySkillsSecond = ` and skill_name = (
//select skill_name from users_skills where user_id = us.user_id and skill_name IN (?) order by skill_name limit 1
//    );`

//	GetUsersBySkills with Errors:
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (repo *UserRepository) GetUsersBySkills(data []skill_entities.Skill) ([]entities.UserWithSkill, error) {
	skills := getSkillNameFromSkills(data)
	usersWithSkills := &[]entities.UserWithSkill{}
	queryFirst, args, err := sqlx.In(queryGetUsersBySkillsFirst, skills)
	if err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}
	//querySecond, _, err := sqlx.In(queryGetUsersBySkillsSecond, skills)
	//if err != nil {
	//	return nil, postgresql_utilits.NewDBError(err)
	//}
	query := queryFirst

	//query := queryFirst + querySecond
	query = repo.store.Rebind(query)
	var queryArgs []interface{}
	queryArgs = append(queryArgs, args...)
	//queryArgs = append(queryArgs, args...)

	if err = repo.store.Select(usersWithSkills, query, queryArgs...); err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}

	return *usersWithSkills, nil
}
