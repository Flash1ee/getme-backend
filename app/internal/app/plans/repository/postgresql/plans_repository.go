package plans_repository_postgresql

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"getme-backend/internal/app/plans/entities"
	skill_entities "getme-backend/internal/app/skill/entities"
	postgresql_utilits "getme-backend/internal/pkg/utilits/postgresql"
)

type PlanRepository struct {
	store *sqlx.DB
}

func NewPlanRepository(store *sqlx.DB) *PlanRepository {
	return &PlanRepository{
		store: store,
	}
}

const queryCreatePlan = `
INSERT INTO plans (name, about, mentor_id, mentee_id) VALUES 
(?, ?, ?, ?) RETURNING id
`
const queryCreatePlanSkills = `
INSERT INTO plans_skills(plan_id, skill_name) VALUES (:plan_id, :skill_name)`

const querySetOfferAccepted = `
UPDATE offers set status = false where id = ?;`

//Create with Errors:
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (repo *PlanRepository) Create(offerID int64, skills []skill_entities.Skill, plan entities.Plan) (*entities.Plan, error) {
	queryPlan := repo.store.Rebind(queryCreatePlan)

	tx, err := repo.store.Beginx()
	if err != nil {
		return nil, postgresql_utilits.NewDBError(errors.Wrapf(err, "PlanRepository: Create() can not create transaction"))
	}
	if err = tx.QueryRow(queryPlan, plan.Name, plan.About, plan.MentorID, plan.MenteeID).Scan(&plan.ID); err != nil {
		_ = tx.Rollback()
		return nil, postgresql_utilits.NewDBError(errors.Wrapf(err, "PlanRepository: Create() can not insert plan"))
	}

	args := entities.ToPlansSkills(plan.ID, skills)
	if _, err := tx.NamedExec(queryCreatePlanSkills, args); err != nil {
		_ = tx.Rollback()
		return nil, postgresql_utilits.NewDBError(errors.Wrapf(err, "PlanRepository: Create() can not insert to plans_skills"))
	}

	queryOffer := repo.store.Rebind(querySetOfferAccepted)
	if _, err := tx.Exec(queryOffer, offerID); err != nil {
		_ = tx.Rollback()
		return nil, postgresql_utilits.NewDBError(errors.Wrapf(err, "PlanRepository: Create() can not set offer status = false"))
	}

	if err = tx.Commit(); err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}

	return &plan, nil
}

const queryGetPlansByMentor = `SELECT plans.id, plans.name, about, is_active, progress, mentor_id, mentee_id, skill_name
from plans
         left join plans_skills ps on plans.id = ps.plan_id
         left join skills s on ps.skill_name = s.name
where plans.mentor_id = ?`

//GetByMentor with Errors:
//		postgresql_utilits.NotFound
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (repo *PlanRepository) GetByMentor(mentorID int64) ([]entities.PlanWithSkill, error) {
	query := repo.store.Rebind(queryGetPlansByMentor)
	plans := &[]entities.PlanWithSkill{}

	err := repo.store.Select(plans, query, mentorID)
	if err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}
	//if len(*plans) == 0 {
	//	return nil, postgresql_utilits.NotFound
	//}

	return *plans, nil
}

const queryGetPlansByMentee = `SELECT plans.id, plans.name, about, is_active, progress, mentor_id, mentee_id, skill_name
from plans
         left join plans_skills ps on plans.id = ps.plan_id
         left join skills s on ps.skill_name = s.name
where plans.mentee_id = ?`

//GetByMentee with Errors:
//		postgresql_utilits.NotFound
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (repo *PlanRepository) GetByMentee(menteeID int64) ([]entities.PlanWithSkill, error) {
	query := repo.store.Rebind(queryGetPlansByMentee)
	plans := &[]entities.PlanWithSkill{}

	err := repo.store.Select(plans, query, menteeID)
	if err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}
	//if len(*plans) == 0 {
	//	return nil, postgresql_utilits.NotFound
	//}

	return *plans, nil
}
