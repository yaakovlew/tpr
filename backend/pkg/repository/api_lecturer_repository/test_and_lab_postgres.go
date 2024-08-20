package api_lecturer_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type LecturerTestAndLabPostgres struct {
	db *sqlx.DB
}

func NewLecturerTestAndLabPostgres(db *sqlx.DB) *LecturerTestAndLabPostgres {
	return &LecturerTestAndLabPostgres{
		db: db,
	}
}

func (r *LecturerTestAndLabPostgres) GetAllTests() ([]model.Test, []model.Test, error) {
	var tests []model.Test
	query := fmt.Sprintf("SELECT id, name, task_description, minutes_duration, default_mark FROM %s ORDER BY id", table_name.TestsTable)
	err := r.db.Select(&tests, query)
	if err != nil {
		return nil, nil, err
	}
	var testsEn []model.Test
	query = fmt.Sprintf("SELECT id, name_en as name, task_description_en as name, minutes_duration, default_mark  FROM %s ORDER BY id", table_name.TestsTable)
	err = r.db.Select(&testsEn, query)
	if err != nil {
		return nil, nil, err
	}
	return tests, testsEn, nil
}

func (r *LecturerTestAndLabPostgres) GetAllTestFromSection(sectionId int) ([]model.Test, []model.Test, error) {
	var tests []model.Test
	query := fmt.Sprintf("SELECT id, name, task_description, minutes_duration, default_mark  FROM %s INNER JOIN %s ON %s.id = %s.test_id WHERE section_id = $1 ORDER BY %s.id",
		table_name.TestsTable, table_name.SectionTests, table_name.TestsTable, table_name.SectionTests, table_name.TestsTable)
	err := r.db.Select(&tests, query, sectionId)
	if err != nil {
		return nil, nil, err
	}
	var testsEn []model.Test
	query = fmt.Sprintf("SELECT id, name_en as name, task_description_en as name, minutes_duration, default_mark  FROM %s INNER JOIN %s ON %s.id = %s.test_id WHERE section_id = $1 ORDER BY %s.id",
		table_name.TestsTable, table_name.SectionTests, table_name.TestsTable, table_name.SectionTests, table_name.TestsTable)
	err = r.db.Select(&testsEn, query, sectionId)
	if err != nil {
		return nil, nil, err
	}
	return tests, testsEn, nil
}

func (r *LecturerTestAndLabPostgres) GetAllExternalLab() ([]model.CommonLaboratoryWork, []model.CommonLaboratoryWork, error) {
	var labs []model.CommonLaboratoryWork
	query := fmt.Sprintf("SELECT id, name, task_description, link  FROM %s ORDER BY id", table_name.ExternalLaboratoryTable)
	err := r.db.Select(&labs, query)
	if err != nil {
		return nil, nil, err
	}
	var labsEn []model.CommonLaboratoryWork
	query = fmt.Sprintf("SELECT id, name_en as name, task_description_en as name, link FROM %s ORDER BY id", table_name.ExternalLaboratoryTable)
	err = r.db.Select(&labsEn, query)
	if err != nil {
		return nil, nil, err
	}
	return labs, labsEn, nil
}

