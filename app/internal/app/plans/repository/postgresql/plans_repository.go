package plans_repository_postgresql

import (
	"database/sql"

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

const queryGetByID = `SELECT * from plans where id = ?`

//GetByID with Errors:
//		postgresql_utilits.NotFound
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (repo *PlanRepository) GetByID(id int64) (*entities.Plan, error) {
	res := &entities.Plan{}
	query := repo.store.Rebind(queryGetByID)
	// @TODO проверить на sql.ErrNoRows
	if err := repo.store.Get(res, query, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}
	return res, nil

}

const queryGetPlanByTaskID = `
SELECT * from plans where id = ?;`

//GetPlanByTaskID with Errors:
// 		postgresql_utilits.NotFound
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (repo *PlanRepository) GetPlanByTaskID(taskID int64) (*entities.Plan, error) {
	res := &entities.Plan{}

	query := repo.store.Rebind(queryGetPlanByTaskID)
	if err := repo.store.Get(res, query, taskID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}
	return res, nil
}

const queryGetPlanWithMentorAndTasks = `
select p.id, p.name, p.about, p.is_active, p.progress, p.mentor_id, p.mentee_id,
       u.id, u.first_name, u.last_name, u.nickname, u.about, u.avatar,
       t.id, t.name, t.description, t.deadline, t.status from plans p
           join users u on p.mentor_id = u.id
            left join task t on p.id = t.plan_id
			left join status s on t.status = s.name where p.id = ? and u.id = ?;
`

//GetPlanWithMentorAndTasks with Errors:
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (repo *PlanRepository) GetPlanWithMentorAndTasks(mentorID int64, planID int64) ([]entities.PlanWithUserAndTask, error) {
	//res := make([]entities.PlanWithUserAndTask, 0)
	//
	//query := repo.store.Rebind(queryGetPlanWithMentorAndTasks)
	//
	//err := repo.store.Select(&res, query, planID, mentorID)
	//if err != nil {
	//	return nil, postgresql_utilits.NewDBError(err)
	//}
	res := make([]entities.PlanWithUserAndTask, 0)

	query := repo.store.Rebind(queryGetPlanWithMentorAndTasks)

	rows, err := repo.store.Queryx(query, planID, mentorID)
	if err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}
	for rows.Next() {
		tmp := entities.PlanWithUserAndTask{}
		err := rows.Scan(&tmp.Plan.ID, &tmp.Plan.Name, &tmp.Plan.About, &tmp.IsActive, &tmp.Progress, &tmp.MentorID, &tmp.MenteeID, &tmp.User.ID,
			&tmp.User.FirstName, &tmp.User.LastName, &tmp.User.Nickname, &tmp.User.About, &tmp.User.Avatar,
			&tmp.Task.ID, &tmp.Task.Name, &tmp.Task.Description, &tmp.Task.Deadline, &tmp.Task.Status)
		if err != nil {
			rows.Close()
			return nil, postgresql_utilits.NewDBError(err)
		}
		res = append(res, tmp)
	}
	//if len(*plans) == 0 {
	//	return nil, postgresql_utilits.NotFound
	//}

	return res, nil
}

const queryGetPlanWithMenteeAndTasks = `
select p.id, p.name, p.about, p.is_active, p.progress, p.mentor_id, p.mentee_id,
       u.id, u.first_name, u.last_name, u.nickname, u.about, u.avatar,
       t.id, t.name, t.description, t.deadline, t.status from plans p
           join users u on p.mentee_id = u.id
            left join task t on p.id = t.plan_id
			left join status s on t.status = s.name where p.id = ? and u.id = ?;
`

//GetPlanWithMenteeAndTasks with Errors:
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (repo *PlanRepository) GetPlanWithMenteeAndTasks(menteeID int64, planID int64) ([]entities.PlanWithUserAndTask, error) {
	res := make([]entities.PlanWithUserAndTask, 0)

	query := repo.store.Rebind(queryGetPlanWithMenteeAndTasks)

	err := repo.store.Select(&res, query, planID, menteeID)
	if err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}
	//if len(*plans) == 0 {
	//	return nil, postgresql_utilits.NotFound
	//}

	return res, nil
}
