package api_common

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"

	"backend/pkg/handler/error_response"
	"backend/pkg/handler/middleware"
	"backend/pkg/model"
	"backend/pkg/service"
)

type AuthHandler struct {
	Service service.Authorization
}

func NewAuthHeader(service service.Authorization) *AuthHandler {
	return &AuthHandler{Service: service}
}

// SignUp @Summary SignUp
// @Tags auth
// @Description create account
// @Id create-account
// @Accept json
// @Produce json
// @Param input body model.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /auth/sign-up [post]
func (h *AuthHandler) SignUp(c *gin.Context) {
	var input model.User
	if err := c.BindJSON(&input); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Service.CreateUser(input)
	if err != nil {
		err = errors.New("ошибка регистрации")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// SignIn @Summary SignIn
// @Tags auth
// @Description login
// @Id login
// @Accept json
// @Produce json
// @Param input body signInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /auth/sign-in [post]
func (h *AuthHandler) SignIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.Service.GenerateToken(input.Email, input.Password)
	if err != nil {
		err = errors.New("ошибка авторизации")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	post, err := h.Service.GetPost(input.Email)
	if err != nil {
		err = errors.New("ошибка авторизации")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"post":  post,
		"token": token,
	})
}

// ChangePassword @Summary change password
// @Security ApiKeyAuthCommon
// @Tags personal-data
// @Description change password
// @Id change-password
// @Accept json
// @Produce json
// @Param input body model.ChangePasswordInput true "change password input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/.../personal-data/change-password [put]
func (h *AuthHandler) ChangePassword(c *gin.Context) {
	var password model.ChangePasswordInput
	if err := c.BindJSON(&password); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	err = h.Service.ChangePassword(userId, password)
	if err != nil {
		err = errors.New("ошибка при смене пароля")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ChangePasswordForStudentAndSeminarianFromLecturer @Summary change password for student and seminarian from lecturer
// @Security ApiKeyAuthLecturer
// @Tags student
// @Description change password for student and seminarian from lecturer
// @Id change-password-for-student-and-seminarian-from-lecturer
// @Accept json
// @Produce json
// @Param input body model.LecturerChangePasswordForOtherInput true "change password for student and seminarian from lecturer input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/student/change-password [put]
func (h *AuthHandler) ChangePasswordForStudentAndSeminarianFromLecturer(c *gin.Context) {
	var password model.LecturerChangePasswordForOtherInput
	if err := c.BindJSON(&password); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.ChangePasswordForStudent(password.UserId, password.NewPassword)
	if err != nil {
		err = errors.New("ошибка при смене пароля")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// RestorePasswordLync @Summary restore password lync
// @Tags restore password
// @Description restore password lync
// @Id restore-password-lync
// @Accept json
// @Produce json
// @Param input body model.ForgotPasswordInput true "restore password lync input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /forget-password [post]
func (h *AuthHandler) RestorePasswordLync(c *gin.Context) {
	var email model.ForgotPasswordInput
	if err := c.BindJSON(&email); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	abc := gomail.NewMessage()
	abc.SetHeader("From", viper.GetString("mail"))
	abc.SetHeader("To", email.Email)
	abc.SetHeader("Subject", "Reset Password")
	link, err := h.Service.GeneratePasswordResetLink(email.Email)
	if err != nil {
		err = errors.New("ошибка при генерации ссылки")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	url := link.ToURL()

	abc.SetBody("text/plain", "Link to change password: "+url+"\nIf it's not your response, ignore this message")
	a := gomail.NewDialer("smtp.mail.ru", 587, viper.GetString("mail"), os.Getenv("PASSWORD_MAIL"))
	if err := a.DialAndSend(abc); err != nil {
		err = errors.New("ошибка: отправки сообщения")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// RestorePassword @Summary restore password
// @Security ApiKeyAuthSpecial
// @Tags restore password
// @Description restore password
// @Id restore-password
// @Accept json
// @Produce json
// @Param input body model.RestorePasswordInput true "restore password input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /restore-password [post]
func (h *AuthHandler) RestorePassword(c *gin.Context) {
	var password model.RestorePasswordInput
	if err := c.BindJSON(&password); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	err = h.Service.RestorePassword(userId, password.NewPassword)
	if err != nil {
		err = errors.New("ошибка при востановлении пароля")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
