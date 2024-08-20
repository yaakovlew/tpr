package api_seminarian_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SeminarianGroupPostgres struct {
	db *sqlx.DB
}

func NewSeminarianGroupPostgres(db *sqlx.DB) *SeminarianGroupPostgres {
	return &SeminarianGroupPostgres{db: db}
}

func (r *SeminarianGroupPostgres) GetOwnGroup(userId int, disciplineId int) ([]model.Group, error) {
	var groups []model.Group
	query := fmt.Sprintf("SELECT %s.id, %s.name FROM %s INNER JOIN %s ON %s.id = %s.group_id INNER JOIN %s ON %s.id = %s.group_id WHERE user_id = $1 AND (%s.discipline_id = $2 AND %s.is_archive = false) AND %s.is_archive = false ORDER BY id",
		table_name.GroupsTable, table_name.GroupsTable, table_name.SeminariansGroupsTable, table_name.GroupsTable, table_name.GroupsTable, table_name.SeminariansGroupsTable,
		table_name.CurriculumTable, table_name.GroupsTable, table_name.CurriculumTable, table_name.SeminariansGroupsTable, table_name.GroupsTable, table_name.CurriculumTable)
	err := r.db.Select(&groups, query, userId, disciplineId)
	if err != nil {
		return []model.Group{}, err
	}
	return groups, nil
}

func (r *SeminarianGroupPostgres) GetAllStudentsFromGroup(id int) ([]model.Student, error) {
	var students []model.Student
	query := fmt.Sprintf("SELECT id, name FROM %s INNER JOIN %s ON %s.id = %s.user_id WHERE group_id = $1 ORDER BY %s.surname, %s.name ",
		table_name.UsersTable, table_name.StudentsTable, table_name.UsersTable, table_name.StudentsTable, table_name.UsersTable, table_name.UsersTable)
	err := r.db.Select(&students, query, id)
	if err != nil {
		return []model.Student{}, err
	}
	return students, nil
}

func (r *SeminarianGroupPostgres) CheckAccessForGroup(seminarianId, groupId int) error {
	var groups []int
	errAccess := errors.New("access denied to group")
	query := fmt.Sprintf("SELECT %s.group_id FROM %s INNER JOIN %s ON %s.group_id = %s.group_id WHERE %s.user_id = $1 AND %s.group_id = $2 AND %s.is_archive = false",
		table_name.SeminariansGroupsTable, table_name.SeminariansGroupsTable, table_name.CurriculumTable, table_name.SeminariansGroupsTable, table_name.CurriculumTable,
		table_name.SeminariansGroupsTable, table_name.CurriculumTable, table_name.CurriculumTable)
	err := r.db.Select(&groups, query, seminarianId, groupId)
	if err != nil {
		return errAccess
	}
	if len(groups) == 0 {
		return errAccess
	}
	return nil
}
