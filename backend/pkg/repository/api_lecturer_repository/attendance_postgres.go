package api_lecturer_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type LecturerAttendancePostgres struct {
	db *sqlx.DB
}

func NewLecturerAttendancePostgres(db *sqlx.DB) *LecturerAttendancePostgres {
	return &LecturerAttendancePostgres{db: db}
}

func (r *LecturerAttendancePostgres) GetAllLessons(disciplineId int) ([]model.Lesson, error) {
	var lessons []model.Lesson
	query := fmt.Sprintf("SELECT id, name FROM %s WHERE discipline_id = $1 ORDER BY id", table_name.LessonsTable)
	err := r.db.Select(&lessons, query, disciplineId)
	if err != nil {
		return []model.Lesson{}, err
	}
	return lessons, nil
}

func (r *LecturerAttendancePostgres) AddLesson(disciplineId int, name string) error {
	query := fmt.Sprintf("INSERT INTO %s (discipline_id, name) VALUES ($1, $2)", table_name.LessonsTable)
	_, err := r.db.Exec(query, disciplineId, name)
	return err
}

func (r *LecturerAttendancePostgres) DeleteLesson(lessonId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", table_name.LessonsTable)
	_, err := r.db.Exec(query, lessonId)
	return err
}

func (r *LecturerAttendancePostgres) ChangeLesson(lessonId int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", table_name.LessonsTable)
	_, err := r.db.Exec(query, name, lessonId)
	return err
}

func (r *LecturerAttendancePostgres) AddSeminar(disciplineId, groupId int, date int, name string) error {
	query := fmt.Sprintf("INSERT INTO %s (discipline_id, group_id, name, date) VALUES ($1, $2, $3, $4)", table_name.SeminarsTable)
	_, err := r.db.Exec(query, disciplineId, groupId, name, date)
	return err
}

func (r *LecturerAttendancePostgres) GetAllSeminars(disciplineId, groupId int) ([]model.Seminar, error) {
	var seminar []model.Seminar
	query := fmt.Sprintf("SELECT id, name, date FROM %s WHERE discipline_id = $1 AND group_id = $2 ORDER BY id", table_name.SeminarsTable)
	err := r.db.Select(&seminar, query, disciplineId, groupId)
	if err != nil {
		return []model.Seminar{}, err
	}
	return seminar, nil
}

func (r *LecturerAttendancePostgres) ChangeSeminar(seminarId int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", table_name.SeminarsTable)
	_, err := r.db.Exec(query, name, seminarId)
	return err
}

func (r *LecturerAttendancePostgres) ChangeSeminarDate(seminarId int, date int) error {
	query := fmt.Sprintf("UPDATE %s SET date = $1 WHERE id = $2", table_name.SeminarsTable)
	_, err := r.db.Exec(query, date, seminarId)
	return err
}

func (r *LecturerAttendancePostgres) DeleteSeminar(seminarId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", table_name.SeminarsTable)
	_, err := r.db.Exec(query, seminarId)
	return err
}

func (r *LecturerAttendancePostgres) AddLessonVisiting(lessonId, userId int, isAbsent bool) error {
	var userCheck int
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE lesson_id = $1 AND user_id = $2", table_name.LessonVisitingTable)
	if err := r.db.Get(&userCheck, query, lessonId, userId); err == nil {
		query = fmt.Sprintf("UPDATE %s SET is_absent = $1 WHERE lesson_id = $2 AND user_id = $3", table_name.LessonVisitingTable)
		if _, err := r.db.Exec(query, isAbsent, lessonId, userId); err != nil {
			return err
		}
		return nil
	}
	query = fmt.Sprintf("INSERT INTO %s (lesson_id, user_id, is_absent) VALUES ($1, $2, $3)", table_name.LessonVisitingTable)
	_, err := r.db.Exec(query, lessonId, userId, isAbsent)
	return err
}

func (r *LecturerAttendancePostgres) AddSeminarVisiting(seminarId, userId int, isAbsent bool) error {
	query := fmt.Sprintf("INSERT INTO %s (seminar_id, user_id, is_absent) VALUES ($1, $2, $3)", table_name.SeminarVisitingTable)
	_, err := r.db.Exec(query, seminarId, userId, isAbsent)
	return err
}

func (r *LecturerAttendancePostgres) GetSeminarVisitingGroup(seminarId int) ([]model.SeminarVisitingStudent, error) {
	var seminar []model.SeminarVisitingStudent
	query := fmt.Sprintf("SELECT %s.id AS student_id, %s.name AS student_name, %s.surname AS student_surname,%s.is_absent FROM %s INNER JOIN %s ON %s.id = %s.user_id WHERE seminar_id = $1 ORDER BY %s.seminar_id",
		table_name.UsersTable, table_name.UsersTable, table_name.UsersTable, table_name.SeminarVisitingTable, table_name.UsersTable,
		table_name.SeminarVisitingTable, table_name.UsersTable, table_name.SeminarVisitingTable, table_name.SeminarVisitingTable)
	err := r.db.Select(&seminar, query, seminarId)
	if err != nil {
		return []model.SeminarVisitingStudent{}, err
	}
	return seminar, nil
}