func (r *LecturerTestAndLabPostgres) GetAllLabFromSection(sectionId int) ([]model.LaboratoryWork, []model.LaboratoryWork, error) {
	var laboratory []model.LaboratoryWork
	query := fmt.Sprintf(`SELECT %s.id, %s.name, %s.task_description, %s.default_mark
								FROM %s
								INNER JOIN %s
								ON %s.id = %s.external_laboratory_id
								INNER JOIN %s
								ON %s.id = %s.laboratory_id
								WHERE section_id = $1
								ORDER BY %s.id`,
		table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.ExternalLaboratoryTable, table_name.LaboratoryTable,
		table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.ExternalLaboratoryTable, table_name.LaboratoryTable,
		table_name.SectionLabs, table_name.LaboratoryTable, table_name.SectionLabs, table_name.LaboratoryTable)
	err := r.db.Select(&laboratory, query, sectionId)
	if err != nil {
		return nil, nil, err
	}
	var laboratoryEn []model.LaboratoryWork
	query = fmt.Sprintf(`SELECT %s.id, %s.name_en as name, %s.task_description_en as task_description, %s.default_mark
								FROM %s
								INNER JOIN %s
								ON %s.id = %s.external_laboratory_id
								INNER JOIN %s
								ON %s.id = %s.laboratory_id
								WHERE section_id = $1
								ORDER BY %s.id`,
		table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.ExternalLaboratoryTable, table_name.LaboratoryTable,
		table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.ExternalLaboratoryTable, table_name.LaboratoryTable,
		table_name.SectionLabs, table_name.LaboratoryTable, table_name.SectionLabs, table_name.LaboratoryTable)
	err = r.db.Select(&laboratoryEn, query, sectionId)
	if err != nil {
		return nil, nil, err
	}
	return laboratory, laboratoryEn, nil
}

func (r *LecturerTestAndLabPostgres) ChangeTestName(testId int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", table_name.TestsTable)
	_, err := r.db.Exec(query, name, testId)
	return err
}

func (r *LecturerTestAndLabPostgres) ChangeTestNameEn(testId int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name_en = $1 WHERE id = $2", table_name.TestsTable)
	_, err := r.db.Exec(query, name, testId)
	return err
}

func (r *LecturerTestAndLabPostgres) ChangeTestTaskDescriptionEn(testId int, description string) error {
	query := fmt.Sprintf("UPDATE %s SET task_description_en = $1 WHERE id = $2", table_name.TestsTable)
	_, err := r.db.Exec(query, description, testId)
	return err
}

func (r *LecturerTestAndLabPostgres) ChangeTestTaskDescription(testId int, description string) error {
	query := fmt.Sprintf("UPDATE %s SET task_description = $1 WHERE id = $2", table_name.TestsTable)
	_, err := r.db.Exec(query, description, testId)
	return err
}

func (r *LecturerTestAndLabPostgres) ChangeLabName(labId int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", table_name.ExternalLaboratoryTable)
	_, err := r.db.Exec(query, name, labId)
	return err
}

func (r *LecturerTestAndLabPostgres) ChangeLabNameEn(labId int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name_en = $1 WHERE id = $2", table_name.ExternalLaboratoryTable)
	_, err := r.db.Exec(query, name, labId)
	return err
}

func (r *LecturerTestAndLabPostgres) ChangeLabTaskDescriptionEn(labId int, description string) error {
	query := fmt.Sprintf("UPDATE %s SET task_description_en = $1 WHERE id = $2", table_name.ExternalLaboratoryTable)
	_, err := r.db.Exec(query, description, labId)
	return err
}

func (r *LecturerTestAndLabPostgres) ChangeLabTaskDescription(labId int, description string) error {
	query := fmt.Sprintf("UPDATE %s SET task_description = $1 WHERE id = $2", table_name.ExternalLaboratoryTable)
	_, err := r.db.Exec(query, description, labId)
	return err
}

func (r *LecturerTestAndLabPostgres) ChangeLabDefaultMark(labId, mark int) error {
	query := fmt.Sprintf("UPDATE %s SET default_mark = $1 WHERE id = $2", table_name.LaboratoryTable)
	_, err := r.db.Exec(query, mark, labId)
	return err
}

func (r *LecturerTestAndLabPostgres) ChangeTestDuration(testId, minutes int) error {
	query := fmt.Sprintf("UPDATE %s SET minutes_duration = $1 WHERE id = $2", table_name.TestsTable)
	_, err := r.db.Exec(query, minutes, testId)

	return err
}

