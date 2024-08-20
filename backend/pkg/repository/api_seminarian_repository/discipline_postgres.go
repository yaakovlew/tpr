package api_seminarian_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SeminarianDisciplinePostgres struct {
	db *sqlx.DB
}

func NewSeminarianDisciplinePostgres(db *sqlx.DB) *SeminarianDisciplinePostgres {
	return &SeminarianDisciplinePostgres{db: db}
}

func (r *SeminarianDisciplinePostgres) GetOwnDiscipline(userId int) ([]model.Discipline, []model.Discipline, error) {
	var disciplines []model.Discipline
	query := fmt.Sprintf("SELECT id, name FROM %s INNER JOIN %s ON %s.discipline_id = %s.id WHERE user_id = $1 GROUP BY id ORDER BY id",
		table_name.SeminariansGroupsTable, table_name.DisciplinesTable, table_name.SeminariansGroupsTable, table_name.DisciplinesTable)
	if err := r.db.Select(&disciplines, query, userId); err != nil {
		return nil, nil, err
	}
	var disciplinesEn []model.Discipline
	query = fmt.Sprintf("SELECT id, name_en as name FROM %s INNER JOIN %s ON %s.discipline_id = %s.id WHERE user_id = $1 GROUP BY id ORDER BY id",
		table_name.SeminariansGroupsTable, table_name.DisciplinesTable, table_name.SeminariansGroupsTable, table_name.DisciplinesTable)
	if err := r.db.Select(&disciplinesEn, query, userId); err != nil {
		return nil, nil, err
	}
	return disciplines, disciplinesEn, nil
}

func (r *SeminarianDisciplinePostgres) GetDisciplineSections(disciplineId int) ([]model.Section, []model.Section, error) {
	var sections []model.Section
	query := fmt.Sprintf("SELECT id, name FROM %s WHERE discipline_id = $1 ORDER BY id", table_name.SectionTable)
	err := r.db.Select(&sections, query, disciplineId)
	if err != nil {
		return nil, nil, err
	}
	var sectionsEn []model.Section
	query = fmt.Sprintf("SELECT id, name_en as name FROM %s WHERE discipline_id = $1 ORDER BY id", table_name.SectionTable)
	err = r.db.Select(&sectionsEn, query, disciplineId)
	if err != nil {
		return nil, nil, err
	}
	return sections, sectionsEn, err
}

func (r *SeminarianDisciplinePostgres) CheckAccessForDiscipline(seminarianId, disciplineId int) error {
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

func (r *SeminarianDisciplinePostgres) GetAllInfoAboutDiscipline(id int) (model.DisciplineInfoDoubleLang, model.DisciplineInfoDoubleLang, error) {
	var discipline model.DisciplineInfoDoubleLang
	query := fmt.Sprintf("SELECT id, name, seminar_visiting_mark, lesson_visiting_mark, exam_mark FROM %s WHERE id = $1", table_name.DisciplinesTable)
	if err := r.db.Get(&discipline, query, id); err != nil {
		return model.DisciplineInfoDoubleLang{}, model.DisciplineInfoDoubleLang{}, err
	}

	var disciplineEn model.DisciplineInfoDoubleLang
	query = fmt.Sprintf("SELECT id, name as name, seminar_visiting_mark, lesson_visiting_mark, exam_mark FROM %s WHERE id = $1", table_name.DisciplinesTable)
	if err := r.db.Get(&disciplineEn, query, id); err != nil {
		return model.DisciplineInfoDoubleLang{}, model.DisciplineInfoDoubleLang{}, err
	}

	return discipline, disciplineEn, nil
}
