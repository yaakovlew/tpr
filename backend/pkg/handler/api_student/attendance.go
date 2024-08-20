package api_student

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"backend/pkg/handler/error_response"
	"backend/pkg/handler/middleware"
	"backend/pkg/model"
	"backend/pkg/service"
)

type StudentAttendanceHandler struct {
	service service.StudentAttendance
}

func NewStudentAttendanceHandler(service service.StudentAttendance) *StudentAttendanceHandler {
	return &StudentAttendanceHandler{service: service}
}

// GetAllSeminarVisiting @Summary get seminar visiting
// @Security ApiKeyAuthStudent
// @Tags attendance
// @Description get seminar visiting
// @Id get-seminar-visiting
// @Produce json
// @Param id body int true "discipline id"
// @Success 200 {object} model.SeminarsVisitingResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/student/attendance/seminars/{id} [get]
func (h *StudentAttendanceHandler) GetAllSeminarVisiting(c *gin.Context) {
	userId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	disciplineId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	seminars, err := h.service.GetAllSeminarVisiting(disciplineId, userId)
	if err != nil {
		err = errors.New("ошибка полчения посещаемости семинара")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.SeminarsVisitingResponse{Seminars: seminars})
}

// GetAllLessonVisiting @Summary get lesson visiting
// @Security ApiKeyAuthStudent
// @Tags attendance
// @Description get lesson visiting
// @Id get-lesson-visiting
// @Produce json
// @Param id body int true "discipline id"
// @Success 200 {object} model.LessonVisitingResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/student/attendance/lessons/{id} [get]
func (h *StudentAttendanceHandler) GetAllLessonVisiting(c *gin.Context) {
	userId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	disciplineId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	lessons, err := h.service.GetAllLessonVisiting(disciplineId, userId)
	if err != nil {
		err = errors.New("ошибка полчения посещаемости лекций")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.LessonVisitingResponse{Lessons: lessons})
}

// GetLessons @Summary get lessons
// @Security ApiKeyAuthStudent
// @Tags attendance
// @Description get lessons
// @Id get-lessons-student
// @Produce json
// @Param id body int true "discipline id"
// @Success 200 {object} model.LessonsTableResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/student/attendance/lessons/date/{id} [get]
func (h *StudentAttendanceHandler) GetLessons(c *gin.Context) {
	disciplineId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	userId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	lessonDate, err := h.service.GetAllLessons(userId, disciplineId)
	if err != nil {
		err = errors.New("ошибка получения лекций")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.LessonsTableResponse{
		Lessons: lessonDate,
	})
}

// GetSeminars @Summary get seminars
// @Security ApiKeyAuthStudent
// @Tags attendance
// @Description get seminars
// @Id get-seminars-student
// @Produce json
// @Param id body int true "discipline id"
// @Success 200 {object} model.SeminarsTableResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/student/attendance/seminars/date/{id} [get]
func (h *StudentAttendanceHandler) GetSeminars(c *gin.Context) {
	disciplineId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	userId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	seminars, err := h.service.GetAllSeminars(userId, disciplineId)
	if err != nil {
		err = errors.New("ошибка полчения семинаров")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.SeminarsTableResponse{Seminars: seminars})
}