func (r *LecturerTestAndLabPostgres) CreateTest(test model.TestAdd) error {
	query := fmt.Sprintf("INSERT INTO %s (name, task_description, name_en, task_description_en, minutes_duration, default_mark) VALUES ($1, $2, $3, $4, $5, $6)",
		table_name.TestsTable)
	_, err := r.db.Exec(query, test.Name, test.TaskDescription, test.NameEn, test.TaskDescriptionEn, test.MinutesDuration, test.DefaultMark)
	return err
}

func (r *LecturerTestAndLabPostgres) DeleteTest(testId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", table_name.TestsTable)
	_, err := r.db.Exec(query, testId)
	return err
}

func (r *LecturerTestAndLabPostgres) AddThemeForTest(testId, themeId, count int) error {
	query := fmt.Sprintf("INSERT INTO %s (test_id, theme_id, count) VALUES($1, $2, $3)", table_name.TestsThemesTable)
	_, err := r.db.Exec(query, testId, themeId, count)
	return err
}

func (r *LecturerTestAndLabPostgres) CreateTheme(name string, weight int) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, weight) VALUES($1, $2) RETURNING id", table_name.ThemesTable)
	if err := r.db.QueryRow(query, name, weight).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *LecturerTestAndLabPostgres) GetThemeIdByName(name string) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE name ILIKE $1", table_name.ThemesTable)
	if err := r.db.QueryRow(query, name).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *LecturerTestAndLabPostgres) GetAllThemes(testId int) ([]model.ThemeOutput, error) {
	var themes []model.ThemeOutput
	query := fmt.Sprintf("SELECT id, name, weight, count FROM %s INNER JOIN %s ON %s.id = %s.theme_id  WHERE test_id = $1 ORDER BY id",
		table_name.ThemesTable, table_name.TestsThemesTable, table_name.ThemesTable, table_name.TestsThemesTable)
	err := r.db.Select(&themes, query, testId)
	if err != nil {
		return nil, err
	}
	return themes, nil
}

func (r *LecturerTestAndLabPostgres) GetAllExistThemes() ([]model.Theme, error) {
	var themes []model.Theme
	query := fmt.Sprintf("SELECT id, name, weight FROM %s ORDER BY id", table_name.ThemesTable)
	err := r.db.Select(&themes, query)
	if err != nil {
		return nil, err
	}
	return themes, nil
}

func (r *LecturerTestAndLabPostgres) DeleteTheme(themeId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", table_name.ThemesTable)
	_, err := r.db.Exec(query, themeId)
	return err
}

func (r *LecturerTestAndLabPostgres) ChangeThemeName(themeId int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", table_name.ThemesTable)
	_, err := r.db.Exec(query, name, themeId)
	return err
}

func (r *LecturerTestAndLabPostgres) ChangeThemeWeight(themeId, weight int) error {
	query := fmt.Sprintf("UPDATE %s SET weight = $1 WHERE id = $2", table_name.ThemesTable)
	_, err := r.db.Exec(query, weight, themeId)
	return err
}

func (r *LecturerTestAndLabPostgres) CreateQuestion(isVariable int, question string, questionEn string) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, name_en, is_variable) VALUES ($1, $2, $3) RETURNING id", table_name.QuestionsTable)
	row := r.db.QueryRow(query, question, questionEn, isVariable)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *LecturerTestAndLabPostgres) AddQuestionForTheme(themeId, questionId int) error {
	query := fmt.Sprintf("INSERT INTO %s (theme_id, question_id) VALUES ($1, $2)", table_name.ThemeQuestionsTable)
	_, err := r.db.Exec(query, themeId, questionId)
	return err
}

func (r *LecturerTestAndLabPostgres) DeleteQuestion(questionId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", table_name.QuestionsTable)
	if _, err := r.db.Exec(query, questionId); err != nil {
		return err
	}

	return nil
}

func (r *LecturerTestAndLabPostgres) ChangeQuestionName(questionId int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", table_name.QuestionsTable)
	_, err := r.db.Exec(query, name, questionId)
	return err
}

