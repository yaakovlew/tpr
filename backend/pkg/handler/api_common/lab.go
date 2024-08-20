package api_common

import (
	"backend/pkg/handler/error_response"
	"backend/pkg/handler/middleware"
	"backend/pkg/model"
	"backend/pkg/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommonLabHandler struct {
	Service service.CommonLab
}

func NewCommonLabHandler(service service.CommonLab) *CommonLabHandler {
	return &CommonLabHandler{Service: service}
}

// WebhookForLab @Summary webhook for lab
// @Tags group
// @Description get all group
// @Id get-all-group
// @Produce json
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /lab [post]
func (h *CommonLabHandler) WebhookForLab(c *gin.Context) {
	var mark model.LabPercentage
	if err := c.BindJSON(&mark); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.Service.ChangeLabDateAndMark(mark.UserId, mark.LabId, mark.Percentage); err != nil {
		err = errors.New("ошибка изменения данных")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *CommonLabHandler) WebhookForGetUser(c *gin.Context) {
	userId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id": userId,
	})
}
