package repository_postgresql

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"

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

const queryGetUserByNickname = `
SELECT id, first_name, last_name, nickname, about, avatar, tg_tag,
       is_searchable, created_at, updated_at from users where nickname = ?;`

func (repo *UserRepository) FindByNickname(nickname string) (*entities_user.User, error) {
	query := repo.store.Rebind(queryGetUserByNickname)

	user := &entities_user.User{}
	if err := repo.store.QueryRow(query, nickname).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Nickname,
		&user.About, &user.Avatar, &user.TgTag, &user.IsSearchable, &user.CreatedAt, &user.UpdatedAt); err != nil {
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

func (repo *UserRepository) CreateFilledUser(data *entities_user.User) (int64, error) {
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

const queryFindByIDWithSkill = `SELECT users.id, first_name, last_name, about, tg_tag, avatar, is_searchable, skill_name from users 
    left join users_skills us on users.id = us.user_id 
    left join skills s on us.skill_name = s.name 
	where users.id = ?
`

//	FindByIDWithSkill with Errors:
//		postgresql_utilits.NotFound
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (repo *UserRepository) FindByIDWithSkill(id int64) (*[]entities_user.UserWithSkill, error) {
	query := repo.store.Rebind(queryFindByIDWithSkill)
	user := &[]entities_user.UserWithSkill{}

	err := repo.store.Select(user, query, id)
	if err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}
	if len(*user) == 0 {
		return nil, postgresql_utilits.NotFound
	}

	return user, nil
}

const queryFindByID = `SELECT users.id, first_name, last_name, about, tg_tag, avatar, is_searchable from users 
	where users.id = ?
`

//	FindByID with Errors:
//		postgresql_utilits.NotFound
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (repo *UserRepository) FindByID(id int64) (*entities_user.User, error) {
	query := repo.store.Rebind(queryFindByID)
	user := &entities_user.User{}

	err := repo.store.Get(user, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}

	return user, nil
}

const queryIsMentor = ` and is_searchable = true;`

//	FindMentorByID with Errors:
//		postgresql_utilits.NotFound
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (repo *UserRepository) FindMentorByID(id int64) (*[]entities_user.UserWithSkill, error) {
	query := repo.store.Rebind(queryFindByID + queryIsMentor)

	usersWithSkills := &[]entities_user.UserWithSkill{}
	err := repo.store.Select(usersWithSkills, query, id)
	if err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}
	if len(*usersWithSkills) == 0 {
		return nil, postgresql_utilits.NotFound
	}

	return usersWithSkills, nil
}

const queryUpdateUser = `update users set
    first_name = COALESCE(NULLIF(TRIM(?), ''), first_name),
    last_name = COALESCE(NULLIF(TRIM(?), ''), last_name),
    about = COALESCE(NULLIF(TRIM(?), ''), about),
    tg_tag = COALESCE(NULLIF(TRIM(?), ''), tg_tag)
WHERE id = ? returning first_name, last_name, nickname, about, tg_tag, avatar, is_searchable;`

const queryUpdateUserDeleteOldSkills = `DELETE FROM users_skills WHERE user_id = ?`

const queryUpdateUserAddSkills = `
INSERT INTO users_skills(user_id, skill_name) VALUES (:user_id, :skill_name)`

func (repo *UserRepository) UpdateUser(user *entities_user.UserWithSkills) (*entities_user.UserWithSkills, error) {
	query := repo.store.Rebind(queryUpdateUser)
	queryDelete := repo.store.Rebind(queryUpdateUserDeleteOldSkills)
	userFromDB := &entities_user.UserWithSkills{}

	tx, err := repo.store.Beginx()
	if err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}

	err = tx.QueryRowx(query, user.FirstName.String, user.LastName.String, user.About.String, user.TgTag.String, user.ID.Int64).
		Scan(&userFromDB.FirstName, &userFromDB.LastName, &userFromDB.Nickname, &userFromDB.About, &userFromDB.TgTag, &userFromDB.Avatar, &userFromDB.IsSearchable)
	if err != nil {
		_ = tx.Rollback()
		return nil, postgresql_utilits.NewDBError(err)
	}

	if len(user.Skills) != 0 {
		if _, err = tx.Exec(queryDelete, user.ID); err != nil {
			_ = tx.Rollback()
			return nil, postgresql_utilits.NewDBError(err)
		}

		args := entities_user.ToUsersSkills(user.ID.Int64, user.Skills)
		if _, err = tx.NamedExec(queryUpdateUserAddSkills, args); err != nil {
			_ = tx.Rollback()
			return nil, postgresql_utilits.NewDBError(err)
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}

	return userFromDB, nil
}

