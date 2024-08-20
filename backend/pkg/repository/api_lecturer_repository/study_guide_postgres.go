package api_lecturer_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type LecturerStudyGuidePostgres struct {
	db *sqlx.DB
}

func NewLecturerStudyGuidePostgres(db *sqlx.DB) *LecturerStudyGuidePostgres {
	return &LecturerStudyGuidePostgres{db: db}
}

func (r *LecturerStudyGuidePostgres) AddStudyGuideHeader(name, nameEn, description, descriptionEn string) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, description, name_en, description_en) VALUES ($1, $2, $3, $4) RETURNING id", table_name.DigitalGuideTable)
	row := r.db.QueryRow(query, name, description, nameEn, descriptionEn)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *LecturerStudyGuidePostgres) GetStudyGuideHeader() ([]model.StudyGuideHeader, []model.StudyGuideHeader, error) {
	var digitalGuide []model.StudyGuideHeader
	query := fmt.Sprintf("SELECT id, name, description FROM %s ORDER BY id", table_name.DigitalGuideTable)
	err := r.db.Select(&digitalGuide, query)
	if err != nil {
		return nil, nil, err
	}
	var digitalGuideEn []model.StudyGuideHeader
	query = fmt.Sprintf("SELECT id, name, description FROM %s ORDER BY id", table_name.DigitalGuideTable)
	err = r.db.Select(&digitalGuideEn, query)
	if err != nil {
		return nil, nil, err
	}
	return digitalGuide, digitalGuideEn, nil
}

func (r *LecturerStudyGuidePostgres) DeleteStudyGuideHeader(digitalGuideId int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING id", table_name.DigitalGuideTable)
	row := r.db.QueryRow(query, digitalGuideId)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *LecturerStudyGuidePostgres) ChangeNameDigitalGuideHeader(digitalGuideId int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", table_name.DigitalGuideTable)
	_, err := r.db.Exec(query, name, digitalGuideId)
	return err
}

func (r *LecturerStudyGuidePostgres) ChangeNameDigitalGuideHeaderEn(digitalGuideId int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name_en = $1 WHERE id = $2", table_name.DigitalGuideTable)
	_, err := r.db.Exec(query, name, digitalGuideId)
	return err
}

func (r *LecturerStudyGuidePostgres) ChangeDescriptionDigitalGuideHeaderEn(digitalGuideId int, description string) error {
	query := fmt.Sprintf("UPDATE %s SET description_en = $1 WHERE id = $2", table_name.DigitalGuideTable)
	_, err := r.db.Exec(query, description, digitalGuideId)
	return err
}

func (r *LecturerStudyGuidePostgres) ChangeDescriptionDigitalGuideHeader(digitalGuideId int, description string) error {
	query := fmt.Sprintf("UPDATE %s SET description = $1 WHERE id = $2", table_name.DigitalGuideTable)
	_, err := r.db.Exec(query, description, digitalGuideId)
	return err
}

func (r *LecturerStudyGuidePostgres) GetDigitalDiscipline(disciplineId int) ([]model.DigitalDiscipline, error) {
	var digitalDiscipline []model.DigitalDiscipline
	query := fmt.Sprintf("SELECT * FROM %s WHERE discipline_id = $1 ORDER BY digital_material_id ", table_name.DigitalDisciplineTable)
	err := r.db.Select(&digitalDiscipline, query, disciplineId)
	if err != nil {
		return nil, err
	}
	return digitalDiscipline, nil
}

func (r *LecturerStudyGuidePostgres) AddDigitalDiscipline(digitalMaterialId, disciplineId int) error {
	query := fmt.Sprintf("INSERT INTO %s (digital_material_id, discipline_id) VALUES ($1, $2)", table_name.DigitalDisciplineTable)
	_, err := r.db.Exec(query, digitalMaterialId, disciplineId)
	return err
}

func (r *LecturerStudyGuidePostgres) DeleteDigitalDiscipline(digitalMaterialId, disciplineId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE digital_material_id = $1 AND discipline_id = $2", table_name.DigitalDisciplineTable)
	_, err := r.db.Exec(query, digitalMaterialId, disciplineId)
	return err
}

func (r *LecturerStudyGuidePostgres) GetFilesIdFromDigital(digitalId int) ([]model.FileId, error) {
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

func (r *LecturerStudyGuidePostgres) AddFileToDigital(path string, digitalId int) error {
	query := fmt.Sprintf("INSERT INTO %s (name, digital_material_id) VALUES($1, $2)", table_name.FilesPathTable)
	_, err := r.db.Exec(query, path, digitalId)
	return err
}

func (r *LecturerStudyGuidePostgres) DeleteFileFromDigital(fileId int) (string, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING name", table_name.FilesPathTable)
	var path string
	row := r.db.QueryRow(query, fileId)
	err := row.Scan(&path)
	if err != nil {
		return "", err
	}
	return path, nil
}

func (r *LecturerStudyGuidePostgres) GetFile(fileId int) (string, error) {
	var path string
	query := fmt.Sprintf("SELECT name FROM %s WHERE id = $1", table_name.FilesPathTable)
	err := r.db.Get(&path, query, fileId)
	if err != nil {
		return "", err
	}
	return path, nil
}
