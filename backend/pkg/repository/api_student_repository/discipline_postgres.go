package api_student_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type StudentDisciplinePostgres struct {
	db *sqlx.DB
}

func NewStudentDisciplinePostgres(db *sqlx.DB) *StudentDisciplinePostgres {
	return &StudentDisciplinePostgres{db: db}
}

func (r *StudentDisciplinePostgres) GetAllUserDiscipline(id int) ([]model.Discipline, []model.Discipline, error) {
	var groupId int
	query := fmt.Sprintf("SELECT group_id FROM %s WHERE user_id = $1", table_name.StudentsTable)
	err := r.db.Get(&groupId, query, id)
	if err != nil {
		return nil, nil, errors.New("fatal error users group not found")
	}
	var disciplines []model.Discipline
	query = fmt.Sprintf("SELECT %s.id, %s.name FROM %s INNER JOIN %s ON %s.id = %s.discipline_id WHERE group_id = $1 AND is_archive = false ORDER BY %s.id",
		table_name.DisciplinesTable, table_name.DisciplinesTable, table_name.DisciplinesTable,
		table_name.CurriculumTable, table_name.DisciplinesTable, table_name.CurriculumTable, table_name.DisciplinesTable)
	err = r.db.Select(&disciplines, query, groupId)
	if err != nil {
		return nil, nil, err //errors.New("fatal error users disciplines not found")
	}
	var disciplinesEn []model.Discipline
	query = fmt.Sprintf("SELECT %s.id, %s.name_en as name FROM %s INNER JOIN %s ON %s.id = %s.discipline_id WHERE group_id = $1 AND is_archive = false ORDER BY %s.id",
		table_name.DisciplinesTable, table_name.DisciplinesTable, table_name.DisciplinesTable,
		table_name.CurriculumTable, table_name.DisciplinesTable, table_name.CurriculumTable, table_name.DisciplinesTable)
	err = r.db.Select(&disciplinesEn, query, groupId)
	if err != nil {
		return nil, nil, err //errors.New("fatal error users disciplines not found")
	}
	return disciplines, disciplinesEn, nil
}

func (r *StudentDisciplinePostgres) GetDisciplineSections(disciplineId int) ([]model.Section, []model.Section, error) {
	var sections []model.Section
	query := fmt.Sprintf("SELECT id, name FROM %s WHERE discipline_id = $1 ORDER BY id", table_name.SectionTable)
	err := r.db.Select(&sections, query, disciplineId)
	if err != nil {
		return nil, nil, err
	}
	var sectionsEn []model.Section
	query = fmt.Sprintf("SELECT id, name_en as name FROM %s WHERE discipline_id = $1 ORDER BY id", table_name.SectionTable)
	err = r.db.Select(&sectionsEn, query, disciplineId)
	if err != nil {
		return nil, nil, err
	}
	return sections, sectionsEn, err
}

func (r *StudentDisciplinePostgres) CheckAccessForDiscipline(studentId, disciplineId int) error {
	var group int
	errAccess := errors.New("access denied to discipline")
	query := fmt.Sprintf("SELECT group_id FROM %s INNER JOIN %s ON %s.group_id = %s.group_id WHERE user_id = $1 AND discipline_id = $2 AND is_archive = false",
		table_name.StudentsTable, table_name.CurriculumTable, table_name.StudentsTable, table_name.CurriculumTable)
	err := r.db.Get(&group, query, studentId, disciplineId)
	if err != nil {
		return errAccess
	}
	if group == 0 {
		return errAccess
	}
	return nil
}
