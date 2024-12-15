package api_lecturer

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"backend/pkg/handler/error_response"
	"backend/pkg/model"
	"backend/pkg/service"
)

type LecturerMarksHandler struct {
	service service.LecturerMarks
}

func NewLecturerMarksHandler(service service.LecturerMarks) *LecturerMarksHandler {
	return &LecturerMarksHandler{service: service}
}

// ChangeTestMark @Summary change test mark
// @Security ApiKeyAuthLecturer
// @Tags mark
// @Description change test mark
// @Id change-test-mark
// @Accept json
// @Produce json
// @Param input body model.ChangeTestMarkInput true "change test mark for student input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/mark/test  [put]
func (h *LecturerMarksHandler) ChangeTestMark(c *gin.Context) {
	var mark model.ChangeTestMarkInput
	if err := c.BindJSON(&mark); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.ChangeTestMark(mark.UserId, mark.TestId, mark.Mark); err != nil {
		err = errors.New("ошибка изменения оценки за тест")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ChangeLaboratoryMark @Summary change laboratory mark
// @Security ApiKeyAuthLecturer
// @Tags mark
// @Description change laboratory mark
// @Id change-laboratory-mark
// @Accept json
// @Produce json
// @Param input body model.ChangeLaboratoryMarkInput true "change laboratory mark for student input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/mark/laboratory  [post]
func (h *LecturerMarksHandler) ChangeLaboratoryMark(c *gin.Context) {
	var mark model.ChangeLaboratoryMarkInput
	if err := c.BindJSON(&mark); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.ChangeLaboratoryMark(mark.UserId, mark.LaboratoryId, mark.Mark); err != nil {
		err = errors.New("ошибка изменения оценки за лабораторную работу")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetTestMarksFromGroup @Summary get test marks from group
// @Security ApiKeyAuthLecturer
// @Tags mark
// @Description get test marks from group
// @Id get-test-marks-from-group
// @Accept json
// @Produce json
// @Param group_id query string true "group_id"
// @Param test_id query string true "test_id"
// @Success 200 {object} model.GroupTestMarksResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/mark/test  [get]
func (h *LecturerMarksHandler) GetTestMarksFromGroup(c *gin.Context) {
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

	marks, err := h.service.GetTestMarksFromGroup(groupId, testId)
	if err != nil {
		err = errors.New("ошибка получения оценок студентов группы за тест")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.GroupTestMarksResponse{
		Marks: marks,
	})
}

// GetLaboratoryMarksFromGroup @Summary get laboratory marks from group
// @Security ApiKeyAuthLecturer
// @Tags mark
// @Description get laboratory marks from group
// @Id get-laboratory-marks-from-group
// @Accept json
// @Produce json
// @Param group_id query string true "group_id"
// @Param laboratory_id query string true "laboratory_id"
// @Success 200 {object} model.GroupLaboratoryMarksResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/mark/laboratory  [get]
func (h *LecturerMarksHandler) GetLaboratoryMarksFromGroup(c *gin.Context) {
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

	marks, err := h.service.GetLaboratoryMarksFromGroup(groupId, laboratoryId)
	if err != nil {
		err = errors.New("ошибка получения оценок студентов группы за лабораторную работу")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.GroupLaboratoryMarksResponse{
		Marks: marks,
	})
}

// GiveExamMark @Summary give exam mark
// @Security ApiKeyAuthLecturer
// @Tags mark
// @Description give exam mark
// @Id give-exam-mark-lecturer
// @Accept json
// @Produce json
// @Param input body model.ExamMarkInput true "give exam mark"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/mark/exam  [post]
func (h *LecturerMarksHandler) GiveExamMark(c *gin.Context) {
	var exam model.ExamMarkInput
	if err := c.BindJSON(&exam); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.service.GiveExamMark(exam.UserId, exam.DisciplineId, exam.Mark)
	if err != nil {
		err = errors.New("ошибка изменения оценки за экзамен/зачет")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetExamMark @Summary get exam mark
// @Security ApiKeyAuthLecturer
// @Tags mark
// @Description get exam mark
// @Id get-exam-mark-lecturer
// @Accept json
// @Produce json
// @Param group_id query string true "group_id"
// @Param discipline_id query string true "discipline_id"
// @Success 200 {object} model.ExamMarkResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/mark/exam  [get]
func (h *LecturerMarksHandler) GetExamMark(c *gin.Context) {
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
	marks, err := h.service.GetAllMarksForExam(groupId, disciplineId)
	if err != nil {
		err = errors.New("ошибка получения оценок за экзамен/зачет студентов группы")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.ExamMarkResponse{Marks: marks})
}
