package api_lecturer_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type LecturerReport struct {
	db *sqlx.DB
}

func NewLecturerReport(db *sqlx.DB) *LecturerReport {
	return &LecturerReport{db: db}
}

func (r *LecturerReport) GetThemesFromDiscipline(disciplineId int) ([]model.Section, error) {
	var sections []model.Section
	query := fmt.Sprintf("SELECT id, name FROM %s WHERE discipline_id = $1 ORDER BY id", table_name.SectionTable)
	err := r.db.Select(&sections, query, disciplineId)
	if err != nil {
		return nil, err
	}
	return sections, nil
}

func (r *LecturerReport) GetAllTests(sectionId int) ([]model.TestsReport, error) {
	var test []model.TestsReport
	query := fmt.Sprintf("SELECT test_id FROM %s WHERE section_id = $1 ORDER BY test_id", table_name.SectionTests)
	err := r.db.Select(&test, query, sectionId)
	if err != nil {
		return nil, err
	}
	return test, nil
}

func (r *LecturerReport) GetAllLabs(sectionId int) ([]model.LabsReport, error) {
	var lab []model.LabsReport
	query := fmt.Sprintf("SELECT laboratory_id FROM %s WHERE section_id = $1 ORDER BY laboratory_id", table_name.SectionLabs)
	err := r.db.Select(&lab, query, sectionId)
	if err != nil {
		return nil, err
	}
	return lab, nil
}

