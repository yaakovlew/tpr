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

type LecturerAttendanceHandler struct {
	Service service.LecturerAttendance
}

func NewLecturerAttendanceHandler(service service.LecturerAttendance) *LecturerAttendanceHandler {
	return &LecturerAttendanceHandler{Service: service}
}

// GetAllLessons @Summary get lessons
// @Security ApiKeyAuthLecturer
// @Tags attendance
// @Description get lessons
// @Id get-lessons
// @Produce json
// @Param id body int true "discipline id"
// @Success 200 {object} model.LessonsResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/attendance/lesson/{id}  [get]
func (h *LecturerAttendanceHandler) GetAllLessons(c *gin.Context) {
	disciplineId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	lessons, err := h.Service.GetAllLessons(disciplineId)
	if err != nil {
		err = errors.New("ошибка получение лекций")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.LessonsResponse{Lessons: lessons})
}

// AddLesson @Summary add lessons
// @Security ApiKeyAuthLecturer
// @Tags attendance
// @Description add lessons
// @Id add-lessons
// @Accept json
// @Param input body model.AddLessonInput true "lesson"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/attendance/lesson  [post]
func (h *LecturerAttendanceHandler) AddLesson(c *gin.Context) {
	var lesson model.AddLessonInput
	if err := c.BindJSON(&lesson); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.AddLesson(lesson.DisciplineId, lesson.Name)
	if err != nil {
		err = errors.New("ошибка добавления лекции")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteLesson @Summary delete lessons
// @Security ApiKeyAuthLecturer
// @Tags attendance
// @Description delete lessons
// @Id delete-lessons
// @Produce json
// @Param id body int true "lesson id"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/attendance/lesson/{id}  [delete]
func (h *LecturerAttendanceHandler) DeleteLesson(c *gin.Context) {
	lessonId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.Service.DeleteLesson(lessonId)
	if err != nil {
		err = errors.New("ошибка удаления лекции")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ChangeLesson @Summary delete lessons
// @Security ApiKeyAuthLecturer
// @Tags attendance
// @Description change lessons
// @Id change-lessons
// @Accept json
// @Param input body model.ChangeLessonInput true "lesson name"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/attendance/lesson  [put]
func (h *LecturerAttendanceHandler) ChangeLesson(c *gin.Context) {
	var lesson model.ChangeLessonInput
	if err := c.BindJSON(&lesson); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.ChangeLesson(lesson.Id, lesson.Name)
	if err != nil {
		err = errors.New("ошибка изменения лекции")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetAllSeminars @Summary get seminars
// @Security ApiKeyAuthLecturer
// @Tags attendance
// @Description get seminars
// @Id get-seminars
// @Accept json
// @Produce json
// @Param discipline_id query string true "discipline_id"
// @Param group_id query string true "group_id"
// @Success 200 {object} model.SeminarResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/attendance/seminar  [get]
func (h *LecturerAttendanceHandler) GetAllSeminars(c *gin.Context) {
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
	seminars, err := h.Service.GetAllSeminars(disciplineId, groupId)
	if err != nil {
		err = errors.New("ошибка получения семинаров")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.SeminarResponse{Seminars: seminars})
}

// ChangeSeminar @Summary change seminar
// @Security ApiKeyAuthLecturer
// @Tags attendance
// @Description change seminar
// @Id change-seminar
// @Accept json
// @Param seminar_id query string true "seminar id"
// @Param name query string true "seminar name"
// @Param date query string true "seminar date"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/attendance/seminar  [put]
func (h *LecturerAttendanceHandler) ChangeSeminar(c *gin.Context) {
	seminar := c.Query("seminar_id")
	name := c.Query("name")
	date := c.Query("date")
	seminarId, err := strconv.Atoi(seminar)
	if err != nil {
		err = errors.New("ошибка изменения семинара")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if name != "" {
		err := h.Service.ChangeSeminar(seminarId, name)
		if err != nil {
			err = errors.New("ошибка изменения семинара")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if date != "" {
		dateInt, err := strconv.Atoi(date)
		if err != nil {
			err = errors.New("ошибка изменения даты семинара")
			error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		err = h.Service.ChangeSeminarDate(seminarId, dateInt)
		if err != nil {
			err = errors.New("ошибка изменения даты семинара")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteSeminar @Summary delete seminar
// @Security ApiKeyAuthLecturer
// @Tags attendance
// @Description delete seminar
// @Id delete-seminar
// @Produce json
// @Param id body int true "seminar id"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/attendance/seminar/{id}  [delete]
func (h *LecturerAttendanceHandler) DeleteSeminar(c *gin.Context) {
	seminarId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.Service.DeleteSeminar(seminarId)
	if err != nil {
		err = errors.New("ошибка удаления семинара")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// AddSeminar @Summary add seminar
// @Security ApiKeyAuthLecturer
// @Tags attendance
// @Description add seminar
// @Id add-seminar
// @Accept json
// @Produce json
// @Param input body model.AddSeminarInput true "seminar input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/attendance/seminar  [post]
func (h *LecturerAttendanceHandler) AddSeminar(c *gin.Context) {
	var seminar model.AddSeminarInput
	if err := c.BindJSON(&seminar); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.AddSeminar(seminar.DisciplineId, seminar.GroupId, seminar.Date, seminar.Name)
	if err != nil {
		err = errors.New("ошибка добавления семинара")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetLessonVisitingGroup @Summary get lesson visiting group
// @Security ApiKeyAuthLecturer
// @Tags attendance
// @Description get lesson visiting from group
// @Id get-lesson-visiting-group
// @Accept json
// @Produce json
// @Param lesson_id query string true "lesson id"
// @Param group_id query string true "group_id"
// @Success 200 {object} model.LessonVisitingStudentResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/attendance/lesson/visiting  [get]
func (h *LecturerAttendanceHandler) GetLessonVisitingGroup(c *gin.Context) {
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

	students, err := h.Service.GetLessonVisitingGroup(lessonId, groupId)
	if err != nil {
		err = errors.New("ошибка полчения посещений лекции")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.LessonVisitingStudentResponse{Students: students})
}

// GetSeminarVisitingGroup @Summary get seminar visiting group
// @Security ApiKeyAuthLecturer
// @Tags attendance
// @Description get seminar visiting from group
// @Id get-seminar-visiting-group
// @Produce json
// @Param id body int true "seminar id"
// @Success 200 {object} model.SeminarVisitingStudentResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/attendance/seminar/visiting/{id}  [get]
func (h *LecturerAttendanceHandler) GetSeminarVisitingGroup(c *gin.Context) {
	seminarId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	students, err := h.Service.GetSeminarVisitingGroup(seminarId)
	if err != nil {
		err = errors.New("ошибка получения посещений семинара")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.SeminarVisitingStudentResponse{Students: students})
}

// AddLessonVisiting @Summary add lesson visiting
// @Security ApiKeyAuthLecturer
// @Tags attendance
// @Description add lesson visiting
// @Id add-lesson-visiting
// @Accept json
// @Produce json
// @Param input body model.AddLessonVisitingInput true "add lesson visiting input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/attendance/lesson/visiting  [post]
func (h *LecturerAttendanceHandler) AddLessonVisiting(c *gin.Context) {
	var lesson model.AddLessonVisitingInput
	if err := c.BindJSON(&lesson); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.AddLessonVisiting(lesson.LessonId, lesson.UserId, *lesson.IsAbsent)
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// AddSeminarVisiting @Summary add seminar visiting
// @Security ApiKeyAuthLecturer
// @Tags attendance
// @Description add seminar visiting
// @Id add-seminar-visiting
// @Accept json
// @Produce json
// @Param input body model.AddSeminarVisitingInput true "add seminar visiting input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/attendance/seminar/visiting  [post]
func (h *LecturerAttendanceHandler) AddSeminarVisiting(c *gin.Context) {
	var seminar model.AddSeminarVisitingInput
	if err := c.BindJSON(&seminar); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.AddSeminarVisiting(seminar.SeminarId, seminar.UserId, *seminar.IsAbsent)
	if err != nil {
		err = errors.New("ошибка добавления посещений семинара")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ChangeSeminarVisiting @Summary change seminar visiting
// @Security ApiKeyAuthLecturer
// @Tags attendance
// @Description change seminar visiting
// @Id change-seminar-visiting
// @Accept json
// @Produce json
// @Param input body model.AddSeminarVisitingInput true "change seminar visiting input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/attendance/seminar/visiting  [put]
func (h *LecturerAttendanceHandler) ChangeSeminarVisiting(c *gin.Context) {
	var seminar model.AddSeminarVisitingInput
	err := c.BindJSON(&seminar)
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.Service.ChangeSeminarVisiting(seminar.SeminarId, seminar.UserId, *seminar.IsAbsent)
	if err != nil {
		err = errors.New("ошибка изменения посещений семинара")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ChangeLessonVisiting @Summary change lesson visiting
// @Security ApiKeyAuthLecturer
// @Tags attendance
// @Description change lesson visiting
// @Id change-lesson-visiting
// @Accept json
// @Produce json
// @Param input body model.AddLessonVisitingInput true "change lesson visiting input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/attendance/lesson/visiting  [put]
func (h *LecturerAttendanceHandler) ChangeLessonVisiting(c *gin.Context) {
	var lesson model.AddLessonVisitingInput
	if err := c.BindJSON(&lesson); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.ChangeLessonVisiting(lesson.LessonId, lesson.UserId, *lesson.IsAbsent)
	if err != nil {
		err = errors.New("ошибка изменения посещения лекции")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetLessonDate @Summary get lesson date
// @Security ApiKeyAuthLecturer
// @Tags attendance
// @Description get lesson date
// @Id get-lesson-date
// @Accept json
// @Produce json
// @Param lesson_id query string true "lesson id"
// @Param group_id query string true "group_id"
// @Success 200 {object} model.LessonDate
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/attendance/lesson/date  [get]
func (h *LecturerAttendanceHandler) GetLessonDate(c *gin.Context) {
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
	lessonDate, err := h.Service.GetLessonDate(lessonId, groupId)
	if err != nil {
		err = errors.New("ошибка получения даты лекции")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, lessonDate)
}

// ChangeLessonDate @Summary change lesson date
// @Security ApiKeyAuthLecturer
// @Tags attendance
// @Description change lesson date
// @Id change-lesson-date-lecturer
// @Accept json
// @Produce json
// @Param input body model.LessonDateToChange true "change lesson date input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/attendance/lesson/date  [put]
func (h *LecturerAttendanceHandler) ChangeLessonDate(c *gin.Context) {
	var lesson model.LessonDateToChange
	if err := c.BindJSON(&lesson); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if lesson.Date != 0 {
		if err := h.Service.ChangeLessonDate(lesson.LessonId, lesson.GroupId, lesson.Date); err != nil {
			err = errors.New("ошибка изменения даты лекции")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if lesson.Description != "" {
		if err := h.Service.ChangeLessonDateDescription(lesson.LessonId, lesson.GroupId, lesson.Description); err != nil {
			err = errors.New("ошибка изменения даты лекции")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{})
}

// AddLessonDate @Summary add lesson date
// @Security ApiKeyAuthLecturer
// @Tags attendance
// @Description add lesson date
// @Id add-lesson-date-lecturer
// @Accept json
// @Produce json
// @Param input body model.LessonDateToChange true "add lesson date input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/attendance/lesson/date  [post]
func (h *LecturerAttendanceHandler) AddLessonDate(c *gin.Context) {
	var lesson model.LessonDateToChange
	if err := c.BindJSON(&lesson); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.AddLessonDate(lesson.LessonId, lesson.GroupId, lesson.Date, lesson.Description)
	if err != nil {
		err = errors.New("ошибка добавления даты лекции")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteLessonDate @Summary delete lesson date
// @Security ApiKeyAuthLecturer
// @Tags attendance
// @Description delete lesson date
// @Id delete-lesson-date-lecturer
// @Accept json
// @Produce json
// @Param lesson_id query string true "lesson id"
// @Param group_id query string true "group_id"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/attendance/lesson/date  [delete]
func (h *LecturerAttendanceHandler) DeleteLessonDate(c *gin.Context) {
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
	err = h.Service.DeleteLessonDate(lessonId, groupId)
	if err != nil {
		err = errors.New("ошибка удаления даты лекции")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetTableLessons @Summary get table lessons
// @Security ApiKeyAuthLecturer
// @Tags attendance
// @Description get table lessons
// @Id get-table-lessons-lecturer
// @Produce json
// @Param discipline_id query string true "discipline_id"
// @Param group_id query string false "group_id"
// @Param lesson_id query string false "lesson_id"
// @Success 200 {object} model.LessonsTableResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/attendance/lesson [get]
func (h *LecturerAttendanceHandler) GetTableLessons(c *gin.Context) {
	disciplineId, err := strconv.Atoi(c.Query("discipline_id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	group := c.Query("group_id")
	lesson := c.Query("lesson_id")
	if lesson != "" && group != "" {
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
		lessonDate, err := h.Service.GetLessonDate(lessonId, groupId)
		if err != nil {
			err = errors.New("ошибка получения даты лекции")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, model.LessonsTableResponse{
			Lessons: []model.LessonDate{lessonDate},
		})
		return
	} else if group != "" {
		groupId, err := strconv.Atoi(group)
		if err != nil {
			err = errors.New("ошибка получения группы")
			error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		lessonDate, err := h.Service.GetTableLessons(disciplineId)
		if err != nil {
			err = errors.New("ошибка получения расписания")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		var lessonDateForGroup []model.LessonDate
		for _, lesson := range lessonDate {
			if lesson.GroupId == groupId {
				lessonDateForGroup = append(lessonDateForGroup, lesson)
			}
		}

		c.JSON(http.StatusOK, model.LessonsTableResponse{
			Lessons: lessonDateForGroup,
		})
		return
	} else {
		lessonDate, err := h.Service.GetTableLessons(disciplineId)
		if err != nil {
			err = errors.New("ошибка получения расписания")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, model.LessonsTableResponse{
			Lessons: lessonDate,
		})
		return
	}
}

// GetTableSeminars @Summary get table seminars
// @Security ApiKeyAuthLecturer
// @Tags attendance
// @Description get seminars table
// @Id get-seminars-table-lecturer
// @Produce json
// @Param id body int true "discipline id"
// @Success 200 {object} model.SeminarsTableResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/attendance/seminars/table/{id} [get]
func (h *LecturerAttendanceHandler) GetTableSeminars(c *gin.Context) {
	disciplineId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	seminars, err := h.Service.GetTableSeminars(disciplineId)
	if err != nil {
		err = errors.New("ошибка получения расписания")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.SeminarsTableResponse{
		Seminars: seminars,
	})
}