func (r *LecturerTestAndLabPostgres) ChangeQuestionNameEn(questionId int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name_en = $1 WHERE id = $2", table_name.QuestionsTable)
	_, err := r.db.Exec(query, name, questionId)
	return err
}

// TODO: do check that is_variable == 2
func (r *LecturerTestAndLabPostgres) AddAnswerForQuestion(questionId int, name string, nameEn string, isRight bool) error {
	query := fmt.Sprintf("INSERT INTO %s (name, name_en, is_right, question_id) VALUES ($1, $2, $3, $4)", table_name.AnswersTable)
	_, err := r.db.Exec(query, name, nameEn, isRight, questionId)
	return err
}

func (r *LecturerTestAndLabPostgres) DeleteAnswer(answerId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", table_name.AnswersTable)
	_, err := r.db.Exec(query, answerId)
	return err
}

func (r *LecturerTestAndLabPostgres) ChangeAnswerName(answerId int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", table_name.AnswersTable)
	_, err := r.db.Exec(query, name, answerId)
	return err
}

func (r *LecturerTestAndLabPostgres) ChangeAnswerNameEn(answerId int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name_en = $1 WHERE id = $2", table_name.AnswersTable)
	_, err := r.db.Exec(query, name, answerId)
	return err
}

func (r *LecturerTestAndLabPostgres) ChangeAnswerRight(answerId int, isRight bool) error {
	query := fmt.Sprintf("UPDATE %s SET is_right = $1 WHERE id = $2", table_name.AnswersTable)
	_, err := r.db.Exec(query, isRight, answerId)
	return err
}

func (r *LecturerTestAndLabPostgres) GetAnswers(questionId int) ([]model.Answer, []model.Answer, error) {
	var answers []model.Answer
	query := fmt.Sprintf("SELECT id, name, is_right FROM %s WHERE question_id = $1 ORDER BY id", table_name.AnswersTable)
	err := r.db.Select(&answers, query, questionId)
	if err != nil {
		return nil, nil, err
	}
	var answersEn []model.Answer
	query = fmt.Sprintf("SELECT id, name_en as name, is_right FROM %s WHERE question_id = $1 ORDER BY id", table_name.AnswersTable)
	err = r.db.Select(&answersEn, query, questionId)
	if err != nil {
		return nil, nil, err
	}
	return answers, answersEn, nil
}

func (r *LecturerTestAndLabPostgres) GetQuestions(themeId int) ([]model.Question, []model.Question, error) {
	var questions []model.Question
	query := fmt.Sprintf("SELECT id, name, is_variable FROM %s INNER JOIN %s ON %s.id = %s.question_id WHERE theme_id = $1 ORDER BY id",
		table_name.QuestionsTable, table_name.ThemeQuestionsTable, table_name.QuestionsTable, table_name.ThemeQuestionsTable)
	err := r.db.Select(&questions, query, themeId)
	if err != nil {
		return nil, nil, err
	}
	var questionsEn []model.Question
	query = fmt.Sprintf("SELECT id, name_en AS name, is_variable FROM %s INNER JOIN %s ON %s.id = %s.question_id WHERE theme_id = $1 ORDER BY id",
		table_name.QuestionsTable, table_name.ThemeQuestionsTable, table_name.QuestionsTable, table_name.ThemeQuestionsTable)
	err = r.db.Select(&questionsEn, query, themeId)
	if err != nil {
		return nil, nil, err
	}
	return questions, questionsEn, nil
}

func (r *LecturerTestAndLabPostgres) CreateTestDate(userId, testId, date int) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, test_id, closed_date) VALUES($1, $2, $3)", table_name.TestsDateTable)
	_, err := r.db.Exec(query, userId, testId, date)
	return err
}

func (r *LecturerTestAndLabPostgres) UpdateTestDate(userId, testId, date int) error {
	query := fmt.Sprintf("UPDATE %s SET closed_date = $1 WHERE user_id = $2 AND test_id = $3", table_name.TestsDateTable)
	_, err := r.db.Exec(query, date, userId, testId)
	return err
}

