package repository_postgresql

import (
	"database/sql"
	"testing"
	"time"

	skill_entities "getme-backend/internal/app/skill/entities"
	"getme-backend/internal/app/user/entities"
)

func getRepositoryData(t *testing.T) *entities_user.User {
	t.Helper()

	return &entities_user.User{
		ID: 1,
		FirstName: sql.NullString{
			String: "Vasiliy",
		},
		LastName: sql.NullString{
			String: "Alexeev",
		},
		Nickname: "vasax",
		About: sql.NullString{
			String: "this is some information",
		},
		Avatar: sql.NullString{
			String: "/img/1.png",
		},
		Email: sql.NullString{
			String: "vasyugan@gmai.com",
		},
		IsSearchable: false,
		CreatedAt:    time.Now().Add(-3600),
		UpdatedAt:    time.Now(),
	}

}
func getSkillNameFromSkills(data []skill_entities.Skill) []string {
	res := make([]string, 0, len(data))
	for _, val := range data {
		res = append(res, val.Name)
	}
	return res
}
