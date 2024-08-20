package api_seminarian_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SeminarianAttendancePostgres struct {
	db *sqlx.DB
}

func NewSeminarianAttendancePostgres(db *sqlx.DB) *SeminarianAttendancePostgres {
	return &SeminarianAttendancePostgres{db: db}
}

func (r *SeminarianAttendancePostgres) GetAllLessons(disciplineId int) ([]model.Lesson, error) {
	var lessons []model.Lesson
	query := fmt.Sprintf("SELECT id, name FROM %s WHERE discipline_id = $1 ORDER BY id", table_name.LessonsTable)
	err := r.db.Select(&lessons, query, disciplineId)
	if err != nil {
		return []model.Lesson{}, err
	}
	return lessons, nil
}

func (r *SeminarianAttendancePostgres) AddSeminar(disciplineId, groupId int, name string, date int) error {
	query := fmt.Sprintf("INSERT INTO %s (discipline_id, group_id, name, date) VALUES ($1, $2, $3, $4)", table_name.SeminarsTable)
	_, err := r.db.Exec(query, disciplineId, groupId, name, date)
	return err
}

func (r *SeminarianAttendancePostgres) GetAllSeminars(disciplineId, groupId int) ([]model.Seminar, error) {
	var seminar []model.Seminar
	query := fmt.Sprintf("SELECT id, name FROM %s WHERE discipline_id = $1 AND group_id = $2 ORDER BY id", table_name.SeminarsTable)
	err := r.db.Select(&seminar, query, disciplineId, groupId)
	if err != nil {
		return []model.Seminar{}, err
	}
	return seminar, nil
}

func (r *SeminarianAttendancePostgres) ChangeSeminar(seminarId int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", table_name.SeminarsTable)
	_, err := r.db.Exec(query, name, seminarId)
	return err
}

func (r *SeminarianAttendancePostgres) DeleteSeminar(seminarId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", table_name.SeminarsTable)
	_, err := r.db.Exec(query, seminarId)
	return err
}

func (r *SeminarianAttendancePostgres) AddLessonVisiting(lessonId, userId int, isAbsent bool) error {
	query := fmt.Sprintf("INSERT INTO %s (lesson_id, user_id, is_absent) VALUES ($1, $2, $3)", table_name.LessonVisitingTable)
	_, err := r.db.Exec(query, lessonId, userId, isAbsent)
	return err
}

func (r *SeminarianAttendancePostgres) AddSeminarVisiting(seminarId, userId int, isAbsent bool) error {
	query := fmt.Sprintf("INSERT INTO %s (seminar_id, user_id, is_absent) VALUES ($1, $2, $3)", table_name.SeminarVisitingTable)
	_, err := r.db.Exec(query, seminarId, userId, isAbsent)
	return err
}

func (r *SeminarianAttendancePostgres) GetSeminarVisitingGroup(seminarId int) ([]model.SeminarVisitingStudent, error) {
	var seminar []model.SeminarVisitingStudent
	query := fmt.Sprintf("SELECT %s.id AS student_id, %s.name AS student_name, %s.is_absent FROM %s INNER JOIN %s ON %s.id = %s.user_id WHERE seminar_id = $1 ORDER BY %s.surname, %s.name",
		table_name.UsersTable, table_name.UsersTable, table_name.SeminarVisitingTable, table_name.UsersTable,
		table_name.SeminarVisitingTable, table_name.UsersTable, table_name.SeminarVisitingTable, table_name.UsersTable, table_name.UsersTable)
	err := r.db.Select(&seminar, query, seminarId)
	if err != nil {
		return []model.SeminarVisitingStudent{}, err
	}
	return seminar, nil
}

func (r *SeminarianAttendancePostgres) GetLessonVisitingGroup(lessonId, groupId int) ([]model.LessonVisitingStudent, error) {
	var lesson []model.LessonVisitingStudent
	query := fmt.Sprintf("SELECT %s.id AS student_id, %s.name AS student_name, %s.is_absent FROM %s INNER JOIN %s ON %s.id = %s.user_id INNER JOIN %s ON %s.id = %s.user_id WHERE lesson_id = $1 AND group_id = $2 ORDER BY %s.surname, %s.name",
		table_name.UsersTable, table_name.UsersTable, table_name.LessonVisitingTable, table_name.UsersTable,
		table_name.LessonVisitingTable, table_name.UsersTable, table_name.LessonVisitingTable,
		table_name.StudentsTable, table_name.UsersTable, table_name.StudentsTable, table_name.UsersTable, table_name.UsersTable)
	err := r.db.Select(&lesson, query, lessonId, groupId)
	if err != nil {
		return []model.LessonVisitingStudent{}, err
	}
	return lesson, nil
}

func (r *SeminarianAttendancePostgres) ChangeSeminarVisiting(seminarId, userId int, isAbsent bool) error {
	query := fmt.Sprintf("UPDATE %s SET is_absent = $1 WHERE seminar_id = $2 AND user_id = $3", table_name.SeminarVisitingTable)
	_, err := r.db.Exec(query, isAbsent, seminarId, userId)
	return err
}