func (r *LecturerAttendancePostgres) GetLessonVisitingGroup(lessonId, groupId int) ([]model.LessonVisitingStudent, error) {
	var lesson []model.LessonVisitingStudent
	query := fmt.Sprintf("SELECT %s.id AS student_id, %s.name AS student_name, %s.surname AS student_surname, %s.is_absent FROM %s INNER JOIN %s ON %s.id = %s.user_id INNER JOIN %s ON %s.id = %s.user_id WHERE lesson_id = $1 AND group_id = $2 ORDER BY %s.lesson_id",
		table_name.UsersTable, table_name.UsersTable, table_name.UsersTable, table_name.LessonVisitingTable, table_name.UsersTable,
		table_name.LessonVisitingTable, table_name.UsersTable, table_name.LessonVisitingTable, table_name.StudentsTable, table_name.UsersTable, table_name.StudentsTable, table_name.LessonVisitingTable)
	err := r.db.Select(&lesson, query, lessonId, groupId)
	if err != nil {
		return []model.LessonVisitingStudent{}, err
	}
	return lesson, nil
}

func (r *LecturerAttendancePostgres) ChangeSeminarVisiting(seminarId, userId int, isAbsent bool) error {
	query := fmt.Sprintf("UPDATE %s SET is_absent = $1 WHERE seminar_id = $2 AND user_id = $3", table_name.SeminarVisitingTable)
	_, err := r.db.Exec(query, isAbsent, seminarId, userId)
	return err
}

func (r *LecturerAttendancePostgres) ChangeLessonVisiting(lessonId, userId int, isAbsent bool) error {
	query := fmt.Sprintf("UPDATE %s SET is_absent = $1 WHERE lesson_id = $2 AND user_id = $3", table_name.LessonVisitingTable)
	_, err := r.db.Exec(query, isAbsent, lessonId, userId)
	return err
}

func (r *LecturerAttendancePostgres) DeleteLessonDate(lessonId int, groupId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE lesson_id = $1 AND group_id = $2", table_name.LessonDateTable)
	_, err := r.db.Exec(query, lessonId, groupId)
	return err
}

func (r *LecturerAttendancePostgres) AddLessonDate(lessonId, groupId, date int, description string) error {
	query := fmt.Sprintf("INSERT INTO %s (lesson_id, group_id, date, description) VALUES ($1, $2, $3, $4)", table_name.LessonDateTable)
	_, err := r.db.Exec(query, lessonId, groupId, date, description)
	return err
}

func (r *LecturerAttendancePostgres) ChangeLessonDate(lessonId, groupId, date int) error {
	query := fmt.Sprintf("UPDATE %s SET date = $1 WHERE lesson_id = $2 AND group_id = $3", table_name.LessonDateTable)
	_, err := r.db.Exec(query, date, lessonId, groupId)
	return err
}

func (r *LecturerAttendancePostgres) ChangeLessonDateDescription(lessonId, groupId int, description string) error {
	query := fmt.Sprintf("UPDATE %s SET description = $1 WHERE lesson_id = $2 AND group_id = $3", table_name.LessonDateTable)
	_, err := r.db.Exec(query, description, lessonId, groupId)
	return err
}

func (r *LecturerAttendancePostgres) GetLessonDate(lessonId, groupId int) (model.LessonDate, error) {
	var lessonDate model.LessonDate
	query := fmt.Sprintf("SELECT %s.lesson_id, %s.name, %s.description, %s.date, %s.group_id FROM %s INNER JOIN %s ON %s.id = %s.lesson_id WHERE lesson_id = $1 AND group_id = $2 ORDER BY %s.lesson_id",
		table_name.LessonDateTable, table_name.LessonDateTable, table_name.LessonsTable, table_name.LessonDateTable, table_name.LessonDateTable,
		table_name.LessonsTable, table_name.LessonDateTable, table_name.LessonsTable, table_name.LessonDateTable, table_name.LessonDateTable,
	)

	if err := r.db.Get(&lessonDate, query, lessonId, groupId); err != nil {
		return model.LessonDate{}, err
	}
	return lessonDate, nil
}

func (r *LecturerAttendancePostgres) GetTableLessons(disciplineId int) ([]model.LessonDate, error) {
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

func (r *LecturerAttendancePostgres) GetTableSeminars(disciplineId int) ([]model.SeminarDate, error) {
	var seminars []model.SeminarDate
	query := fmt.Sprintf("SELECT id, name, date FROM %s WHERE discipline_id = $1 ORDER BY id",
		table_name.SeminarsTable)
	if err := r.db.Select(&seminars, query, disciplineId); err != nil {
		return nil, err
	}
	return seminars, nil
}
