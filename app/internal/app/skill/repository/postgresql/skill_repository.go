package repository_postgresql

import (
	"github.com/jmoiron/sqlx"

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
// 			repository.DefaultErrDB
func (repo *SkillRepository) GetAllSkills() ([]entities.Skill, error) {
	res := make([]entities.Skill, 0, defaultCountSkills)

	query := repo.store.Rebind(queryGetAllSkills)

	err := repo.store.Select(&res, query)
	if err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}
	return res, nil
}
