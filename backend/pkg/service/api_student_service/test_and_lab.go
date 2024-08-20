package api_student_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type StudentTestAndLabService struct {
	repo repository.StudentTestAndLab
}

func NewStudentTestAndLabService(repo repository.StudentTestAndLab) *StudentTestAndLabService {
	return &StudentTestAndLabService{
		repo: repo,
	}
}

func (s *StudentTestAndLabService) GetAllTestFromSection(userId, sectionId int) ([]model.Test, []model.Test, error) {
	if err := s.repo.CheckAccessForSection(userId, sectionId); err != nil {
		return nil, nil, err
	}
	return s.repo.GetAllTestFromSection(userId, sectionId)
}

func (s *StudentTestAndLabService) GetAllLab(userId, sectionId int) ([]model.LaboratoryWork, []model.LaboratoryWork, error) {
	if err := s.repo.CheckAccessForSection(userId, sectionId); err != nil {
		return nil, nil, err
	}
	return s.repo.GetAllLab(userId, sectionId)
}

func (s *StudentTestAndLabService) getQuestionsNumber(testId int) ([]model.QuestionWithTheme, int, error) {
	return s.repo.GetThemesQuestions(testId)
}

func (s *StudentTestAndLabService) GetAllDoneTests(userId int) ([]model.TestWithClosedDate, []model.TestWithClosedDate, error) {
	return s.repo.GetAllDoneTests(userId, time.Now().Unix())
}

func (s *StudentTestAndLabService) GetAllOpenedTests(userId int) ([]model.TestWithClosedDate, []model.TestWithClosedDate, error) {
	return s.repo.GetAllOpenedTests(userId, time.Now().Unix())
}

func (s *StudentTestAndLabService) GetQuestionsForTest(userId, testId int) ([]model.QuestionsWithAnswers, []model.QuestionsWithAnswers, error) {
	if !s.repo.CheckAccessToOpenTest(userId, testId) {
		err := errors.New("отказано в доступе")
		return nil, nil, err
	}

	questionsId, count, err := s.getQuestionsNumber(testId)
	if err != nil {
		err := errors.New("ошибка сервера")
		return nil, nil, err
	}

	questionsRu, questionsEn, err := s.repo.GetQuestionsAndAnswersForTest(questionsId)
	if err != nil {
		err := errors.New("ошибка сервера")
		return nil, nil, err
	}

	if len(questionsRu) != count {
		return nil, nil, fmt.Errorf("ошибка формирования вопросов")
	}

	return questionsRu, questionsEn, nil
}

func (s *StudentTestAndLabService) CheckAnswers(userId, testId int, answers []model.QuestionAndAnswerResponse) (int, []model.QuestionPercentage, error) {
	if !s.repo.CheckAccessToOpenTest(userId, testId) {
		err := errors.New("отказано в доступе")
		return 0, nil, err
	}
	maxPoints, err := s.repo.GetTestMaxPoint(testId)
	if err != nil {
		return 0, nil, err
	}
	maxMark, err := s.repo.GetTestMark(testId)
	if err != nil {
		return 0, nil, err
	}
	var points float64 = 0
	for _, answer := range answers {
		rightAnswers, countAllAnswers, err := s.repo.GetQuestionAnswers(answer.QuestionId)
		if err != nil {
			continue
		}
		equalCount := s.checkEqualCount(countAllAnswers, answer.Answer, rightAnswers)
		points = points + equalCount*float64(s.repo.GetWeightFromQuestionTheme(answer.ThemeId))
	}

	currentPoint := int(math.Round(points * float64(maxMark) / float64(maxPoints)))

	if err := s.repo.DoneOfTest(userId, testId); err != nil {
		return 0, nil, err
	}

	done, err := s.MakeReportForTest(userId, testId, currentPoint, answers)
	if err != nil {
		return 0, nil, err
	}

	if err := s.repo.SaveResultOfTest(userId, testId, currentPoint); err != nil {
		return 0, nil, err
	}

	return currentPoint, done, nil
}

func (s *StudentTestAndLabService) checkEqualCount(countAllAnswers int, currentCounts []string, rightCounts model.GetQuestionWithRightAnswer) float64 {
	var resultCount float64 = 0
	var point float64 = 1 / float64(len(rightCounts.Answers))
	for _, count := range currentCounts {
		var checkCurrentCount float64 = 0
		count = strings.ToLower(count)
		flag := false
		for _, rightCount := range rightCounts.Answers {
			if strings.ToLower(rightCount.Name) == count || strings.ToLower(rightCount.NameEn) == count {
				flag = true
				checkCurrentCount = checkCurrentCount + point
				continue
			}
		}
		if !flag {
			checkCurrentCount = checkCurrentCount - point
		}
		resultCount = resultCount + checkCurrentCount
	}
	if resultCount >= 0 {
		return resultCount
	} else {
		return 0
	}
}

