package api_student_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"strconv"
	"time"
)

type StudentTestAndLabPostgres struct {
	db *sqlx.DB
}

func NewStudentTestAndLabPostgres(db *sqlx.DB) *StudentTestAndLabPostgres {
	return &StudentTestAndLabPostgres{db: db}
}

func (r *StudentTestAndLabPostgres) GetAllTestFromSection(userId, sectionId int) ([]model.Test, []model.Test, error) {
	var tests []model.Test
	query := fmt.Sprintf("SELECT id, name, task_description, minutes_duration, default_mark  FROM %s INNER JOIN %s ON %s.id = %s.test_id INNER JOIN %s ON %s.id = %s.test_id WHERE section_id = $1 AND user_id = $2 ORDER BY %s.id",
		table_name.TestsTable, table_name.TestsDateTable, table_name.TestsTable,
		table_name.TestsDateTable, table_name.SectionTests, table_name.TestsTable,
		table_name.SectionTests, table_name.TestsTable)
	err := r.db.Select(&tests, query, sectionId, userId)
	if err != nil {
		return nil, nil, err
	}
	var testsEn []model.Test
	query = fmt.Sprintf("SELECT id, name_en as name, task_description_en as task_description, minutes_duration, default_mark  FROM %s INNER JOIN %s ON %s.id = %s.test_id INNER JOIN %s ON %s.id = %s.test_id WHERE section_id = $1 AND user_id = $2 ORDER BY %s.id",
		table_name.TestsTable, table_name.TestsDateTable, table_name.TestsTable,
		table_name.TestsDateTable, table_name.SectionTests, table_name.TestsTable,
		table_name.SectionTests, table_name.TestsTable)
	err = r.db.Select(&testsEn, query, sectionId, userId)
	if err != nil {
		return nil, nil, err
	}
	return tests, testsEn, nil
}

func (r *StudentTestAndLabPostgres) GetAllOpenedTests(userId int, timeNow int64) ([]model.TestWithClosedDate, []model.TestWithClosedDate, error) {
	var tests []model.TestWithClosedDate
	query := fmt.Sprintf("SELECT id, name, task_description, minutes_duration, default_mark, closed_date FROM %s INNER JOIN %s ON %s.id = %s.test_id WHERE (user_id = $1 AND is_done = false) AND (closed_date > $2 OR closed_date = 0) ORDER BY %s.id",
		table_name.TestsTable, table_name.TestsDateTable, table_name.TestsTable, table_name.TestsDateTable, table_name.TestsTable)
	err := r.db.Select(&tests, query, userId, timeNow)
	if err != nil {
		return nil, nil, err
	}
	var testsEn []model.TestWithClosedDate
	query = fmt.Sprintf("SELECT id, name_en as name, task_description_en as task_description, minutes_duration, default_mark, closed_date  FROM %s INNER JOIN %s ON %s.id = %s.test_id WHERE(user_id = $1 AND is_done = false) AND (closed_date > $2 OR closed_date = 0) ORDER BY %s.id",
		table_name.TestsTable, table_name.TestsDateTable, table_name.TestsTable, table_name.TestsDateTable, table_name.TestsTable)
	err = r.db.Select(&testsEn, query, userId, timeNow)
	if err != nil {
		return nil, nil, err
	}
	return tests, testsEn, nil
}

