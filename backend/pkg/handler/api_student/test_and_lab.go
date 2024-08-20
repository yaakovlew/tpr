package api_student

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"backend/pkg/handler/error_response"
	"backend/pkg/handler/middleware"
	"backend/pkg/model"
	"backend/pkg/service"
)

const (
	tokenForLab = "lab-token"
)

type StudentTestAndLabHandler struct {
	Service service.StudentTestAndLab
}

func NewStudentTestAndLabHandler(service service.StudentTestAndLab) *StudentTestAndLabHandler {
	return &StudentTestAndLabHandler{Service: service}
}

// GetAllTestFromSection @Summary get all test from section
// @Security ApiKeyAuthStudent
// @Tags test and lab
// @Description get all test from section
// @Id get-all-test-from-section-student
// @Produce json
// @Param id body int true "get all test from section id"
// @Success 200 {object} model.TestResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/student/disciplines/test/{id} [get]
func (h *StudentTestAndLabHandler) GetAllTestFromSection(c *gin.Context) {
	disciplineId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	userId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	tests, testsEn, err := h.Service.GetAllTestFromSection(userId, disciplineId)
	if err != nil {
		err = errors.New("ошибка получения тестов")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.TestResponse{Ru: tests, En: testsEn})
}

// GetAllLabFromSection @Summary get all labs from section
// @Security ApiKeyAuthStudent
// @Tags test and lab
// @Description get all labs from section
// @Id get-all-labs-from-section-student
// @Produce json
// @Param id body int true "get all labs from section id"
// @Success 200 {object} model.LabsResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/student/disciplines/laboratory/{id} [get]
func (h *StudentTestAndLabHandler) GetAllLabFromSection(c *gin.Context) {
	disciplineId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	labs, labsEn, err := h.Service.GetAllLab(userId, disciplineId)
	if err != nil {
		err = errors.New("ошибка получения лабораторных работ")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.LabsResponse{Ru: labs, En: labsEn})
}

// GetQuestionsForTest @Summary get question for test
// @Security ApiKeyAuthStudent
// @Tags test and lab
// @Description get question for test
// @Id get-question-for-test-student
// @Produce json
// @Param id body int true "test id"
// @Success 200 {object} model.QuestionsForTestResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/student/test/{id} [get]
func (h *StudentTestAndLabHandler) GetQuestionsForTest(c *gin.Context) {
	testId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	questionsRu, questionsEn, err := h.Service.GetQuestionsForTest(userId, testId)
	if err != nil {
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model.QuestionsForTestResponse{Ru: questionsRu, En: questionsEn})
}