func (r *LecturerTestAndLabPostgres) GetAllQuestions() ([]model.QuestionWithAmountAnswers, []model.QuestionWithAmountAnswers, error) {
	var questions []model.QuestionWithAmountAnswers
	query := fmt.Sprintf("SELECT id, name, is_variable FROM %s ORDER BY id", table_name.QuestionsTable)
	err := r.db.Select(&questions, query)
	if err != nil {
		return nil, nil, err
	}
	var questionsEn []model.QuestionWithAmountAnswers
	query = fmt.Sprintf("SELECT id, name_en AS name, is_variable FROM %s ORDER BY id", table_name.QuestionsTable)
	err = r.db.Select(&questionsEn, query)
	if err != nil {
		return nil, nil, err
	}

	for i, question := range questions {
		var questionAmount int
		query := fmt.Sprintf("SELECT COUNT(*) AS count FROM %s WHERE question_id = $1  GROUP BY (question_id)", table_name.AnswersTable)
		if err := r.db.Get(&questionAmount, query, question.QuestionId); err != nil {
			questionAmount = 0
		}
		questions[i].AmountAnswers = questionAmount
	}

	for i, question := range questionsEn {
		var questionAmount int
		query := fmt.Sprintf("SELECT COUNT(*) AS count FROM %s WHERE question_id = $1  GROUP BY (question_id)", table_name.AnswersTable)
		if err := r.db.Get(&questionAmount, query, question.QuestionId); err != nil {
			questionAmount = 0
		}
		questionsEn[i].AmountAnswers = questionAmount
	}

	return questions, questionsEn, nil
}

func (r *LecturerTestAndLabPostgres) GetQuestionsByName(name string) ([]model.Question, error) {
	var questions []model.Question
	query := fmt.Sprintf("SELECT id, name, is_variable FROM %s WHERE name ilike $1 ORDER BY id", table_name.QuestionsTable)
	err := r.db.Select(&questions, query, name+"%")
	if err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *LecturerTestAndLabPostgres) DeleteQuestionFromTheme(themeId, questionId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE theme_id = $1 AND question_id = $2", table_name.ThemeQuestionsTable)
	_, err := r.db.Exec(query, themeId, questionId)
	return err
}

func (r *LecturerTestAndLabPostgres) ChangeThemeTestCount(testId, themeId, count int) error {
	query := fmt.Sprintf("UPDATE %s SET count = $1 WHERE test_id = $2 AND theme_id = $3", table_name.TestsThemesTable)
	_, err := r.db.Exec(query, count, testId, themeId)
	return err
}

func (r *LecturerTestAndLabPostgres) DeleteThemeFromTest(testId, themeId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE test_id = $1 AND theme_id = $2", table_name.TestsThemesTable)
	_, err := r.db.Exec(query, testId, themeId)
	return err
}

func (r *LecturerTestAndLabPostgres) checkExistTestDate(userId, testId int) bool {
	query := fmt.Sprintf("SELECT closed_date FROM %s WHERE user_id = $1 AND test_id = $2", table_name.TestsDateTable)
	row := r.db.QueryRow(query, userId, testId)
	var date int
	if err := row.Scan(&date); err != nil {
		return false
	}
	return true
}

