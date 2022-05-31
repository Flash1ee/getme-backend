package entities_user

import (
	"database/sql"
	"time"
)

type User struct {
	ID           int64          `db:"id"`
	FirstName    sql.NullString `db:"first_name"`
	LastName     sql.NullString `db:"last_name"`
	Nickname     string         `db:"nickname"`
	TgTag        string         `db:"tg_tag"`
	About        sql.NullString `db:"about"`
	Avatar       sql.NullString `db:"avatar"`
	Email        sql.NullString `db:"email"`
	IsSearchable bool           `db:"is_searchable"`
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at"`
}

type UserWithSkill struct {
	User
	Skill sql.NullString `db:"skill_name"`
}

type UserWithSkills struct {
	User
	Skills []string `db:"skill_name"`
}

type UserWithOfferID struct {
	User
	OfferID int64 `db:"offer_id"`
}

type UserSkills struct {
	ID        int64  `db:"id"`
	UserId    int64  `db:"user_id"`
	SkillName string `db:"skill_name"`
}

func ToUsersSkills(userId int64, skills []string) []UserSkills {
	res := make([]UserSkills, 0, len(skills))
	for _, val := range skills {
		res = append(res, UserSkills{
			UserId:    userId,
			SkillName: val,
		})
	}
	return res
}