// PassTest @Summary pass test
// @Security ApiKeyAuthStudent
// @Tags test and lab
// @Description pass test
// @Id pass-test-student
// @Produce json
// @Param id body int true "test id"
// @Param answers body model.QuestionAndAnswersParser true "answers"
// @Success 200 {object} model.TestMarkResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/student/test/{id} [post]
func (h *StudentTestAndLabHandler) PassTest(c *gin.Context) {
	var jsonInput map[string]interface{}

	if err := c.BindJSON(&jsonInput); err != nil {
		logrus.Info("ERROR PARSE JSON")
		return
	}
	data, err := json.Marshal(jsonInput)
	if err != nil {
		logrus.Info("ERROR MARSHAL JSON")
		return
	}

	userId, err := middleware.GetUserId(c)
	if err != nil {
		logrus.Info(fmt.Sprintf("ERROR DON't KNOWN USER TEST:%s USER:%d DATA:%s", c.Param("id"), userId, string(data)))
		return
	}
	testId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Info(fmt.Sprintf("ERROR DON't KNOWN TEST:%s USER:%d DATA:%s", c.Param("id"), userId, string(data)))
		err := fmt.Errorf("ошибка запроса: неверные данные теста: %v", err)
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var checkAnswers model.QuestionAndAnswersParser
	if err := json.Unmarshal(data, &checkAnswers); err != nil {
		logrus.Info(fmt.Sprintf("ERROR PARSE DATA FOR TEST:%s USER:%d DATA:%s", c.Param("id"), userId, string(data)))
		err := fmt.Errorf("ошибка запроса: неверные данные: %v", err)
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	mark, done, err := h.Service.CheckAnswers(userId, testId, checkAnswers.Answers)
	if err != nil {
		logrus.Info(fmt.Sprintf("ERROR CHECK TEST:%s USER:%d DATA:%s", c.Param("id"), userId, string(data)))
		err := fmt.Errorf("ошибка проверки теста: %v", err)
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model.TestMarkResponse{Mark: mark, QuestionPercentage: done})
}

// GetTestMark @Summary get test mark
// @Security ApiKeyAuthStudent
// @Tags test and lab
// @Description get test mark
// @Id get-test-mark-student
// @Produce json
// @Param id body int true "test id"
// @Success 200 {object} model.TestResult
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/student/test/mark/{id} [get]
func (h *StudentTestAndLabHandler) GetTestMark(c *gin.Context) {
	testId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	mark, err := h.Service.GetResultOfTest(userId, testId)
	if err != nil {
		err = errors.New("ошибка сервера: возможно отсутствуют данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, mark)
}

// GetLabMark @Summary get lab mark
// @Security ApiKeyAuthStudent
// @Tags test and lab
// @Description get lab mark
// @Id get-lab-mark-student
// @Produce json
// @Param id body int true "lab id"
// @Success 200 {object} model.LabResult
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/student/laboratory-work/mark/{id} [get]
func (h *StudentTestAndLabHandler) GetLabMark(c *gin.Context) {
	labId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	mark, err := h.Service.GetResultOfTest(userId, labId)
	if err != nil {
		err = errors.New("ошибка сервера: возможно отсутствуют данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, mark)
}

// GetReportForTest @Summary get report for test
// @Security ApiKeyAuthStudent
// @Tags test and lab
// @Description get report for test
// @Id get-report-for-test-student
// @Produce json
// @Param id body int true "test id"
// @Success 200 {file} file
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/student/test/report/{id} [get]
func (h *StudentTestAndLabHandler) GetReportForTest(c *gin.Context) {
	testId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := middleware.GetUserId(c)
	if err != nil {
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

// GetAllDoneTests @Summary get all test done tests
// @Security ApiKeyAuthStudent
// @Tags test and lab
// @Description get all done tests
// @Id get-all-done-test-student
// @Produce json
// @Param is_done query string false "is_done test"
// @Success 200 {object} model.TestResponseWithClosedDate
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/student/test [get]
func (h *StudentTestAndLabHandler) GetAllDoneTests(c *gin.Context) {
	isDone := c.Query("is_done")
	userId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	if isDone == "1" {
		tests, testsEn, err := h.Service.GetAllDoneTests(userId)
		if err != nil {
			err = errors.New("ошибка получения тестов")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, model.TestResponseWithClosedDate{Ru: tests, En: testsEn})
		return
	} else {
		tests, testsEn, err := h.Service.GetAllOpenedTests(userId)
		if err != nil {
			err = errors.New("ошибка получения тестов")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, model.TestResponseWithClosedDate{Ru: tests, En: testsEn})
		return
	}
}

// GetAllDoneLabs @Summary get all done labs
// @Security ApiKeyAuthStudent
// @Tags test and lab
// @Description get all done labs
// @Id get-all-done-labs-student
// @Produce json
// @Param is_done query string true "is_done lab"
// @Success 200 {object} model.LabsResponseWithClosedDate
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/student/laboratory-work/done [get]
func (h *StudentTestAndLabHandler) GetAllDoneLabs(c *gin.Context) {
	isDone := c.Query("is_done")
	userId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	if isDone == "1" {
		labs, labsEn, err := h.Service.GetAllDoneLabs(userId)
		if err != nil {
			err = errors.New("ошибка получения лабораторных работ")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, model.LabsResponseWithClosedDate{Ru: labs, En: labsEn})
		return
	} else {
		labs, labsEn, err := h.Service.GetAllOpenedLabs(userId)
		if err != nil {
			err = errors.New("ошибка получения лабораторных работ")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, model.LabsResponseWithClosedDate{Ru: labs, En: labsEn})
		return
	}
}

// OpenLab @Summary open lab
// @Security ApiKeyAuthStudent
// @Tags test and lab
// @Description open lab
// @Id open-lab-student
// @Produce json
// @Success 200 {object} model.LabsResponseWithClosedDate
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/student/laboratory-work/{id} [get]
func (h *StudentTestAndLabHandler) OpenLab(c *gin.Context) {
	labId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}

	link, err := h.Service.GetLinkForLab(userId, labId)
	if err != nil {
		err = errors.New("ошибка открытия лабораторной работы")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Writer.Header().Set(tokenForLab, link.Token)
	url := link.Link + "?user_id=" + strconv.Itoa(userId) + "&laboratory_id=" + strconv.Itoa(labId)
	c.Redirect(http.StatusMovedPermanently, url)
}