func (r *SeminarianAttendancePostgres) ChangeLessonVisiting(lessonId, userId int, isAbsent bool) error {
	query := fmt.Sprintf("UPDATE %s SET is_absent = $1 WHERE lesson_id = $2 AND user_id = $3", table_name.LessonVisitingTable)
	_, err := r.db.Exec(query, isAbsent, lessonId, userId)
	return err
}

func (r *SeminarianAttendancePostgres) CheckAccessForGroup(seminarianId, groupId, disciplineId int) error {
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

func (r *SeminarianAttendancePostgres) ChangeSeminarDate(seminarId int, date int) error {
	query := fmt.Sprintf("UPDATE %s SET date = $1 WHERE id = $2", table_name.SeminarsTable)
	_, err := r.db.Exec(query, date, seminarId)
	return err
}

func (r *SeminarianAttendancePostgres) GetLessonDate(lessonId, groupId int) (model.LessonDate, error) {
	var lessonDate model.LessonDate
	query := fmt.Sprintf("SELECT * FROM %s WHERE lesson_id = $1 AND group_id = $2", table_name.LessonDateTable)
	if err := r.db.Get(&lessonDate, query, lessonId, groupId); err != nil {
		return model.LessonDate{}, err
	}
	return lessonDate, nil
}

func (r *SeminarianAttendancePostgres) CheckAccessForStudent(seminarianId, userId, disciplineId int) error {
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

func (r *SeminarianAttendancePostgres) CheckAccessForDiscipline(seminarianId, disciplineId int) error {
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

func (r *SeminarianAttendancePostgres) CheckAccessForSeminar(seminarianId, seminarId int) error {
	var seminar model.AllInfoSeminar
	errAccess := errors.New("access denied to seminar")
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", table_name.SeminarsTable)
	err := r.db.Get(&seminar, query, seminarId)
	if err != nil {
		return errAccess
	}
	if err := r.CheckAccessForDiscipline(seminarianId, seminar.DisciplineId); err != nil {
		return err
	}
	if err := r.CheckAccessForGroup(seminarianId, seminar.GroupId, seminar.DisciplineId); err != nil {
		return err
	}
	return nil
}

func (r *SeminarianAttendancePostgres) CheckAccessForLesson(seminarianId, lessonId, groupId int) error {
	var disciplineId int
	errAccess := errors.New("access denied to lesson")
	query := fmt.Sprintf("SELECT discipline_id FROM %s WHERE id = $1", table_name.LessonsTable)
	if err := r.db.Get(&disciplineId, query, lessonId); err != nil {
		return errAccess
	}
	if err := r.CheckAccessForGroup(seminarianId, groupId, disciplineId); err != nil {
		return errAccess
	}
	return nil
}

func (r *SeminarianAttendancePostgres) CheckAccessForStudentToLesson(seminarianId, userId, lessonId int) error {
	var discipline int
	errAccess := errors.New("access denied to lesson")
	query := fmt.Sprintf("SELECT discipline_id FROM %s WHERE id = $1", table_name.LessonsTable)
	if err := r.db.Get(&discipline, query, lessonId); err != nil {
		return errAccess
	}
	if err := r.CheckAccessForStudent(seminarianId, userId, discipline); err != nil {
		return err
	}
	return nil
}

func (r *SeminarianAttendancePostgres) CheckAccessForStudentToSeminar(seminarianId, userId, seminarId int) error {
	var discipline int
	errAccess := errors.New("access denied to seminar")
	query := fmt.Sprintf("SELECT discipline_id FROM %s WHERE id = $1", table_name.SeminarsTable)
	if err := r.db.Get(&discipline, query, seminarId); err != nil {
		return errAccess
	}
	if err := r.CheckAccessForStudent(seminarianId, userId, discipline); err != nil {
		return err
	}
	return nil
}

func (r *SeminarianAttendancePostgres) GetTableLessons(disciplineId int) ([]model.LessonDate, error) {
	var lessons []model.LessonDate
	query := fmt.Sprintf("SELECT %s.lesson_id, %s.description, %s.name, %s.date, %s.group_id FROM %s INNER JOIN %s ON %s.id = %s.lesson_id WHERE discipline_id = $1 ORDER BY %s.lesson_id",
		table_name.LessonDateTable, table_name.LessonDateTable, table_name.LessonsTable, table_name.LessonDateTable, table_name.LessonDateTable,
		table_name.LessonsTable, table_name.LessonDateTable, table_name.LessonsTable, table_name.LessonDateTable, table_name.LessonDateTable,
	)
	if err := r.db.Select(&lessons, query, disciplineId); err != nil {
		return nil, err
	}
	return lessons, nil
}

func (r *SeminarianAttendancePostgres) GetTableSeminars(disciplineId int) ([]model.SeminarDate, error) {
	var seminars []model.SeminarDate
	query := fmt.Sprintf("SELECT id, name, date FROM %s WHERE discipline_id = $1 ORDER BY id",
		table_name.SeminarsTable)
	if err := r.db.Select(&seminars, query, disciplineId); err != nil {
		return nil, err
	}
	return seminars, nil
}
