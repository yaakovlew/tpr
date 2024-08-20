package api_seminarian

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/pkg/handler/error_response"
	"backend/pkg/handler/middleware"
	"backend/pkg/service"
)

type SeminarianPersonalDataHandler struct {
	service service.SeminarianPersonalData
}

func NewSeminarianPersonalDataHandler(service service.SeminarianPersonalData) *SeminarianPersonalDataHandler {
	return &SeminarianPersonalDataHandler{service: service}
}

type userData struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

// GetPersonalData @Summary get personal data
// @Security ApiKeyAuthSeminarian
// @Tags personal-data
// @Description get personal data
// @Id get-personal-data-seminarian
// @Accept json
// @Produce json
// @Success 200 {object} model.UserData
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/personal-data  [get]
func (h *SeminarianPersonalDataHandler) GetPersonalData(c *gin.Context) {
	userId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	user, err := h.service.GetPersonalData(userId)
	if err != nil {
		err = errors.New("ошибка полчения пресональных данных")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, userData{
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
	})
}

// UpdatePersonalData @Summary update name
// @Security ApiKeyAuthSeminarian
// @Tags personal-data
// @Description update personal data
// @Id update-personal-data-seminarian
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param surname query string false "surname"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/personal-data  [put]
func (h *SeminarianPersonalDataHandler) UpdatePersonalData(c *gin.Context) {
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
