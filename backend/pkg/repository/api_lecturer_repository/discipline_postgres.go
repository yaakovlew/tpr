package api_lecturer_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type LecturerDisciplinePostgres struct {
	db *sqlx.DB
}

func NewLecturerDisciplinePostgres(db *sqlx.DB) *LecturerDisciplinePostgres {
	return &LecturerDisciplinePostgres{db: db}
}

func (r *LecturerDisciplinePostgres) GetGroupForDiscipline(id int) ([]model.Group, error) {
	var groups []model.Group
	query := fmt.Sprintf("SELECT %s.id, %s.name FROM %s INNER JOIN %s ON %s.id = %s.group_id WHERE discipline_id = $1 AND %s.is_archive = false ORDER BY %s.id",
		table_name.GroupsTable, table_name.GroupsTable, table_name.GroupsTable, table_name.CurriculumTable, table_name.GroupsTable, table_name.CurriculumTable, table_name.CurriculumTable, table_name.GroupsTable)
	err := r.db.Select(&groups, query, id)
	if err != nil {
		return []model.Group{}, err
	}
	return groups, nil
}

func removeElement(arr []model.Group, index int) []model.Group {
	copy(arr[index:], arr[index+1:])
	return arr[:len(arr)-1]
}

func findElement(allGroups, groups []model.Group) []model.Group {
	for _, group := range groups {
		id := group.Id
		for i, v := range allGroups {
			if v.Id == id {
				allGroups = removeElement(allGroups, i)
			}
		}
	}
	return allGroups
}

func (r *LecturerDisciplinePostgres) GetGroupsAvailableToAddForDiscipline(id int) ([]model.Group, error) {
	var groups []model.Group
	query := fmt.Sprintf("SELECT %s.id, %s.name FROM %s INNER JOIN %s ON %s.id = %s.group_id WHERE discipline_id = $1 AND %s.is_archive = false ORDER BY %s.id",
		table_name.GroupsTable, table_name.GroupsTable, table_name.GroupsTable, table_name.CurriculumTable, table_name.GroupsTable, table_name.CurriculumTable,
		table_name.CurriculumTable, table_name.GroupsTable)
	err := r.db.Select(&groups, query, id)
	if err != nil {
		return []model.Group{}, err
	}
	var allGroups []model.Group
	query = fmt.Sprintf("SELECT id, name FROM %s ORDER BY id", table_name.GroupsTable)
	err = r.db.Select(&allGroups, query)
	if err != nil {
		return []model.Group{}, err
	}
	return findElement(allGroups, groups), nil
}

func (r *LecturerDisciplinePostgres) AddDiscipline(discipline model.AddNewDiscipline) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, name_en, seminar_visiting_mark, lesson_visiting_mark, exam_mark) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		table_name.DisciplinesTable)
	if err := r.db.QueryRow(query, discipline.Name, discipline.NameEn, discipline.SeminarVisitingMark, discipline.LessonVisitingMark, discipline.ExamMark).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *LecturerDisciplinePostgres) ChangeSeminarMarks(id, seminarMarks int) error {
	query := fmt.Sprintf("UPDATE %s SET seminar_visiting_mark = $1 WHERE id = $2", table_name.DisciplinesTable)
	_, err := r.db.Exec(query, seminarMarks, id)
	return err
}

func (r *LecturerDisciplinePostgres) ChangeLessonMarks(id, lessonMarks int) error {
	query := fmt.Sprintf("UPDATE %s SET lesson_visiting_mark = $1 WHERE id = $2", table_name.DisciplinesTable)
	_, err := r.db.Exec(query, lessonMarks, id)
	return err
}

func (r *LecturerDisciplinePostgres) ChangeExamMark(id, examMark int) error {
	query := fmt.Sprintf("UPDATE %s SET exam_mark = $1 WHERE id = $2", table_name.DisciplinesTable)
	_, err := r.db.Exec(query, examMark, id)
	return err
}

func (r *LecturerDisciplinePostgres) GetAllDisciplines() ([]model.Discipline, []model.Discipline, error) {
	var disciplines []model.Discipline
	query := fmt.Sprintf("SELECT id, name FROM %s ORDER BY id", table_name.DisciplinesTable)
	if err := r.db.Select(&disciplines, query); err != nil {
		return nil, nil, err
	}
	var disciplinesEn []model.Discipline
	query = fmt.Sprintf("SELECT id, name_en as name FROM %s ORDER BY id", table_name.DisciplinesTable)
	if err := r.db.Select(&disciplinesEn, query); err != nil {
		return nil, nil, err
	}
	return disciplines, disciplinesEn, nil
}

func (r *LecturerDisciplinePostgres) GetAllInfoAboutDiscipline(id int) (model.DisciplineInfo, error) {
	var discipline model.DisciplineInfo
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", table_name.DisciplinesTable)
	if err := r.db.Get(&discipline, query, id); err != nil {
		return model.DisciplineInfo{}, err
	}
	return discipline, nil
}

func (r *LecturerDisciplinePostgres) DeleteDiscipline(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", table_name.DisciplinesTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *LecturerDisciplinePostgres) AddGroupToDiscipline(groupId, disciplineId int) error {
	queryToFineGroup := fmt.Sprintf("SELECT COUNT(*) AS count FROM %s WHERE discipline_id = $1 AND group_id = $2",
		table_name.CurriculumTable)
	var count int
	if err := r.db.QueryRow(queryToFineGroup, disciplineId, groupId).Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		queryToUpdate := fmt.Sprintf("UPDATE %s SET is_archive = false WHERE discipline_id = $1 AND group_id = $2", table_name.CurriculumTable)
		if _, err := r.db.Exec(queryToUpdate, disciplineId, groupId); err != nil {
			return err
		}
		return nil
	} else {
		query := fmt.Sprintf("INSERT INTO %s (discipline_id, group_id, is_archive) VALUES ($1, $2, false)", table_name.CurriculumTable)
		_, err := r.db.Exec(query, disciplineId, groupId)
		return err
	}
}

func (r *LecturerDisciplinePostgres) ArchiveGroupToDiscipline(groupId, disciplineId int) error {
	query := fmt.Sprintf("UPDATE %s SET is_archive = true WHERE discipline_id = $1 AND group_id = $2", table_name.CurriculumTable)
	_, err := r.db.Exec(query, disciplineId, groupId)
	return err
}

func (r *LecturerDisciplinePostgres) DeleteGroupFromDiscipline(groupId, disciplineId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE group_id = $1 AND discipline_id = $2", table_name.CurriculumTable)
	_, err := r.db.Exec(query, groupId, disciplineId)
	return err
}

func (r *LecturerDisciplinePostgres) ChangeDiscipline(disciplineId int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", table_name.DisciplinesTable)
	_, err := r.db.Exec(query, name, disciplineId)
	return err
}

func (r *LecturerDisciplinePostgres) ChangeDisciplineEn(disciplineId int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name_en = $1 WHERE id = $2", table_name.DisciplinesTable)
	_, err := r.db.Exec(query, name, disciplineId)
	return err
}
