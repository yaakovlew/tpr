package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"backend/pkg/handler/error_response"
	"backend/pkg/service"
)

const (
	authorizationHeader = "Authorization"
	userCTX             = "userId"
	labHeader           = "lab-mark-token"
)

type MiddlewareHandler struct {
	service service.Authorization
}

func NewMiddlewareHandler(service service.Authorization) *MiddlewareHandler {
	return &MiddlewareHandler{service: service}
}
func (h *MiddlewareHandler) UserIdentityStudent(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		error_response.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		error_response.NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userId, userPost, err := h.service.ParseToken(headerParts[1])
	if err != nil {
		err = errors.New("jwt is invalid type")
		error_response.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	if userPost != "student" {
		err := errors.New("don't have access")
		error_response.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCTX, userId)
}

func (h *MiddlewareHandler) UserIdentitySeminarian(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		error_response.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		error_response.NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userId, userPost, err := h.service.ParseToken(headerParts[1])
	if err != nil {
		err = errors.New("jwt is invalid type")
		error_response.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	if userPost != "seminarian" {
		err := errors.New("don't have access")
		error_response.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCTX, userId)
}

func (h *MiddlewareHandler) UserIdentityLecturer(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		error_response.NewErrorResponse(c, http.StatusUnauthorized, errors.New("empty auth header").Error())
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		error_response.NewErrorResponse(c, http.StatusUnauthorized, errors.New("invalid auth header").Error())
		return
	}
	userId, userPost, err := h.service.ParseToken(headerParts[1])
	if err != nil {
		err = errors.New("jwt is invalid type")
		error_response.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	if userPost != "lecturer" {
		err := errors.New("don't have access")
		error_response.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCTX, userId)
}

func (h *MiddlewareHandler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		error_response.NewErrorResponse(c, http.StatusUnauthorized, errors.New("empty auth header").Error())
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		error_response.NewErrorResponse(c, http.StatusUnauthorized, errors.New("invalid auth header").Error())
		return
	}

	userId, err := h.service.ParseTokenForRestorePassword(headerParts[1])
	if err != nil {
		err = errors.New("jwt is invalid type")
		error_response.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCTX, userId)
}

func GetUserId(c *gin.Context) (int, error) {
	userId, ok := c.Get(userCTX)
	if !ok {
		error_response.NewErrorResponse(c, http.StatusInternalServerError, errors.New("user not found").Error())
		return 0, errors.New("user id not found")
	}
	id, ok := userId.(int)
	if !ok {
		error_response.NewErrorResponse(c, http.StatusInternalServerError, errors.New("user id is invalid type").Error())
		return 0, errors.New("user id is of invalid type")
	}
	return id, nil
}

func (h *MiddlewareHandler) CheckHeaderForWebhook(c *gin.Context) {
	header := c.GetHeader(labHeader)
	if header == "" {
		error_response.NewErrorResponse(c, http.StatusUnauthorized, errors.New("empty header").Error())
		return
	}

	token := os.Getenv("LABS_HEADER")
	if header != token {
		error_response.NewErrorResponse(c, http.StatusUnauthorized, errors.New("not correct header").Error())
		return
	}
}
