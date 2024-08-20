package api_seminarian_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SeminarianStudyGuidePostgres struct {
	db *sqlx.DB
}

func NewSeminarianStudyGuidePostgres(db *sqlx.DB) *SeminarianStudyGuidePostgres {
	return &SeminarianStudyGuidePostgres{db: db}
}

func (r *SeminarianStudyGuidePostgres) GetDigitalDiscipline(disciplineId int) ([]model.DigitalDisciplineWithInfo, []model.DigitalDisciplineWithInfo, error) {
	var digitalDiscipline []model.DigitalDisciplineWithInfo
	query := fmt.Sprintf("SELECT %s.name, %s.description, %s.id AS digital_material_id FROM %s  INNER JOIN %s ON digital_material_id = id WHERE discipline_id = $1 ORDER BY digital_material_id",
		table_name.DigitalGuideTable, table_name.DigitalGuideTable, table_name.DigitalGuideTable, table_name.DigitalGuideTable, table_name.DigitalDisciplineTable)
	err := r.db.Select(&digitalDiscipline, query, disciplineId)
	if err != nil {
		return nil, nil, err
	}
	var digitalDisciplineEn []model.DigitalDisciplineWithInfo
	query = fmt.Sprintf("SELECT %s.name_en as name, %s.description_en as name, %s.id AS digital_material_id FROM %s  INNER JOIN %s ON digital_material_id = id WHERE discipline_id = $1 ORDER BY digital_material_id",
		table_name.DigitalGuideTable, table_name.DigitalGuideTable, table_name.DigitalGuideTable, table_name.DigitalGuideTable, table_name.DigitalDisciplineTable)
	err = r.db.Select(&digitalDisciplineEn, query, disciplineId)
	if err != nil {
		return nil, nil, err
	}
	return digitalDiscipline, digitalDisciplineEn, nil
}

func (r *SeminarianStudyGuidePostgres) GetFile(fileId int) (string, error) {
	var path string
	query := fmt.Sprintf("SELECT name FROM %s WHERE id = $1 ORDER BY id", table_name.FilesPathTable)
	err := r.db.Get(&path, query, fileId)
	if err != nil {
		return "", err
	}
	return path, nil
}

func (r *SeminarianStudyGuidePostgres) GetFilesIdFromDigital(digitalId int) ([]model.FileId, error) {
	var files []model.FileId
	query := fmt.Sprintf("SELECT %s.id, %s.name FROM %s INNER JOIN %s ON %s.digital_material_id = %s.id WHERE %s.id = $1 ORDER BY %s.id",
		table_name.FilesPathTable, table_name.FilesPathTable, table_name.FilesPathTable, table_name.DigitalGuideTable, table_name.FilesPathTable,
		table_name.DigitalGuideTable, table_name.DigitalGuideTable, table_name.FilesPathTable)
	err := r.db.Select(&files, query, digitalId)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (r *SeminarianStudyGuidePostgres) CheckAccessForDiscipline(seminarianId, disciplineId int) error {
	var discipline []int
	errAccess := errors.New("access denied to discipline")
	query := fmt.Sprintf("SELECT discipline_id FROM %s WHERE user_id = $1 AND discipline_id = $2", table_name.SeminariansGroupsTable)
	err := r.db.Select(&discipline, query, seminarianId, disciplineId)
	if err != nil {
		return errAccess
	}
	if len(discipline) == 0 {
		return errAccess
	}
	return nil
}

func (r *SeminarianStudyGuidePostgres) CheckAccessForLesson(seminarianId, lessonId int) error {
	errAccess := errors.New("access denied to lesson")
	var disciplines []int
	query := fmt.Sprintf("SELECT discipline_id FROM %s WHERE digital_material_id = $1", table_name.DigitalDisciplineTable)
	err := r.db.Select(&disciplines, query, lessonId)
	if err != nil {
		return errAccess
	}
	for _, value := range disciplines {
		if err := r.CheckAccessForDiscipline(seminarianId, value); err == nil {
			return nil
		}
	}
	return errAccess
}

func (r *SeminarianStudyGuidePostgres) CheckAccessForFile(seminarianId, fileId int) error {
	errAccess := errors.New("access denied to file")
	var lessonId int
	query := fmt.Sprintf("SELECT digital_material_id FROM %s WHERE id = $1", table_name.FilesPathTable)
	err := r.db.Get(&lessonId, query, fileId)
	if err != nil {
		return errAccess
	}
	if err := r.CheckAccessForLesson(seminarianId, lessonId); err != nil {
		return errAccess
	}
	return nil
}
