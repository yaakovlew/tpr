package api_seminarian

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"backend/pkg/handler/error_response"
	"backend/pkg/handler/middleware"
	"backend/pkg/model"
	"backend/pkg/service"
)

type SeminarianMarksHandler struct {
	service service.SeminarianMark
}

func NewSeminarianMarksHandler(service service.SeminarianMark) *SeminarianMarksHandler {
	return &SeminarianMarksHandler{service: service}
}

// GetTestMarksFromGroup @Summary get test marks from group
// @Security ApiKeyAuthSeminarian
// @Tags mark
// @Description get test marks from group
// @Id get-test-marks-from-group-seminarian
// @Accept json
// @Produce json
// @Param group_id query string true "group_id"
// @Param test_id query string true "test_id"
// @Success 200 {object} model.GroupTestMarksResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/mark/test  [get]
func (h *SeminarianMarksHandler) GetTestMarksFromGroup(c *gin.Context) {
	group := c.Query("group_id")
	test := c.Query("test_id")
	groupId, err := strconv.Atoi(group)
	if err != nil {
		err = errors.New("ошибка получения группы")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	testId, err := strconv.Atoi(test)
	if err != nil {
		err = errors.New("ошибка получения теста")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	marks, err := h.service.GetTestMarksFromGroup(seminarianId, groupId, testId)
	if err != nil {
		err = errors.New("ошибка получения оценок за тест")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.GroupTestMarksResponse{
		Marks: marks,
	})
}

// GetLaboratoryMarksFromGroup @Summary get laboratory marks from group
// @Security ApiKeyAuthSeminarian
// @Tags mark
// @Description get laboratory marks from group
// @Id get-laboratory-marks-from-group-seminarian
// @Accept json
// @Produce json
// @Param group_id query string true "group_id"
// @Param laboratory_id query string true "laboratory_id"
// @Success 200 {object} model.GroupLaboratoryMarksResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/mark/laboratory  [get]
func (h *SeminarianMarksHandler) GetLaboratoryMarksFromGroup(c *gin.Context) {
	group := c.Query("group_id")
	laboratory := c.Query("laboratory_id")
	groupId, err := strconv.Atoi(group)
	if err != nil {
		err = errors.New("ошибка получения группы")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	laboratoryId, err := strconv.Atoi(laboratory)
	if err != nil {
		err = errors.New("ошибка получения лабораторной")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	marks, err := h.service.GetLaboratoryMarksFromGroup(seminarianId, groupId, laboratoryId)
	if err != nil {
		err = errors.New("ошибка получения оценок лабораторной работы")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.GroupLaboratoryMarksResponse{
		Marks: marks,
	})
}

// GiveExamMark @Summary give exam mark
// @Security ApiKeyAuthSeminarian
// @Tags mark
// @Description give exam mark
// @Id give-exam-mark-seminarian
// @Accept json
// @Produce json
// @Param input body model.ExamMarkInput true "give exam mark"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/mark/exam  [post]
func (h *SeminarianMarksHandler) GiveExamMark(c *gin.Context) {
	var exam model.ExamMarkInput
	if err := c.BindJSON(&exam); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	err = h.service.GiveExamMark(seminarianId, exam.UserId, exam.DisciplineId, exam.Mark)
	if err != nil {
		err = errors.New("ошибка изменения оценки за экзамен/зачет")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetExamMark @Summary get exam mark
// @Security ApiKeyAuthSeminarian
// @Tags mark
// @Description get exam mark
// @Id get-exam-mark-seminarian
// @Accept json
// @Produce json
// @Param group_id query string true "group_id"
// @Param discipline_id query string true "discipline_id"
// @Success 200 {object} model.ExamMarkResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/mark/exam  [get]
func (h *SeminarianMarksHandler) GetExamMark(c *gin.Context) {
	group := c.Query("group_id")
	discipline := c.Query("discipline_id")
	groupId, err := strconv.Atoi(group)
	if err != nil {
		err = errors.New("ошибка получения группы")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	disciplineId, err := strconv.Atoi(discipline)
	if err != nil {
		err = errors.New("ошибка получения дисциплины")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	marks, err := h.service.GetAllMarksForExam(seminarianId, groupId, disciplineId)
	if err != nil {
		err = errors.New("ошибка полчения оценки за экзамен/зачет")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.ExamMarkResponse{Marks: marks})
}
