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

type SeminarianDisciplineHandler struct {
	Service service.SeminarianDiscipline
}

func NewSeminarianDisciplineHandler(service service.SeminarianDiscipline) *SeminarianDisciplineHandler {
	return &SeminarianDisciplineHandler{Service: service}
}

// GetOwnDiscipline @Summary get own disciplines
// @Security ApiKeyAuthSeminarian
// @Tags discipline
// @Description get all disciplines
// @Id get-all-disciplines-seminarian
// @Produce json
// @Success 200 {object} model.DisciplinesResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/discipline  [get]
func (h *SeminarianDisciplineHandler) GetOwnDiscipline(c *gin.Context) {
	userId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	disciplines, disciplinesEn, err := h.Service.GetOwnDiscipline(userId)
	if err != nil {
		err = errors.New("ошибка полчения дисциплин")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.DisciplinesResponse{Ru: disciplines, En: disciplinesEn})
}

// GetAllInfoAboutDiscipline @Summary get all info about discipline
// @Security ApiKeyAuthSeminarian
// @Tags discipline
// @Description get all info about discipline
// @Id get-all-info-about-discipline-seminarian
// @Produce json
// @Param id body int true "get all info about discipline input"
// @Success 200 {object} model.DisciplineInfo
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/discipline/{id}  [get]
func (h *SeminarianDisciplineHandler) GetAllInfoAboutDiscipline(c *gin.Context) {
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

	discipline, disciplineEn, err := h.Service.GetAllInfoAboutDiscipline(userId, disciplineId)
	if err != nil {
		err = errors.New("ошибка получения информации о дисциплине")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"discipline_ru": discipline,
		"discipline_en": disciplineEn,
	})
}

// GetDisciplineSections @Summary get discipline sections
// @Security ApiKeyAuthSeminarian
// @Tags section
// @Description get discipline sections
// @Id get-discipline-sections-seminarian
// @Produce json
// @Param id body int true "discipline id"
// @Success 200 {object} model.SectionsResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/discipline/section/{id}  [get]
func (h *SeminarianDisciplineHandler) GetDisciplineSections(c *gin.Context) {
	disciplineId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	sections, sectionsEn, err := h.Service.GetDisciplineSections(seminarianId, disciplineId)
	if err != nil {
		err = errors.New("ошибка полчения разделов")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.SectionsResponse{Ru: sections, En: sectionsEn})
}
