package api_lecturer_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type LecturerStudentPostgres struct {
	db *sqlx.DB
}

func NewLecturerStudentPostgres(db *sqlx.DB) *LecturerStudentPostgres {
	return &LecturerStudentPostgres{db: db}
}

func (r *LecturerStudentPostgres) ChangeGroupForStudent(userId, groupId int) error {
	query := fmt.Sprintf("UPDATE %s SET group_id = $1 WHERE user_id = $2", table_name.StudentsTable)
	_, err := r.db.Exec(query, groupId, userId)
	return err
}

func (r *LecturerStudentPostgres) GetAllStudents() ([]model.StudentWithGroup, error) {
	var students []model.StudentWithGroup
	query := fmt.Sprintf("SELECT %s.id, %s.name, %s.surname, %s.name AS group_name FROM %s INNER JOIN %s ON %s.user_id = %s.id INNER JOIN %s ON %s.id = %s.group_id ORDER BY %s.id",
		table_name.UsersTable, table_name.UsersTable, table_name.UsersTable, table_name.GroupsTable, table_name.StudentsTable, table_name.UsersTable, table_name.StudentsTable,
		table_name.UsersTable, table_name.GroupsTable, table_name.GroupsTable, table_name.StudentsTable, table_name.UsersTable)
	err := r.db.Select(&students, query)
	if err != nil {
		return []model.StudentWithGroup{}, err
	}
	return students, nil
}

func (r *LecturerStudentPostgres) DeleteUser(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", table_name.StudentsTable)
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	if _, err := tx.Exec(query, userId); err != nil {
		tx.Rollback()
		return err
	}
	query = fmt.Sprintf("DELETE FROM %s WHERE id = $1", table_name.UsersTable)
	if _, err := tx.Exec(query, userId); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