func (r *LecturerTestAndLabPostgres) OpenTestForStudent(studentId, testId int, date int64) error {
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

func (r *LecturerTestAndLabPostgres) checkExistTestMark(studentId, testId int) bool {
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

func (r *LecturerTestAndLabPostgres) GetOpenedTestForStudent(studentId, testId int) (model.OpenedTest, error) {
	var test model.OpenedTest
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND test_id = $2", table_name.TestsDateTable)
	if err := r.db.Get(&test, query, studentId, testId); err != nil {
		return model.OpenedTest{}, err
	}

	return test, nil
}

func (r *LecturerTestAndLabPostgres) CloseOpenedTestForStudent(studentId, testId int) error {
	query := fmt.Sprintf("UPDATE %s SET is_done = true WHERE user_id = $1 AND test_id = $2", table_name.TestsDateTable)
	if _, err := r.db.Exec(query, studentId, testId); err != nil {
		return err
	}

	return nil
}

func (r *LecturerTestAndLabPostgres) ChangeTestMarkForStudent(studentId, testId, mark int) error {
	query := fmt.Sprintf("UPDATE %s SET mark = $1 WHERE user_id = $2 AND test_id = $3", table_name.TestsMarkTable)

	if _, err := r.db.Exec(query, mark, studentId, testId); err != nil {
		return err
	}

	return nil
}

func (r *LecturerTestAndLabPostgres) GetTestMarkForStudent(studentId, testId int) (int, error) {
	var mark int
	query := fmt.Sprintf("SELECT mark FROM %s WHERE user_id = $1 AND test_id = $2", table_name.TestsMarkTable)

	if err := r.db.Get(&mark, query, studentId, testId); err != nil {
		return 0, err
	}

	return mark, nil
}

func (r *LecturerTestAndLabPostgres) GetQuestionWithoutEnglishVersion() ([]model.Question, error) {
	var questions []model.Question
	query := fmt.Sprintf("SELECT id, name, is_variable FROM %s WHERE name_en = $1 ORDER BY id", table_name.QuestionsTable)
	err := r.db.Select(&questions, query, "")
	if err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *LecturerTestAndLabPostgres) GetQuestionWithoutTheme() ([]model.Question, error) {
	var questions []model.Question
	query := fmt.Sprintf("SELECT id, name, is_variable FROM %s WHERE id NOT IN (SELECT %s.id FROM %s INNER JOIN %s ON %s.id = %s.question_id) ORDER BY id",
		table_name.QuestionsTable, table_name.QuestionsTable, table_name.QuestionsTable, table_name.ThemeQuestionsTable, table_name.QuestionsTable, table_name.ThemeQuestionsTable)
	err := r.db.Select(&questions, query)
	if err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *LecturerTestAndLabPostgres) GetUsersWithOpenedTest(testId int, timeNow int64) ([]model.StudentWithGroupWithClosedDate, error) {
	var students []model.StudentWithGroupWithClosedDate
	query := fmt.Sprintf("SELECT %s.id, %s.name, %s.surname, %s.name AS group_name, closed_date FROM %s INNER JOIN %s ON %s.user_id = %s.id INNER JOIN %s ON %s.test_id = %s.id WHERE test_id = $1 AND is_done = false AND (closed_date >= $2 OR closed_date = 0) ORDER BY closed_date DESC",
		table_name.UsersTable, table_name.UsersTable, table_name.UsersTable,
		table_name.UsersTable, table_name.UsersTable, table_name.TestsDateTable,
		table_name.TestsDateTable, table_name.UsersTable, table_name.TestsTable,
		table_name.TestsDateTable, table_name.TestsTable)
	err := r.db.Select(&students, query, testId, timeNow)
	if err != nil {
		return []model.StudentWithGroupWithClosedDate{}, err
	}
	return students, nil
}

func (r *LecturerTestAndLabPostgres) GetUsersWithDoneTests(testId int, timeNow int64) ([]model.StudentWithGroupWithClosedDate, error) {
	var students []model.StudentWithGroupWithClosedDate
	query := fmt.Sprintf(`SELECT %s.id, %s.name, %s.surname, %s.name AS group_name, closed_date
				FROM %s
				INNER JOIN %s
				ON %s.user_id = %s.id
				INNER JOIN %s ON %s.test_id = %s.id
				WHERE (test_id = $1 AND is_done = true) OR (closed_date < $2 AND closed_date != 0)
				ORDER BY closed_date DESC`,
		table_name.UsersTable, table_name.UsersTable, table_name.UsersTable,
		table_name.UsersTable, table_name.UsersTable, table_name.TestsDateTable,
		table_name.TestsDateTable, table_name.UsersTable, table_name.TestsTable,
		table_name.TestsDateTable, table_name.TestsTable)
	err := r.db.Select(&students, query, testId, timeNow)
	if err != nil {
		return []model.StudentWithGroupWithClosedDate{}, err
	}
	return students, nil
}

func (r *LecturerTestAndLabPostgres) GetUsersWithOpenedLab(labId int, timeNow int64) ([]model.StudentWithGroupWithClosedDate, error) {
	var students []model.StudentWithGroupWithClosedDate
	query := fmt.Sprintf("SELECT %s.id, %s.name, %s.surname, %s.name AS group_name, closed_date FROM %s INNER JOIN %s ON %s.user_id = %s.id INNER JOIN %s ON %s.laboratory_id = %s.id WHERE laboratory_id = $1 AND is_done = false AND (closed_date >= $2 OR closed_date = 0) ORDER BY closed_date DESC",
		table_name.UsersTable, table_name.UsersTable, table_name.UsersTable,
		table_name.UsersTable, table_name.UsersTable, table_name.LaboratoryDateTable,
		table_name.LaboratoryDateTable, table_name.UsersTable, table_name.LaboratoryTable,
		table_name.LaboratoryDateTable, table_name.LaboratoryTable)
	err := r.db.Select(&students, query, labId, timeNow)
	if err != nil {
		return []model.StudentWithGroupWithClosedDate{}, err
	}
	return students, nil
}

func (r *LecturerTestAndLabPostgres) GetUsersWithDoneLaboratory(labId int, timeNow int64) ([]model.StudentWithGroupWithClosedDate, error) {
	var students []model.StudentWithGroupWithClosedDate
	query := fmt.Sprintf(`SELECT %s.id, %s.name, %s.surname, %s.name AS group_name, closed_date
				FROM %s
				INNER JOIN %s
				ON %s.user_id = %s.id
				INNER JOIN %s ON %s.laboratory_id = %s.id
				WHERE (laboratory_id = $1 AND is_done = true) OR (closed_date < $2 AND closed_date != 0)
				ORDER BY closed_date DESC`,
		table_name.UsersTable, table_name.UsersTable, table_name.UsersTable,
		table_name.UsersTable, table_name.UsersTable, table_name.LaboratoryDateTable,
		table_name.LaboratoryDateTable, table_name.UsersTable, table_name.LaboratoryTable,
		table_name.LaboratoryDateTable, table_name.LaboratoryTable)
	err := r.db.Select(&students, query, labId, timeNow)
	if err != nil {
		return []model.StudentWithGroupWithClosedDate{}, err
	}
	return students, nil
}

func (r *LecturerTestAndLabPostgres) GetThemesByQuestion(questionId int) ([]model.Theme, error) {
	var themes []model.Theme
	query := fmt.Sprintf("SELECT %s.id, %s.name, %s.weight FROM %s INNER JOIN %s ON %s.theme_id = %s.id WHERE %s.question_id = $1 ORDER BY %s.id",
		table_name.ThemesTable, table_name.ThemesTable, table_name.ThemesTable, table_name.ThemesTable, table_name.ThemeQuestionsTable,
		table_name.ThemeQuestionsTable, table_name.ThemesTable, table_name.ThemeQuestionsTable, table_name.ThemesTable)
	err := r.db.Select(&themes, query, questionId)
	if err != nil {
		return nil, err
	}
	return themes, nil
}

func (r *LecturerTestAndLabPostgres) CreateExternalLab(lab model.LaboratoryWorkInputWithoutId) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, name_en, task_description, task_description_en, link, token) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", table_name.ExternalLaboratoryTable)
	var id int
	if err := r.db.QueryRow(query, lab.Name, lab.NameEn, lab.TaskDescription, lab.TaskDescriptionEn, lab.Linc, lab.Token).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *LecturerTestAndLabPostgres) DeleteExternalLab(labId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", table_name.ExternalLaboratoryTable)
	if _, err := r.db.Exec(query, labId); err != nil {
		return err
	}

	return nil
}

