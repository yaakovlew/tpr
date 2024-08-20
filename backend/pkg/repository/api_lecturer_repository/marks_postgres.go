package api_lecturer_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type LecturerMarksPostgres struct {
	db *sqlx.DB
}

func NewLecturerMarksPostgres(db *sqlx.DB) *LecturerMarksPostgres {
	return &LecturerMarksPostgres{db: db}
}

func (r *LecturerMarksPostgres) ChangeTestMark(userId, testId, mark int) error {
	query := fmt.Sprintf("UPDATE %s SET mark = $1 WHERE user_id = $2 AND test_id = $3", table_name.TestsMarkTable)
	_, err := r.db.Exec(query, mark, userId, testId)
	return err
}

func (r *LecturerMarksPostgres) ChangeLaboratoryMark(userId, laboratoryId, mark int) error {
	query := fmt.Sprintf("UPDATE %s SET mark = $1 WHERE user_id = $2 AND laboratory_id = $3", table_name.LaboratoryMarkTable)
	_, err := r.db.Exec(query, mark, userId, laboratoryId)
	return err
}

func (r *LecturerMarksPostgres) GetTestMarksFromGroup(groupId, testId int) ([]model.GroupTestMarks, error) {
	var marks []model.GroupTestMarks
	query := fmt.Sprintf("SELECT %s.user_id, %s.name AS user_name, %s.surname AS user_surname, %s.mark FROM %s INNER JOIN %s ON %s.user_id = %s.user_id INNER JOIN %s ON %s.user_id = %s.id WHERE group_id = $1 AND test_id = $2 ORDER BY %s.surname, %s.name",
		table_name.TestsMarkTable, table_name.UsersTable, table_name.UsersTable,
		table_name.TestsMarkTable, table_name.TestsMarkTable, table_name.StudentsTable, table_name.TestsMarkTable,
		table_name.StudentsTable, table_name.UsersTable, table_name.TestsMarkTable,
		table_name.UsersTable, table_name.UsersTable, table_name.UsersTable)
	err := r.db.Select(&marks, query, groupId, testId)
	if err != nil {
		return []model.GroupTestMarks{}, err
	}
	return marks, nil
}

func (r *LecturerMarksPostgres) GetLaboratoryMarksFromGroup(groupId, laboratoryId int) ([]model.GroupLaboratoryMarks, error) {
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

func (r *LecturerMarksPostgres) GiveExamMark(userId, disciplineId, mark int) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, discipline_id, mark) VALUES ($1, $2, $3)", table_name.ExamTable)
	_, err := r.db.Exec(query, userId, disciplineId, mark)
	return err
}

func (r *LecturerMarksPostgres) ChangeExamMark(userId, disciplineId, mark int) error {
	query := fmt.Sprintf("UPDATE %s SET mark = $1 WHERE user_id = $2 AND discipline_id = $3", table_name.ExamTable)
	_, err := r.db.Exec(query, mark, userId, disciplineId)
	return err
}

func (r *LecturerMarksPostgres) GetAllMarksForExam(groupId, disciplineId int) ([]model.ExamMark, error) {
	var exam []model.ExamMark
	query := fmt.Sprintf(`SELECT %s.id AS user_id, %s.name AS user_name, %s.surname AS user_surname, %s.mark AS mark
					FROM %s 
					INNER JOIN %s
					ON %s.user_id = %s.user_id
					INNER JOIN %s
					ON %s.group_id = %s.group_id
					INNER JOIN %s
					ON %s.user_id = %s.id
					WHERE %s.group_id = $1 AND %s.discipline_id = $2 ORDER BY %s.surname, %s.name`,
		table_name.UsersTable, table_name.UsersTable, table_name.UsersTable, table_name.ExamTable,
		table_name.ExamTable, table_name.StudentsTable, table_name.ExamTable,
		table_name.StudentsTable, table_name.CurriculumTable, table_name.StudentsTable, table_name.CurriculumTable,
		table_name.UsersTable, table_name.StudentsTable,
		table_name.UsersTable, table_name.StudentsTable, table_name.ExamTable,
		table_name.UsersTable, table_name.UsersTable)
	err := r.db.Select(&exam, query, groupId, disciplineId)
	if err != nil {
		return nil, err
	}
	return exam, nil
}

func (r *LecturerMarksPostgres) CheckExistMark(userId, disciplineId int) error {
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

func (r *LecturerMarksPostgres) MaxExamMark(disciplineId int) int {
	var maxMark int
	query := fmt.Sprintf("SELECT exam_mark FROM %s WHERE id = $1", table_name.DisciplinesTable)
	if err := r.db.Get(&maxMark, query, disciplineId); err != nil {
		return 0
	}
	return maxMark
}
