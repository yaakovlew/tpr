package api_lecturer

import (
	"bufio"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"backend/pkg/handler/error_response"
	"backend/pkg/model"
	"backend/pkg/service"
)

type LecturerTestAndLabHandler struct {
	Service service.LecturerTestAndLab
}

func NewLecturerTestAndLabHandler(service service.LecturerTestAndLab) *LecturerTestAndLabHandler {
	return &LecturerTestAndLabHandler{Service: service}
}

// GetAllTestFromSection @Summary get all test from section
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description get all test from section
// @Id get-all-test-from-section
// @Produce json
// @Param id body int true "get all test from section id"
// @Success 200 {object} model.TestResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline/section/test/{id} [get]
func (h *LecturerTestAndLabHandler) GetAllTestFromSection(c *gin.Context) {
	sectionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	tests, testsEn, err := h.Service.GetAllTestFromSection(sectionId)
	if err != nil {
		err = errors.New("ошибка получения тестов")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.TestResponse{Ru: tests, En: testsEn})
}

// ChangeExternalLab @Summary change lab
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description change lab
// @Id change-lab
// @Accept json
// @Produce json
// @Param laboratory_id query string true "laboratory_id"
// @Param name query string false "name"
// @Param description query string false "description"
// @Param name_en query string false "name_en"
// @Param description_en query string false "name_en"
// @Param token query string false "token"
// @Param linc query string false "linc"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/laboratory-work [put]
func (h *LecturerTestAndLabHandler) ChangeExternalLab(c *gin.Context) {
	laboratoryId, err := strconv.Atoi(c.Query("laboratory_id"))
	if err != nil {
		err = errors.New("ошибка изменения лабораторной работы")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	name := c.Query("name")
	description := c.Query("description")
	nameEn := c.Query("name_en")
	descriptionEn := c.Query("description_en")
	token := c.Query("token")
	link := c.Query("linc")

	if token != "" {
		err = h.Service.ChangeLabToken(laboratoryId, token)
		if err != nil {
			err = errors.New("ошибка изменения лабораторной работы")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if link != "" {
		err = h.Service.ChangeLabLinc(laboratoryId, link)
		if err != nil {
			err = errors.New("ошибка изменения лабораторной работы")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if name != "" {
		err = h.Service.ChangeLabName(laboratoryId, name)
		if err != nil {
			err = errors.New("ошибка изменения лабораторной работы")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if nameEn != "" {
		err = h.Service.ChangeLabNameEn(laboratoryId, nameEn)
		if err != nil {
			err = errors.New("ошибка изменения лабораторной работы")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if description != "" {
		err = h.Service.ChangeLabTaskDescription(laboratoryId, description)
		if err != nil {
			err = errors.New("ошибка изменения лабораторной работы")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if descriptionEn != "" {
		err = h.Service.ChangeLabTaskDescriptionEn(laboratoryId, descriptionEn)
		if err != nil {
			err = errors.New("ошибка изменения лабораторной работы")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{})
}

// ChangeTest @Summary change test
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description change test duration
// @Id change-test-duration
// @Accept json
// @Produce json
// @Param test_id query string true "test_id"
// @Param duration query string false "duration"
// @Param name query string false "name"
// @Param name_en query string false "name_en"
// @Param description query string false "description"
// @Param description_en query string false "description_en"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test [put]
func (h *LecturerTestAndLabHandler) ChangeTest(c *gin.Context) {
	test := c.Query("test_id")
	duration := c.Query("duration")
	name := c.Query("name")
	nameEn := c.Query("name_en")
	description := c.Query("description")
	descriptionEn := c.Query("description_en")
	testId, err := strconv.Atoi(test)
	if err != nil {
		err = errors.New("ошибка изменения теста")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if description != "" {
		err := h.Service.ChangeTestTaskDescription(testId, description)
		if err != nil {
			err = errors.New("ошибка изменения описания теста")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if descriptionEn != "" {
		err := h.Service.ChangeTestTaskDescriptionEn(testId, descriptionEn)
		if err != nil {
			err = errors.New("ошибка изменения описания теста")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if name != "" {
		err := h.Service.ChangeTestName(testId, name)
		if err != nil {
			err = errors.New("ошибка изменения названия теста")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if nameEn != "" {
		err := h.Service.ChangeTestNameEn(testId, nameEn)
		if err != nil {
			err = errors.New("ошибка изменения названия теста")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if duration != "" {
		durationInt, err := strconv.Atoi(duration)
		if err != nil {
			err = errors.New("ошибка изменения теста")
			error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		err = h.Service.ChangeTestDuration(testId, durationInt)
		if err != nil {
			err = errors.New("ошибка изменения теста")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetAllTests @Summary get all test
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description get all test
// @Id get-all-test
// @Produce json
// @Success 200 {object} model.TestResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test [get]
func (h *LecturerTestAndLabHandler) GetAllTests(c *gin.Context) {
	tests, testsEn, err := h.Service.GetAllTests()
	if err != nil {
		err = errors.New("ошибка получения тестов")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.TestResponse{Ru: tests, En: testsEn})
}

// CreateTest @Summary create test
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description create test
// @Id create-test
// @Accept json
// @Produce json
// @Param input body model.TestAdd true "create test input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test [post]
func (h *LecturerTestAndLabHandler) CreateTest(c *gin.Context) {
	var test model.TestAdd
	if err := c.BindJSON(&test); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.CreateTest(test)
	if err != nil {
		err = errors.New("ошибка добавления теста")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteTest @Summary delete test
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description delete test
// @Id delete-test
// @Produce json
// @Param id body int true "deleted test id"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/{id} [delete]
func (h *LecturerTestAndLabHandler) DeleteTest(c *gin.Context) {
	testId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.Service.DeleteTest(testId)
	if err != nil {
		err = errors.New("ошибка удаления теста")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// AddThemeForTest @Summary add theme for test
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description add theme for test
// @Id add-theme-for-test
// @Produce json
// @Param input body model.AddedThemeForTestInput true "added theme input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme/add [post]
func (h *LecturerTestAndLabHandler) AddThemeForTest(c *gin.Context) {
	var theme model.AddedThemeForTestInput
	if err := c.BindJSON(&theme); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.AddThemeForTest(theme.TestId, theme.ThemeId, theme.Count)
	if err != nil {
		err = errors.New("ошибка добавления темы к тесту")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ChangeThemeTestCount @Summary change theme test count
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description change theme test count
// @Id change-theme-test-count
// @Produce json
// @Param input body model.AddedThemeForTestInput true "added theme input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme/count [put]
func (h *LecturerTestAndLabHandler) ChangeThemeTestCount(c *gin.Context) {
	var theme model.AddedThemeForTestInput
	if err := c.BindJSON(&theme); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.ChangeThemeTestCount(theme.TestId, theme.ThemeId, theme.Count)
	if err != nil {
		err = errors.New("ошибка изменения количества вопросов по теме")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// CreateTheme @Summary create theme
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description create theme
// @Id create-theme
// @Produce json
// @Param input body model.ThemeInput true "create theme input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme [post]
func (h *LecturerTestAndLabHandler) CreateTheme(c *gin.Context) {
	var theme model.ThemeInput
	if err := c.BindJSON(&theme); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if _, err := h.Service.CreateTheme(theme.Name, theme.Weight); err != nil {
		err = errors.New("ошибка добавления темы")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteAnswer @Summary delete answer
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description delete answer
// @Id delete-answer
// @Produce json
// @Param id body int true "deleted answer id"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme/question/answer/{id} [delete]
func (h *LecturerTestAndLabHandler) DeleteAnswer(c *gin.Context) {
	answerId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.Service.DeleteAnswer(answerId)
	if err != nil {
		err = errors.New("ошибка удаления вопроса")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ChangeTheme @Summary change theme
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description change theme
// @Id change-theme
// @Accept json
// @Produce json
// @Param theme_id query string true "theme_id"
// @Param name query string true "name"
// @Param weight query string true "weight"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme [put]
func (h *LecturerTestAndLabHandler) ChangeTheme(c *gin.Context) {
	themeId, err := strconv.Atoi(c.Query("theme_id"))
	if err != nil {
		err = errors.New("ошибка изменения раздела теста")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	name := c.Query("name")
	weight := c.Query("weight")

	if weight != "" {
		weightInt, err := strconv.Atoi(weight)
		if err != nil {
			err = errors.New("ошибка изменения раздела теста")
			error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		err = h.Service.ChangeThemeWeight(themeId, weightInt)
		if err != nil {
			err = errors.New("ошибка изменения баалов за раздел теста")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if name != "" {
		err := h.Service.ChangeThemeName(themeId, name)
		if err != nil {
			err = errors.New("ошибка изменения раздела теста")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteTheme @Summary delete theme
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description delete theme
// @Id delete-theme
// @Produce json
// @Param id body int true "deleted theme id"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme/{id} [delete]
func (h *LecturerTestAndLabHandler) DeleteTheme(c *gin.Context) {
	themeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.Service.DeleteTheme(themeId)
	if err != nil {
		err = errors.New("ошибка удаления раздела теста")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteThemeFromTest @Summary delete theme from test
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description delete theme from test
// @Id delete-theme-from-test
// @Produce json
// @Param theme_id query string true "theme_id"
// @Param test_id query string true "test_id"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme [delete]
func (h *LecturerTestAndLabHandler) DeleteThemeFromTest(c *gin.Context) {
	theme := c.Query("theme_id")
	test := c.Query("test_id")
	themeId, err := strconv.Atoi(theme)
	if err != nil {
		err = errors.New("ошибка получения темы")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	testId, err := strconv.Atoi(test)
	if err != nil {
		err = errors.New("ошибка получения теста")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.Service.DeleteThemeFromTest(testId, themeId)
	if err != nil {
		err = errors.New("ошибка удаления раздела теста")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetAllThemes @Summary get all themes
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description get all themes
// @Id get-all-themes
// @Produce json
// @Param id body int true "get all themes test id"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme/{id} [get]
func (h *LecturerTestAndLabHandler) GetAllThemes(c *gin.Context) {
	testId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	themes, err := h.Service.GetAllThemes(testId)
	if err != nil {
		fmt.Println(err)
		err = errors.New("ошибка получения разделов теста")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.ThemesResponseOutput{Themes: themes})
}

// GetAllExistThemes @Summary get all exist themes
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description get all exist themes
// @Id get-all-exist-themes
// @Produce json
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme [get]
func (h *LecturerTestAndLabHandler) GetAllExistThemes(c *gin.Context) {
	themes, err := h.Service.GetAllExistThemes()
	if err != nil {
		err = errors.New("ошибка получения разделов")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.ThemesResponse{Themes: themes})
}

// AddQuestionForTheme @Summary add question for theme
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description add question for theme
// @Id add-question-for-theme
// @Accept json
// @Produce json
// @Param input body model.AddQuestion true "add question for theme input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme/question [post]
func (h *LecturerTestAndLabHandler) AddQuestionForTheme(c *gin.Context) {
	var question model.AddQuestion
	if err := c.BindJSON(&question); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.AddQuestionForTheme(question.ThemeId, question.QuestionId)
	if err != nil {
		err = errors.New("ошибка добавления вопроса теста")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteQuestionFromTheme @Summary delete question for theme
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description delete question from theme
// @Id delete-question-from-theme
// @Accept json
// @Produce json
// @Param theme_id query string true "theme_id"
// @Param question_id query string true "question_id"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme/question [delete]
func (h *LecturerTestAndLabHandler) DeleteQuestionFromTheme(c *gin.Context) {
	theme := c.Query("theme_id")
	question := c.Query("question_id")
	themeId, err := strconv.Atoi(theme)
	if err != nil {
		err = errors.New("ошибка получения вопроса")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	questionId, err := strconv.Atoi(question)
	if err != nil {
		err = errors.New("ошибка получения теста")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.Service.DeleteQuestionFromTheme(themeId, questionId)
	if err != nil {
		err = errors.New("ошибка удаления вопроса из темы")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetQuestionsByName @Summary get questions by name
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description get questions by name
// @Id get-questions-by-name
// @Produce json
// @Param name query string true "question name"
// @Success 200 {object} model.Questions
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme/question/name [get]
func (h *LecturerTestAndLabHandler) GetQuestionsByName(c *gin.Context) {
	question := c.Param("name")
	questions, err := h.Service.GetQuestionsByName(question)
	if err != nil {
		err = errors.New("ошибка получения вопросов")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.Questions{
		Questions: questions,
	})
}

// GetQuestions @Summary get questions
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description get questions
// @Id get-questions
// @Produce json
// @Param id body int true "get questions from theme id"
// @Success 200 {object} model.QuestionsResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme/question/{id} [get]
func (h *LecturerTestAndLabHandler) GetQuestions(c *gin.Context) {
	themeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	questions, questionsEn, err := h.Service.GetQuestions(themeId)
	if err != nil {
		err = errors.New("ошибка получения вопросов раздела теста")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.QuestionsResponse{Ru: questions, En: questionsEn})
}

// GetAllQuestions @Summary get questions
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description get all questions
// @Id get-all-questions
// @Produce json
// @Success 200 {object} model.QuestionsResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme/question [get]
func (h *LecturerTestAndLabHandler) GetAllQuestions(c *gin.Context) {
	questions, questionsEn, err := h.Service.GetAllQuestions()
	if err != nil {
		err = errors.New("ошибка получения вопросов")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.QuestionsResponseWithAnswersAmount{Ru: questions, En: questionsEn})
}

// DeleteQuestion @Summary delete question
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description delete question
// @Id delete-question
// @Produce json
// @Param id body int true "deleted question id"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme/question/{id} [delete]
func (h *LecturerTestAndLabHandler) DeleteQuestion(c *gin.Context) {
	questionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.Service.DeleteQuestion(questionId)
	if err != nil {
		err = errors.New("ошибка удаления вопроса")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ChangeQuestion @Summary change question
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description change question
// @Id change-question
// @Produce json
// @Accept json
// @Param question_id query string true "question_id"
// @Param name query string true "name"
// @Param name_en query string true "name_en"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme/question [put]
func (h *LecturerTestAndLabHandler) ChangeQuestion(c *gin.Context) {
	questionId, err := strconv.Atoi(c.Query("question_id"))
	if err != nil {
		err = errors.New("ошибка изменения вопроса")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	name := c.Query("name")
	nameEn := c.Query("name_en")

	if name != "" {
		err := h.Service.ChangeQuestionName(questionId, name)
		if err != nil {
			err = errors.New("ошибка изменения вопроса")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if nameEn != "" {
		err := h.Service.ChangeQuestionNameEn(questionId, nameEn)
		if err != nil {
			err = errors.New("ошибка изменения вопроса")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{})
}

// AddAnswerForQuestion @Summary add answer for question
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description add answer for question
// @Id add-answer-for-question
// @Produce json
// @Accept json
// @Param input body model.AddAnswerInput true "add answer for question input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme/question/answer [post]
func (h *LecturerTestAndLabHandler) AddAnswerForQuestion(c *gin.Context) {
	var answer model.AddAnswerInput
	if err := c.BindJSON(&answer); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.AddAnswerForQuestion(answer.Id, answer.Name, answer.NameEn, *answer.IsRight)
	if err != nil {
		err = errors.New("ошибка добавления ответа")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ChangeAnswer @Summary change answer
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description change answer
// @Id change-answer
// @Produce json
// @Accept json
// @Param answer_id query string true "answer_id"
// @Param name query string true "name"
// @Param nameEn query string true "name_en"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme/question/answer [put]
func (h *LecturerTestAndLabHandler) ChangeAnswer(c *gin.Context) {
	answerId, err := strconv.Atoi(c.Query("answer_id"))
	if err != nil {
		err = errors.New("ошибка изменения ответа")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	name := c.Query("name")
	nameEn := c.Query("name_en")
	isRight := c.Query("is_right")

	if name != "" {
		err := h.Service.ChangeAnswerName(answerId, name)
		if err != nil {
			err = errors.New("ошибка изменения ответа")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if nameEn != "" {
		err := h.Service.ChangeAnswerNameEn(answerId, nameEn)
		if err != nil {
			err = errors.New("ошибка изменения ответа")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if isRight != "" {
		isRightCheck := false
		if isRight == "1" {
			isRightCheck = true
		} else if isRight == "0" {
			isRightCheck = false
		} else {
			err = errors.New("ошибка изменения ответа")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		err := h.Service.ChangeAnswerRight(answerId, isRightCheck)
		if err != nil {
			err = errors.New("ошибка изменения ответа")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetAnswers @Summary get answers
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description get answers
// @Id get-answers
// @Produce json
// @Param id body int true "question id"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme/question/answer/{id} [get]
func (h *LecturerTestAndLabHandler) GetAnswers(c *gin.Context) {
	questionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	answers, answersEn, err := h.Service.GetAnswers(questionId)
	if err != nil {
		err = errors.New("ошибка получения ответов")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.AnswersResponse{Ru: answers, En: answersEn})
}

// OpenTest @Summary open test for user
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description open test for user
// @Id open-test-for-user-lecturer
// @Accept json
// @Produce json
// @Param input body model.OpenTestInput true "open test for user"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/activity [post]
func (h *LecturerTestAndLabHandler) OpenTest(c *gin.Context) {
	var test model.OpenTestInput
	if err := c.BindJSON(&test); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.OpenTestForStudent(test.UserId, test.TestId, test.Date)
	if err != nil {
		err = errors.New("ошибка открытия теста")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetOpenedTestForStudent @Summary get opened test for student
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description get opened test for user
// @Id get-opened-test-for-user-lecturer
// @Accept json
// @Produce json
// @Param user_id query string true "user_id"
// @Param test_id query string true "test_id"
// @Success 200 {object} model.OpenedTest
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/activity/get [get]
func (h *LecturerTestAndLabHandler) GetOpenedTestForStudent(c *gin.Context) {
	user := c.Query("user_id")
	test := c.Query("test_id")
	userId, err := strconv.Atoi(user)
	if err != nil {
		err = errors.New("ошибка получения пользователя")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	testId, err := strconv.Atoi(test)
	if err != nil {
		err = errors.New("ошибка получения теста")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	testTake, err := h.Service.GetOpenedTestForStudent(userId, testId)
	if err != nil {
		err = errors.New("ошибка открытия теста")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, testTake)
}

// CloseOpenedTestForStudent @Summary close opened test for student
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description close opened test for student
// @Id close-opened-test-for-student-lecturer
// @Accept json
// @Produce json
// @Param user_id query string true "user_id"
// @Param test_id query string true "test_id"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/activity [delete]
func (h *LecturerTestAndLabHandler) CloseOpenedTestForStudent(c *gin.Context) {
	user := c.Query("user_id")
	test := c.Query("test_id")
	userId, err := strconv.Atoi(user)
	if err != nil {
		err = errors.New("ошибка получения пользователя")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	testId, err := strconv.Atoi(test)
	if err != nil {
		err = errors.New("ошибка получения теста")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.Service.CloseOpenedTestForStudent(userId, testId); err != nil {
		err = errors.New("ошибка закрытия теста")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// CreateQuestion @Summary create question
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description create question
// @Id create-question
// @Accept json
// @Produce json
// @Param input body model.QuestionInput true "create question"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme/question/create [post]
func (h *LecturerTestAndLabHandler) CreateQuestion(c *gin.Context) {
	var question model.QuestionInput
	if err := c.BindJSON(&question); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_, err := h.Service.CreateQuestion(question.IsVariable, question.Name, question.NameEn)
	if err != nil {
		err = errors.New("ошибка добавления вопроса")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetReportForTest @Summary get report for test
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description get report for test
// @Id get-report-for-test-lecturer
// @Produce json
// @Param user_id query string true "user_id"
// @Param test_id query string true "test_id"
// @Success 200 {file} file
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/activity/report [get]
func (h *LecturerTestAndLabHandler) GetReportForTest(c *gin.Context) {
	user := c.Query("user_id")
	test := c.Query("test_id")
	userId, err := strconv.Atoi(user)
	if err != nil {
		err = errors.New("ошибка получения пользователя")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	testId, err := strconv.Atoi(test)
	if err != nil {
		err = errors.New("ошибка получения теста")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	path, err := h.Service.GetPathForReportTest(userId, testId)
	if err != nil {
		err = errors.New("ошибка сервера: возможно отсутствуют данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.FileAttachment(path, "file")
}

// GetMarkTestForStudent @Summary get mark test for student
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description get student mark test
// @Id get-student-mark-test-lecturer
// @Accept json
// @Produce json
// @Param input body model.CloseTestInput true "get student mark test"
// @Success 200 {integer} integer 1
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/mark [get]
func (h *LecturerTestAndLabHandler) GetMarkTestForStudent(c *gin.Context) {
	user := c.Query("user_id")
	test := c.Query("test_id")
	userId, err := strconv.Atoi(user)
	if err != nil {
		err = errors.New("ошибка получения пользователя")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	testId, err := strconv.Atoi(test)
	if err != nil {
		err = errors.New("ошибка получения теста")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	mark, err := h.Service.GetTestMarkForStudent(userId, testId)
	if err != nil {
		err = errors.New("ошибка получения оценки")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"mark": mark,
	})
}

// ChangeMarkTestForStudent @Summary change mark test for student
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description change mark test for student
// @Id change-mark-test-for-student-lecturer
// @Accept json
// @Produce json
// @Param input body model.ChangeTestMark true "change mark test for student"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/mark [put]
func (h *LecturerTestAndLabHandler) ChangeMarkTestForStudent(c *gin.Context) {
	var test model.ChangeTestMark
	if err := c.BindJSON(&test); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.ChangeTestMarkForStudent(test.UserId, test.TestId, test.Mark)
	if err != nil {
		err = errors.New("ошибка изменения оценки")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetQuestionWithoutEnglishVersion @Summary get questions without english version
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description get questions without english version
// @Id get-questions-without-english-version
// @Produce json
// @Success 200 {object} model.QuestionsResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme/question/without-english [get]
func (h *LecturerTestAndLabHandler) GetQuestionWithoutEnglishVersion(c *gin.Context) {
	questions, err := h.Service.GetQuestionWithoutEnglishVersion()
	if err != nil {
		err = errors.New("ошибка получения вопросов")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.QuestionResponse{Questions: questions})
}

// GetQuestionWithoutTheme @Summary get questions without theme
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description get questions without theme
// @Id get-questions-without-theme
// @Produce json
// @Success 200 {object} model.QuestionsResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme/question/without-theme [get]
func (h *LecturerTestAndLabHandler) GetQuestionWithoutTheme(c *gin.Context) {
	questions, err := h.Service.GetQuestionWithoutTheme()
	if err != nil {
		err = errors.New("ошибка получения вопросов")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.QuestionResponse{Questions: questions})
}

// ExportTheme @Summary export theme file
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description export theme file
// @Id export-theme-file
// @Produce json
// @Param input body model.ThemeInputForMultiLanguage true "export theme file"
// @Success 200 {file} file
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme/export [get]
func (h *LecturerTestAndLabHandler) ExportTheme(c *gin.Context) {
	var input model.ThemeInputForMultiLanguage
	if err := c.BindJSON(&input); err != nil {
		err = errors.New("ошибка запроса: неверные данные выбранной темы")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	file, err := os.Create("file.txt")
	if err != nil {
		err = errors.New("ошибка получения файла")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	questionsRu, questionsEn, err := h.Service.GetQuestions(input.ThemeId)
	if err != nil {
		err = errors.New("ошибка получения вопросов")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	writer := bufio.NewWriter(file)

	for _, question := range questionsRu {
		answersRu, answersEn, err := h.Service.GetAnswers(question.QuestionId)
		if err != nil {
			continue
		}
		switch *input.IsMultiLanguage {
		case true:
			var questionMulti model.Question
			for _, questionEn := range questionsEn {
				if question.QuestionId == questionEn.QuestionId {
					questionMulti = questionEn
				}
			}
			if questionMulti == (model.Question{}) {
				_, err = writer.WriteString(fmt.Sprintf("::%d::%s{\n", question.QuestionId, question.Name))
				if err != nil {
					continue
				}
			} else {
				_, err = writer.WriteString(fmt.Sprintf("::%d::%s$$%s{\n", question.QuestionId, question.Name, questionMulti.Name))
				if err != nil {
					continue
				}
			}
			for _, answerRu := range answersRu {
				var answerCheck model.Answer
				for _, answerEn := range answersEn {
					if answerRu.AnswerId == answerEn.AnswerId {
						answerCheck = answerEn
					}
				}
				if answerCheck == (model.Answer{}) {
					switch question.IsVariable {
					case 0:
						_, err = writer.WriteString("	=%100%" + fmt.Sprintf("%s#\n", answerRu.Name))
						if err != nil {
							continue
						}
					case 1:
						if *answerRu.IsRight {
							_, err = writer.WriteString(fmt.Sprintf("	=%s\n", answerRu.Name))
							if err != nil {
								continue
							}
						} else {
							_, err = writer.WriteString(fmt.Sprintf("	~%s\n", answerRu.Name))
							if err != nil {
								continue
							}
						}
					case 2:
						mark, err := h.Service.GetEqualMarkForExport(question.QuestionId)
						if err != nil {
							continue
						}
						if *answerRu.IsRight {
							_, err = writer.WriteString("~%" + fmt.Sprintf("	%f%s%s\n", mark, "%", answerRu.Name))
							if err != nil {
								continue
							}
						} else {
							_, err = writer.WriteString("~%" + fmt.Sprintf("	%f%s%s\n", -mark, "%", answerRu.Name))

							if err != nil {
								continue
							}
						}
					default:
						continue
					}
				} else {
					switch question.IsVariable {
					case 0:
						_, err = writer.WriteString("	=%100%" + fmt.Sprintf("%s$$%s#\n", answerRu.Name, answerCheck.Name))
						if err != nil {
							continue
						}
					case 1:
						if *answerRu.IsRight {
							_, err = writer.WriteString(fmt.Sprintf("	=%s$$%s\n", answerRu.Name, answerCheck.Name))
							if err != nil {
								continue
							}
						} else {
							_, err = writer.WriteString(fmt.Sprintf("	~%s$$%s\n", answerRu.Name, answerCheck.Name))
							if err != nil {
								continue
							}
						}
					case 2:
						mark, err := h.Service.GetEqualMarkForExport(question.QuestionId)
						if err != nil {
							continue
						}
						if *answerRu.IsRight {
							_, err = writer.WriteString("	~%" + fmt.Sprintf("%f%s%s$$%s\n", mark, "%", answerRu.Name, answerCheck.Name))
							if err != nil {
								continue
							}
						} else {
							_, err = writer.WriteString("	~%" + fmt.Sprintf("%f%s%s$$%s\n", -mark, "%", answerRu.Name, answerCheck.Name))
							if err != nil {
								continue
							}
						}
					default:
						continue
					}
				}
			}
		case false:
			_, err = writer.WriteString(fmt.Sprintf("::%d::%s{\n", question.QuestionId, question.Name))
			if err != nil {
				continue
			}
			for _, answerRu := range answersRu {
				switch question.IsVariable {
				case 0:
					_, err = writer.WriteString("	=%100%" + fmt.Sprintf("%s#\n", answerRu.Name))
					if err != nil {
						continue
					}
				case 1:
					if *answerRu.IsRight {
						_, err = writer.WriteString(fmt.Sprintf("	=%s\n", answerRu.Name))
						if err != nil {
							continue
						}
					} else {
						_, err = writer.WriteString(fmt.Sprintf("	~%s\n", answerRu.Name))
						if err != nil {
							continue
						}
					}
				case 2:
					mark, err := h.Service.GetEqualMarkForExport(question.QuestionId)
					if err != nil {
						continue
					}
					if *answerRu.IsRight {
						_, err = writer.WriteString("	~%" + fmt.Sprintf("%f%s%s\n", mark, "%", answerRu.Name))
						if err != nil {
							continue
						}
					} else {
						_, err = writer.WriteString("	~%" + fmt.Sprintf("%f%s%s\n", -mark, "%", answerRu.Name))
						if err != nil {
							continue
						}
					}
				default:
					continue
				}
			}
		}
		_, err = writer.WriteString("}\n\n")
		if err != nil {
			continue
		}
	}
	writer.Flush()
	c.FileAttachment("./file.txt", "file.txt")
	if err := os.Remove("file.txt"); err != nil {
		err = errors.New("ошибка получения файла")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

// ImportQuestions @Summary import questions
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description import questions
// @Id import-questions
// @Produce json
// @Success 200 {file} file
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme/question/import [post]
func (h *LecturerTestAndLabHandler) ImportQuestions(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		err = errors.New("ошибка загрузки файла")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ext := filepath.Ext(file.Filename)
	if ext != ".txt" {
		error_response.NewErrorResponse(c, http.StatusBadRequest, errors.New("неправильный формат файла").Error())
		return
	}

	f, err := file.Open()
	if err != nil {
		error_response.NewErrorResponse(c, http.StatusBadRequest, errors.New("ошибка открытия файла").Error())
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var question model.QuestionInput
	var answers []model.AddAnswerInput
	i := 0
	themeId := 0
	for scanner.Scan() {
		str := scanner.Text()
		deleteWord1 := "[html]"
		deleteWord2 := "[moodle]"
		deleteWord3 := "<p>"
		deleteWord4 := "</p>"
		str = strings.Replace(str, deleteWord1, "", -1)
		str = strings.Replace(str, deleteWord2, "", -1)
		str = strings.Replace(str, deleteWord3, "", -1)
		str = strings.Replace(str, deleteWord4, "", -1)
		i++
		if i == 1 {
			if strings.Contains(str, "Тема:") {
				str = strings.Replace(str, "Тема:", "", -1)
				checkThemeId, err := h.Service.GetThemeIdByName(strings.TrimSpace(str))
				if err != nil || checkThemeId == 0 {
					if id, err := h.Service.CreateTheme(strings.TrimSpace(str), 1); err != nil {
						continue
					} else {
						themeId = id
					}
				} else {
					themeId = checkThemeId
				}
				continue
			}
		}
		if len(str) < 1 {
			continue
		}
		if len(str) == 1 {
			if str[0] == '}' {
				id, err := h.Service.CreateQuestion(question.IsVariable, question.Name, question.NameEn)
				if err != nil {
					continue
				}
				for _, answer := range answers {
					if err := h.Service.AddAnswerForQuestion(id, answer.Name, answer.NameEn, *answer.IsRight); err != nil {
						continue
					}
				}
				if themeId != 0 {
					if err := h.Service.AddQuestionForTheme(themeId, id); err != nil {
						continue
					}
				}
				question = model.QuestionInput{}
				answers = []model.AddAnswerInput{}
				continue
			} else {
				continue
			}
		}
		if str[:2] == "::" {
			resultStr := strings.Replace(str, "{", "", -1)
			questionNames := strings.Split(resultStr, "$$")
			question.Name = strings.Split(questionNames[0], "::")[2]
			question.Name = question.Name[:len(question.Name)]
			if len(questionNames) == 2 {
				question.NameEn = strings.Replace(questionNames[1], "{", "", -1)
				question.NameEn = strings.Replace(question.NameEn, "#", "", -1)
			}
		}

		if len(str) < 7 {
			if len(str) < 4 {
				if len(str) < 3 {
					if str[1] == '~' {
						question.IsVariable = 1
						var answer model.AddAnswerInput
						isRight := false
						answer.IsRight = &isRight
						check := strings.Split(str, "$$")
						if len(check) >= 2 {
							if string(check[1][len(check[1])-1]) == "," {
								check[1] = check[1][:len(check[1])-1]
							}
							answer.NameEn = check[1]
						}
						check = strings.Split(check[0], "~")
						if string(check[1][len(check[1])]) == "," {
							check[1] = check[1][:(len(check[1]) - 1)]
						}
						answer.Name = check[1]
						answers = append(answers, answer)
					} else if str[1] == '=' {
						question.IsVariable = 1
						var answer model.AddAnswerInput
						isRight := true
						answer.IsRight = &isRight
						check := strings.Split(str, "$$")
						if len(check) >= 2 {
							if string(check[1][len(check[1])-1]) == "," {
								check[1] = check[1][:len(check[1])-1]
							}
							answer.NameEn = check[1]
						}
						check = strings.Split(check[0], "=")
						if string(check[1][len(check[1])-1]) == "," {
							check[1] = check[1][:len(check[1])-1]
						}
						answer.Name = check[1]
						answers = append(answers, answer)
					}
				} else {
					if str[1:3] == "~%" {
						question.IsVariable = 2
						var answer model.AddAnswerInput
						isRight := true
						answer.IsRight = &isRight
						check := strings.Split(str, "$$")
						if len(check) >= 2 {
							if string(check[1][len(check[1])-1]) == "," {
								check[1] = check[1][:len(check[1])-1]
							}
							answer.NameEn = check[1]
						}
						check = strings.Split(check[0], "~")
						check[1] = strings.Split(check[1], "%")[2]
						if string(check[1][len(check[1])-1]) == "," {
							check[1] = check[1][:len(check[1])-1]
						}
						answer.Name = check[1]
						answers = append(answers, answer)
					} else if str[1] == '~' {
						question.IsVariable = 1
						var answer model.AddAnswerInput
						isRight := false
						answer.IsRight = &isRight
						check := strings.Split(str, "$$")
						if len(check) >= 2 {
							if string(check[1][len(check[1])-1]) == "," {
								check[1] = check[1][:len(check[1])-1]
							}
							answer.NameEn = check[1]
						}
						check = strings.Split(check[0], "~")
						if check[1][(len(check[1])-1):(len(check[1]))] == "," {
							check[1] = check[1][:(len(check[1]) - 1)]
						}
						answer.Name = check[1]
						answers = append(answers, answer)
					} else if str[1] == '=' {
						question.IsVariable = 1
						var answer model.AddAnswerInput
						isRight := true
						answer.IsRight = &isRight
						check := strings.Split(str, "$$")
						if len(check) >= 2 {
							if string(check[1][len(check[1])-1]) == "," {
								check[1] = check[1][:len(check[1])-1]
							}
							answer.NameEn = check[1]
						}
						check = strings.Split(check[0], "=")
						if string(check[1][len(check[1])-1]) == "," {
							check[1] = check[1][:len(check[1])-1]
						}
						answer.Name = check[1]
						answers = append(answers, answer)
					}
				}
			} else {
				if str[1:4] == "~%-" {
					question.IsVariable = 2
					var answer model.AddAnswerInput
					isRight := false
					answer.IsRight = &isRight
					check := strings.Split(str, "$$")
					if len(check) >= 2 {
						answer.NameEn = check[1]
					}
					check = strings.Split(check[0], "~")
					check[1] = strings.Split(check[1], "%")[2]
					answer.Name = check[1]
					answers = append(answers, answer)
				} else if str[1:3] == "~%" {
					question.IsVariable = 2
					var answer model.AddAnswerInput
					isRight := true
					answer.IsRight = &isRight
					check := strings.Split(str, "$$")
					if len(check) >= 2 {
						answer.NameEn = check[1]
					}
					check = strings.Split(check[0], "~")
					check[1] = strings.Split(check[1], "%")[2]
					answer.Name = check[1]
					answers = append(answers, answer)
				} else if str[1] == '~' {
					question.IsVariable = 1
					var answer model.AddAnswerInput
					isRight := false
					answer.IsRight = &isRight
					check := strings.Split(str, "$$")
					if len(check) >= 2 {
						answer.NameEn = check[1]
					}
					check = strings.Split(check[0], "~")
					if check[1][(len(check[1])-1):(len(check[1]))] == "," {
						check[1] = check[1][:(len(check[1]) - 1)]
					}
					answer.Name = check[1]
					answers = append(answers, answer)
				} else if str[1] == '=' {
					question.IsVariable = 1
					var answer model.AddAnswerInput
					isRight := true
					answer.IsRight = &isRight
					check := strings.Split(str, "$$")
					if len(check) >= 2 {
						answer.NameEn = check[1]
					}
					check = strings.Split(check[0], "=")
					answer.Name = check[1]
					answers = append(answers, answer)
				}
			}
		} else {
			if str[1:7] == "=%100%" {
				question.IsVariable = 0
				var answer model.AddAnswerInput
				isRight := true
				answer.IsRight = &isRight
				check := strings.Split(str, "$$")
				if len(check) >= 2 {
					answer.NameEn = check[1]
				}
				check = strings.Split(check[0], "=%100%")
				answer.Name = check[1]
				answers = append(answers, answer)
			} else if str[1:4] == "~%-" {
				question.IsVariable = 2
				var answer model.AddAnswerInput
				isRight := false
				answer.IsRight = &isRight
				check := strings.Split(str, "$$")
				if len(check) >= 2 {
					answer.NameEn = check[1]
				}
				check = strings.Split(check[0], "~")
				check[1] = strings.Split(check[1], "%")[2]
				answer.Name = check[1]
				answers = append(answers, answer)
			} else if str[1:3] == "~%" {
				question.IsVariable = 2
				var answer model.AddAnswerInput
				isRight := true
				answer.IsRight = &isRight
				check := strings.Split(str, "$$")
				if len(check) >= 2 {
					answer.NameEn = check[1]
				}
				check = strings.Split(check[0], "~")
				check[1] = strings.Split(check[1], "%")[2]
				answer.Name = check[1]
				answers = append(answers, answer)
			} else if str[1] == '~' {
				question.IsVariable = 1
				var answer model.AddAnswerInput
				isRight := false
				answer.IsRight = &isRight
				check := strings.Split(str, "$$")
				if len(check) >= 2 {
					answer.NameEn = check[1]
				}
				check = strings.Split(check[0], "~")
				if check[1][(len(check[1])-1):(len(check[1]))] == "," {
					check[1] = check[1][:(len(check[1]) - 1)]
				}
				answer.Name = check[1]
				answers = append(answers, answer)
			} else if str[1] == '=' {
				question.IsVariable = 1
				var answer model.AddAnswerInput
				isRight := true
				answer.IsRight = &isRight
				check := strings.Split(str, "$$")
				if len(check) >= 2 {
					answer.NameEn = check[1]
				}
				check = strings.Split(check[0], "=")
				if string(check[1][len(check[1])-1]) == "," {
					check[1] = check[1][:len(check[1])-1]
				}
				answer.Name = check[1]
				answers = append(answers, answer)
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetUsersWithDoneTests @Summary get all students with done test
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description get all students with done test
// @Id get-all-students-with-done-test
// @Produce json
// @Param test_id query string true "test_id"
// @Param is_done query string true "is_done"
// @Success 200 {object} model.StudentWithGroupWithClosedDateResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/students/done/{id} [get]
func (h *LecturerTestAndLabHandler) GetUsersWithDoneTests(c *gin.Context) {
	testId, err := strconv.Atoi(c.Query("test_id"))
	if err != nil {
		err = errors.New("ошибка получения теста")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	isDone := c.Query("is_done")
	if isDone == "1" {
		users, err := h.Service.GetUsersWithDoneTests(testId)
		if err != nil {
			err = errors.New("ошибка получения студентов")
			error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, model.StudentWithGroupWithClosedDateResponse{Students: users})
		return
	} else {
		users, err := h.Service.GetUsersWithOpenedTest(testId)
		if err != nil {
			err = errors.New("ошибка получения студентов")
			error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, model.StudentWithGroupWithClosedDateResponse{Students: users})
		return
	}
}

// GetAllThemesByQuestion @Summary get all themes by question
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description get all themes by question
// @Id get-all-themes-by-question
// @Produce json
// @Param id body int true "get all themes by question"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/test/theme/question/by/{id} [get]
func (h *LecturerTestAndLabHandler) GetAllThemesByQuestion(c *gin.Context) {
	questionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	themes, err := h.Service.GetThemesByQuestion(questionId)
	if err != nil {
		fmt.Println(err)
		err = errors.New("ошибка получения разделов теста")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.ThemesResponse{Themes: themes})
}

// DeleteExternalLab @Summary delete lab
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description delete lab
// @Id delete-lab
// @Produce json
// @Param id body int true "deleted lab id"
// @Success 200 {integer} integer 1
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/laboratory-work/{id} [delete]
func (h *LecturerTestAndLabHandler) DeleteExternalLab(c *gin.Context) {
	labId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.Service.DeleteExternalLab(labId); err != nil {
		fmt.Println(err)
		err = errors.New("ошибка удаления лабораторной работы")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ChangeLabLinc @Summary change lab linc
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description change lab linc
// @Id change-lab-linc
// @Accept json
// @Produce json
// @Param input body model.ChangeLabLincInput true "change lab linc input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/laboratory-work/linc [put]
func (h *LecturerTestAndLabHandler) ChangeLabLinc(c *gin.Context) {
	var laboratory model.ChangeLabLincInput
	if err := c.BindJSON(&laboratory); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.ChangeLabLinc(laboratory.LaboratoryId, laboratory.Linc)
	if err != nil {
		err = errors.New("ошибка изменения ссылки на лабораторную работу")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ChangeLabToken @Summary change lab token
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description change lab token
// @Id change-lab-token
// @Accept json
// @Produce json
// @Param input body model.ChangeLabLincInput true "change lab token input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/laboratory-work/token [put]
func (h *LecturerTestAndLabHandler) ChangeLabToken(c *gin.Context) {
	var laboratory model.ChangeLabTokenInput
	if err := c.BindJSON(&laboratory); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.ChangeLabToken(laboratory.LaboratoryId, laboratory.Token)
	if err != nil {
		err = errors.New("ошибка изменения токена доступа для лабораторной работы")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetExternalLabInfo @Summary get lab info
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description get lab info
// @Id get-lab-info
// @Accept json
// @Produce json
// @Param id body int true "lab id"
// @Success 200 {object} model.LaboratoryWorkResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/laboratory-work/{id} [get]
func (h *LecturerTestAndLabHandler) GetExternalLabInfo(c *gin.Context) {
	labId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	lab, err := h.Service.GetExternalLabInfo(labId)
	if err != nil {
		err = errors.New("ошибка получения лабораторной работы")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, lab)
}

// GetUsersWithDoneLabs @Summary get all students with done lab
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description get all students with done lab
// @Id get-all-students-with-done-lab
// @Produce json
// @Param laboratory_id query string true "laboratory_id"
// @Param is_done query string false "is_done"
// @Success 200 {object} model.StudentWithGroupWithClosedDateResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/laboratory-work/students/done/{id} [get]
func (h *LecturerTestAndLabHandler) GetUsersWithDoneLabs(c *gin.Context) {
	laboratoryId, err := strconv.Atoi(c.Query("laboratory_id"))
	if err != nil {
		err = errors.New("ошибка получения лабораторной работы")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	isDone := c.Query("is_done")
	if isDone == "1" {
		users, err := h.Service.GetUsersWithDoneLaboratory(laboratoryId)
		if err != nil {
			err = errors.New("ошибка получения студентов")
			error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, model.StudentWithGroupWithClosedDateResponse{Students: users})
		return
	} else {
		users, err := h.Service.GetUsersWithOpenedLab(laboratoryId)
		if err != nil {
			err = errors.New("ошибка получения студентов")
			error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, model.StudentWithGroupWithClosedDateResponse{Students: users})
		return
	}
}

// OpenLab @Summary open lab for user
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description open lab for user
// @Id open-lab-for-user-lecturer
// @Accept json
// @Produce json
// @Param input body model.OpenLabInput true "open lab for user"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/laboratory-work/activity [post]
func (h *LecturerTestAndLabHandler) OpenLab(c *gin.Context) {
	var lab model.OpenLabInput
	if err := c.BindJSON(&lab); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.OpenLabForStudent(lab.UserId, lab.LabId, lab.Date)
	if err != nil {
		err = errors.New("ошибка открытия лабораторной работы")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// CloseOpenedLabForStudent @Summary close opened lab for student
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description close opened lab for student
// @Id close-opened-lab-for-student-lecturer
// @Accept json
// @Produce json
// @Param user_id query string true "user_id"
// @Param lab_id query string true "lab_id"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/laboratory-work/activity [delete]
func (h *LecturerTestAndLabHandler) CloseOpenedLabForStudent(c *gin.Context) {
	user := c.Query("user_id")
	lab := c.Query("laboratory_id")
	userId, err := strconv.Atoi(user)
	if err != nil {
		err = errors.New("ошибка получения пользователя")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	labId, err := strconv.Atoi(lab)
	if err != nil {
		err = errors.New("ошибка получения лабораторной работы")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.Service.CloseOpenedLabForStudent(userId, labId); err != nil {
		err = errors.New("ошибка закрытия лабораторной работы")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetMarkLabForStudent @Summary get mark lab for student
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description get student mark lab
// @Id get-student-mark-lab-lecturer
// @Accept json
// @Produce json
// @Param input body model.CloseLabInput true "get student mark lab"
// @Success 200 {integer} integer 1
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/laboratory-work/mark [get]
func (h *LecturerTestAndLabHandler) GetMarkLabForStudent(c *gin.Context) {
	user := c.Query("user_id")
	lab := c.Query("laboratory_id")
	userId, err := strconv.Atoi(user)
	if err != nil {
		err = errors.New("ошибка получения пользователя")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	labId, err := strconv.Atoi(lab)
	if err != nil {
		err = errors.New("ошибка получения оценки лабораторной работы")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	mark, err := h.Service.GetLabMarkForStudent(userId, labId)
	if err != nil {
		err = errors.New("ошибка получения оценки")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"mark": mark,
	})
}

// ChangeMarkLabForStudent @Summary change mark lab for student
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description change mark lab for student
// @Id change-mark-lab-for-student-lecturer
// @Accept json
// @Produce json
// @Param input body model.ChangeLabMark true "change mark lab for student"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/laboratory-work/mark [put]
func (h *LecturerTestAndLabHandler) ChangeMarkLabForStudent(c *gin.Context) {
	var test model.ChangeLabMark
	if err := c.BindJSON(&test); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.ChangeLabMarkForStudent(test.UserId, test.LabId, test.Mark)
	if err != nil {
		err = errors.New("ошибка изменения оценки")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// CreateExternalLab @Summary create lab
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description create lab
// @Id create-lab
// @Produce json
// @Param lab body model.LaboratoryWorkInputWithoutId true "create lab input"
// @Success 200 {int} id
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/laboratory-work/external [post]
func (h *LecturerTestAndLabHandler) CreateExternalLab(c *gin.Context) {
	var lab model.LaboratoryWorkInputWithoutId
	if err := c.BindJSON(&lab); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	labId, err := h.Service.CreateExternalLab(lab)
	if err != nil {
		err = errors.New("ошибка создания лабораторной работы")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"laboratory_id": labId,
	})
}

// GetAllExternalLab @Summary get all lab
// @Security ApiKeyAuthLecturer
// @Tags test and lab
// @Description get all labs
// @Id get-all-lab
// @Produce json
// @Success 200 {object} model.LabsResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/laboratory-work/external [get]
func (h *LecturerTestAndLabHandler) GetAllExternalLab(c *gin.Context) {
	labs, labsEn, err := h.Service.GetAllExternalLab()
	if err != nil {
		err = errors.New("ошибка получения лабораторных работ")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Ru": labs,
		"En": labsEn,
	})
}