func (r *LecturerTestAndLabPostgres) ChangeLabLinc(labId int, linc string) error {
	query := fmt.Sprintf("UPDATE %s SET link = $1 WHERE id = $2", table_name.ExternalLaboratoryTable)
	if _, err := r.db.Exec(query, linc, labId); err != nil {
		return err
	}

	return nil
}

func (r *LecturerTestAndLabPostgres) ChangeLabToken(labId int, token string) error {
	query := fmt.Sprintf("UPDATE %s SET token = $1 WHERE id = $2", table_name.ExternalLaboratoryTable)
	if _, err := r.db.Exec(query, token, labId); err != nil {
		return err
	}

	return nil
}

func (r *LecturerTestAndLabPostgres) GetExternalLabInfo(labId int) (model.LaboratoryWorkResponse, error) {
	query := fmt.Sprintf("SELECT id, name, name_en, task_description, task_description_en, link FROM %s WHERE id = $1",
		table_name.ExternalLaboratoryTable)
	var lab model.LaboratoryWorkResponse
	if err := r.db.Get(&lab, query, labId); err != nil {
		return model.LaboratoryWorkResponse{}, err
	}

	return lab, nil
}

func (r *LecturerTestAndLabPostgres) GetLabToken(labId int) (string, error) {
	query := fmt.Sprintf("SELECT token FROM %s WHERE id = $1", table_name.ExternalLaboratoryTable)
	var token string
	if err := r.db.Get(&token, query, labId); err != nil {
		return "", err
	}

	return token, nil
}

