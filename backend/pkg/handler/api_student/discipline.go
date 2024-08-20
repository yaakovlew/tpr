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

type StudentDisciplineHandler struct {
	service service.StudentDiscipline
}

func NewStudentDisciplineHandler(service service.StudentDiscipline) *StudentDisciplineHandler {
	return &StudentDisciplineHandler{service: service}
}

// GetUserDiscipline @Summary get user disciplines
// @Security ApiKeyAuthStudent
// @Tags discipline
// @Description get user disciplines
// @Id get-user-disciplines
// @Produce json
// @Success 200 {object} model.DisciplinesResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/student/disciplines [get]
func (h *StudentDisciplineHandler) GetUserDiscipline(c *gin.Context) {
	userId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	disciplines, disciplinesEn, err := h.service.GetAllUserDiscipline(userId)
	if err != nil {
		err = errors.New("ошибка полчения дисциплин")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.DisciplinesResponse{Ru: disciplines, En: disciplinesEn})
}

// GetDisciplineSections @Summary get discipline sections
// @Security ApiKeyAuthStudent
// @Tags section
// @Description get discipline sections
// @Id get-discipline-sections-student
// @Produce json
// @Param id body int true "discipline id"
// @Success 200 {object} model.SectionsResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/student/disciplines/section/{id}  [get]
func (h *StudentDisciplineHandler) GetDisciplineSections(c *gin.Context) {
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
	sections, sectionsEn, err := h.service.GetDisciplineSections(userId, disciplineId)
	if err != nil {
		err = errors.New("ошибка получения разделов")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.SectionsResponse{Ru: sections, En: sectionsEn})
}
