package api_lecturer_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type LecturerSectionPostgres struct {
	db *sqlx.DB
}

func NewLecturerSectionPostgres(db *sqlx.DB) *LecturerSectionPostgres {
	return &LecturerSectionPostgres{db: db}
}

func (r *LecturerSectionPostgres) AddSection(name string, nameEn string, disciplineId int) error {
	query := fmt.Sprintf("INSERT INTO %s (name, name_en, discipline_id) VALUES ($1, $2, $3)", table_name.SectionTable)
	_, err := r.db.Exec(query, name, nameEn, disciplineId)
	return err
}

func (r *LecturerSectionPostgres) GetDisciplineSections(disciplineId int) ([]model.Section, []model.Section, error) {
	var sections []model.Section
	query := fmt.Sprintf("SELECT id, name FROM %s WHERE discipline_id = $1 ORDER BY id", table_name.SectionTable)
	err := r.db.Select(&sections, query, disciplineId)
	if err != nil {
		return nil, nil, err
	}
	var sectionsEn []model.Section
	query = fmt.Sprintf("SELECT id, name FROM %s WHERE discipline_id = $1 ORDER BY id", table_name.SectionTable)
	err = r.db.Select(&sectionsEn, query, disciplineId)
	if err != nil {
		return nil, nil, err
	}
	return sections, sectionsEn, err
}

func (r *LecturerSectionPostgres) DeleteSection(sectionId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", table_name.SectionTable)
	_, err := r.db.Exec(query, sectionId)
	return err
}

func (r *LecturerSectionPostgres) ChangeSectionName(sectionId int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", table_name.SectionTable)
	_, err := r.db.Exec(query, name, sectionId)
	return err
}

func (r *LecturerSectionPostgres) ChangeSectionNameEn(sectionId int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name_en = $1 WHERE id = $2", table_name.SectionTable)
	_, err := r.db.Exec(query, name, sectionId)
	return err
}

func (r *LecturerSectionPostgres) AddTestToSection(sectionId, testId int) error {
	query := fmt.Sprintf("INSERT INTO %s (section_id, test_id) VALUES ($1, $2)", table_name.SectionTests)
	_, err := r.db.Exec(query, sectionId, testId)
	return err
}

func (r *LecturerSectionPostgres) DeleteTestFromSection(sectionId, testId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE section_id = $1 AND test_id = $2", table_name.SectionTests)
	_, err := r.db.Exec(query, sectionId, testId)
	return err
}

func (r *LecturerSectionPostgres) AddLabToSection(labId, sectionId, defaultMark int) error {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (external_laboratory_id, default_mark) VALUES($1, $2) RETURNING id", table_name.LaboratoryTable)
	if err := r.db.QueryRow(query, labId, defaultMark).Scan(&id); err != nil {
		return err
	}
	query = fmt.Sprintf("INSERT INTO %s (laboratory_id, section_id) VALUES($1, $2)", table_name.SectionLabs)
	_, err := r.db.Exec(query, id, sectionId)
	return err
}

func (r *LecturerSectionPostgres) DeleteLabFromSection(labId, sectionId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE laboratory_id = $1 AND section_id = $2", table_name.SectionLabs)
	_, err := r.db.Exec(query, labId, sectionId)
	return err
}

func (r *LecturerSectionPostgres) GetLabFromSection(sectionId int) ([]model.LaboratoryWorkWithExternal, error) {
	var labs []model.LaboratoryWorkWithExternal

	query := fmt.Sprintf(`SELECT %s.id as laboratory_id,  %s.external_laboratory_id, %s.name, %s.task_description, %s.link, %s.default_mark 
								FROM %s
								JOIN %s
								ON %s.laboratory_id = %s.id
								JOIN %s
								ON %s.external_laboratory_id = %s.id
								WHERE %s.section_id = $1`,
		table_name.LaboratoryTable, table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.ExternalLaboratoryTable,
		table_name.ExternalLaboratoryTable, table_name.LaboratoryTable,
		table_name.SectionLabs, table_name.LaboratoryTable, table_name.SectionLabs, table_name.LaboratoryTable,
		table_name.ExternalLaboratoryTable, table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.SectionLabs)

	if err := r.db.Select(&labs, query, sectionId); err != nil {
		return nil, err
	}

	return labs, nil
}
