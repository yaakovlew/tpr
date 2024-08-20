package api_seminarian

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

type SeminarianAttendanceHandler struct {
	Service service.SeminarianAttendance
}

func NewSeminarianAttendanceHandler(service service.SeminarianAttendance) *SeminarianAttendanceHandler {
	return &SeminarianAttendanceHandler{Service: service}
}

type LecturerAttendanceHandler struct {
	Service service.LecturerAttendance
}

func NewLecturerAttendanceHandler(service service.LecturerAttendance) *LecturerAttendanceHandler {
	return &LecturerAttendanceHandler{Service: service}
}

// GetAllLessons @Summary get lessons
// @Security ApiKeyAuthSeminarian
// @Tags attendance
// @Description get lessons
// @Id get-lessons-seminarian
// @Produce json
// @Param id body int true "discipline id"
// @Success 200 {object} model.LessonsResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/attendance/lesson/{id}  [get]
func (h *SeminarianAttendanceHandler) GetAllLessons(c *gin.Context) {
	disciplineId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	lessons, err := h.Service.GetAllLessons(seminarianId, disciplineId)
	if err != nil {
		err = errors.New("ошибка получения лекций")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.LessonsResponse{Lessons: lessons})
}

// GetAllSeminars @Summary get seminars
// @Security ApiKeyAuthSeminarian
// @Tags attendance
// @Description get seminars
// @Id get-seminars-seminarian
// @Accept json
// @Produce json
// @Param discipline_id query string true "discipline_id"
// @Param group_id query string true "group_id"
// @Success 200 {object} model.SeminarResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/attendance/seminar  [get]
func (h *SeminarianAttendanceHandler) GetAllSeminars(c *gin.Context) {
	discipline := c.Query("discipline_id")
	group := c.Query("group_id")
	disciplineId, err := strconv.Atoi(discipline)
	if err != nil {
		err = errors.New("ошибка получения дисциплины")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	groupId, err := strconv.Atoi(group)
	if err != nil {
		err = errors.New("ошибка получения группы")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	seminars, err := h.Service.GetAllSeminars(seminarianId, disciplineId, groupId)
	if err != nil {
		err = errors.New("ошибка получения семинаров")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.SeminarResponse{Seminars: seminars})
}

// ChangeSeminar @Summary change seminar
// @Security ApiKeyAuthSeminarian
// @Tags attendance
// @Description change seminar
// @Id change-seminar-seminarian
// @Accept json
// @Param seminar_id query string true "seminar_id"
// @Param name query string false "name"
// @Param date query string false "date"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/attendance/seminar  [put]
func (h *SeminarianAttendanceHandler) ChangeSeminar(c *gin.Context) {
	seminar := c.Query("seminar_id")
	name := c.Query("name")
	date := c.Query("date")
	seminarId, err := strconv.Atoi(seminar)
	if err != nil {
		err = errors.New("ошибка изменения семинара")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	if name != "" {
		err = h.Service.ChangeSeminar(seminarianId, seminarId, name)
		if err != nil {
			err = errors.New("ошибка изменения семинара")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if date != "" {
		dateInt, err := strconv.Atoi(date)
		if err != nil {
			err = errors.New("ошибка изменения семинара")
			error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		err = h.Service.ChangeSeminarDate(seminarianId, seminarId, dateInt)
		if err != nil {
			err = errors.New("ошибка изменения даты семианара")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteSeminar @Summary delete seminar
// @Security ApiKeyAuthSeminarian
// @Tags attendance
// @Description delete seminar
// @Id delete-seminar-seminarian
// @Produce json
// @Param id body int true "seminar id"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/attendance/seminar/{id}  [delete]
func (h *SeminarianAttendanceHandler) DeleteSeminar(c *gin.Context) {
	seminarId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	err = h.Service.DeleteSeminar(seminarianId, seminarId)
	if err != nil {
		err = errors.New("ошибка удаления семинара")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// AddSeminar @Summary add seminar
// @Security ApiKeyAuthSeminarian
// @Tags attendance
// @Description add seminar
// @Id add-seminar-seminarian
// @Accept json
// @Produce json
// @Param input body model.AddSeminarInput true "seminar input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/attendance/seminar  [post]
func (h *SeminarianAttendanceHandler) AddSeminar(c *gin.Context) {
	var seminar model.AddSeminarInput
	if err := c.BindJSON(&seminar); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	err = h.Service.AddSeminar(seminarianId, seminar.DisciplineId, seminar.GroupId, seminar.Name, seminar.Date)
	if err != nil {
		err = errors.New("ошибка добавления семинара")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetLessonVisitingGroup @Summary get lesson visiting group
// @Security ApiKeyAuthSeminarian
// @Tags attendance
// @Description get lesson visiting from group
// @Id get-lesson-visiting-group-seminarian
// @Accept json
// @Produce json
// @Param lesson_id query string true "lesson_id"
// @Param group_id query string true "group_id"
// @Success 200 {object} model.LessonVisitingStudentResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/attendance/lesson/visiting  [get]
func (h *SeminarianAttendanceHandler) GetLessonVisitingGroup(c *gin.Context) {
	lesson := c.Query("lesson_id")
	group := c.Query("group_id")
	lessonId, err := strconv.Atoi(lesson)
	if err != nil {
		err = errors.New("ошибка получения лекции")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	groupId, err := strconv.Atoi(group)
	if err != nil {
		err = errors.New("ошибка получения группы")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	students, err := h.Service.GetLessonVisitingGroup(seminarianId, lessonId, groupId)
	if err != nil {
		err = errors.New("ошибка получения посещения лекций")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.LessonVisitingStudentResponse{Students: students})
}

// GetSeminarVisitingGroup @Summary get seminar visiting group
// @Security ApiKeyAuthSeminarian
// @Tags attendance
// @Description get seminar visiting from group
// @Id get-seminar-visiting-group-seminarian
// @Produce json
// @Param id body int true "seminar id"
// @Success 200 {object} model.SeminarVisitingStudentResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/attendance/seminar/visiting/{id}  [get]
func (h *SeminarianAttendanceHandler) GetSeminarVisitingGroup(c *gin.Context) {
	seminarId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	students, err := h.Service.GetSeminarVisitingGroup(seminarianId, seminarId)
	if err != nil {
		err = errors.New("ошибка получения посещений семинара")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.SeminarVisitingStudentResponse{Students: students})
}

// AddLessonVisiting @Summary add lesson visiting
// @Security ApiKeyAuthSeminarian
// @Tags attendance
// @Description add lesson visiting
// @Id add-lesson-visiting-seminarian
// @Accept json
// @Produce json
// @Param input body model.AddLessonVisitingInput true "add lesson visiting input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/attendance/lesson/visiting  [post]
func (h *SeminarianAttendanceHandler) AddLessonVisiting(c *gin.Context) {
	var lesson model.AddLessonVisitingInput
	if err := c.BindJSON(&lesson); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	err = h.Service.AddLessonVisiting(seminarianId, lesson.LessonId, lesson.UserId, *lesson.IsAbsent)
	if err != nil {
		err = errors.New("ошибка добавления посещения лекций")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// AddSeminarVisiting @Summary add seminar visiting
// @Security ApiKeyAuthSeminarian
// @Tags attendance
// @Description add seminar visiting
// @Id add-seminar-visiting-seminarian
// @Accept json
// @Produce json
// @Param input body model.AddSeminarVisitingInput true "add seminar visiting input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/attendance/seminar/visiting  [post]
func (h *SeminarianAttendanceHandler) AddSeminarVisiting(c *gin.Context) {
	var seminar model.AddSeminarVisitingInput
	if err := c.BindJSON(&seminar); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	err = h.Service.AddSeminarVisiting(seminarianId, seminar.SeminarId, seminar.UserId, *seminar.IsAbsent)
	if err != nil {
		err = errors.New("ошибка доавбления посещения сеимнара")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ChangeSeminarVisiting @Summary change seminar visiting
// @Security ApiKeyAuthSeminarian
// @Tags attendance
// @Description change seminar visiting
// @Id change-seminar-visiting-seminarian
// @Accept json
// @Produce json
// @Param input body model.AddSeminarVisitingInput true "change seminar visiting input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/attendance/seminar/visiting  [put]
func (h *SeminarianAttendanceHandler) ChangeSeminarVisiting(c *gin.Context) {
	var seminar model.AddSeminarVisitingInput
	err := c.BindJSON(&seminar)
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	err = h.Service.ChangeSeminarVisiting(seminarianId, seminar.SeminarId, seminar.UserId, *seminar.IsAbsent)
	if err != nil {
		err = errors.New("ошибка изменения посещаемости семинара")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ChangeLessonVisiting @Summary change lesson visiting
// @Security ApiKeyAuthSeminarian
// @Tags attendance
// @Description change lesson visiting
// @Id change-lesson-visiting-seminarian
// @Accept json
// @Produce json
// @Param input body model.AddLessonVisitingInput true "change lesson visiting input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/attendance/lesson/visiting  [put]
func (h *SeminarianAttendanceHandler) ChangeLessonVisiting(c *gin.Context) {
	var lesson model.AddLessonVisitingInput
	if err := c.BindJSON(&lesson); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	err = h.Service.ChangeLessonVisiting(seminarianId, lesson.LessonId, lesson.UserId, *lesson.IsAbsent)
	if err != nil {
		err = errors.New("ошибка изменения посещаемости лекции")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetLessonDate @Summary get lesson date
// @Security ApiKeyAuthSeminarian
// @Tags attendance
// @Description get lesson date
// @Id get-lesson-date-seminarian
// @Accept json
// @Produce json
// @Param input body model.LessonDateInput true "get lesson date input"
// @Success 200 {object} model.LessonDate
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/attendance/lesson/date  [get]
func (h *SeminarianAttendanceHandler) GetLessonDate(c *gin.Context) {
	var lesson model.LessonDateInput
	if err := c.BindJSON(&lesson); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	lessonDate, err := h.Service.GetLessonDate(seminarianId, lesson.LessonId, lesson.GroupId)
	if err != nil {
		err = errors.New("ошибка полчения даты лекции")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, lessonDate)
}

// GetTableLessons @Summary get table lessons
// @Security ApiKeyAuthSeminarian
// @Tags attendance
// @Description get table lessons
// @Id get-table-lessons-seminarian
// @Produce json
// @Param discipline_id query string true "discipline_id"
// @Param group_id query string false "group_id"
// @Success 200 {object} model.LessonsTableResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/attendance/lesson/table [get]
func (h *SeminarianAttendanceHandler) GetTableLessons(c *gin.Context) {
	disciplineId, err := strconv.Atoi(c.Query("discipline_id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	group := c.Query("group_id")
	if group != "" {
		groupId, err := strconv.Atoi(group)
		if err != nil {
			err = errors.New("ошибка запроса: неверные данные")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		lessonDate, err := h.Service.GetTableLessonsByGroup(seminarianId, disciplineId, groupId)
		if err != nil {
			err = errors.New("ошибка получения расписания лекций")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, model.LessonsTableResponse{
			Lessons: lessonDate,
		})
	} else {
		lessonDate, err := h.Service.GetTableLessons(seminarianId, disciplineId)
		if err != nil {
			err = errors.New("ошибка получения расписания лекций")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, model.LessonsTableResponse{
			Lessons: lessonDate,
		})
	}
}

// GetTableSeminars @Summary get seminars
// @Security ApiKeyAuthSeminarian
// @Tags attendance
// @Description get seminars
// @Id get-seminars-table-seminarian
// @Produce json
// @Param discipline_id query string true "discipline_id"
// @Param group_id query string true "group_id"
// @Success 200 {object} model.SeminarsTableResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/attendance/seminar/table [get]
func (h *SeminarianAttendanceHandler) GetTableSeminars(c *gin.Context) {
	discipline := c.Query("discipline_id")
	group := c.Query("group_id")
	disciplineId, err := strconv.Atoi(discipline)
	if err != nil {
		err = errors.New("ошибка получения дисциплины")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	groupId, err := strconv.Atoi(group)
	if err != nil {
		err = errors.New("ошибка получения группы")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	seminars, err := h.Service.GetTableSeminars(seminarianId, disciplineId, groupId)
	if err != nil {
		err = errors.New("ошибка полчения расписания семинаров")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.SeminarsTableResponse{
		Seminars: seminars,
	})
}
