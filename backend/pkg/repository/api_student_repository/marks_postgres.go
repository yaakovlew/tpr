package api_student_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type StudentsMarksPostgres struct {
	db *sqlx.DB
}

func NewStudentMarksPostgres(db *sqlx.DB) *StudentsMarksPostgres {
	return &StudentsMarksPostgres{db: db}
}

func (r *StudentsMarksPostgres) GetAllTestsMarks(id, disciplineId int) ([]model.TestsMarks, error) {
	var testMarks []model.TestsMarks
	query := fmt.Sprintf("SELECT %s.name, %s.mark FROM %s INNER JOIN %s ON %s.id = %s.test_id INNER JOIN %s ON %s.test_id = %s.id INNER JOIN %s ON %s.id = %s.section_id WHERE user_id = $1 AND discipline_id = $2 ORDER BY %s.id",
		table_name.TestsTable, table_name.TestsMarkTable, table_name.TestsTable,
		table_name.TestsMarkTable, table_name.TestsTable, table_name.TestsMarkTable,
		table_name.SectionTests, table_name.SectionTests, table_name.TestsTable,
		table_name.SectionTable, table_name.SectionTable, table_name.SectionTests,
		table_name.TestsTable)
	err := r.db.Select(&testMarks, query, id, disciplineId)
	return testMarks, err
}

func (r *StudentsMarksPostgres) GetAllLaboratoryMarks(id, disciplineId int) ([]model.LaboratoryMarks, error) {
	var laboratoryMarks []model.LaboratoryMarks
	query := fmt.Sprintf(`SELECT %s.name, %s.mark FROM %s
                        INNER JOIN %s
                        ON %s.external_laboratory_id = %s.id
                        INNER JOIN %s
                        ON %s.id = %s.laboratory_id
                        INNER JOIN %s
                        ON %s.laboratory_id = %s.id
                        INNER JOIN %s 
                        ON %s.id = %s.section_id
                        WHERE user_id = $1 AND discipline_id = $2
                        ORDER BY %s.id`,
		table_name.ExternalLaboratoryTable, table_name.LaboratoryMarkTable, table_name.LaboratoryTable,
		table_name.ExternalLaboratoryTable, table_name.LaboratoryTable, table_name.ExternalLaboratoryTable,
		table_name.LaboratoryMarkTable, table_name.LaboratoryTable, table_name.LaboratoryMarkTable,
		table_name.SectionLabs, table_name.SectionLabs, table_name.LaboratoryTable,
		table_name.SectionTable, table_name.SectionTable, table_name.SectionLabs,
		table_name.LaboratoryTable)
	err := r.db.Select(&laboratoryMarks, query, id, disciplineId)
	return laboratoryMarks, err
}

func (r *StudentsMarksPostgres) GetExamMark(userId, disciplineId int) (int, error) {
	var mark int
	query := fmt.Sprintf("SELECT mark FROM %s WHERE user_id = $1 AND discipline_id = $2", table_name.ExamTable)
	err := r.db.Get(&mark, query, userId, disciplineId)
	if err != nil {
		return 0, err
	}
	return mark, nil
}
