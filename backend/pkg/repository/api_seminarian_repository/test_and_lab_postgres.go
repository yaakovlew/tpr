package api_seminarian_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SeminarianTestAndLabPostgres struct {
	db *sqlx.DB
}

func NewSeminarianTestAndLabPostgres(db *sqlx.DB) *SeminarianTestAndLabPostgres {
	return &SeminarianTestAndLabPostgres{db: db}
}

func (r *SeminarianTestAndLabPostgres) GetAllTestFromSection(sectionId int) ([]model.Test, []model.Test, error) {
	var tests []model.Test
	query := fmt.Sprintf("SELECT id, name, task_description, minutes_duration, default_mark  FROM %s INNER JOIN %s ON %s.id = %s.test_id WHERE section_id = $1 ORDER BY id",
		table_name.TestsTable, table_name.SectionTests, table_name.TestsTable, table_name.SectionTests)
	err := r.db.Select(&tests, query, sectionId)
	if err != nil {
		return nil, nil, err
	}
	var testsEn []model.Test
	query = fmt.Sprintf("SELECT id, name_en as name, task_description_en as task_description, minutes_duration, default_mark  FROM %s INNER JOIN %s ON %s.id = %s.test_id WHERE section_id = $1 ORDER BY id",
		table_name.TestsTable, table_name.SectionTests, table_name.TestsTable, table_name.SectionTests)
	err = r.db.Select(&testsEn, query, sectionId)
	if err != nil {
		return nil, nil, err
	}
	return tests, testsEn, nil
}

func (r *SeminarianTestAndLabPostgres) GetAllLab(sectionId int) ([]model.LaboratoryWork, []model.LaboratoryWork, error) {
	var laboratory []model.LaboratoryWork
	query := fmt.Sprintf(`SELECT %s.id, %s.name, %s.task_description, %s.link, %s.default_mark
								FROM %s
								INNER JOIN %s
								ON %s.id = %s.laboratory_id
								INNER JOIN %s
								ON %s.external_laboratory_id = %s.id
								WHERE section_id = $1 ORDER BY %s.id`,
		table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.ExternalLaboratoryTable, table_name.ExternalLaboratoryTable, table_name.LaboratoryTable,
		table_name.LaboratoryTable, table_name.SectionLabs, table_name.LaboratoryTable, table_name.SectionLabs,
		table_name.ExternalLaboratoryTable, table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.LaboratoryTable)
	err := r.db.Select(&laboratory, query, sectionId)
	if err != nil {
		return nil, nil, err
	}
	var laboratoryEn []model.LaboratoryWork
	query = fmt.Sprintf(`SELECT %s.id, %s.name_en as name, %s.task_description_en as task_description , %s.link, %s.default_mark
								FROM %s
								INNER JOIN %s
								ON %s.id = %s.laboratory_id
								INNER JOIN %s
								ON %s.external_laboratory_id = %s.id
								WHERE section_id = $1 ORDER BY %s.id`,
		table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.ExternalLaboratoryTable, table_name.ExternalLaboratoryTable, table_name.LaboratoryTable,
		table_name.LaboratoryTable, table_name.SectionLabs, table_name.LaboratoryTable, table_name.SectionLabs,
		table_name.ExternalLaboratoryTable, table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.LaboratoryTable)
	err = r.db.Select(&laboratoryEn, query, sectionId)
	if err != nil {
		return nil, nil, err
	}
	return laboratory, laboratoryEn, nil
}

func (r *SeminarianTestAndLabPostgres) CreateTestDate(userId, testId, date int) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, test_id, closed_date) VALUES($1, $2, $3)", table_name.TestsDateTable)
	_, err := r.db.Exec(query, userId, testId, date)
	return err
}

func (r *SeminarianTestAndLabPostgres) UpdateTestDate(userId, testId, date int) error {
	query := fmt.Sprintf("UPDATE %s SET closed_date = $1 WHERE user_id = $2 AND test_id = $3", table_name.TestsDateTable)
	_, err := r.db.Exec(query, date, userId, testId)
	return err
}

func (r *SeminarianTestAndLabPostgres) CheckExistTestDate(userId, testId int) bool {
	query := fmt.Sprintf("SELECT closed_date FROM %s WHERE user_id = $1 AND test_id = $2", table_name.TestsDateTable)
	row := r.db.QueryRow(query, userId, testId)
	var date int
	if err := row.Scan(&date); err != nil {
		return false
	}
	return true
}

