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

type SeminarianStudyGuideHandler struct {
	Service service.SeminarianStudyGuide
}

func NewSeminarianStudyGuideHandler(service service.SeminarianStudyGuide) *SeminarianStudyGuideHandler {
	return &SeminarianStudyGuideHandler{Service: service}
}

// GetAllLessonsForDiscipline @Summary get all digital lessons for discipline
// @Security ApiKeyAuthSeminarian
// @Tags digital lesson
// @Description get all lessons for discipline
// @Id get-all-lessons-for-discipline-seminarian
// @Produce json
// @Success 200 {object} model.DigitalDisciplinesInfoResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/material/{id}  [get]
func (h *SeminarianStudyGuideHandler) GetAllLessonsForDiscipline(c *gin.Context) {
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
	lessons, lessonsEn, err := h.Service.GetDigitalDiscipline(userId, disciplineId)
	if err != nil {
		err = errors.New("ошибка полчения материалов")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.DigitalDisciplinesInfoResponse{
		Ru: lessons,
		En: lessonsEn,
	})
}

// GetGuideFile @Summary get guide file
// @Security ApiKeyAuthSeminarian
// @Tags material
// @Description get guide file
// @Id get-guide-file-seminarian
// @Produce json
// @Param id body int true "get guide file id"
// @Success 200 {file} file
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/material/download/{id} [get]
func (h *SeminarianStudyGuideHandler) GetGuideFile(c *gin.Context) {
	fileId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	path, err := h.Service.GetFile(userId, fileId)
	if err != nil {
		err = errors.New("ошибка получения материалов")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.FileAttachment(path, "file")
}

// GetGuide @Summary get guide
// @Security ApiKeyAuthSeminarian
// @Tags material
// @Description get guide
// @Id get-guide-seminarian
// @Produce json
// @Param id body int true "get lesson id"
// @Success 200 {object} model.FilesResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/material/files/{id} [get]
func (h *SeminarianStudyGuideHandler) GetGuide(c *gin.Context) {
	lessonId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	files, err := h.Service.GetFilesIdFromDigital(userId, lessonId)
	if err != nil {
		err = errors.New("ошибка полчения материалов")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.FilesResponse{Files: files})
}
