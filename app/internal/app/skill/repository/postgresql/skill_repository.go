package skill_repository_postgresql

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"getme-backend/internal/app/skill/entities"
	postgresql_utilits "getme-backend/internal/pkg/utilits/postgresql"
)

const defaultCountSkills = 10

type SkillRepository struct {
	store *sqlx.DB
}

func NewSkillRepository(st *sqlx.DB) *SkillRepository {
	return &SkillRepository{
		store: st,
	}
}

const queryGetAllSkills = `select name from skills;`

// GetAllSkills Errors:
// 		app.GeneralError with Errors:
// 			postgresql_utilits.DefaultErrDB
func (repo *SkillRepository) GetAllSkills() ([]entities.Skill, error) {
	res := make([]entities.Skill, 0, defaultCountSkills)

	query := repo.store.Rebind(queryGetAllSkills)

	err := repo.store.Select(&res, query)
	if err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}
	return res, nil
}

const checkExistsSkill = "SELECT count(*) from skills where name = ?;"

// CheckExists Errors:
//		postgresql_utilits.Exists
//		postgresql_utilits.NotFound
// 		app.GeneralError with Errors:
// 			postgresql_utilits.DefaultErrDB
func (repo *SkillRepository) CheckExists(skillName string) error {
	cnt := int64(0)
	query := repo.store.Rebind(checkExistsSkill)
	if err := repo.store.Get(&cnt, query, skillName); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return postgresql_utilits.NotFound
		}
		return postgresql_utilits.NewDBError(errors.Wrap(err, "CheckExists skill_repository"))
	}

	if cnt != 0 {
		return postgresql_utilits.Exists
	}
	return postgresql_utilits.NotFound
}
