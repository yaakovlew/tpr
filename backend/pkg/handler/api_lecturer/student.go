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

type LecturerStudentsHandler struct {
	Service service.LecturerStudents
}

func NewLecturerStudentsHandler(service service.LecturerStudents) *LecturerStudentsHandler {
	return &LecturerStudentsHandler{Service: service}
}

// ChangeGroupForStudent @Summary change group for student
// @Security ApiKeyAuthLecturer
// @Tags student
// @Description change group for student
// @Id change-group-for-student
// @Accept json
// @Produce json
// @Param input body model.StudentChangeGroup true "change group for student"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/student [put]
func (h *LecturerStudentsHandler) ChangeGroupForStudent(c *gin.Context) {
	var changeGroup model.StudentChangeGroup
	if err := c.BindJSON(&changeGroup); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.ChangeGroupForStudent(changeGroup.UserId, changeGroup.GroupId)
	if err != nil {
		err = errors.New("ошибка изменения группы студента")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetAllStudents @Summary get all students
// @Security ApiKeyAuthLecturer
// @Tags student
// @Description get all students
// @Id get-all-students
// @Produce json
// @Success 200 {object} model.StudentWithGroupResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/student [get]
func (h *LecturerStudentsHandler) GetAllStudents(c *gin.Context) {
	students, err := h.Service.GetAllStudents()
	if err != nil {
		err = errors.New("ошибка получения студентов")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.StudentWithGroupResponse{Students: students})
}

// DeleteUser @Summary delete user
// @Security ApiKeyAuthLecturer
// @Tags student
// @Description delete user
// @Id delete-user
// @Produce json
// @Param id body int true "delete user input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/student/{id}  [delete]
func (h *LecturerStudentsHandler) DeleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.Service.DeleteUser(userId); err != nil {
		err = errors.New("ошибка удаления пользователя")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
