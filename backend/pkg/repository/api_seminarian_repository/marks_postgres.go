package api_seminarian_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SeminarianMarksPostgres struct {
	db *sqlx.DB
}

func NewSeminarianMarksPostgres(db *sqlx.DB) *SeminarianMarksPostgres {
	return &SeminarianMarksPostgres{db: db}
}

func (r *SeminarianMarksPostgres) GetTestMarksFromGroup(groupId, testId int) ([]model.GroupTestMarks, error) {
	var marks []model.GroupTestMarks
	query := fmt.Sprintf("SELECT %s.user_id, %s.name AS user_name, %s.surname AS user_surname, %s.mark FROM %s INNER JOIN %s ON %s.user_id = %s.user_id INNER JOIN %s ON %s.user_id = %s.id WHERE group_id = $1 AND test_id = $2 ORDER BY %s.surname, %s.name",
		table_name.TestsMarkTable, table_name.UsersTable, table_name.UsersTable, table_name.TestsMarkTable, table_name.TestsMarkTable, table_name.StudentsTable, table_name.TestsMarkTable,
		table_name.StudentsTable, table_name.UsersTable, table_name.TestsMarkTable, table_name.UsersTable, table_name.UsersTable, table_name.UsersTable)
	err := r.db.Select(&marks, query, groupId, testId)
	if err != nil {
		return []model.GroupTestMarks{}, err
	}
	return marks, nil
}

func (r *SeminarianMarksPostgres) GetLaboratoryMarksFromGroup(groupId, laboratoryId int) ([]model.GroupLaboratoryMarks, error) {
	var marks []model.GroupLaboratoryMarks
	query := fmt.Sprintf("SELECT %s.user_id, %s.name AS user_name, %s.surname AS user_surname, %s.mark FROM %s INNER JOIN %s ON %s.user_id = %s.user_id INNER JOIN %s ON %s.user_id = %s.id WHERE group_id = $1 AND laboratory_id = $2 ORDER BY %s.surname, %s.name",
		table_name.LaboratoryMarkTable, table_name.UsersTable, table_name.UsersTable, table_name.LaboratoryMarkTable, table_name.LaboratoryMarkTable, table_name.StudentsTable, table_name.LaboratoryMarkTable,
		table_name.StudentsTable, table_name.UsersTable, table_name.LaboratoryMarkTable, table_name.UsersTable, table_name.UsersTable, table_name.UsersTable)
	err := r.db.Select(&marks, query, groupId, laboratoryId)
	if err != nil {
		return []model.GroupLaboratoryMarks{}, err
	}
	return marks, nil
}

func (r *SeminarianMarksPostgres) CheckAccessForGroup(seminarianId, groupId, disciplineId int) error {
	query := fmt.Sprintf("SELECT group_id FROM %s WHERE user_id = $1 AND group_id = $2 AND discipline_id = $3",
		table_name.SeminariansGroupsTable)
	var group int
	err := r.db.Get(&group, query, seminarianId, groupId, disciplineId)
	if err != nil {
		errorAccess := errors.New("access denied to group")
		return errorAccess
	}
	return nil
}

func (r *SeminarianMarksPostgres) CheckAccessForTest(seminarianId, groupId, testId int) error {
	var disciplineId []int
	errAccess := errors.New("access denied to test")
	query := fmt.Sprintf("SELECT discipline_id FROM %s INNER JOIN %s ON %s.section_id = %s.id INNER JOIN %s ON %s.discipline_id = %s.id WHERE test_id = $1",
		table_name.SectionTests, table_name.SectionTable, table_name.SectionTests, table_name.SectionTable, table_name.DisciplinesTable, table_name.SectionTable, table_name.DisciplinesTable)
	if err := r.db.Select(&disciplineId, query, testId); err != nil {
		return errAccess
	}
	for _, value := range disciplineId {
		if err := r.CheckAccessForGroup(seminarianId, groupId, value); err == nil {
			return nil
		}
	}
	return errAccess
}

