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

type LecturerDisciplineHandler struct {
	Service service.LecturerDiscipline
}

func NewLecturerDisciplineHandler(service service.LecturerDiscipline) *LecturerDisciplineHandler {
	return &LecturerDisciplineHandler{Service: service}
}

// GetAllGroupForDiscipline @Summary get all groups for discipline
// @Security ApiKeyAuthLecturer
// @Tags discipline
// @Description get all groups for discipline
// @Id get-all-groups-for-discipline
// @Produce json
// @Param discipline_id query string true "discipline_id"
// @Param is_add query string false "is_add"
// @Success 200 {object} model.GroupResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline/group  [get]
func (h *LecturerDisciplineHandler) GetAllGroupForDiscipline(c *gin.Context) {
	disciplineId, err := strconv.Atoi(c.Query("discipline_id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	isAdd := c.Query("is_add")
	if isAdd == "1" {
		groups, err := h.Service.GetGroupsAvailableToAddForDiscipline(disciplineId)
		if err != nil {
			err = errors.New("ошибка получения групп для добавления")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, model.GroupResponse{Groups: groups})
		return
	} else {
		groups, err := h.Service.GetGroupForDiscipline(disciplineId)
		if err != nil {
			err = errors.New("ошибка получения групп дисциплины")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, model.GroupResponse{Groups: groups})
		return
	}
}

// AddDiscipline @Summary add discipline
// @Security ApiKeyAuthLecturer
// @Tags discipline
// @Description add discipline
// @Id add-discipline
// @Accept json
// @Produce json
// @Param input body model.AddNewDiscipline true "discipline input"
// @Success 200 {int} int
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline  [post]
func (h *LecturerDisciplineHandler) AddDiscipline(c *gin.Context) {
	var inputDiscipline model.AddNewDiscipline
	if err := c.BindJSON(&inputDiscipline); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Service.AddDiscipline(inputDiscipline)
	if err != nil {
		err = errors.New("ошибка добавления дисциплины")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"discipline_id": id,
	})
}

// GetAllDisciplines @Summary get all disciplines
// @Security ApiKeyAuthLecturer
// @Tags discipline
// @Description get all disciplines
// @Id get-all-disciplines-lecturer
// @Produce json
// @Success 200 {object} model.DisciplinesResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline  [get]
func (h *LecturerDisciplineHandler) GetAllDisciplines(c *gin.Context) {
	disciplines, disciplinesEn, err := h.Service.GetAllDisciplines()
	if err != nil {
		err = errors.New("ошибка получения дисциплин")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.DisciplinesResponse{Ru: disciplines, En: disciplinesEn})
}

// ChangeSeminarMarks @Summary change seminar visiting mark
// @Security ApiKeyAuthLecturer
// @Tags discipline
// @Description change seminar visiting mark
// @Id change-seminar-visiting-mark
// @Accept json
// @Produce json
// @Param input body model.InputSeminarMark true "change seminar visiting input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline/mark/seminar  [put]
func (h *LecturerDisciplineHandler) ChangeSeminarMarks(c *gin.Context) {
	var semMark model.InputSeminarMark
	if err := c.BindJSON(&semMark); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.ChangeSeminarMarks(semMark.DisciplineId, semMark.Mark)
	if err != nil {
		err = errors.New("ошибка изменения баллов за посещения семинаров")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ChangeLessonMarks @Summary change lesson visiting mark
// @Security ApiKeyAuthLecturer
// @Tags discipline
// @Description change lesson visiting mark
// @Id change-lesson-visiting-mark
// @Accept json
// @Produce json
// @Param input body model.InputLessonMark true "change lesson visiting input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline/mark/lesson  [put]
func (h *LecturerDisciplineHandler) ChangeLessonMarks(c *gin.Context) {
	var lessonMark model.InputLessonMark
	if err := c.BindJSON(&lessonMark); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.ChangeLessonMarks(lessonMark.DisciplineId, lessonMark.Mark)
	if err != nil {
		err = errors.New("ошибка изменения баллов за посещения лекций")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ChangeExamMark @Summary change exam mark
// @Security ApiKeyAuthLecturer
// @Tags discipline
// @Description change exam mark
// @Id change-exam-mark
// @Accept json
// @Produce json
// @Param input body model.InputLessonMark true "change exam mark input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline/mark/exam  [put]
func (h *LecturerDisciplineHandler) ChangeExamMark(c *gin.Context) {
	var examMark model.InputExamMark
	if err := c.BindJSON(&examMark); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.ChangeExamMark(examMark.DisciplineId, examMark.Mark)
	if err != nil {
		err = errors.New("ошибка изменения баллов за экзамен")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetAllInfoAboutDiscipline @Summary get all info about discipline
// @Security ApiKeyAuthLecturer
// @Tags discipline
// @Description get all info about discipline
// @Id get-all-info-about-discipline
// @Produce json
// @Param id body int true "get all info about discipline input"
// @Success 200 {object} model.DisciplineInfo
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline/{id}  [get]
func (h *LecturerDisciplineHandler) GetAllInfoAboutDiscipline(c *gin.Context) {
	disciplineId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	discipline, err := h.Service.GetAllInfoAboutDiscipline(disciplineId)
	if err != nil {
		err = errors.New("ошибка получения информации о дисциплине")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, discipline)
}

// DeleteDiscipline @Summary delete discipline
// @Security ApiKeyAuthLecturer
// @Tags discipline
// @Description delete discipline
// @Id delete-discipline
// @Produce json
// @Param id body int true "delete discipline input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline/{id}  [delete]
func (h *LecturerDisciplineHandler) DeleteDiscipline(c *gin.Context) {
	disciplineId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.Service.DeleteDiscipline(disciplineId)
	if err != nil {
		err = errors.New("ошибка удаления дисциплины")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// AddGroupToDiscipline @Summary add group to discipline
// @Security ApiKeyAuthLecturer
// @Tags discipline
// @Description add group to discipline
// @Id add-group-to-discipline
// @Accept json
// @Produce json
// @Param input body model.InputGroupToDiscipline true "add group to discipline input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline/group  [post]
func (h *LecturerDisciplineHandler) AddGroupToDiscipline(c *gin.Context) {
	var groupToDiscipline model.InputGroupToDiscipline
	if err := c.BindJSON(&groupToDiscipline); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.AddGroupToDiscipline(groupToDiscipline.GroupId, groupToDiscipline.DisciplineId)
	if err != nil {
		err = errors.New("ошибка добавления группы к дисциплине")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteGroupFromDiscipline @Summary delete group from discipline
// @Security ApiKeyAuthLecturer
// @Tags discipline
// @Description delete group from discipline
// @Id delete-group-from-discipline
// @Accept json
// @Produce json
// @Param discipline_id query string true "discipline id"
// @Param group_id query string true "group_id"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline/group  [delete]
func (h *LecturerDisciplineHandler) DeleteGroupFromDiscipline(c *gin.Context) {
	group := c.Query("group_id")
	discipline := c.Query("discipline_id")
	groupId, err := strconv.Atoi(group)
	if err != nil {
		err = errors.New("ошибка получения группы")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	disciplineId, err := strconv.Atoi(discipline)
	if err != nil {
		err = errors.New("ошибка получения дисциплины")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.Service.DeleteGroupFromDiscipline(groupId, disciplineId); err != nil {
		err = errors.New("ошибка удаления группы от дисциплины")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ChangeDiscipline @Summary change discipline
// @Security ApiKeyAuthLecturer
// @Tags discipline
// @Description change discipline
// @Id change-discipline
// @Accept json
// @Produce json
// @Param discipline_id query string true "discipline_id"
// @Param name query string true "name"
// @Param name_en query string true "name_en"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline  [put]
func (h *LecturerDisciplineHandler) ChangeDiscipline(c *gin.Context) {
	discipline := c.Query("discipline_id")
	name := c.Query("name")
	nameEn := c.Query("name_en")
	disciplineId, err := strconv.Atoi(discipline)
	if err != nil {
		err = errors.New("ошибка изменения названия дисциплины")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if name != "" {
		err := h.Service.ChangeDiscipline(disciplineId, name)
		if err != nil {
			err = errors.New("ошибка изменения названия дисциплины")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if nameEn != "" {
		err := h.Service.ChangeDisciplineEn(disciplineId, nameEn)
		if err != nil {
			err = errors.New("ошибка изменения названия дисциплины")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ArchiveGroupToDiscipline @Summary archive group to discipline
// @Security ApiKeyAuthLecturer
// @Tags discipline
// @Description archive group to discipline
// @Id archive-group-to-discipline
// @Accept json
// @Produce json
// @Param input body model.InputGroupToDiscipline true "archive group to discipline input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline/archive  [put]
func (h *LecturerDisciplineHandler) ArchiveGroupToDiscipline(c *gin.Context) {
	var groupToDiscipline model.InputGroupToDiscipline
	if err := c.BindJSON(&groupToDiscipline); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.ArchiveGroupToDiscipline(groupToDiscipline.GroupId, groupToDiscipline.DisciplineId)
	if err != nil {
		err = errors.New("ошибка добавления в архив дисциплины группы")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
