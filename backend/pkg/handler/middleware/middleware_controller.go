package middleware

import (
	"backend/pkg/service"
	"github.com/gin-gonic/gin"
)

type MiddleWare interface {
	UserIdentityStudent(c *gin.Context)
	UserIdentitySeminarian(c *gin.Context)
	UserIdentityLecturer(c *gin.Context)
	UserIdentity(c *gin.Context)
	CheckHeaderForWebhook(c *gin.Context)
}

type MiddleWareController struct {
	MiddleWare
}

func NewMiddleWareController(service *service.Service) *MiddleWareController {
	return &MiddleWareController{
		MiddleWare: NewMiddlewareHandler(service.Authorization),
	}
}
