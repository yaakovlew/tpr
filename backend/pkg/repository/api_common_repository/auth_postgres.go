package api_common_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, surname, email ,password) VALUES($1, $2, $3, $4) RETURNING id", table_name.UsersTable)
	tx, err := r.db.Begin()
	defer tx.Rollback()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	row := tx.QueryRow(query, user.Name, user.Surname, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	err = r.AddPost(id, user, tx)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return id, nil
}

func (r *AuthPostgres) CheckGroup(groupName string) (int, error) {
	var group model.Group
	query := fmt.Sprintf("SELECT * FROM %s VALUES WHERE name = $1", table_name.GroupsTable)
	err := r.db.Get(&group, query, groupName)
	if err != nil {
		return 0, err
	}
	return group.Id, nil
}

func (r *AuthPostgres) AddPost(id int, user model.User, tx *sql.Tx) error {
	switch user.Post {
	case "student":
		query := fmt.Sprintf("INSERT INTO %s (user_id, group_id) VALUES($1, $2)", table_name.StudentsTable)
		groupId, err := r.CheckGroup(user.GroupName)
		if err != nil {
			err := errors.New("this group is not found")
			return err
		}
		_, err = tx.Exec(query, id, groupId)
		if err != nil {
			return err
		} else {
			return nil
		}
	case "seminarian":
		query := fmt.Sprintf("INSERT INTO %s (user_id) VALUES($1)", table_name.SeminariansTable)
		_, err := tx.Exec(query, id)
		if err != nil {
			return err
		} else {
			return nil
		}
	default:
		err := errors.New("not found this post")
		return err
	}
}

func (r *AuthPostgres) GetUser(email string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE email = $1", table_name.UsersTable)
	err := r.db.Get(&user, query, email)
	if err != nil {
		return model.User{}, err
	}
	user.Post, err = r.GetPost(user.Id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

type seminarian struct {
	GroupId *int `db:"group_id"`
	UserId  int  `db:"user_id"`
}

type lecturer struct {
	DisciplineId *int `db:"discipline_id"`
	UserId       int  `db:"user_id"`
}

type student struct {
	UserId  int `db:"user_id"`
	GroupId int `db:"group_id"`
}

func (r *AuthPostgres) GetPost(id int) (string, error) {
	var findSeminarian seminarian
	var findLecturer lecturer
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", table_name.SeminariansTable)
	err := r.db.Get(&findSeminarian, query, id)
	if err != nil {
		query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", table_name.LecturersTable)
		err := r.db.Get(&findLecturer, query, id)
		if err != nil {
			var findStudent student
			query = fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", table_name.StudentsTable)
			err = r.db.Get(&findStudent, query, id)
			if err != nil {
				err = errors.New("not found user in DB")
				return "", err
			} else {
				return "student", nil
			}
		} else {
			return "lecturer", nil
		}
	} else {
		return "seminarian", nil
	}
}

func (r *AuthPostgres) GetUserMail(userId int) (string, error) {
	var mail string
	query := fmt.Sprintf("SELECT email FROM %s WHERE id = $1", table_name.UsersTable)
	err := r.db.Get(&mail, query, userId)
	if err != nil {
		return "", err
	}
	return mail, nil
}

func (r *AuthPostgres) ChangePassword(userId int, password string) error {
	query := fmt.Sprintf("UPDATE %s SET password = $1 WHERE id = $2", table_name.UsersTable)
	_, err := r.db.Exec(query, password, userId)
	return err
}
