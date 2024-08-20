package api_student

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/pkg/handler/error_response"
	"backend/pkg/handler/middleware"
	"backend/pkg/model"
	"backend/pkg/service"
)

type StudentPersonalDataHandler struct {
	service service.StudentPersonalData
}

func NewStudentPersonalDataHandler(service service.StudentPersonalData) *StudentPersonalDataHandler {
	return &StudentPersonalDataHandler{service: service}
}

// GetPersonalData @Summary get personal data
// @Security ApiKeyAuthStudent
// @Tags personal-data
// @Description get personal data
// @Id get-personal-data-student
// @Accept json
// @Produce json
// @Success 200 {object} model.UserData
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/student/personal-data  [get]
func (h *StudentPersonalDataHandler) GetPersonalData(c *gin.Context) {
	userId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	user, err := h.service.GetPersonalData(userId)
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.UserData{
		Name:      user.Name,
		Surname:   user.Surname,
		Email:     user.Email,
		GroupName: user.GroupName,
	})
}

// UpdatePersonalData @Summary update name
// @Security ApiKeyAuthStudent
// @Tags personal-data
// @Description update name
// @Id update-name-student
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param surname query string false "surname"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/student/personal-data  [put]
func (h *StudentPersonalDataHandler) UpdatePersonalData(c *gin.Context) {
	name := c.Query("name")
	surname := c.Query("surname")
	userId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	if name != "" {
		if err := h.service.UpdateName(userId, name); err != nil {
			err = errors.New("ошибка изменения имени")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if surname != "" {
		if err := h.service.UpdateSurname(userId, surname); err != nil {
			err = errors.New("ошибка изменения имени")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{})
}