func (r *SeminarianMarksPostgres) CheckAccessForLaboratory(seminarianId, groupId, labId int) error {
	var disciplineId []int
	errAccess := errors.New("access denied to laboratory work")
	query := fmt.Sprintf("SELECT discipline_id FROM %s INNER JOIN %s ON %s.section_id = %s.id INNER JOIN %s ON %s.discipline_id = %s.id WHERE laboratory_id = $1",
		table_name.SectionLabs, table_name.SectionTable, table_name.SectionLabs, table_name.SectionTable, table_name.DisciplinesTable, table_name.SectionTable, table_name.DisciplinesTable)
	if err := r.db.Select(&disciplineId, query, labId); err != nil {
		return errAccess
	}
	for _, value := range disciplineId {
		if err := r.CheckAccessForGroup(seminarianId, groupId, value); err == nil {
			return nil
		}
	}
	return errAccess
}

func (r *SeminarianMarksPostgres) GiveExamMark(userId, disciplineId, mark int) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, discipline_id, mark) VALUES ($1, $2, $3)", table_name.ExamTable)
	_, err := r.db.Exec(query, userId, disciplineId, mark)
	return err
}

func (r *SeminarianMarksPostgres) ChangeExamMark(userId, disciplineId, mark int) error {
	query := fmt.Sprintf("UPDATE %s SET mark = $1 WHERE user_id = $1 AND discipline_id = $2", table_name.ExamTable)
	_, err := r.db.Exec(query, mark, userId, disciplineId)
	return err
}

func (r *SeminarianMarksPostgres) GetAllMarksForExam(groupId, disciplineId int) ([]model.ExamMark, error) {
	var exam []model.ExamMark
	query := fmt.Sprintf("SELECT %s.id AS user_id, %s.name AS user_name, %s.surname AS user_surname, %s.mark AS mark  FROM %s INNER JOIN %s ON %s.user_id = %s.user_id INNER JOIN %s ON %s.group_id = %s.group_id INNER JOIN %s ON %s.user_id = %s.id WHERE %s.group_id = $1 AND %s.discipline_id = $2",
		table_name.UsersTable, table_name.UsersTable, table_name.UsersTable, table_name.ExamTable,
		table_name.ExamTable, table_name.StudentsTable, table_name.ExamTable,
		table_name.StudentsTable, table_name.CurriculumTable, table_name.StudentsTable, table_name.CurriculumTable,
		table_name.UsersTable, table_name.StudentsTable, table_name.UsersTable, table_name.StudentsTable, table_name.CurriculumTable)
	err := r.db.Select(&exam, query, groupId, disciplineId)
	if err != nil {
		return nil, err
	}
	return exam, nil
}

func (r *SeminarianMarksPostgres) CheckExistMark(userId, disciplineId int) error {
	var user int
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE user_id =$1 AND discipline_id = $2", table_name.ExamTable)
	err := r.db.Get(&user, query, userId, disciplineId)
	if err != nil {
		return err
	}
	if user != userId {
		return err
	}
	return nil
}

func (r *SeminarianMarksPostgres) CheckAccessForStudent(seminarianId, userId, disciplineId int) error {
	if err := r.CheckAccessForDiscipline(seminarianId, disciplineId); err != nil {
		return err
	}
	var groupId int
	errAccess := errors.New("access denied to student")
	query := fmt.Sprintf("SELECT %s.group_id FROM %s INNER JOIN %s ON %s.group_id = %s.group_id WHERE %s.user_id = $1 AND discipline_id = $2",
		table_name.StudentsTable, table_name.StudentsTable, table_name.CurriculumTable, table_name.StudentsTable, table_name.CurriculumTable, table_name.StudentsTable)
	err := r.db.Get(&groupId, query, userId, disciplineId)
	if err != nil {
		return errAccess
	}
	if err := r.CheckAccessForGroup(seminarianId, groupId, disciplineId); err != nil {
		return errAccess
	}
	return nil
}

func (r *SeminarianMarksPostgres) CheckAccessForDiscipline(seminarianId, disciplineId int) error {
	var groups []int
	errAccess := errors.New("access denied to discipline")
	query := fmt.Sprintf("SELECT group_id FROM %s WHERE user_id = $1 AND discipline_id = $2", table_name.SeminariansGroupsTable)
	err := r.db.Select(&groups, query, seminarianId, disciplineId)
	if err != nil {
		return errAccess
	}
	if len(groups) == 0 {
		return errAccess
	}
	return nil
}

func (r *SeminarianMarksPostgres) MaxExamMark(disciplineId int) int {
	var maxMark int
	query := fmt.Sprintf("SELECT exam_mark FROM %s WHERE id = $1", table_name.DisciplinesTable)
	if err := r.db.Get(&maxMark, query, disciplineId); err != nil {
		return 0
	}
	return maxMark
}