func (r *LecturerReport) GetAllStudents(groupId int) ([]model.StudentReport, error) {
	var student []model.StudentReport
	query := fmt.Sprintf("SELECT %s.id AS student_id, %s.name AS student_name, %s.surname AS student_surname FROM %s INNER JOIN %s ON %s.id = %s.user_id WHERE %s.group_id = $1 ORDER BY %s.surname, %s.name",
		table_name.UsersTable, table_name.UsersTable, table_name.UsersTable, table_name.UsersTable, table_name.StudentsTable, table_name.UsersTable,
		table_name.StudentsTable, table_name.StudentsTable, table_name.UsersTable, table_name.UsersTable)
	err := r.db.Select(&student, query, groupId)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (r *LecturerReport) GetStudentMarkTest(studentId, testId int) int {
	var mark int
	query := fmt.Sprintf("SELECT mark FROM %s WHERE user_id = $1 AND test_id = $2", table_name.TestsMarkTable)
	err := r.db.Get(&mark, query, studentId, testId)
	if err != nil {
		return 0
	}
	return mark
}

func (r *LecturerReport) GetStudentMarkLab(studentId, labId int) int {
	var mark int
	query := fmt.Sprintf("SELECT mark FROM %s WHERE user_id = $1 AND laboratory_id = $2", table_name.LaboratoryMarkTable)
	err := r.db.Get(&mark, query, studentId, labId)
	if err != nil {
		return 0
	}
	return mark
}

func (r *LecturerReport) GetMarkFromSection(userId, sectionId int) (int, error) {
	mark := 0
	tests, err := r.GetAllTests(sectionId)
	if err != nil {
		return 0, err
	}
	labs, err := r.GetAllLabs(sectionId)
	if err != nil {
		return 0, err
	}
	for _, value := range tests {
		mark = mark + r.GetStudentMarkTest(userId, value.TestId)
	}
	for _, value := range labs {
		mark = mark + r.GetStudentMarkLab(userId, value.LabId)
	}
	var disciplineId int
	query := fmt.Sprintf("SELECT discipline_id FROM %s WHERE id = $1", table_name.SectionTable)
	if err := r.db.Get(&disciplineId, query, sectionId); err != nil {
		return 0, err
	}
	countSection := r.GetSummaryCountSection(disciplineId)
	if countSection == 0 {
		err := errors.New("count section is equal 0")
		return 0, err
	}
	markForLesson := r.GetAttendanceLesson(userId, disciplineId)
	if countSection == 0 {
		err := errors.New("count section is equal 0")
		return 0, err
	}
	markForSeminar := r.GetAttendanceSeminar(userId, disciplineId)
	if countSection == 0 {
		err := errors.New("count section is equal 0")
		return 0, err
	}
	mark = mark + (markForLesson+markForSeminar)/countSection
	return mark, nil
}

func (r *LecturerReport) GetSummaryCountSection(disciplineId int) int {
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE discipline_id = $1", table_name.SectionTable)
	if err := r.db.Get(&count, query, disciplineId); err != nil {
		return 0
	}
	return count
}

func (r *LecturerReport) GetMarkFromExam(userId, disciplineId int) int {
	var mark int
	query := fmt.Sprintf("SELECT mark FROM %s WHERE user_id = $1 AND discipline_id = $2", table_name.ExamTable)
	err := r.db.Get(&mark, query, userId, disciplineId)
	if err != nil {
		return 0
	}
	return mark
}

func (r *LecturerReport) GetSummaryMarkFromSections(userId, disciplineId int) int {
	sections, err := r.GetThemesFromDiscipline(disciplineId)
	if err != nil {
		return 0
	}
	sum := 0
	for _, section := range sections {
		mark, _ := r.GetMarkFromSection(userId, section.SectionId)
		sum = sum + mark
	}
	return sum
}

func (r *LecturerReport) GetAttendanceLesson(studentId, disciplineId int) int {
	var mark int
	query := fmt.Sprintf("SELECT lesson_visiting_mark FROM %s WHERE id = $1", table_name.DisciplinesTable)
	if err := r.db.Get(&mark, query, disciplineId); err != nil {
		return 0
	}
	var countVisiting int
	query = fmt.Sprintf("SELECT COUNT(*) FROM %s INNER JOIN %s ON %s.lesson_id = %s.id WHERE user_id = $1 AND discipline_id = $2 AND is_absent = true",
		table_name.LessonVisitingTable, table_name.LessonsTable, table_name.LessonVisitingTable, table_name.LessonsTable)
	if err := r.db.Get(&countVisiting, query, studentId, disciplineId); err != nil {
		return 0
	}
	var countNotVisiting int
	query = fmt.Sprintf("SELECT COUNT(*) FROM %s INNER JOIN %s ON %s.lesson_id = %s.id WHERE user_id = $1 AND discipline_id = $2 AND is_absent = false",
		table_name.LessonVisitingTable, table_name.LessonsTable, table_name.LessonVisitingTable, table_name.LessonsTable)
	if err := r.db.Get(&countNotVisiting, query, studentId, disciplineId); err != nil {
		return 0
	}
	if (countNotVisiting + countVisiting) == 0 {
		return 0
	}
	return mark * countVisiting / (countNotVisiting + countVisiting)
}

func (r *LecturerReport) GetAttendanceSeminar(studentId, disciplineId int) int {
	var mark int
	query := fmt.Sprintf("SELECT seminar_visiting_mark FROM %s WHERE id = $1", table_name.DisciplinesTable)
	if err := r.db.Get(&mark, query, disciplineId); err != nil {
		return 0
	}
	var countVisiting int
	query = fmt.Sprintf("SELECT COUNT(*) FROM %s INNER JOIN %s ON %s.seminar_id = %s.id WHERE user_id = $1 AND discipline_id = $2 AND is_absent = true",
		table_name.SeminarVisitingTable, table_name.SeminarsTable, table_name.SeminarVisitingTable, table_name.SeminarsTable)
	if err := r.db.Get(&countVisiting, query, studentId, disciplineId); err != nil {
		return 0
	}
	var countNotVisiting int
	query = fmt.Sprintf("SELECT COUNT(*) FROM %s INNER JOIN %s ON %s.seminar_id = %s.id WHERE user_id = $1 AND discipline_id = $2 AND is_absent = false",
		table_name.SeminarVisitingTable, table_name.SeminarsTable, table_name.SeminarVisitingTable, table_name.SeminarsTable)
	if err := r.db.Get(&countNotVisiting, query, studentId, disciplineId); err != nil {
		return 0
	}
	if (countNotVisiting + countVisiting) == 0 {
		return 0
	}
	return mark * countVisiting / (countNotVisiting + countVisiting)
}

func (r *LecturerReport) GetResult(studentId, disciplineId int) (string, string) {
	result := "зачтено"
	sections, err := r.GetThemesFromDiscipline(disciplineId)
	if err != nil {
		return "не зачтено", "F"
	}
	for _, section := range sections {
		tests, err := r.GetAllTests(section.SectionId)
		if err != nil {
			return "не зачтено", "F"
		}
		for _, test := range tests {
			if !r.GetMarkFromTest(studentId, test.TestId) {
				result = "не зачтено"
			}
		}
		labs, err := r.GetAllLabs(section.SectionId)
		if err != nil {
			return "не зачтено", "F"
		}
		for _, lab := range labs {
			if !r.GetMarkFromLab(studentId, lab.LabId) {
				result = "не зачтено"
			}
		}
	}
	mark := r.GetSummaryMarkFromSections(studentId, disciplineId) + r.GetMarkFromExam(studentId, disciplineId)
	if mark >= 90 {
		return result, "A"
	} else if mark >= 85 {
		return result, "B"
	} else if mark >= 75 {
		return result, "C"
	} else if mark >= 70 {
		return result, "D"
	} else if mark >= 65 {
		return result, "D"
	} else if mark >= 60 {
		return result, "E"
	} else {
		return "не зачтено", "F"
	}
}

func (r *LecturerReport) GetSectionsResult(studentId, disciplineId int) string {
	result := "а"
	sections, err := r.GetThemesFromDiscipline(disciplineId)
	if err != nil {
		return "н/а"
	}
	for _, section := range sections {
		tests, err := r.GetAllTests(section.SectionId)
		if err != nil {
			return "н/а"
		}
		for _, test := range tests {
			if !r.GetMarkFromTest(studentId, test.TestId) {
				result = "н/а"
			}
		}
		labs, err := r.GetAllLabs(section.SectionId)
		if err != nil {
			return "н/а"
		}
		for _, lab := range labs {
			if !r.GetMarkFromLab(studentId, lab.LabId) {
				result = "н/а"
			}
		}
	}
	return result
}

func (r *LecturerReport) GetMarkFromTest(studentId, testId int) bool {
	var mark int
	query := fmt.Sprintf("SELECT mark FROM %s WHERE user_id = $1 AND test_id = $2", table_name.TestsMarkTable)
	if err := r.db.Get(&mark, query, studentId, testId); err != nil {
		return false
	}
	var defaultMark int
	query = fmt.Sprintf("SELECT default_mark FROM %s WHERE id = $1", table_name.TestsTable)
	if err := r.db.Get(&defaultMark, query, studentId, testId); err != nil {
		return false
	}
	if float64(mark) < 0.6*float64(defaultMark) {
		return false
	}
	return true
}

func (r *LecturerReport) GetMarkFromLab(studentId, laboratoryId int) bool {
	var mark int
	query := fmt.Sprintf("SELECT mark FROM %s WHERE user_id = $1 AND laboratory_id = $2", table_name.LaboratoryMarkTable)
	if err := r.db.Get(&mark, query, studentId, laboratoryId); err != nil {
		return false
	}
	var defaultMark int
	query = fmt.Sprintf("SELECT default_mark FROM %s WHERE id = $1", table_name.LaboratoryTable)
	if err := r.db.Get(&defaultMark, query, studentId, laboratoryId); err != nil {
		return false
	}
	if float64(mark) < 0.6*float64(defaultMark) {
		return false
	}
	return true
}
