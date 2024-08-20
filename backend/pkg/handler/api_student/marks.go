package api_student

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

type StudentMarksHandler struct {
	service service.StudentMarks
}

func NewStudentMarksHandler(service service.StudentMarks) *StudentMarksHandler {
	return &StudentMarksHandler{service: service}
}

// GetAllTestMarks @Summary get test marks
// @Security ApiKeyAuthStudent
// @Tags marks
// @Description get test marks
// @Id get-test-marks
// @Produce json
// @Param input body int true "discipline id"
// @Success 200 {object} model.GetAllTestsMarksResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/student/marks/test/{id}  [get]
func (h *StudentMarksHandler) GetAllTestMarks(c *gin.Context) {
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
	testMarks, err := h.service.GetAllTestsMarks(userId, disciplineId)
	if err != nil {
		err = errors.New("ошибка полчения оценок за тест")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model.GetAllTestsMarksResponse{
		Tests: testMarks,
	})
}

// GetAllLaboratoryMarks @Summary get laboratory marks
// @Security ApiKeyAuthStudent
// @Tags marks
// @Description get laboratory marks
// @Id get-laboratory-marks
// @Produce json
// @Param input body int true "discipline id"
// @Success 200 {object} model.GetAllLaboratoryMarksResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/student/marks/laboratory-work/{id}  [get]
func (h *StudentMarksHandler) GetAllLaboratoryMarks(c *gin.Context) {
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
	laboratoryMarks, err := h.service.GetAllLaboratoryMarks(userId, disciplineId)
	if err != nil {
		err = errors.New("ошибка получения оценок за лабораторную работу")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model.GetAllLaboratoryMarksResponse{
		LaboratoryWorks: laboratoryMarks,
	})
}

// GetExamMark @Summary get exam mark
// @Security ApiKeyAuthStudent
// @Tags marks
// @Description get exam mark
// @Id get-exam-mark
// @Produce json
// @Param input body int true "discipline id"
// @Success 200 {object} model.MarkResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/student/marks/exam/{id}  [get]
func (h *StudentMarksHandler) GetExamMark(c *gin.Context) {
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
	mark, err := h.service.GetExamMark(userId, disciplineId)
	if err != nil {
		err = errors.New("ошибка полчения оценки за экзамен/зачет")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model.MarkResponse{Mark: mark})
}
