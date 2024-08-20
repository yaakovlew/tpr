package api_common

import (
	"backend/pkg/service"
	"github.com/gin-gonic/gin"
)

type Auth interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
	ChangePassword(c *gin.Context)
	RestorePasswordLync(c *gin.Context)
	ChangePasswordForStudentAndSeminarianFromLecturer(c *gin.Context)
	RestorePassword(c *gin.Context)
}

type Group interface {
	GetAllGroup(c *gin.Context)
}

type Lab interface {
	WebhookForLab(c *gin.Context)
	WebhookForGetUser(c *gin.Context)
}

type CommonController struct {
	Auth
	Group
	Lab
}

func NewCommonController(service *service.Service) *CommonController {
	return &CommonController{
		Lab:   NewCommonLabHandler(service.CommonLab),
		Auth:  NewAuthHeader(service.Authorization),
		Group: NewCommonGroupHandler(service.CommonGroup),
	}
}
