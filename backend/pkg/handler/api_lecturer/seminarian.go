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

type LecturerSeminarianHandler struct {
	Service service.LecturerSeminarian
}

func NewLecturerSeminarianHandler(service service.LecturerSeminarian) *LecturerSeminarianHandler {
	return &LecturerSeminarianHandler{Service: service}
}

// GetSeminarianFromGroupAndDiscipline @Summary get seminarian from group and discipline
// @Security ApiKeyAuthLecturer
// @Tags seminarian
// @Description get seminarian from group and discipline
// @Id get-seminarian-from-group-and-discipline
// @Accept json
// @Produce json
// @Param group_id query string true "group_id"
// @Param discipline_id query string true "discipline_id"
// @Success 200 {object} model.SeminarianResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/group/seminarian [get]
func (h *LecturerSeminarianHandler) GetSeminarianFromGroupAndDiscipline(c *gin.Context) {
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

	seminarians, err := h.Service.GetSeminarianFromGroupsAndDiscipline(groupId, disciplineId)
	if err != nil {
		err = errors.New("ошибка получения семинаристов группы дисциплины")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.SeminarianResponse{Seminarians: seminarians})
}

// AddSeminarianToGroup @Summary add seminarian to group
// @Security ApiKeyAuthLecturer
// @Tags seminarian
// @Description add seminarian to group
// @Id add-seminarian-to-group
// @Accept json
// @Produce json
// @Param input body model.SeminarianAddToGroupInput true "add seminarian to group"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/seminarian  [post]
func (h *LecturerSeminarianHandler) AddSeminarianToGroup(c *gin.Context) {
	var seminarian model.SeminarianAddToGroupInput
	if err := c.BindJSON(&seminarian); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.AddSeminarian(seminarian.Id, seminarian.GroupId, seminarian.DisciplineId)
	if err != nil {
		err = errors.New("ошибка добавления семинариста")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetAllSeminarians @Summary get all seminarians
// @Security ApiKeyAuthLecturer
// @Tags seminarian
// @Description get all seminarians
// @Id get-all-seminarians
// @Produce json
// @Success 200 {object} model.SeminarianResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/seminarian  [get]
func (h *LecturerSeminarianHandler) GetAllSeminarians(c *gin.Context) {
	seminarians, err := h.Service.GetAllSeminarians()
	if err != nil {
		err = errors.New("ошибка получения всех семинаристов")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.SeminarianResponse{Seminarians: seminarians})
}

// DeleteSeminarianFromGroupAndDiscipline @Summary delete seminarian from group and discipline
// @Security ApiKeyAuthLecturer
// @Tags seminarian
// @Description delete seminarian from group and discipline
// @Id delete-seminarian-from-group-and-discipline
// @Accept json
// @Produce json
// @Param seminarian_id query string true "seminarian id"
// @Param group_id query string true "group_id"
// @Param discipline_id query string true "discipline id"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/group/seminarian [delete]
func (h *LecturerSeminarianHandler) DeleteSeminarianFromGroupAndDiscipline(c *gin.Context) {
	group := c.Query("group_id")
	discipline := c.Query("discipline_id")
	seminarian := c.Query("seminarian_id")
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
	seminarianId, err := strconv.Atoi(seminarian)
	if err != nil {
		err = errors.New("ошибка получения семинариста")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.Service.DeleteSeminarianFromGroupAndDiscipline(seminarianId, groupId, disciplineId); err != nil {
		err = errors.New("ошибка удаления семинариста от группы")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
