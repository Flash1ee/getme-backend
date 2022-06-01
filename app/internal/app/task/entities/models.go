package entities

import (
	"database/sql"
	"time"
)

type Task struct {
	ID          sql.NullInt64  `json:"id"`
	Name        sql.NullString `json:"name"`
	Description sql.NullString `json:"description"`
	Deadline    time.Time      `json:"deadline"`
	Status      sql.NullString `json:"status"`
	PlanID      sql.NullInt64  `json:"plan_id"`
	CreatedAt   time.Time      `json:"created_at"`
}
