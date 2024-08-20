package api_lecturer_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type LecturerTestAndLabService struct {
	repo repository.LecturerTestAndLab
}

func NewLecturerTestAndLabService(repo repository.LecturerTestAndLab) *LecturerTestAndLabService {
	return &LecturerTestAndLabService{repo: repo}
}

func (s *LecturerTestAndLabService) GetAllTests() ([]model.Test, []model.Test, error) {
	return s.repo.GetAllTests()
}

func (s *LecturerTestAndLabService) GetAllTestFromSection(sectionId int) ([]model.Test, []model.Test, error) {
	return s.repo.GetAllTestFromSection(sectionId)
}

func (s *LecturerTestAndLabService) GetAllLabFromSection(sectionId int) ([]model.LaboratoryWork, []model.LaboratoryWork, error) {
	return s.repo.GetAllLabFromSection(sectionId)
}

func (s *LecturerTestAndLabService) ChangeLabTaskDescriptionEn(labId int, description string) error {
	return s.repo.ChangeLabTaskDescriptionEn(labId, description)
}

func (s *LecturerTestAndLabService) ChangeLabNameEn(labId int, name string) error {
	return s.repo.ChangeLabNameEn(labId, name)
}

func (s *LecturerTestAndLabService) ChangeTestTaskDescriptionEn(testId int, description string) error {
	return s.repo.ChangeTestTaskDescriptionEn(testId, description)
}

func (s *LecturerTestAndLabService) ChangeTestNameEn(testId int, name string) error {
	return s.repo.ChangeTestNameEn(testId, name)
}

func (s *LecturerTestAndLabService) ChangeTestDuration(testId, minutes int) error {
	return s.repo.ChangeTestDuration(testId, minutes)
}

func (s *LecturerTestAndLabService) ChangeLabDefaultMark(labId, mark int) error {
	return s.repo.ChangeLabDefaultMark(labId, mark)
}

func (s *LecturerTestAndLabService) ChangeLabTaskDescription(labId int, description string) error {
	return s.repo.ChangeLabTaskDescription(labId, description)
}

func (s *LecturerTestAndLabService) ChangeLabName(labId int, name string) error {
	return s.repo.ChangeLabName(labId, name)
}

func (s *LecturerTestAndLabService) ChangeTestTaskDescription(testId int, description string) error {
	return s.repo.ChangeTestTaskDescription(testId, description)
}

func (s *LecturerTestAndLabService) ChangeTestName(testId int, name string) error {
	return s.repo.ChangeTestName(testId, name)
}

func (s *LecturerTestAndLabService) CreateTest(test model.TestAdd) error {
	return s.repo.CreateTest(test)
}

func (s *LecturerTestAndLabService) DeleteTest(testId int) error {
	return s.repo.DeleteTest(testId)
}

func (s *LecturerTestAndLabService) DeleteAnswer(answerId int) error {
	return s.repo.DeleteAnswer(answerId)
}

func (s *LecturerTestAndLabService) ChangeThemeName(themeId int, name string) error {
	return s.repo.ChangeThemeName(themeId, name)
}

func (s *LecturerTestAndLabService) DeleteTheme(themeId int) error {
	return s.repo.DeleteTheme(themeId)
}

func (s *LecturerTestAndLabService) GetAllThemes(testId int) ([]model.ThemeOutput, error) {
	return s.repo.GetAllThemes(testId)
}

func (s *LecturerTestAndLabService) GetQuestions(themeId int) ([]model.Question, []model.Question, error) {
	return s.repo.GetQuestions(themeId)
}

func (s *LecturerTestAndLabService) DeleteQuestion(questionId int) error {
	return s.repo.DeleteQuestion(questionId)
}

func (s *LecturerTestAndLabService) ChangeQuestionName(questionId int, name string) error {
	return s.repo.ChangeQuestionName(questionId, name)
}

func (s *LecturerTestAndLabService) ChangeQuestionNameEn(questionId int, name string) error {
	return s.repo.ChangeQuestionNameEn(questionId, name)
}

func (s *LecturerTestAndLabService) ChangeThemeWeight(themeId, weight int) error {
	return s.repo.ChangeThemeWeight(themeId, weight)
}

func (s *LecturerTestAndLabService) AddAnswerForQuestion(questionId int, name string, nameEn string, isRight bool) error {
	return s.repo.AddAnswerForQuestion(questionId, name, nameEn, isRight)
}

