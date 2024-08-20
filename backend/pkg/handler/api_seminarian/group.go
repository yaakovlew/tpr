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

type SeminarianGroupHandler struct {
	service service.SeminarianGroup
}

func NewSeminarianGroupHandler(service service.SeminarianGroup) *SeminarianGroupHandler {
	return &SeminarianGroupHandler{service: service}
}

type studentsGroupResponse struct {
	Students []model.Student `json:"students"`
}

// GetAllStudentsFromGroup @Summary get all students from group
// @Security ApiKeyAuthSeminarian
// @Tags group
// @Description get all students from group
// @Id get-all-students-from-group-seminarian
// @Produce json
// @Param id body int true "get all students from group input"
// @Success 200 {object} studentsGroupResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/discipline/group/students/{id}  [get]
func (h *SeminarianGroupHandler) GetAllStudentsFromGroup(c *gin.Context) {
	groupId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	students, err := h.service.GetAllStudentsFromGroup(seminarianId, groupId)
	if err != nil {
		err = errors.New("ошибка получения студентов группы")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, studentsGroupResponse{Students: students})
}

// GetOwnGroup @Summary get own group
// @Security ApiKeyAuthSeminarian
// @Tags group
// @Description get own groups
// @Id get-own-groups-seminarian
// @Produce json
// @Success 200 {object} model.GroupResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/discipline/group/{id}  [get]
func (h *SeminarianGroupHandler) GetOwnGroup(c *gin.Context) {
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
	groups, err := h.service.GetOwnGroup(userId, disciplineId)
	if err != nil {
		err = errors.New("ошибка получения групп")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.GroupResponse{Groups: groups})
}