func (s *StudentTestAndLabService) MakeReportForTest(userId, testId, currentPoint int, answers []model.QuestionAndAnswerResponse) ([]model.QuestionPercentage, error) {
	var done []model.QuestionPercentage
	file, err := os.Create(viper.GetString("test") + "/" + strconv.Itoa(userId) + "-" + strconv.Itoa(testId) + ".txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	personalData, err := s.repo.GetPersonalData(userId)
	if err != nil {
		return nil, err
	}

	testRu, _, err := s.repo.GetTest(testId)
	if err != nil {
		return nil, err
	}

	_, err = writer.WriteString(testRu.Name + "\n")
	if err != nil {
		return nil, err
	}

	_, err = writer.WriteString(personalData.GroupName + " " + personalData.Surname + " " + personalData.Name + "\n")
	if err != nil {
		return nil, err
	}

	_, err = writer.WriteString("Оценка:   " + strconv.Itoa(currentPoint) + "/" + strconv.Itoa(testRu.DefaultMark) + "\n\n")
	if err != nil {
		return nil, err
	}

	for i, answer := range answers {
		questionRu, _, err := s.repo.GetQuestion(answer.QuestionId)
		if err != nil {
			continue
		}

		_, err = writer.WriteString(strconv.Itoa(i+1) + ") " + questionRu.Name + "\n")
		if err != nil {
			continue
		}
		if len(answer.Answer) == 0 {
			_, err = writer.WriteString("Выбранный ответ: ()\n")
			if err != nil {
				continue
			}
			_, err = writer.WriteString(fmt.Sprintf("Процент выполнения:   %v%s \n\n", 0, "%"))
			if err != nil {
				continue
			}
			continue
		}
		switch questionRu.IsVariable {
		case 0:

			allAnswers, err := s.repo.GetAnswers(answer.QuestionId)
			if err != nil {
				continue
			}
			_, err = writer.WriteString("Выбранный ответ: ")
			if err != nil {
				continue
			}
			for _, userAnswer := range answer.Answer {
				for _, allAnswer := range allAnswers {
					if userAnswer == allAnswer.Name || userAnswer == allAnswer.NameEn {
						_, err = writer.WriteString("   (" + allAnswer.Name + ")   ")
						if err != nil {
							continue
						}
					} else {
						_, err = writer.WriteString("   (" + userAnswer + ")   ")
						if err != nil {
							continue
						}
					}
				}
			}
			_, err = writer.WriteString("\n")
			if err != nil {
				continue
			}
		case 1:
			_, err = writer.WriteString("Выбранный ответ: " + "   (" + answer.Answer[0] + ")\n")
			if err != nil {
				continue
			}
		case 2:
			_, err = writer.WriteString("Выбранный ответ: " + "   (" + answer.Answer[0] + ")\n")
			if err != nil {
				continue
			}
		case 3:
			_, err = writer.WriteString("Выбранный ответ: " + "   (" + answer.Answer[0] + ")\n")
			if err != nil {
				continue
			}
		case 4:
			_, err = writer.WriteString("Выбранный ответ: " + "   (" + answer.Answer[0] + ")\n")
			if err != nil {
				continue
			}
		default:
			_, err = writer.WriteString("Выбранный ответ: " + "   (" + answer.Answer[0] + ")\n")
			if err != nil {
				continue
			}
		}
		rightAnswers, countAllAnswers, err := s.repo.GetQuestionAnswers(answer.QuestionId)
		if err != nil {
			continue
		}
		equalCount := s.checkEqualCount(countAllAnswers, answer.Answer, rightAnswers)
		pointsForQuestion := equalCount * float64(s.repo.GetWeightFromQuestionTheme(answer.ThemeId))
		_, err = writer.WriteString(fmt.Sprintf("Процент выполнения:   %v%s \n\n", pointsForQuestion/float64(s.repo.GetWeightFromQuestionTheme(answer.ThemeId))*100, "%"))
		if err != nil {
			continue
		}
		done = append(done, model.QuestionPercentage{QuestionId: answer.QuestionId, Percentage: pointsForQuestion / float64(s.repo.GetWeightFromQuestionTheme(answer.ThemeId)) * 100})
	}
	writer.Flush()

	return done, nil
}

func (s *StudentTestAndLabService) GetResultOfTest(userId, testId int) (model.TestResult, error) {
	return s.repo.GetResultOfTest(userId, testId)
}

func (s *StudentTestAndLabService) GetPathForReportTest(userId, testId int) (string, error) {
	path := viper.GetString("test") + "/" + strconv.Itoa(userId) + "-" + strconv.Itoa(testId) + ".txt"

	return path, nil
}

func (s *StudentTestAndLabService) GetAllDoneLabs(userId int) ([]model.LabWithClosedDate, []model.LabWithClosedDate, error) {
	return s.repo.GetAllDoneLabs(userId, time.Now().Unix())
}

func (s *StudentTestAndLabService) GetAllOpenedLabs(userId int) ([]model.LabWithClosedDate, []model.LabWithClosedDate, error) {
	return s.repo.GetAllOpenedLabs(userId, time.Now().Unix())
}

func (s *StudentTestAndLabService) GetResultOfLab(userId, labId int) (model.LabResult, error) {
	return s.repo.GetResultOfLab(userId, labId)
}

func (s *StudentTestAndLabService) GetLinkForLab(userId, labId int) (model.LabCredential, error) {
	access, err := s.repo.CheckAccessToOpenLab(userId, labId)
	if err != nil {
		return model.LabCredential{}, err
	}

	if !access {
		return model.LabCredential{}, err
	}

	link, err := s.repo.GetLabLink(labId)
	if err != nil {
		return model.LabCredential{}, err
	}

	return link, nil
}
