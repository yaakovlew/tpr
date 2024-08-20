package api_lecturer

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"backend/pkg/handler/error_response"
	"backend/pkg/model"
	"backend/pkg/service"
)

type LecturerStudyGuideHandler struct {
	Service service.LecturerStudyGuide
}

func NewLecturerStudyGuideHandler(service service.LecturerStudyGuide) *LecturerStudyGuideHandler {
	return &LecturerStudyGuideHandler{Service: service}
}

// AddStudyGuideHeader @Summary add study guide header
// @Security ApiKeyAuthLecturer
// @Tags material
// @Description add study guide header
// @Id add-study-guide-header
// @Accept json
// @Produce json
// @Param input body model.AddStudyGuideHeaderInput true "add study guide header"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/material [post]
func (h *LecturerStudyGuideHandler) AddStudyGuideHeader(c *gin.Context) {
	var studyGuide model.AddStudyGuideHeaderInput
	if err := c.BindJSON(&studyGuide); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Service.AddStudyGuideHeader(studyGuide.Name, studyGuide.NameEn, studyGuide.Description, studyGuide.DescriptionEn)
	if err != nil {
		err = errors.New("ошибка добавления названия материала")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	dirName := strconv.Itoa(id)
	err = os.Mkdir(viper.GetString("destination")+dirName, 0755)
	if err != nil {
		err = errors.New("ошибка добавления заголовка материала")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetStudyGuideHeader @Summary get study guide header
// @Security ApiKeyAuthLecturer
// @Tags material
// @Description get study guide header
// @Id get-study-guide-header
// @Produce json
// @Success 200 {object} model.StudyGuideHeaderResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/material [get]
func (h *LecturerStudyGuideHandler) GetStudyGuideHeader(c *gin.Context) {
	guides, guidesEn, err := h.Service.GetStudyGuideHeader()
	if err != nil {
		err = errors.New("ошибка получения заголовков материала")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.StudyGuideHeaderResponse{Ru: guides, En: guidesEn})
}

// DeleteStudyGuideHeader @Summary delete study guide header
// @Security ApiKeyAuthLecturer
// @Tags material
// @Description delete study guide header
// @Id delete-study-guide-header
// @Accept json
// @Produce json
// @Param id body int true "delete study guide from id"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/material/{id} [delete]
func (h *LecturerStudyGuideHandler) DeleteStudyGuideHeader(c *gin.Context) {
	studyGuideId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	_, err = h.Service.DeleteStudyGuideHeader(studyGuideId)
	if err != nil {
		err = errors.New("ошибка удаления заголовка материала")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ChangeDigitalGuideHeader @Summary change name digital guide header
// @Security ApiKeyAuthLecturer
// @Tags material
// @Description change name digital guide header
// @Id change-name-digital-guide-header
// @Accept json
// @Produce json
// @Param digitalGuide query string true "digital_guide_id"
// @Param name query string true "name"
// @Param nameEn query string true "name_en"
// @Param description query string true "description"
// @Param descriptionEn query string true "description_en"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/material [put]
func (h *LecturerStudyGuideHandler) ChangeDigitalGuideHeader(c *gin.Context) {
	digitalGuide := c.Query("digital_guide_id")
	name := c.Query("name")
	nameEn := c.Query("name_en")
	description := c.Query("description")
	descriptionEn := c.Query("description_en")

	digitalGuideId, err := strconv.Atoi(digitalGuide)
	if err != nil {
		err = errors.New("ошибка изменения заголовка")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if name != "" {
		err := h.Service.ChangeNameDigitalGuideHeader(digitalGuideId, name)
		if err != nil {
			err = errors.New("ошибка изменения заголовка")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if nameEn != "" {
		err := h.Service.ChangeNameDigitalGuideHeaderEn(digitalGuideId, nameEn)
		if err != nil {
			err = errors.New("ошибка изменения заголовка")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if description != "" {
		err := h.Service.ChangeDescriptionDigitalGuideHeader(digitalGuideId, description)
		if err != nil {
			err = errors.New("ошибка изменения описания")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if descriptionEn != "" {
		err := h.Service.ChangeDescriptionDigitalGuideHeaderEn(digitalGuideId, descriptionEn)
		if err != nil {
			err = errors.New("ошибка изменения описания")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{})
}

// UploadFile @Summary upload file
// @Security ApiKeyAuthLecturer
// @Tags material
// @Description upload file
// @Id upload-file
// @Produce json
// @Param id body int true "upload file"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/material/upload/:id [post]
func (h *LecturerStudyGuideHandler) UploadFile(c *gin.Context) {
	lessonId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		err = errors.New("ошибка загрузки материала")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	path := fmt.Sprintf("%s/%d/%s", viper.GetString("destination"), lessonId, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		err = errors.New("ошибка загрузки материала")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.Service.AddFileToDigital(path, lessonId)
	if err != nil {
		err = errors.New("ошибка загрузки")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetGuideFile @Summary get guide file
// @Security ApiKeyAuthLecturer
// @Tags material
// @Description get guide file
// @Id get-guide-file
// @Produce json
// @Param id body int true "get guide file id"
// @Success 200 {file} file
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/material/download/{id} [get]
func (h *LecturerStudyGuideHandler) GetGuideFile(c *gin.Context) {
	fileId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	path, err := h.Service.GetFile(fileId)
	if err != nil {
		err = errors.New("ошибка получения материала")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.FileAttachment(path, "file")
}

// GetGuide @Summary get guide
// @Security ApiKeyAuthLecturer
// @Tags material
// @Description get guide
// @Id get-guide-lecturer
// @Produce json
// @Param id body int true "get lesson id"
// @Success 200 {object} model.FilesResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/material/{id} [get]
func (h *LecturerStudyGuideHandler) GetGuide(c *gin.Context) {
	lessonId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	files, err := h.Service.GetFilesIdFromDigital(lessonId)
	if err != nil {
		err = errors.New("ошибка получения материалов")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.FilesResponse{Files: files})
}

// DeleteGuide @Summary delete guide
// @Security ApiKeyAuthLecturer
// @Tags material
// @Description delete guide
// @Id delete-guide
// @Accept json
// @Produce json
// @Param input body model.DeleteGuideInput true "delete guide"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/material/file/{id} [delete]
func (h *LecturerStudyGuideHandler) DeleteGuide(c *gin.Context) {
	guideId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	path, err := h.Service.DeleteFileFromDigital(guideId)
	if err != nil {
		err = errors.New("ошибка удаления файла лекции")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = os.Remove(path)
	if err != nil {
		err = errors.New("ошибка удаления файла лекции")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetDigitalDiscipline @Summary get digital discipline
// @Security ApiKeyAuthLecturer
// @Tags material
// @Description get digital discipline
// @Id get-digital-discipline
// @Produce json
// @Param id body int true "get digital discipline"
// @Success 200 {object} model.DigitalDisciplinesResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/material/digital-lesson/{id} [get]
func (h *LecturerStudyGuideHandler) GetDigitalDiscipline(c *gin.Context) {
	disciplineId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	digitalDisciplines, err := h.Service.GetDigitalDiscipline(disciplineId)
	if err != nil {
		err = errors.New("ошибка получения материалов")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.DigitalDisciplinesResponse{DigitalGuides: digitalDisciplines})
}

// AddDigitalDiscipline @Summary add digital discipline
// @Security ApiKeyAuthLecturer
// @Tags material
// @Description add digital discipline
// @Id add-digital-discipline
// @Produce json
// @Param input body model.DigitalDiscipline true "add digital discipline"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/material/digital-lesson [post]
func (h *LecturerStudyGuideHandler) AddDigitalDiscipline(c *gin.Context) {
	var addDigitalDiscipline model.DigitalDiscipline
	if err := c.BindJSON(&addDigitalDiscipline); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.AddDigitalDiscipline(addDigitalDiscipline.DigitalMaterialId, addDigitalDiscipline.DisciplineId)
	if err != nil {
		err = errors.New("ошибка добавления материала")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteDigitalDiscipline @Summary delete digital discipline
// @Security ApiKeyAuthLecturer
// @Tags material
// @Description delete digital discipline
// @Id delete-digital-discipline
// @Produce json
// @Param discipline_id query string true "discipline_id"
// @Param digital_material_id query string true "digital_material_id"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/material/digital-lesson [delete]
func (h *LecturerStudyGuideHandler) DeleteDigitalDiscipline(c *gin.Context) {
	discipline := c.Query("discipline_id")
	digitalMaterial := c.Query("digital_material_id")
	disciplineId, err := strconv.Atoi(discipline)
	if err != nil {
		err = errors.New("ошибка получения дисциплины")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	digitalMaterialId, err := strconv.Atoi(digitalMaterial)
	if err != nil {
		err = errors.New("ошибка получения материала")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.Service.DeleteDigitalDiscipline(digitalMaterialId, disciplineId)
	if err != nil {
		err = errors.New("ошибка удаления материала")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
