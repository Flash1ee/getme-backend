package entities

import "time"

type Offer struct {
	ID        int64     `db:"id"`
	SkillName string    `db:"skill_name"`
	Status    bool      `db:"status"`
	MentorID  int64     `db:"mentor_id"`
	MenteeID  int64     `db:"mentee_id"`
	CreatedAt time.Time `db:"created_at"`
}
