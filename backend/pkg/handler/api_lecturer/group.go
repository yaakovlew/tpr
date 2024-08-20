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

type LecturerGroupHandler struct {
	service service.LecturerGroup
}

func NewLecturerGroupHandler(service service.LecturerGroup) *LecturerGroupHandler {
	return &LecturerGroupHandler{service: service}
}

type addedGroupName struct {
	GroupName string `json:"group_name" binding:"required"`
}

// AddGroup @Summary add group
// @Security ApiKeyAuthLecturer
// @Tags group
// @Description add group
// @Id add-group
// @Accept json
// @Produce json
// @Param input body addedGroupName true "add group input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/group  [post]
func (h *LecturerGroupHandler) AddGroup(c *gin.Context) {
	var groupAdded addedGroupName
	if err := c.BindJSON(&groupAdded); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.AddGroup(groupAdded.GroupName); err != nil {
		err = errors.New("ошибка добавления группы")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteGroup @Summary delete group
// @Security ApiKeyAuthLecturer
// @Tags group
// @Description delete group
// @Id delete-group
// @Produce json
// @Param id body int true "delete group input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/group/{id}  [delete]
func (h *LecturerGroupHandler) DeleteGroup(c *gin.Context) {
	groupId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := h.service.DeleteGroup(groupId); err != nil {
		err = errors.New("ошибка удаления группы")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

type studentsGroupResponse struct {
	Students []model.Student `json:"students"`
}

// GetAllStudentsFromGroup @Summary get all students from group
// @Security ApiKeyAuthLecturer
// @Tags group
// @Description get all students from group
// @Id get-all-students-from-group
// @Produce json
// @Param id body int true "get all students from group input"
// @Success 200 {object} studentsGroupResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/group/students/{id}  [get]
func (h *LecturerGroupHandler) GetAllStudentsFromGroup(c *gin.Context) {
	groupId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	students, err := h.service.GetAllStudentsFromGroup(groupId)
	if err != nil {
		err = errors.New("ошибка получения студентов группы")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, studentsGroupResponse{Students: students})
}

// GetAllGroups @Summary get all groups
// @Security ApiKeyAuthLecturer
// @Tags group
// @Description get all groups
// @Id get-all-groups
// @Produce json
// @Success 200 {object} model.GroupResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/group  [get]
func (h *LecturerGroupHandler) GetAllGroups(c *gin.Context) {
	groups, err := h.service.GetAllGroups()
	if err != nil {
		err = errors.New("ошибка получения групп")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.GroupResponse{Groups: groups})
}

// GetGroupsDisciplines @Summary get groups discipline
// @Security ApiKeyAuthLecturer
// @Tags group
// @Description get groups from discipline
// @Id get-groups-disciplines
// @Produce json
// @Param id body int true "get groups from discipline id"
// @Success 200 {object} model.DisciplinesResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/group/discipline/{id}  [get]
func (h *LecturerGroupHandler) GetGroupsDisciplines(c *gin.Context) {
	groupId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	disciplines, disciplinesEn, err := h.service.GetGroupsDisciplines(groupId)
	if err != nil {
		err = errors.New("ошибка получения дисциплин группы")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.DisciplinesResponse{Ru: disciplines, En: disciplinesEn})
}

// ChangeGroupName @Summary change group name
// @Security ApiKeyAuthLecturer
// @Tags group
// @Description change group name
// @Id change-group-name
// @Accept json
// @Produce json
// @Param input body model.GroupChangeNameInput true "change group name"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/group  [put]
func (h *LecturerGroupHandler) ChangeGroupName(c *gin.Context) {
	var groupName model.GroupChangeNameInput
	if err := c.BindJSON(&groupName); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.service.ChangeName(groupName.GroupId, groupName.Name)
	if err != nil {
		err = errors.New("ошибка изменения названия группы")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// AddGroupInArchive @Summary add group to archive
// @Security ApiKeyAuthLecturer
// @Tags group
// @Description add groups to archive
// @Id add-groups-archive
// @Produce json
// @Param group_id query string true "group_id"
// @Param is_archive query string true "is_archive"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/group/archive/{id}  [post]
func (h *LecturerGroupHandler) AddGroupInArchive(c *gin.Context) {
	groupId, err := strconv.Atoi(c.Query("group_id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	isArchive := c.Query("is_archive")
	if isArchive == "1" {
		if err := h.service.AddGroupInArchive(groupId); err != nil {
			err = errors.New("ошибка добавления группы")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{})
		return
	} else {
		if err := h.service.DeleteGroupFromArchive(groupId); err != nil {
			err = errors.New("ошибка востановления группы из архива")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{})
		return
	}
}
