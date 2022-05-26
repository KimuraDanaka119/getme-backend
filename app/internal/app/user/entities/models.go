package entities

import (
	"database/sql"
	"time"
)

type User struct {
	ID           int64          `db:"id"`
	FirstName    sql.NullString `db:"first_name"`
	LastName     sql.NullString `db:"last_name"`
	Nickname     string         `db:"nickname"`
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
