package api_common

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/pkg/handler/error_response"
	"backend/pkg/model"
	"backend/pkg/service"
)

type CommonGroupHandler struct {
	Service service.CommonGroup
}

func NewCommonGroupHandler(service service.CommonGroup) *CommonGroupHandler {
	return &CommonGroupHandler{Service: service}
}

// GetAllGroup @Summary GetAllGroup
// @Tags group
// @Description get all group
// @Id get-all-group-common
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /group [get]
func (h *CommonGroupHandler) GetAllGroup(c *gin.Context) {
	groups, err := h.Service.GetAllGroups()
	if err != nil {
		fmt.Println(err)
		err = errors.New("ошибка получения групп")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	notArchiveGroups := []model.Group{}
	for _, group := range groups {
		if !group.IsArchive {
			notArchiveGroups = append(notArchiveGroups, group)
		}
	}
	c.JSON(http.StatusOK, model.GroupResponse{Groups: notArchiveGroups})
}