func (s *LecturerTestAndLabService) ChangeAnswerName(answerId int, name string) error {
	return s.repo.ChangeAnswerName(answerId, name)
}

func (s *LecturerTestAndLabService) ChangeAnswerNameEn(answerId int, name string) error {
	return s.repo.ChangeAnswerNameEn(answerId, name)
}

func (s *LecturerTestAndLabService) ChangeAnswerRight(answerId int, isRight bool) error {
	return s.repo.ChangeAnswerRight(answerId, isRight)
}

func (s *LecturerTestAndLabService) GetAnswers(questionId int) ([]model.Answer, []model.Answer, error) {
	return s.repo.GetAnswers(questionId)
}

func (s *LecturerTestAndLabService) AddQuestionForTheme(themeId, questionId int) error {
	return s.repo.AddQuestionForTheme(themeId, questionId)
}

func (s *LecturerTestAndLabService) CreateQuestion(isVariable int, question string, questionEn string) (int, error) {
	if isVariable < 0 || isVariable > 4 {
		err := errors.New("not correct count")
		return 0, err
	}
	return s.repo.CreateQuestion(isVariable, question, questionEn)
}

func (s *LecturerTestAndLabService) GetAllQuestions() ([]model.QuestionWithAmountAnswers, []model.QuestionWithAmountAnswers, error) {
	return s.repo.GetAllQuestions()
}

func (s *LecturerTestAndLabService) GetThemesByQuestion(questionId int) ([]model.Theme, error) {
	return s.repo.GetThemesByQuestion(questionId)
}

func (s *LecturerTestAndLabService) GetAllExistThemes() ([]model.Theme, error) {
	return s.repo.GetAllExistThemes()
}

func (s *LecturerTestAndLabService) AddThemeForTest(testId, themeId, count int) error {
	return s.repo.AddThemeForTest(testId, themeId, count)
}

func (s *LecturerTestAndLabService) CreateTheme(name string, weight int) (int, error) {
	return s.repo.CreateTheme(name, weight)
}

func (s *LecturerTestAndLabService) GetThemeIdByName(name string) (int, error) {
	return s.repo.GetThemeIdByName(name)
}

func (s *LecturerTestAndLabService) DeleteQuestionFromTheme(themeId, questionId int) error {
	return s.repo.DeleteQuestionFromTheme(themeId, questionId)
}

func (s *LecturerTestAndLabService) ChangeThemeTestCount(testId, themeId, count int) error {
	return s.repo.ChangeThemeTestCount(testId, themeId, count)
}

func (s *LecturerTestAndLabService) DeleteThemeFromTest(testId, themeId int) error {
	return s.repo.DeleteThemeFromTest(testId, themeId)
}

func (s *LecturerTestAndLabService) OpenTestForStudent(studentId, testId int, date int64) error {
	return s.repo.OpenTestForStudent(studentId, testId, date)
}

func (s *LecturerTestAndLabService) GetOpenedTestForStudent(studentId, testId int) (model.OpenedTest, error) {
	return s.repo.GetOpenedTestForStudent(studentId, testId)
}

func (s *LecturerTestAndLabService) CloseOpenedTestForStudent(studentId, testId int) error {
	return s.repo.CloseOpenedTestForStudent(studentId, testId)
}

func (s *LecturerTestAndLabService) GetPathForReportTest(userId, testId int) (string, error) {
	path := viper.GetString("test") + "/" + strconv.Itoa(userId) + "-" + strconv.Itoa(testId) + ".txt"

	return path, nil
}

func (s *LecturerTestAndLabService) ChangeTestMarkForStudent(studentId, testId, mark int) error {
	return s.repo.ChangeTestMarkForStudent(studentId, testId, mark)
}

func (s *LecturerTestAndLabService) GetTestMarkForStudent(studentId, testId int) (int, error) {
	return s.repo.GetTestMarkForStudent(studentId, testId)
}

func (s *LecturerTestAndLabService) GetQuestionWithoutEnglishVersion() ([]model.Question, error) {
	return s.repo.GetQuestionWithoutEnglishVersion()
}

func (s *LecturerTestAndLabService) GetQuestionWithoutTheme() ([]model.Question, error) {
	return s.repo.GetQuestionWithoutTheme()
}

func (s *LecturerTestAndLabService) GetQuestionsByName(name string) ([]model.Question, error) {
	return s.repo.GetQuestionsByName(name)
}

func (s *LecturerTestAndLabService) GetUsersWithDoneTests(testId int) ([]model.StudentWithGroupWithClosedDate, error) {
	return s.repo.GetUsersWithDoneTests(testId, time.Now().Unix())
}