func (r *StudentTestAndLabPostgres) GetAllLab(userId, sectionId int) ([]model.LaboratoryWork, []model.LaboratoryWork, error) {
	var laboratory []model.LaboratoryWork
	query := fmt.Sprintf("SELECT %s.id, %s.name, %s.task_description, %s.lync, %s.default_mark FROM %s INNER JOIN %s ON %s.external_laboratory_id = %s.id INNER JOIN %s ON %s.id = %s.laboratory_id INNER JOIN %s ON %s.id = %s.laboratory_id WHERE section_id = $1 AND user_id = $2 ORDER BY %s.id",
		table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.ExternalLaboratoryTable, table_name.ExternalLaboratoryTable,
		table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.ExternalLaboratoryTable, table_name.LaboratoryTable,
		table_name.LaboratoryDateTable, table_name.LaboratoryTable,
		table_name.LaboratoryDateTable, table_name.SectionLabs, table_name.LaboratoryTable,
		table_name.SectionLabs, table_name.LaboratoryTable, table_name.LaboratoryTable)
	err := r.db.Select(&laboratory, query, sectionId, userId)
	if err != nil {
		return nil, nil, err
	}
	var laboratoryEn []model.LaboratoryWork
	query = fmt.Sprintf("SELECT %s.id, %s.name_en as name, %s.task_description_en as task_description, %s.lync, %s.default_mark FROM %s INNER JOIN %s ON %s.external_laboratory_id = %s.id INNER JOIN %s ON %s.id = %s.laboratory_id INNER JOIN %s ON %s.id = %s.laboratory_id WHERE section_id = $1 AND user_id = $2 ORDER BY %s.id",
		table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.ExternalLaboratoryTable, table_name.ExternalLaboratoryTable,
		table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.ExternalLaboratoryTable, table_name.LaboratoryTable,
		table_name.LaboratoryDateTable, table_name.LaboratoryTable,
		table_name.LaboratoryDateTable, table_name.SectionLabs, table_name.LaboratoryTable,
		table_name.SectionLabs, table_name.LaboratoryTable, table_name.LaboratoryTable)
	err = r.db.Select(&laboratoryEn, query, sectionId, userId)
	if err != nil {
		return nil, nil, err
	}
	return laboratory, laboratoryEn, nil
}

func (r *StudentTestAndLabPostgres) GiveAccessForTest(userId, testId int) (bool, error) {
	query := fmt.Sprintf("SELECT closed_date FROM %s WHERE user_id = $1 AND test_id = $2", table_name.TestsDateTable)
	var date int
	errAccess := errors.New("access to test denied")
	row := r.db.QueryRow(query, userId, testId)
	if err := row.Scan(&date); err != nil {
		return false, errAccess
	}
	currentTime := time.Now()
	unixTime := currentTime.Unix()
	if int(unixTime) < date {
		return true, nil
	} else {
		return false, errAccess
	}
}

