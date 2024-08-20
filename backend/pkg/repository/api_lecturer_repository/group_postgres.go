package api_lecturer_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type LecturerGroupPostgres struct {
	db *sqlx.DB
}

func NewLecturerGroupPostgres(db *sqlx.DB) *LecturerGroupPostgres {
	return &LecturerGroupPostgres{db: db}
}

func (r *LecturerGroupPostgres) AddGroup(name string) error {
	query := fmt.Sprintf("INSERT INTO %s (name, is_archive) VALUES( $1 , false )", table_name.GroupsTable)
	_, err := r.db.Exec(query, name)
	return err
}

func (r *LecturerGroupPostgres) ChangeName(groupId int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", table_name.GroupsTable)
	_, err := r.db.Exec(query, name, groupId)
	return err
}

func (r *LecturerGroupPostgres) DeleteGroup(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", table_name.GroupsTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *LecturerGroupPostgres) GetAllStudentsFromGroup(id int) ([]model.Student, error) {
	var students []model.Student
	query := fmt.Sprintf("SELECT id, name, surname FROM %s INNER JOIN %s ON %s.id = %s.user_id WHERE group_id = $1 ORDER BY surname, name",
		table_name.UsersTable, table_name.StudentsTable, table_name.UsersTable, table_name.StudentsTable)
	err := r.db.Select(&students, query, id)
	if err != nil {
		return []model.Student{}, err
	}
	return students, nil
}

func (r *LecturerGroupPostgres) GetAllGroups() ([]model.Group, error) {
	var groups []model.Group
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id", table_name.GroupsTable)
	err := r.db.Select(&groups, query)
	if err != nil {
		return []model.Group{}, err
	}
	return groups, nil
}

func (r *LecturerGroupPostgres) GetGroupsDisciplines(groupId int) ([]model.Discipline, []model.Discipline, error) {
	var disciplines []model.Discipline
	query := fmt.Sprintf("SELECT id, name FROM %s INNER JOIN %s ON %s.discipline_id = %s.id WHERE group_id = $1 AND %s.is_archive = false ORDER BY %s.id",
		table_name.CurriculumTable, table_name.DisciplinesTable, table_name.CurriculumTable, table_name.DisciplinesTable, table_name.CurriculumTable, table_name.DisciplinesTable)
	err := r.db.Select(&disciplines, query, groupId)
	if err != nil {
		return nil, nil, err
	}
	var disciplinesEn []model.Discipline
	query = fmt.Sprintf("SELECT id, name_en as name FROM %s INNER JOIN %s ON %s.discipline_id = %s.id WHERE group_id = $1 AND %s.is_archive = false ORDER BY %s.id",
		table_name.CurriculumTable, table_name.DisciplinesTable, table_name.CurriculumTable, table_name.DisciplinesTable, table_name.CurriculumTable, table_name.DisciplinesTable)
	err = r.db.Select(&disciplinesEn, query, groupId)
	if err != nil {
		return nil, nil, err
	}
	return disciplines, disciplinesEn, nil
}

func (r *LecturerGroupPostgres) AddGroupInArchive(groupId int) error {
	query := fmt.Sprintf("UPDATE %s SET is_archive = true WHERE id = $1", table_name.GroupsTable)
	_, err := r.db.Exec(query, groupId)
	return err
}

func (r *LecturerGroupPostgres) DeleteGroupFromArchive(groupId int) error {
	query := fmt.Sprintf("UPDATE %s SET is_archive = false WHERE id = $1", table_name.GroupsTable)
	_, err := r.db.Exec(query, groupId)
	return err
}
