package api_student_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type StudentAttendancePostgres struct {
	db *sqlx.DB
}

func NewStudentAttendancePostgres(db *sqlx.DB) *StudentAttendancePostgres {
	return &StudentAttendancePostgres{db: db}
}

func (r *StudentAttendancePostgres) GetAllSeminarVisiting(disciplineId, userId int) ([]model.SeminarVisiting, error) {
	var seminars []model.SeminarVisiting
	query := fmt.Sprintf("SELECT %s.id, %s.name, %s.is_absent FROM %s INNER JOIN %s ON %s.id = %s.seminar_id WHERE user_id = $1 AND discipline_id = $2 ORDER BY %s.id",
		table_name.SeminarsTable, table_name.SeminarsTable, table_name.SeminarVisitingTable, table_name.SeminarsTable,
		table_name.SeminarVisitingTable, table_name.SeminarsTable, table_name.SeminarVisitingTable, table_name.SeminarsTable)
	err := r.db.Select(&seminars, query, userId, disciplineId)
	if err != nil {
		return []model.SeminarVisiting{}, err
	}
	return seminars, nil
}

func (r *StudentAttendancePostgres) GetAllLessonVisiting(disciplineId, userId int) ([]model.LessonVisiting, error) {
	var lessons []model.LessonVisiting
	query := fmt.Sprintf("SELECT %s.id, %s.name, %s.is_absent FROM %s INNER JOIN %s ON %s.id = %s.lesson_id WHERE user_id = $1 AND discipline_id = $2 ORDER BY %s.id",
		table_name.LessonsTable, table_name.LessonsTable,
		table_name.LessonVisitingTable, table_name.LessonsTable, table_name.LessonVisitingTable,
		table_name.LessonsTable, table_name.LessonVisitingTable, table_name.LessonsTable)
	err := r.db.Select(&lessons, query, userId, disciplineId)
	if err != nil {
		return []model.LessonVisiting{}, err
	}
	return lessons, nil
}

func (r *StudentAttendancePostgres) GetLessonDate(lessonId, groupId int) (model.LessonDate, error) {
	var lessonDate model.LessonDate
	query := fmt.Sprintf("SELECT * FROM %s WHERE lesson_id = $1 AND group_id = $2", table_name.LessonDateTable)
	if err := r.db.Get(&lessonDate, query, lessonId, groupId); err != nil {
		return model.LessonDate{}, err
	}
	return lessonDate, nil
}

func (r *StudentAttendancePostgres) GetAllLessons(userId, disciplineId int) ([]model.LessonDate, error) {
	var groupId model.GroupId
	query := fmt.Sprintf("SELECT group_id FROM %s WHERE user_id = $1", table_name.StudentsTable)
	if err := r.db.Get(&groupId, query, userId); err != nil {
		err = errors.New("access denied")
		return nil, err
	}
	if groupId.Id == 0 {
		err := errors.New("access denied")
		return nil, err
	}
	var lessons []model.LessonDate
	query = fmt.Sprintf("SELECT %s.lesson_id, %s.description, %s.name, %s.date, %s.group_id FROM %s INNER JOIN %s ON %s.id = %s.lesson_id WHERE discipline_id = $1 AND group_id = $2 ORDER BY %s.lesson_id",
		table_name.LessonDateTable, table_name.LessonDateTable, table_name.LessonsTable, table_name.LessonDateTable, table_name.LessonDateTable,
		table_name.LessonsTable, table_name.LessonDateTable, table_name.LessonsTable, table_name.LessonDateTable, table_name.LessonDateTable,
	)
	if err := r.db.Select(&lessons, query, disciplineId, groupId.Id); err != nil {
		return nil, err
	}
	return lessons, nil
}

func (r *StudentAttendancePostgres) GetAllSeminars(userId, disciplineId int) ([]model.SeminarDate, error) {
	var groupId model.GroupId
	query := fmt.Sprintf("SELECT group_id FROM %s WHERE user_id = $1", table_name.StudentsTable)
	if err := r.db.Get(&groupId, query, userId); err != nil {
		err = errors.New("access denied")
		return nil, err
	}
	if groupId.Id == 0 {
		err := errors.New("access denied")
		return nil, err
	}
	var seminars []model.SeminarDate
	query = fmt.Sprintf("SELECT id, name, date FROM %s WHERE discipline_id = $1 AND group_id = $2 ORDER BY id",
		table_name.SeminarsTable)
	if err := r.db.Select(&seminars, query, disciplineId, groupId.Id); err != nil {
		return nil, err
	}
	return seminars, nil
}