func (r *SeminarianTestAndLabPostgres) CheckAccessForDiscipline(seminarianId, disciplineId int) error {
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

func (r *SeminarianTestAndLabPostgres) CheckAccessForSection(seminarianId, sectionId int) error {
	var disciplineId int
	errAccess := errors.New("access denied to section")
	query := fmt.Sprintf("SELECT discipline_id FROM %s WHERE id = $1", table_name.SectionTable)
	if err := r.db.Get(&disciplineId, query, sectionId); err != nil {
		return errAccess
	}
	if err := r.CheckAccessForDiscipline(seminarianId, disciplineId); err != nil {
		return errAccess
	}
	return nil
}

func (r *SeminarianTestAndLabPostgres) CheckAccessToOpenTest(seminarianId, userId, testId int) error {
	var disciplineId []int
	errAccess := errors.New("access denied to open test")
	query := fmt.Sprintf("SELECT discipline_id FROM %s INNER JOIN %s ON %s.section_id = %s.id INNER JOIN %s ON %s.discipline_id = %s.id WHERE test_id = $1",
		table_name.SectionTests, table_name.SectionTable, table_name.SectionTests, table_name.SectionTable, table_name.DisciplinesTable, table_name.SectionTable, table_name.DisciplinesTable)
	if err := r.db.Select(&disciplineId, query, testId); err != nil {
		return errAccess
	}
	for _, value := range disciplineId {
		if err := r.CheckAccessForStudent(seminarianId, userId, value); err == nil {
			return nil
		}
	}
	return errAccess
}

func (r *SeminarianTestAndLabPostgres) CheckAccessToOpenLab(seminarianId, userId, labId int) error {
	var disciplineId []int
	errAccess := errors.New("access denied to open test")
	query := fmt.Sprintf("SELECT discipline_id FROM %s INNER JOIN %s ON %s.section_id = %s.id INNER JOIN %s ON %s.discipline_id = %s.id WHERE laboratory_id = $1",
		table_name.SectionLabs, table_name.SectionTable, table_name.SectionLabs, table_name.SectionTable, table_name.DisciplinesTable, table_name.SectionTable, table_name.DisciplinesTable)
	if err := r.db.Select(&disciplineId, query, labId); err != nil {
		return errAccess
	}
	for _, value := range disciplineId {
		if err := r.CheckAccessForStudent(seminarianId, userId, value); err == nil {
			return nil
		}
	}
	return errAccess
}

func (r *SeminarianTestAndLabPostgres) CheckAccessForStudent(seminarianId, userId, disciplineId int) error {
	if err := r.CheckAccessForDiscipline(seminarianId, disciplineId); err != nil {
		return err
	}
	var groupId int
	errAccess := errors.New("access denied to student")
	query := fmt.Sprintf("SELECT %s.group_id FROM %s INNER JOIN %s ON %s.group_id = %s.group_id WHERE %s.user_id = $1 AND discipline_id = $2 AND %s.is_archive = false",
		table_name.StudentsTable, table_name.StudentsTable, table_name.CurriculumTable, table_name.StudentsTable, table_name.CurriculumTable, table_name.StudentsTable, table_name.CurriculumTable)
	err := r.db.Get(&groupId, query, userId, disciplineId)
	if err != nil {
		return errAccess
	}
	if err := r.CheckAccessForGroup(seminarianId, groupId, disciplineId); err != nil {
		return errAccess
	}
	return nil
}

func (r *SeminarianTestAndLabPostgres) CheckAccessForGroup(seminarianId, groupId, disciplineId int) error {
	query := fmt.Sprintf("SELECT %s.group_id FROM %s INNER JOIN %s ON %s.group_id = %s.group_id WHERE %s.user_id = $1 AND %s.group_id = $2 AND %s.discipline_id = $3 AND %s.is_archive = false",
		table_name.SeminariansGroupsTable, table_name.SeminariansGroupsTable,
		table_name.CurriculumTable, table_name.CurriculumTable,
		table_name.SeminariansGroupsTable, table_name.SeminariansGroupsTable,
		table_name.SeminariansGroupsTable, table_name.SeminariansGroupsTable, table_name.CurriculumTable)
	var group int
	err := r.db.Get(&group, query, seminarianId, groupId, disciplineId)
	if err != nil {
		errorAccess := errors.New("access denied to group")
		return errorAccess
	}
	return nil
}

func (r *SeminarianTestAndLabPostgres) checkExistTestDate(userId, testId int) bool {
	query := fmt.Sprintf("SELECT closed_date FROM %s WHERE user_id = $1 AND test_id = $2", table_name.TestsDateTable)
	row := r.db.QueryRow(query, userId, testId)
	var date int
	if err := row.Scan(&date); err != nil {
		return false
	}
	return true
}

func (r *SeminarianTestAndLabPostgres) checkExistLabDate(userId, labId int) bool {
	query := fmt.Sprintf("SELECT closed_date FROM %s WHERE user_id = $1 AND laboratory_id = $2", table_name.LaboratoryDateTable)
	row := r.db.QueryRow(query, userId, labId)
	var date int
	if err := row.Scan(&date); err != nil {
		return false
	}
	return true
}

func (r *SeminarianTestAndLabPostgres) OpenTestForStudent(studentId, testId int, date int64) error {
	if r.checkExistTestDate(studentId, testId) {
		query := fmt.Sprintf("UPDATE %s SET closed_date = $1, is_done = false WHERE user_id = $2 AND test_id = $3", table_name.TestsDateTable)
		if _, err := r.db.Exec(query, date, studentId, testId); err != nil {
			return err
		}
	} else {
		query := fmt.Sprintf("INSERT INTO %s (user_id, test_id, closed_date, is_done) VALUES ($1, $2, $3, false)", table_name.TestsDateTable)
		if _, err := r.db.Exec(query, studentId, testId, date); err != nil {
			return err
		}
	}

	if !r.checkExistTestMark(studentId, testId) {
		query := fmt.Sprintf("INSERT INTO %s (user_id, test_id, mark) VALUES ($1, $2, $3)", table_name.TestsMarkTable)
		if _, err := r.db.Exec(query, studentId, testId, 0); err != nil {
			return err
		}
	}

	return nil
}

func (r *SeminarianTestAndLabPostgres) checkExistTestMark(studentId, testId int) bool {
	var userId int
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE user_id = $1 AND test_id = $2", table_name.TestsMarkTable)
	if err := r.db.Get(&userId, query, studentId, testId); err != nil {
		return false
	}

	if userId == 0 {
		return false
	}

	return true
}

func (r *SeminarianTestAndLabPostgres) GetOpenedTestForStudent(studentId, testId int) (model.OpenedTest, error) {
	var test model.OpenedTest
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND test_id = $2", table_name.TestsDateTable)
	if err := r.db.Get(&test, query, studentId, testId); err != nil {
		return model.OpenedTest{}, err
	}

	return test, nil
}

func (r *SeminarianTestAndLabPostgres) CloseOpenedTestForStudent(studentId, testId int) error {
	query := fmt.Sprintf("UPDATE %s SET is_done = true WHERE user_id = $1 AND test_id = $2", table_name.TestsDateTable)
	if _, err := r.db.Exec(query, studentId, testId); err != nil {
		return err
	}

	return nil
}

func (r *SeminarianTestAndLabPostgres) GetUsersWithOpenedTest(seminarianId, testId int, timeNow int64) ([]model.StudentWithGroupWithClosedDate, error) {
	var students []model.StudentWithGroupWithClosedDate
	query := fmt.Sprintf(`SELECT %s.id, %s.name, %s.surname, %s.name AS group_name, closed_date, minutes_duration
		FROM %s 
		INNER JOIN %s 
		ON %s.user_id = %s.id
		INNER JOIN %s
		ON %s.test_id = %s.id
		INNER JOIN %s
		ON %s.user_id = %s.id
		INNER JOIN %s
		ON %s.group_id = %s.group_id
		INNER JOIN %s
		ON %s.test_id = %s.id
		INNER JOIN %s
		ON %s.id = %s.section_id 
		INNER JOIN %s
		ON %s.id = %s.discipline_id 
		WHERE %s.id = $1 AND is_done = false AND %s.user_id = $2 AND (closed_date >= $3 OR closed_date = 0)
		ORDER BY closed_date DESC`,
		table_name.UsersTable, table_name.UsersTable, table_name.UsersTable,
		table_name.UsersTable, table_name.UsersTable, table_name.TestsDateTable,
		table_name.TestsDateTable, table_name.UsersTable, table_name.TestsTable,
		table_name.TestsDateTable, table_name.TestsTable,
		table_name.StudentsTable, table_name.StudentsTable, table_name.UsersTable,
		table_name.SeminariansGroupsTable, table_name.StudentsTable, table_name.SeminariansGroupsTable,
		table_name.SectionTests, table_name.SectionTests, table_name.TestsTable,
		table_name.SectionTable, table_name.SectionTable, table_name.SectionTests,
		table_name.DisciplinesTable, table_name.DisciplinesTable, table_name.SectionTable,
		table_name.TestsTable, table_name.SeminariansGroupsTable)
	err := r.db.Select(&students, query, testId, seminarianId, timeNow)
	if err != nil {
		return []model.StudentWithGroupWithClosedDate{}, err
	}
	return students, nil
}

func (r *SeminarianTestAndLabPostgres) GetUsersWithDoneTest(seminarianId, testId int, timeNow int64) ([]model.StudentWithGroupWithClosedDate, error) {
	var students []model.StudentWithGroupWithClosedDate
	query := fmt.Sprintf(`SELECT %s.id, %s.name, %s.surname, %s.name AS group_name, closed_date, minutes_duration
		FROM %s 
		INNER JOIN %s 
		ON %s.user_id = %s.id
		INNER JOIN %s
		ON %s.test_id = %s.id
		INNER JOIN %s
		ON %s.user_id = %s.id
		INNER JOIN %s
		ON %s.group_id = %s.group_id
		INNER JOIN %s
		ON %s.test_id = %s.id
		INNER JOIN %s
		ON %s.id = %s.section_id 
		INNER JOIN %s
		ON %s.id = %s.discipline_id 
		WHERE ((%s.id = $1 AND %s.user_id = $2 AND is_done = true) OR (closed_date < $3 AND closed_date != 0))
		ORDER BY closed_date DESC`,
		table_name.UsersTable, table_name.UsersTable, table_name.UsersTable,
		table_name.UsersTable, table_name.UsersTable, table_name.TestsDateTable,
		table_name.TestsDateTable, table_name.UsersTable, table_name.TestsTable,
		table_name.TestsDateTable, table_name.TestsTable,
		table_name.StudentsTable, table_name.StudentsTable, table_name.UsersTable,
		table_name.SeminariansGroupsTable, table_name.StudentsTable, table_name.SeminariansGroupsTable,
		table_name.SectionTests, table_name.SectionTests, table_name.TestsTable,
		table_name.SectionTable, table_name.SectionTable, table_name.SectionTests,
		table_name.DisciplinesTable, table_name.DisciplinesTable, table_name.SectionTable,
		table_name.TestsTable, table_name.SeminariansGroupsTable)
	err := r.db.Select(&students, query, testId, seminarianId, timeNow)
	if err != nil {
		return []model.StudentWithGroupWithClosedDate{}, err
	}
	return students, nil
}

func (r *SeminarianTestAndLabPostgres) GetOpenedLabForStudent(studentId, labId int) (model.OpenedLab, error) {
	var lab model.OpenedLab
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND laboratory_id = $2", table_name.LaboratoryDateTable)
	if err := r.db.Get(&lab, query, studentId, labId); err != nil {
		return model.OpenedLab{}, err
	}

	return lab, nil
}

func (r *SeminarianTestAndLabPostgres) CloseOpenedLabForStudent(studentId, labId int) error {
	query := fmt.Sprintf("UPDATE %s SET is_done = true WHERE user_id = $1 AND laboratory_id = $2", table_name.LaboratoryDateTable)
	if _, err := r.db.Exec(query, studentId, labId); err != nil {
		return err
	}

	return nil
}

func (r *SeminarianTestAndLabPostgres) GetUsersWithOpenedLab(seminarianId, labId, groupId int, timeNow int64) ([]model.StudentWithGroupWithClosedDate, error) {
	var students []model.StudentWithGroupWithClosedDate
	query := fmt.Sprintf(`SELECT %s.id, %s.name, %s.surname, %s.name AS group_name, %s.closed_date
		FROM %s
		JOIN %s 
		ON %s.group_id = %s.group_id
		JOIN %s
		ON %s.group_id = %s.id
		JOIN %s
		ON %s.id = %s.user_id
		JOIN %s
		ON %s.user_id = %s.user_id
		WHERE %s.laboratory_id = $1 AND is_done = false AND %s.user_id = $2 AND (closed_date >= $3 OR closed_date = 0) AND %s.group_id = $4
		ORDER BY closed_date DESC`,
		table_name.UsersTable, table_name.UsersTable, table_name.UsersTable, table_name.GroupsTable, table_name.LaboratoryDateTable,
		table_name.SeminariansGroupsTable, table_name.StudentsTable, table_name.SeminariansGroupsTable, table_name.StudentsTable,
		table_name.GroupsTable, table_name.SeminariansGroupsTable, table_name.GroupsTable,
		table_name.UsersTable, table_name.UsersTable, table_name.StudentsTable,
		table_name.LaboratoryDateTable, table_name.StudentsTable, table_name.LaboratoryDateTable,
		table_name.LaboratoryDateTable, table_name.SeminariansGroupsTable, table_name.SeminariansGroupsTable)
	err := r.db.Select(&students, query, labId, seminarianId, timeNow, groupId)
	if err != nil {
		return []model.StudentWithGroupWithClosedDate{}, err
	}
	return students, nil
}

func (r *SeminarianTestAndLabPostgres) GetInternalLabInfo(labId int) (int, error) {
	query := fmt.Sprintf("SELECT external_laboratory_id FROM %s WHERE id = $1", table_name.LaboratoryTable)
	var id int
	if err := r.db.Get(&id, query, labId); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *SeminarianTestAndLabPostgres) GetUsersWithDoneLab(seminarianId, labId, groupId int, timeNow int64) ([]model.StudentWithGroupWithClosedDate, error) {
	var students []model.StudentWithGroupWithClosedDate
	query := fmt.Sprintf(`SELECT %s.id, %s.name, %s.surname, %s.name AS group_name, %s.closed_date
		FROM %s
		JOIN %s 
		ON %s.group_id = %s.group_id
		JOIN %s
		ON %s.group_id = %s.id
		JOIN %s
		ON %s.id = %s.user_id
		JOIN %s
		ON %s.user_id = %s.user_id
		WHERE (%s.laboratory_id = $1 AND %s.user_id = $2) AND (is_done = true OR (closed_date < $3 AND closed_date != 0)) AND %s.group_id = $4
		ORDER BY closed_date DESC`,
		table_name.UsersTable, table_name.UsersTable, table_name.UsersTable, table_name.GroupsTable, table_name.LaboratoryDateTable,
		table_name.SeminariansGroupsTable, table_name.StudentsTable, table_name.SeminariansGroupsTable, table_name.StudentsTable,
		table_name.GroupsTable, table_name.SeminariansGroupsTable, table_name.GroupsTable,
		table_name.UsersTable, table_name.UsersTable, table_name.StudentsTable,
		table_name.LaboratoryDateTable, table_name.StudentsTable, table_name.LaboratoryDateTable,
		table_name.LaboratoryDateTable, table_name.SeminariansGroupsTable, table_name.SeminariansGroupsTable)
	err := r.db.Select(&students, query, labId, seminarianId, timeNow, groupId)
	if err != nil {
		return []model.StudentWithGroupWithClosedDate{}, err
	}
	return students, nil
}

func (r *SeminarianTestAndLabPostgres) OpenLabForStudent(studentId, labId int, date int64) error {
	if r.checkExistLabDate(studentId, labId) {
		query := fmt.Sprintf("UPDATE %s SET closed_date = $1, is_done = false WHERE user_id = $2 AND laboratory_id = $3", table_name.LaboratoryDateTable)
		if _, err := r.db.Exec(query, date, studentId, labId); err != nil {
			return err
		}
	} else {
		query := fmt.Sprintf("INSERT INTO %s (user_id, laboratory_id, closed_date, is_done) VALUES ($1, $2, $3, false)", table_name.LaboratoryDateTable)
		if _, err := r.db.Exec(query, studentId, labId, date); err != nil {
			return err
		}
	}

	if !r.checkExistLabMark(studentId, labId) {
		query := fmt.Sprintf("INSERT INTO %s (user_id, laboratory_id, mark) VALUES ($1, $2, $3)", table_name.LaboratoryMarkTable)
		if _, err := r.db.Exec(query, studentId, labId, 0); err != nil {
			return err
		}
	}

	return nil
}

func (r *SeminarianTestAndLabPostgres) checkExistLabMark(studentId, labId int) bool {
	var userId int
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE user_id = $1 AND laboratory_id = $2", table_name.LaboratoryMarkTable)
	if err := r.db.Get(&userId, query, studentId, labId); err != nil {
		return false
	}

	if userId == 0 {
		return false
	}

	return true
}

func (r *SeminarianTestAndLabPostgres) GetLabToken(labId int) (string, error) {
	query := fmt.Sprintf("SELECT token FROM %s WHERE id = $1", table_name.ExternalLaboratoryTable)
	var token string
	if err := r.db.Get(&token, query, labId); err != nil {
		return "", err
	}

	return token, nil
}

func (r *SeminarianTestAndLabPostgres) GetExternalLabInfo(labId int) (model.LaboratoryWorkResponse, error) {
	query := fmt.Sprintf("SELECT id, name, name_en, task_description, task_description_en, link FROM %s WHERE id = $1",
		table_name.ExternalLaboratoryTable)
	var lab model.LaboratoryWorkResponse
	if err := r.db.Get(&lab, query, labId); err != nil {
		return model.LaboratoryWorkResponse{}, err
	}

	return lab, nil
}