func (r *StudentTestAndLabPostgres) CheckAccessForDiscipline(studentId, disciplineId int) error {
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

func (r *StudentTestAndLabPostgres) CheckAccessForSection(studentId, sectionId int) error {
	var disciplineId int
	errAccess := errors.New("access denied to section")
	query := fmt.Sprintf("SELECT discipline_id FROM %s WHERE id = $1", table_name.SectionTable)
	if err := r.db.Get(&disciplineId, query, sectionId); err != nil {
		return errAccess
	}
	if err := r.CheckAccessForDiscipline(studentId, disciplineId); err != nil {
		return errAccess
	}
	return nil
}

func (r *StudentTestAndLabPostgres) GetQuestionsForTest(testId int) ([]int, error) {
	var theme []int
	query := fmt.Sprintf("SELECT id FROM %s WHERE test_id = $1", table_name.ThemesTable)
	err := r.db.Select(&theme, query, testId)
	if err != nil {
		return nil, err
	}
	var questionId []int
	var chosenQuestion []int
	for _, value := range theme {
		questionId = []int{}
		query = fmt.Sprintf("SELECT id FROM %s WHERE theme_id = $1", table_name.QuestionsTable)
		err := r.db.Select(&questionId, query, value)
		if err != nil {
			return nil, err
		}
		if len(questionId) == 0 {
			continue
		}
		question := rand.Intn(len(questionId))
		chosenQuestion = append(chosenQuestion, questionId[question])
	}
	return chosenQuestion, nil
}

func (r *StudentTestAndLabPostgres) CheckAccessToOpenTest(userId, testId int) bool {
	var userAccess model.UserTestAccess
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND test_id = $2", table_name.TestsDateTable)
	if err := r.db.Get(&userAccess, query, userId, testId); err != nil {
		return false
	}
	if userAccess.IsDone {
		return false
	}
	var testCheck model.CheckTimeForTest
	query = fmt.Sprintf("SELECT id, minutes_duration FROM %s WHERE id = $1", table_name.TestsTable)
	if err := r.db.Get(&testCheck, query, testId); err != nil {
		return false
	}
	if time.Now().Unix() >= userAccess.ClosedDate && userAccess.ClosedDate != 0 {
		return false
	} else {
		query = fmt.Sprintf("UPDATE %s SET closed_date = $1 WHERE user_id = $2 AND test_id = $3", table_name.TestsDateTable)
		if _, err := r.db.Exec(query, time.Now().Unix()+60*int64(testCheck.MinutesDuration+3), userId, testId); err != nil {
			return false
		}
	}
	return true
}

func (r *StudentTestAndLabPostgres) CheckAccessToOpenLab(userId, labId int) (bool, error) {
	var userAccess model.UserLabAccess
	accessDeniedErr := fmt.Errorf("access denied")

	currentTime := time.Now().Unix()
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND laboratory_id = $2 AND closed_date >= $3", table_name.LaboratoryDateTable)
	if err := r.db.Get(&userAccess, query, userId, labId, currentTime); err != nil {
		return false, accessDeniedErr
	}
	if userAccess.IsDone {
		return false, accessDeniedErr
	}

	var labCheck model.CheckTimeForLab
	query = fmt.Sprintf("SELECT id, minutes_duration FROM %s WHERE id = $1", table_name.LaboratoryTable)
	if err := r.db.Get(&labCheck, query, labId); err != nil {
		return false, accessDeniedErr
	}

	if userAccess.ClosedDate == 0 {
		query = fmt.Sprintf("UPDATE %s SET closed_date = $1 WHERE user_id = $2 AND laboratory_id = $3", table_name.LaboratoryDateTable)
		if _, err := r.db.Exec(query, time.Now().Unix()+60*int64(labCheck.MinutesDuration), userId, labId); err != nil {
			return false, accessDeniedErr
		}
	} else if time.Now().Unix() >= userAccess.ClosedDate {
		return false, accessDeniedErr
	}
	if userAccess.IsDone {
		return false, accessDeniedErr
	}
	return true, nil
}

func (r *StudentTestAndLabPostgres) GetThemesQuestions(testId int) ([]model.QuestionWithTheme, int, error) {
	var questionsToChoose []model.QuestionWithTheme
	var themes []model.ThemeForChoose
	query := fmt.Sprintf("SELECT theme_id, count FROM %s WHERE test_id = $1", table_name.TestsThemesTable)
	if err := r.db.Select(&themes, query, testId); err != nil {
		return nil, 0, err
	}
	sum := 0
	for _, theme := range themes {
		sum += theme.Count
		var questions []model.QuestionWithTheme
		query := fmt.Sprintf("SELECT question_id, %s.theme_id FROM %s INNER JOIN %s ON %s.theme_id = %s.theme_id INNER JOIN %s ON %s.id = %s.theme_id WHERE %s.test_id = $1 AND %s.theme_id = $2",
			table_name.TestsThemesTable, table_name.TestsThemesTable, table_name.ThemeQuestionsTable, table_name.TestsThemesTable, table_name.ThemeQuestionsTable,
			table_name.ThemesTable, table_name.ThemesTable, table_name.TestsThemesTable,
			table_name.TestsThemesTable, table_name.TestsThemesTable)
		if err := r.db.Select(&questions, query, testId, theme.ThemeId); err != nil {
			return nil, 0, err
		}
		if len(questions) < theme.Count {
			err := errors.New("недостаточно вопросов")
			return nil, 0, err
		}
		length := len(questions)
		for i := 0; i < theme.Count; i++ {
			rand.Seed(time.Now().UnixNano())
			randomNum := rand.Intn(length)
			questionsToChoose = append(questionsToChoose, questions[randomNum])
			questions = append(questions[:randomNum], questions[randomNum+1:]...)
			length--
		}
	}

	return questionsToChoose, sum, nil
}

func (r *StudentTestAndLabPostgres) GetQuestionsAndAnswersForTest(questions []model.QuestionWithTheme) ([]model.QuestionsWithAnswers, []model.QuestionsWithAnswers, error) {
	var questionsAndAnswersRu []model.QuestionsWithAnswers
	var questionsAndAnswersEn []model.QuestionsWithAnswers

	str := "("
	for i := range questions {
		str = str + strconv.Itoa(questions[i].QuestionId)
		if i < len(questions)-1 {
			str = str + ", "
		} else {
			str = str + ")"
		}
	}

	var questionsRu []model.Question
	query := fmt.Sprintf("SELECT id, name, is_variable FROM %s WHERE id IN %s", table_name.QuestionsTable, str)
	if err := r.db.Select(&questionsRu, query); err != nil {
		return nil, nil, err
	}

	var questionsEn []model.Question
	query = fmt.Sprintf("SELECT id, name_en as name, is_variable FROM %s WHERE id IN %s", table_name.QuestionsTable, str)
	if err := r.db.Select(&questionsEn, query); err != nil {
		return nil, nil, err
	}

	var ansRu []model.AnswerDuringTest
	query = fmt.Sprintf("SELECT id, name, question_id FROM %s WHERE question_id IN %s", table_name.AnswersTable, str)
	if err := r.db.Select(&ansRu, query); err != nil {
		return nil, nil, err
	}

	var ansEn []model.AnswerDuringTest
	query = fmt.Sprintf("SELECT id, name_en as name, question_id FROM %s WHERE question_id IN %s", table_name.AnswersTable, str)
	if err := r.db.Select(&ansEn, query); err != nil {
		return nil, nil, err
	}

	mapkaRu := make(map[int][]model.AnswerDuringTest)
	for _, ans := range ansRu {
		mapkaRu[ans.QuestionId] = append(mapkaRu[ans.QuestionId], ans)
	}

	mapkaEn := make(map[int][]model.AnswerDuringTest)
	for _, ans := range ansEn {
		mapkaEn[ans.QuestionId] = append(mapkaEn[ans.QuestionId], ans)
	}

	for _, question := range questionsRu {
		var ans []model.AnswerDuringTest
		switch question.IsVariable {
		case 0:
			ans = nil
		default:
			ans = mapkaRu[question.QuestionId]
		}
		themeId := 0
		for i := range questions {
			if question.QuestionId == questions[i].QuestionId {
				themeId = questions[i].ThemeId
			}
		}
		data := model.QuestionsWithAnswers{
			Question: question,
			ThemeId:  themeId,
			Answers:  ans,
		}

		questionsAndAnswersRu = append(questionsAndAnswersRu, data)
	}

	for _, question := range questionsEn {
		var ans []model.AnswerDuringTest
		switch question.IsVariable {
		case 0:
			ans = nil
		default:
			ans = mapkaEn[question.QuestionId]
		}
		themeId := 0
		for i := range questions {
			if question.QuestionId == questions[i].QuestionId {
				themeId = questions[i].ThemeId
			}
		}
		data := model.QuestionsWithAnswers{
			Question: question,
			ThemeId:  themeId,
			Answers:  ans,
		}

		questionsAndAnswersEn = append(questionsAndAnswersEn, data)
	}

	return questionsAndAnswersRu, questionsAndAnswersEn, nil
}

func (r *StudentTestAndLabPostgres) GetTestMark(testId int) (int, error) {
	var mark int
	query := fmt.Sprintf("SELECT default_mark FROM %s WHERE id = $1", table_name.TestsTable)
	if err := r.db.Get(&mark, query, testId); err != nil {
		return 0, err
	}
	return mark, nil
}

func (r *StudentTestAndLabPostgres) GetWeightFromQuestionTheme(themeId int) int {
	var weight int
	query := fmt.Sprintf("SELECT weight FROM %s WHERE id = $1", table_name.ThemesTable)
	if err := r.db.Get(&weight, query, themeId); err != nil {
		return 0
	}

	return weight
}

func (r *StudentTestAndLabPostgres) GetTestMaxPoint(testId int) (int, error) {
	var points []model.ThemePoint
	query := fmt.Sprintf("SELECT count, weight FROM %s INNER JOIN %s ON %s.theme_id = %s.id WHERE test_id = $1",
		table_name.TestsThemesTable, table_name.ThemesTable, table_name.TestsThemesTable, table_name.ThemesTable)
	if err := r.db.Select(&points, query, testId); err != nil {
		return 0, err
	}
	maxPoints := 0
	for _, point := range points {
		maxPoints = maxPoints + point.Weight*point.Count
	}
	return maxPoints, nil
}

func (r *StudentTestAndLabPostgres) GetQuestionAnswers(questionId int) (model.GetQuestionWithRightAnswer, int, error) {
	var answers model.GetQuestionWithRightAnswer
	query := fmt.Sprintf("SELECT name, name_en FROM %s WHERE question_id = $1 AND is_right = true", table_name.AnswersTable)
	if err := r.db.Select(&answers.Answers, query, questionId); err != nil {
		return model.GetQuestionWithRightAnswer{}, 0, err
	}

	var numberOfAnswer int
	query = fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE question_id = $1", table_name.AnswersTable)
	if err := r.db.Get(&numberOfAnswer, query, questionId); err != nil {
		return model.GetQuestionWithRightAnswer{}, 0, err
	}

	return answers, numberOfAnswer, nil
}

func (r *StudentTestAndLabPostgres) DoneOfTest(userId, testId int) error {
	query := fmt.Sprintf("UPDATE %s SET is_done = true WHERE user_id = $1 AND test_id = $2", table_name.TestsDateTable)
	_, err := r.db.Exec(query, userId, testId)

	return err
}

func (r *StudentTestAndLabPostgres) GetPersonalData(id int) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", table_name.UsersTable)
	err := r.db.Get(&user, query, id)
	if err != nil {
		return model.User{}, err
	}
	query = fmt.Sprintf("SELECT %s.name FROM %s INNER JOIN %s ON %s.group_id = %s.id WHERE user_id = $1",
		table_name.GroupsTable, table_name.StudentsTable, table_name.GroupsTable, table_name.StudentsTable, table_name.GroupsTable)
	err = r.db.Get(&user.GroupName, query, id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *StudentTestAndLabPostgres) GetTest(testId int) (model.Test, model.Test, error) {
	var tests model.Test
	query := fmt.Sprintf("SELECT id, name, task_description, minutes_duration, default_mark  FROM %s INNER JOIN %s ON %s.id = %s.test_id INNER JOIN %s ON %s.id = %s.test_id WHERE %s.id = $1",
		table_name.TestsTable, table_name.TestsDateTable, table_name.TestsTable,
		table_name.TestsDateTable, table_name.SectionTests, table_name.TestsTable,
		table_name.SectionTests, table_name.TestsTable)
	err := r.db.Get(&tests, query, testId)
	if err != nil {
		return model.Test{}, model.Test{}, err
	}
	var testsEn model.Test
	query = fmt.Sprintf("SELECT id, name_en as name, task_description_en as task_description, minutes_duration, default_mark  FROM %s INNER JOIN %s ON %s.id = %s.test_id INNER JOIN %s ON %s.id = %s.test_id WHERE %s.id = $1",
		table_name.TestsTable, table_name.TestsDateTable, table_name.TestsTable,
		table_name.TestsDateTable, table_name.SectionTests, table_name.TestsTable,
		table_name.SectionTests, table_name.TestsTable)
	err = r.db.Get(&testsEn, query, testId)
	if err != nil {
		return model.Test{}, model.Test{}, err
	}
	return tests, testsEn, nil
}

func (r *StudentTestAndLabPostgres) GetQuestion(id int) (model.Question, model.Question, error) {
	var questionRu model.Question
	query := fmt.Sprintf("SELECT id, name, is_variable FROM %s WHERE id = $1", table_name.QuestionsTable)
	if err := r.db.Get(&questionRu, query, id); err != nil {
		return model.Question{}, model.Question{}, err
	}

	var questionEn model.Question
	query = fmt.Sprintf("SELECT id, name_en as name, is_variable FROM %s WHERE id = $1", table_name.QuestionsTable)
	if err := r.db.Get(&questionEn, query, id); err != nil {
		return model.Question{}, model.Question{}, err
	}

	return questionRu, questionEn, nil
}

func (r *StudentTestAndLabPostgres) GetAnswers(id int) ([]model.AllAnswersName, error) {
	var answers []model.AllAnswersName
	query := fmt.Sprintf("SELECT name, name_en FROM %s WHERE question_id = $1", table_name.AnswersTable)
	if err := r.db.Select(&answers, query, id); err != nil {
		return nil, err
	}

	return answers, nil
}

func (r *StudentTestAndLabPostgres) GetResultOfTest(userId, testId int) (model.TestResult, error) {
	var test model.TestResult
	query := fmt.Sprintf("SELECT * FROM %s WHERE test_id = $1 AND user_id = $2", table_name.TestsMarkTable)
	if err := r.db.Get(&test, query, testId, userId); err != nil {
		return model.TestResult{}, err
	}

	return test, nil
}

func (r *StudentTestAndLabPostgres) SaveResultOfTest(userId, testId, mark int) error {
	testRes, _ := r.GetResultOfTest(userId, testId)
	if testRes == (model.TestResult{}) {
		query := fmt.Sprintf("INSERT INTO %s (user_id, test_id, mark) VALUES ($1, $2, $3)", table_name.TestsMarkTable)
		if _, err := r.db.Exec(query, userId, testId, mark); err != nil {
			return err
		}
		return nil
	} else {
		query := fmt.Sprintf("UPDATE %s SET mark = $1 WHERE test_id = $2 AND user_id = $3", table_name.TestsMarkTable)
		if _, err := r.db.Exec(query, mark, testId, userId); err != nil {
			return err
		}
	}
	return nil
}

func (r *StudentTestAndLabPostgres) GetAllDoneTests(userId int, timeNow int64) ([]model.TestWithClosedDate, []model.TestWithClosedDate, error) {
	var tests []model.TestWithClosedDate
	query := fmt.Sprintf(`SELECT id, name, task_description, minutes_duration, default_mark, closed_date 
			FROM %s 
			INNER JOIN %s 
			ON %s.id = %s.test_id
			WHERE %s.user_id = $1 AND (is_done = true OR (closed_date < $2 AND closed_date != 0))
			ORDER BY %s.id`,
		table_name.TestsTable, table_name.TestsDateTable, table_name.TestsTable,
		table_name.TestsDateTable, table_name.TestsDateTable, table_name.TestsTable)
	err := r.db.Select(&tests, query, userId, timeNow)
	if err != nil {
		return nil, nil, err
	}
	var testsEn []model.TestWithClosedDate
	query = fmt.Sprintf(`SELECT id, name_en as name, task_description_en as task_description, minutes_duration, default_mark, closed_date
			FROM %s
			INNER JOIN %s
			ON %s.id = %s.test_id
			WHERE %s.user_id = $1 AND (is_done = true OR (closed_date < $2 AND closed_date != 0))
			ORDER BY %s.id`,
		table_name.TestsTable, table_name.TestsDateTable, table_name.TestsTable,
		table_name.TestsDateTable, table_name.TestsDateTable, table_name.TestsTable)
	err = r.db.Select(&testsEn, query, userId, timeNow)
	if err != nil {
		return nil, nil, err
	}
	return tests, testsEn, nil
}

func (r *StudentTestAndLabPostgres) GetAllDoneLabs(userId int, timeNow int64) ([]model.LabWithClosedDate, []model.LabWithClosedDate, error) {
	var labs []model.LabWithClosedDate
	query := fmt.Sprintf(`SELECT %s.id, %s.name, %s.task_description, %s.default_mark, closed_date 
			FROM %s 
			INNER JOIN %s 
			ON %s.external_laboratory_id = %s.id
			INNER JOIN %s 
			ON %s.id = %s.laboratory_id
			WHERE %s.user_id = $1 AND (is_done = true OR (closed_date < $2 AND closed_date != 0))
			ORDER BY %s.id`,
		table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.ExternalLaboratoryTable,
		table_name.LaboratoryTable, table_name.LaboratoryTable, table_name.ExternalLaboratoryTable,
		table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.LaboratoryDateTable,
		table_name.LaboratoryTable, table_name.LaboratoryDateTable, table_name.LaboratoryDateTable, table_name.LaboratoryTable)
	err := r.db.Select(&labs, query, userId, timeNow)
	if err != nil {
		return nil, nil, err
	}
	var labsEn []model.LabWithClosedDate
	query = fmt.Sprintf(`SELECT %s.id, %s.name_en as name, %s.task_description_en as task_description, %s.default_mark, closed_date 
			FROM %s 
			INNER JOIN %s 
			ON %s.external_laboratory_id = %s.id
			INNER JOIN %s 
			ON %s.id = %s.laboratory_id
			WHERE %s.user_id = $1 AND (is_done = true OR (closed_date < $2 AND closed_date != 0))
			ORDER BY %s.id`,
		table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.ExternalLaboratoryTable,
		table_name.LaboratoryTable, table_name.LaboratoryTable, table_name.ExternalLaboratoryTable,
		table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.LaboratoryDateTable,
		table_name.LaboratoryTable, table_name.LaboratoryDateTable, table_name.LaboratoryDateTable, table_name.LaboratoryTable)
	err = r.db.Select(&labsEn, query, userId, timeNow)
	if err != nil {
		return nil, nil, err
	}
	return labs, labsEn, nil
}

func (r *StudentTestAndLabPostgres) GetAllOpenedLabs(userId int, timeNow int64) ([]model.LabWithClosedDate, []model.LabWithClosedDate, error) {
	var tests []model.LabWithClosedDate
	query := fmt.Sprintf(`SELECT %s.id, %s.name, %s.task_description, default_mark, closed_date
								FROM %s
								INNER JOIN %s
								ON %s.external_laboratory_id = %s.id
								INNER JOIN %s
								ON %s.id = %s.laboratory_id
								WHERE 
								user_id = $1 AND (is_done = false AND (closed_date > $2 OR closed_date = 0))
								ORDER BY %s.id`,
		table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.ExternalLaboratoryTable,
		table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.LaboratoryTable, table_name.ExternalLaboratoryTable,
		table_name.LaboratoryDateTable, table_name.LaboratoryTable, table_name.LaboratoryDateTable, table_name.LaboratoryTable)
	if err := r.db.Select(&tests, query, userId, timeNow); err != nil {
		return nil, nil, err
	}
	var testsEn []model.LabWithClosedDate
	query = fmt.Sprintf(`SELECT %s.id, %s.name_en as name, %s.task_description_en as task_description, default_mark, closed_date
								FROM %s
								INNER JOIN %s
								ON %s.external_laboratory_id = %s.id
								INNER JOIN %s
								ON %s.id = %s.laboratory_id
								WHERE 
								user_id = $1 AND (is_done = false AND (closed_date > $2 OR closed_date = 0))
								ORDER BY %s.id`,
		table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.ExternalLaboratoryTable,
		table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.LaboratoryTable, table_name.ExternalLaboratoryTable,
		table_name.LaboratoryDateTable, table_name.LaboratoryTable, table_name.LaboratoryDateTable, table_name.LaboratoryTable)
	if err := r.db.Select(&testsEn, query, userId, timeNow); err != nil {
		return nil, nil, err
	}
	return tests, testsEn, nil
}

func (r *StudentTestAndLabPostgres) GetResultOfLab(userId, labId int) (model.LabResult, error) {
	var lab model.LabResult
	query := fmt.Sprintf("SELECT * FROM %s WHERE laboratory_id = $1 AND user_id = $2", table_name.LaboratoryMarkTable)
	if err := r.db.Get(&lab, query, labId, userId); err != nil {
		return model.LabResult{}, err
	}

	return lab, nil
}

func (r *StudentTestAndLabPostgres) GetLabLink(labId int) (model.LabCredential, error) {
	var link model.LabCredential

	query := fmt.Sprintf("SELECT %s.link, %s.token FROM %s INNER JOIN %s ON %s.external_laboratory_id = %s.id WHERE %s.id = $1",
		table_name.ExternalLaboratoryTable, table_name.ExternalLaboratoryTable, table_name.LaboratoryTable,
		table_name.ExternalLaboratoryTable, table_name.LaboratoryTable, table_name.ExternalLaboratoryTable, table_name.LaboratoryTable)
	if err := r.db.Get(&link, query, labId); err != nil {
		return model.LabCredential{}, err
	}

	return link, nil
}