func (r *LecturerTestAndLabPostgres) GetLabInfo(labId int) (int, error) {
	query := fmt.Sprintf("SELECT external_laboratory_id FROM %s WHERE id = $1", table_name.LaboratoryTable)
	var id int
	if err := r.db.Get(&id, query, labId); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *LecturerTestAndLabPostgres) checkExistLabDate(userId, labId int) bool {
	query := fmt.Sprintf("SELECT closed_date FROM %s WHERE user_id = $1 AND laboratory_id = $2", table_name.LaboratoryDateTable)
	row := r.db.QueryRow(query, userId, labId)
	var date int
	if err := row.Scan(&date); err != nil {
		return false
	}
	return true
}

func (r *LecturerTestAndLabPostgres) OpenLabForStudent(studentId, labId int, date int64) error {
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

func (r *LecturerTestAndLabPostgres) checkExistLabMark(studentId, labId int) bool {
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

func (r *LecturerTestAndLabPostgres) CloseOpenedLabForStudent(studentId, labId int) error {
	query := fmt.Sprintf("UPDATE %s SET is_done = true WHERE user_id = $1 AND laboratory_id = $2", table_name.LaboratoryDateTable)
	if _, err := r.db.Exec(query, studentId, labId); err != nil {
		return err
	}

	return nil
}

func (r *LecturerTestAndLabPostgres) ChangeLabMarkForStudent(studentId, labId, mark int) error {
	query := fmt.Sprintf("UPDATE %s SET mark = $1 WHERE user_id = $2 AND laboratory_id = $3", table_name.LaboratoryMarkTable)

	if _, err := r.db.Exec(query, mark, studentId, labId); err != nil {
		return err
	}

	return nil
}

func (r *LecturerTestAndLabPostgres) GetLabMarkForStudent(studentId, labId int) (int, error) {
	var mark int
	query := fmt.Sprintf("SELECT mark FROM %s WHERE user_id = $1 AND laboratory_id = $2", table_name.LaboratoryMarkTable)

	if err := r.db.Get(&mark, query, studentId, labId); err != nil {
		return 0, err
	}

	return mark, nil
}
