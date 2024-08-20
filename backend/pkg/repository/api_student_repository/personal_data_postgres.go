package api_student_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type StudentPersonalDataPostgres struct {
	db *sqlx.DB
}

func NewStudentPersonalDataPostgres(db *sqlx.DB) *StudentPersonalDataPostgres {
	return &StudentPersonalDataPostgres{db: db}
}

func (r *StudentPersonalDataPostgres) GetPersonalData(id int) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", table_name.UsersTable)
	err := r.db.Get(&user, query, id)
	if err != nil {
		return model.User{}, err
	}
	query = fmt.Sprintf("SELECT %s.name FROM %s INNER JOIN %s ON %s.group_id = %s.id WHERE user_id = $1",
		table_name.GroupsTable, table_name.StudentsTable, table_name.GroupsTable, table_name.StudentsTable, table_name.GroupsTable)
	err = r.db.Get(&user.GroupName, query, id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *StudentPersonalDataPostgres) UpdateName(id int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1 WHERE id = $2", table_name.UsersTable)
	_, err := r.db.Exec(query, name, id)
	return err
}

func (r *StudentPersonalDataPostgres) UpdateSurname(id int, surname string) error {
	query := fmt.Sprintf("UPDATE %s SET surname=$1 WHERE id = $2", table_name.UsersTable)
	_, err := r.db.Exec(query, surname, id)
	return err
}
