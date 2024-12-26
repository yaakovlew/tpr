package api_lecturer_repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"backend/pkg/model"
	"backend/pkg/repository/table_name"
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

func (r *LecturerMarksPostgres) GetAttendanceMarksFromGroup(disciplineId, groupID int) ([]model.AttendanceMark, error) {
	var marks []model.AttendanceMark
	query := fmt.Sprintf(`WITH exist_lessons AS (
    SELECT COUNT(*) AS count 
    FROM %s 
    WHERE %s.discipline_id = $1
),
exist_seminars AS (
    SELECT COUNT(*) AS count 
    FROM %s 
    JOIN %s ON %s.id = %s.seminar_id
    WHERE %s.discipline_id = $2 AND %s.group_id = $3
),
exist_lesson_table AS (
    SELECT 
        %s.id, 
        %s.name, 
        %s.surname,
				lesson_visiting_mark, 
        COUNT(*) AS exist 
    FROM %s
    JOIN %s ON %s.id = %s.user_id
    JOIN %s ON %s.group_id = %s.group_id
    JOIN %s ON %s.id = %s.discipline_id
    JOIN %s ON %s.discipline_id = %s.id
    JOIN %s ON %s.id = %s.lesson_id
    WHERE %s.discipline_id = $4 
      AND %s.group_id = $5 
      AND %s.is_absent = false
    GROUP BY %s.id, %s.name, %s.surname, %s.lesson_visiting_mark
),
exist_seminar_table AS (
    SELECT 
        %s.id, 
        %s.name, 
        %s.surname,
				seminar_visiting_mark, 
        COUNT(*) AS exist 
    FROM %s
    JOIN %s ON %s.id = %s.user_id
    JOIN %s ON %s.group_id = %s.group_id
    JOIN %s ON %s.id = %s.discipline_id
    JOIN %s ON %s.discipline_id = %s.id
    JOIN %s ON %s.seminar_id = %s.id
    WHERE %s.discipline_id = $6 
      AND %s.group_id = $7 
      AND %s.is_absent = false
    GROUP BY %s.id, %s.name, %s.surname, %s.seminar_visiting_mark
)
SELECT 
    %s.id AS user_id, 
    %s.name AS user_name, 
    %s.surname AS user_surname, 
    ROUND(
        CASE 
					WHEN exist_seminars.count = 0 THEN 0
					WHEN exist_seminar_table.exist IS NULL THEN 0
          ELSE exist_seminar_table.seminar_visiting_mark * exist_seminar_table.exist::float / exist_seminars.count
        END
    ) AS seminar,
    ROUND(
        CASE 
					WHEN exist_lessons.count = 0 THEN 0
					WHEN exist_lesson_table.exist IS NULL THEN 0
          ELSE exist_lesson_table.lesson_visiting_mark * exist_lesson_table.exist::float / exist_lessons.count
        END
    ) AS lesson
FROM %s
LEFT JOIN exist_seminar_table ON exist_seminar_table.id = %s.id
LEFT JOIN exist_lesson_table ON exist_lesson_table.id = exist_seminar_table.id
JOIN %s ON %s.id = %s.user_id
JOIN %s ON %s.group_id = %s.group_id
JOIN %s ON %s.id = %s.discipline_id
CROSS JOIN exist_seminars
CROSS JOIN exist_lessons
WHERE %s.discipline_id = $8 AND %s.group_id = $9`,
		table_name.LessonsTable, table_name.LessonsTable,

		table_name.SeminarsTable, table_name.SeminarVisitingTable, table_name.SeminarsTable,
		table_name.SeminarVisitingTable, table_name.SeminarsTable, table_name.SeminarsTable,

		table_name.UsersTable, table_name.UsersTable, table_name.UsersTable, table_name.UsersTable,
		table_name.StudentsTable, table_name.UsersTable, table_name.StudentsTable,
		table_name.CurriculumTable, table_name.StudentsTable, table_name.CurriculumTable,
		table_name.DisciplinesTable, table_name.DisciplinesTable, table_name.CurriculumTable,
		table_name.LessonsTable, table_name.LessonsTable, table_name.DisciplinesTable,
		table_name.LessonVisitingTable, table_name.LessonsTable, table_name.LessonVisitingTable,
		table_name.CurriculumTable, table_name.CurriculumTable, table_name.LessonVisitingTable,
		table_name.UsersTable, table_name.UsersTable, table_name.UsersTable, table_name.DisciplinesTable,

		table_name.UsersTable, table_name.UsersTable, table_name.UsersTable, table_name.UsersTable,
		table_name.StudentsTable, table_name.UsersTable, table_name.StudentsTable,
		table_name.CurriculumTable, table_name.StudentsTable, table_name.CurriculumTable,
		table_name.DisciplinesTable, table_name.DisciplinesTable, table_name.CurriculumTable,
		table_name.SeminarsTable, table_name.SeminarsTable, table_name.DisciplinesTable,
		table_name.SeminarVisitingTable, table_name.SeminarVisitingTable, table_name.SeminarsTable,
		table_name.CurriculumTable, table_name.CurriculumTable, table_name.SeminarVisitingTable,
		table_name.UsersTable, table_name.UsersTable, table_name.UsersTable, table_name.DisciplinesTable,

		table_name.UsersTable, table_name.UsersTable, table_name.UsersTable, table_name.UsersTable, table_name.UsersTable,
		table_name.StudentsTable, table_name.UsersTable, table_name.StudentsTable,
		table_name.CurriculumTable, table_name.StudentsTable, table_name.CurriculumTable,
		table_name.DisciplinesTable, table_name.DisciplinesTable, table_name.CurriculumTable,
		table_name.CurriculumTable, table_name.CurriculumTable,
	)

	if err := r.db.Select(&marks, query, disciplineId, disciplineId, groupID, disciplineId, groupID, disciplineId, groupID, disciplineId, groupID); err != nil {
		return nil, err
	}

	return marks, nil
}
