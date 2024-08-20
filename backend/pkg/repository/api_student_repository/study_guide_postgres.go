package api_student_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type StudentStudyGuidePostgres struct {
	db *sqlx.DB
}

func NewStudentStudyGuidePostgres(db *sqlx.DB) *StudentStudyGuidePostgres {
	return &StudentStudyGuidePostgres{db: db}
}

func (r *StudentStudyGuidePostgres) GetDigitalDiscipline(disciplineId int) ([]model.DigitalDisciplineWithInfo, []model.DigitalDisciplineWithInfo, error) {
	var digitalDiscipline []model.DigitalDisciplineWithInfo
	query := fmt.Sprintf("SELECT %s.name, %s.description, %s.id AS digital_material_id FROM %s  INNER JOIN %s ON digital_material_id = id WHERE discipline_id = $1 ORDER BY digital_material_id",
		table_name.DigitalGuideTable, table_name.DigitalGuideTable, table_name.DigitalGuideTable, table_name.DigitalGuideTable, table_name.DigitalDisciplineTable)
	err := r.db.Select(&digitalDiscipline, query, disciplineId)
	if err != nil {
		return nil, nil, err
	}
	var digitalDisciplineEn []model.DigitalDisciplineWithInfo
	query = fmt.Sprintf("SELECT %s.name_en as name, %s.description_en as description, %s.id AS digital_material_id FROM %s  INNER JOIN %s ON digital_material_id = id WHERE discipline_id = $1 ORDER BY digital_material_id",
		table_name.DigitalGuideTable, table_name.DigitalGuideTable, table_name.DigitalGuideTable, table_name.DigitalGuideTable, table_name.DigitalDisciplineTable)
	err = r.db.Select(&digitalDisciplineEn, query, disciplineId)
	if err != nil {
		return nil, nil, err
	}
	return digitalDiscipline, digitalDisciplineEn, nil
}

func (r *StudentStudyGuidePostgres) GetFile(fileId int) (string, error) {
	var path string
	query := fmt.Sprintf("SELECT name FROM %s WHERE id = $1", table_name.FilesPathTable)
	err := r.db.Get(&path, query, fileId)
	if err != nil {
		return "", err
	}
	return path, nil
}

func (r *StudentStudyGuidePostgres) GetFilesIdFromDigital(digitalId int) ([]model.FileId, error) {
	var files []model.FileId
	query := fmt.Sprintf("SELECT %s.id, %s.name FROM %s INNER JOIN %s ON %s.digital_material_id = %s.id WHERE %s.id = $1 ORDER BY %s.id",
		table_name.FilesPathTable, table_name.FilesPathTable, table_name.FilesPathTable, table_name.DigitalGuideTable,
		table_name.FilesPathTable, table_name.DigitalGuideTable, table_name.DigitalGuideTable, table_name.FilesPathTable)
	err := r.db.Select(&files, query, digitalId)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (r *StudentStudyGuidePostgres) CheckAccessForDiscipline(studentId, disciplineId int) error {
	var group int
	errAccess := errors.New("access denied to discipline")
	query := fmt.Sprintf("SELECT %s.group_id FROM %s INNER JOIN %s ON %s.group_id = %s.group_id WHERE user_id = $1 AND discipline_id = $2",
		table_name.StudentsTable, table_name.StudentsTable, table_name.CurriculumTable, table_name.StudentsTable, table_name.CurriculumTable)
	err := r.db.Get(&group, query, studentId, disciplineId)
	if err != nil {
		return errAccess
	}
	if group == 0 {
		return errAccess
	}
	return nil
}

func (r *StudentStudyGuidePostgres) CheckAccessForLesson(studentId, lessonId int) error {
	errAccess := errors.New("access denied to lesson")
	var disciplines []int
	query := fmt.Sprintf("SELECT discipline_id FROM %s WHERE digital_material_id = $1", table_name.DigitalDisciplineTable)
	err := r.db.Select(&disciplines, query, lessonId)
	if err != nil {
		return errAccess
	}
	for _, value := range disciplines {
		if err := r.CheckAccessForDiscipline(studentId, value); err == nil {
			return nil
		}
	}
	return errAccess
}

func (r *StudentStudyGuidePostgres) CheckAccessForFile(studentId, fileId int) error {
	errAccess := errors.New("access denied to file")
	var lessonId int
	query := fmt.Sprintf("SELECT digital_material_id FROM %s WHERE id = $1", table_name.FilesPathTable)
	err := r.db.Get(&lessonId, query, fileId)
	if err != nil {
		return errAccess
	}
	if err := r.CheckAccessForLesson(studentId, lessonId); err != nil {
		return errAccess
	}
	return nil
}