func (s *LecturerTestAndLabService) GetUsersWithOpenedTest(testId int) ([]model.StudentWithGroupWithClosedDate, error) {
	return s.repo.GetUsersWithOpenedTest(testId, time.Now().Unix())
}

func (s *LecturerTestAndLabService) CreateExternalLab(lab model.LaboratoryWorkInputWithoutId) (int, error) {
	return s.repo.CreateExternalLab(lab)
}

func (s *LecturerTestAndLabService) DeleteExternalLab(labId int) error {
	return s.repo.DeleteExternalLab(labId)
}

func (s *LecturerTestAndLabService) ChangeLabLinc(labId int, linc string) error {
	return s.repo.ChangeLabLinc(labId, linc)
}

func (s *LecturerTestAndLabService) ChangeLabToken(labId int, token string) error {
	return s.repo.ChangeLabToken(labId, token)
}

func (s *LecturerTestAndLabService) GetExternalLabInfo(labId int) (model.LaboratoryWorkResponse, error) {
	return s.repo.GetExternalLabInfo(labId)
}

func (s *LecturerTestAndLabService) GetAllExternalLab() ([]model.CommonLaboratoryWork, []model.CommonLaboratoryWork, error) {
	return s.repo.GetAllExternalLab()
}

func (s *LecturerTestAndLabService) GetEqualMarkForExport(questionId int) (float64, error) {
	answersRu, _, err := s.repo.GetAnswers(questionId)
	if err != nil {
		return 0, err
	}
	count := 0
	for _, answerRu := range answersRu {
		if *answerRu.IsRight {
			count++
		}
	}
	if len(answersRu) == 0 {
		err := errors.New("не найдены ответы")
		return 0, err
	}
	return float64(count) / float64(len(answersRu)), nil
}

func (s *LecturerTestAndLabService) GetUsersWithOpenedLab(labId int) ([]model.StudentWithGroupWithClosedDate, error) {
	return s.repo.GetUsersWithOpenedLab(labId, time.Now().Unix())
}

func (s *LecturerTestAndLabService) GetUsersWithDoneLaboratory(labId int) ([]model.StudentWithGroupWithClosedDate, error) {
	return s.repo.GetUsersWithDoneLaboratory(labId, time.Now().Unix())
}

func (s *LecturerTestAndLabService) OpenLabForStudent(studentId, labId int, date int64) error {
	externalLab, err := s.repo.GetLabInfo(labId)
	if err != nil {
		return err
	}
	token, err := s.repo.GetLabToken(externalLab)
	if err != nil {
		return err
	}
	lab, err := s.GetExternalLabInfo(externalLab)
	if err != nil {
		return err
	}

	if err := s.sendRequestToOpenLab(lab, studentId, labId, token, true); err != nil {
		return err
	}

	return s.repo.OpenLabForStudent(studentId, labId, date)
}

func (s *LecturerTestAndLabService) CloseOpenedLabForStudent(studentId, labId int) error {
	externalLab, err := s.repo.GetLabInfo(labId)
	if err != nil {
		return err
	}
	token, err := s.repo.GetLabToken(externalLab)
	if err != nil {
		return err
	}
	lab, err := s.GetExternalLabInfo(externalLab)
	if err != nil {
		return err
	}
	if err := s.sendRequestToOpenLab(lab, studentId, labId, token, false); err != nil {
		return err
	}

	return s.repo.CloseOpenedLabForStudent(studentId, labId)
}

func (s *LecturerTestAndLabService) ChangeLabMarkForStudent(studentId, labId, mark int) error {
	return s.repo.ChangeLabMarkForStudent(studentId, labId, mark)
}

func (s *LecturerTestAndLabService) GetLabMarkForStudent(studentId, labId int) (int, error) {
	return s.repo.GetLabMarkForStudent(studentId, labId)
}

func (s *LecturerTestAndLabService) sendRequestToOpenLab(lab model.LaboratoryWorkResponse, userId, labId int, token string, isOpen bool) error {
	url := fmt.Sprintf("%s/%s?user_id=%d&is_open=%t&lab_id=%d", lab.Link, "open", userId, isOpen, labId)
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return err
	}
	req.Header.Add("lab-token", token)
	req.Header.Add("lecturer-token", os.Getenv("LECTURER_HEADER"))

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if _, err := ioutil.ReadAll(res.Body); err != nil {
		return err
	}

	return nil
}