const queryGetUsersBySkillsAll = `SELECT users.id, first_name, last_name, nickname, about, avatar, is_searchable, skill_name from users
JOIN users_skills as us on us.user_id = users.id `

const queryGetUsersBySkill = queryGetUsersBySkillsAll + `where us.skill_name IN (?)`

const queryConditionIsMentor = ` and is_searchable = true`

//	GetUsersBySkills with Errors:
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (repo *UserRepository) GetUsersBySkills(data []skill_entities.Skill) ([]entities_user.UserWithSkill, error) {
	skills := getSkillNameFromSkills(data)
	usersWithSkills := &[]entities_user.UserWithSkill{}
	if len(skills) == 0 {
		query := repo.store.Rebind(queryGetUsersBySkillsAll + queryConditionIsMentor)
		if err := repo.store.Select(usersWithSkills, query); err != nil {
			return nil, postgresql_utilits.NewDBError(err)
		}
		return *usersWithSkills, nil
	}

	query, args, err := sqlx.In(queryGetUsersBySkill+queryConditionIsMentor, skills)
	if err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}
	query = repo.store.Rebind(query)
	var queryArgs []interface{}
	queryArgs = append(queryArgs, args...)

	if err = repo.store.Select(usersWithSkills, query, queryArgs...); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NewDBError(err)
		}
	}
	return *usersWithSkills, nil

}

const queryGetMenteeByOffers = `
SELECT offers.id as offer_id, users.id as id, first_name, last_name, tg_tag, about, avatar, is_searchable from users join getme_db.public.offers
    on users.id = offers.mentee_id and offers.mentor_id = ? and offers.status = true;`

//	GetMenteeByMentorWithOfferID with Errors:
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (r *UserRepository) GetMenteeByMentorWithOfferID(mentorID int64) ([]entities_user.UserWithOfferID, error) {
	users := &[]entities_user.UserWithOfferID{}

	query := r.store.Rebind(queryGetMenteeByOffers)

	if err := r.store.Select(users, query, mentorID); err != nil {
		return nil, postgresql_utilits.NewDBError(errors.Wrap(err, fmt.Sprintf("UserRepository: GetMenteeByMentorWithOfferID(%v)", mentorID)))
	}

	return *users, nil
}

const queryUpdateUserStatus = `
update users set is_searchable = not is_searchable where users.id = ? returning is_searchable;
`

//UpdateMentorStatus with Errors:
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (r *UserRepository) UpdateMentorStatus(mentorID int64) (bool, error) {
	query := r.store.Rebind(queryUpdateUserStatus)
	isMentor := false
	if err := r.store.QueryRow(query, mentorID).Scan(&isMentor); err != nil {
		return false, postgresql_utilits.NewDBError(errors.Wrapf(err,
			"UserRepository: UpdateMentorStatus(%v) can not update user status", mentorID))
	}
	return isMentor, nil
}

const querySetUserStatus = `
update users set is_searchable = ? where users.id = ? returning is_searchable;
`

//SetMentorStatus with Errors:
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (r *UserRepository) SetMentorStatus(mentorID int64, status bool) (bool, error) {
	query := r.store.Rebind(querySetUserStatus)
	isMentor := false
	if err := r.store.QueryRow(query, status, mentorID).Scan(&isMentor); err != nil {
		return false, postgresql_utilits.NewDBError(errors.Wrapf(err,
			"UserRepository: SetMentorStatus(%v, %v) can not set user status", mentorID, status))
	}
	return isMentor, nil
}
