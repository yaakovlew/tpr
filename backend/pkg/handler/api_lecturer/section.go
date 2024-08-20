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

type LecturerSectionHandler struct {
	Service service.LecturerSection
}

func NewLecturerSectionHandler(service service.LecturerSection) *LecturerSectionHandler {
	return &LecturerSectionHandler{Service: service}
}

// AddSection @Summary add section
// @Security ApiKeyAuthLecturer
// @Tags section
// @Description add section
// @Id add-section
// @Accept json
// @Produce json
// @Param input body model.AddSectionInput true "add section input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline/section  [post]
func (h *LecturerSectionHandler) AddSection(c *gin.Context) {
	var section model.AddSectionInput
	if err := c.BindJSON(&section); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.AddSection(section.Name, section.NameEn, section.DisciplineId)
	if err != nil {
		err = errors.New("ошибка добавления раздела")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetDisciplineSections @Summary get discipline sections
// @Security ApiKeyAuthLecturer
// @Tags section
// @Description get discipline sections
// @Id get-discipline-sections
// @Produce json
// @Param id body int true "discipline id"
// @Success 200 {object} model.SectionsResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline/section/{id}  [get]
func (h *LecturerSectionHandler) GetDisciplineSections(c *gin.Context) {
	disciplineId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	sections, sectionEn, err := h.Service.GetDisciplineSections(disciplineId)
	if err != nil {
		err = errors.New("ошибка получения разделов")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.SectionsResponse{Ru: sections, En: sectionEn})
}

// DeleteSection @Summary delete section
// @Security ApiKeyAuthLecturer
// @Tags section
// @Description delete section
// @Id delete-section
// @Produce json
// @Param id body int true "section id"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline/section/{id}  [delete]
func (h *LecturerSectionHandler) DeleteSection(c *gin.Context) {
	sectionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Service.DeleteSection(sectionId)
	if err != nil {
		err = errors.New("ошибка удаления раздела")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ChangeSectionName @Summary change section name
// @Security ApiKeyAuthLecturer
// @Tags section
// @Description change section name
// @Id change-section-name
// @Accept json
// @Produce json
// @Param input body model.ChangeSectionNameInput true "change section name input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline/section [put]
func (h *LecturerSectionHandler) ChangeSectionName(c *gin.Context) {
	var section model.ChangeSectionNameInput
	if err := c.BindJSON(&section); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.ChangeSectionName(section.Id, section.Name)
	if err != nil {
		err = errors.New("ошибка изменения названия раздела")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// ChangeSectionNameEn @Summary change english section name
// @Security ApiKeyAuthLecturer
// @Tags section
// @Description change english section name
// @Id change-english-section-name
// @Accept json
// @Produce json
// @Param input body model.ChangeSectionNameInput true "change section name input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline/section [put]
func (h *LecturerSectionHandler) ChangeSectionNameEn(c *gin.Context) {
	var section model.ChangeSectionNameInput
	if err := c.BindJSON(&section); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.ChangeSectionNameEn(section.Id, section.Name)
	if err != nil {
		err = errors.New("ошибка изменения названия раздела")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// AddTestToSection @Summary add test to section
// @Security ApiKeyAuthLecturer
// @Tags section
// @Description add test to section
// @Id add-test-to-section
// @Accept json
// @Produce json
// @Param input body model.TestToSectionInput true "add test to section input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline/section/test [post]
func (h *LecturerSectionHandler) AddTestToSection(c *gin.Context) {
	var sectionAndTest model.TestToSectionInput
	if err := c.BindJSON(&sectionAndTest); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.AddTestToSection(sectionAndTest.SectionId, sectionAndTest.TestId)
	if err != nil {
		err = errors.New("ошибка добавления теста к разделу")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteTestFromSection @Summary delete test from section
// @Security ApiKeyAuthLecturer
// @Tags section
// @Description delete test from section
// @Id delete-test-from-section
// @Accept json
// @Produce json
// @Param input body model.TestToSectionInput true "delete test from section input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline/section/test [delete]
func (h *LecturerSectionHandler) DeleteTestFromSection(c *gin.Context) {
	section := c.Query("section_id")
	test := c.Query("test_id")

	sectionId, err := strconv.Atoi(section)
	if err != nil {
		err = errors.New("ошибка получения раздела")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	testId, err := strconv.Atoi(test)
	if err != nil {
		err = errors.New("ошибка получения теста")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.Service.DeleteTestFromSection(sectionId, testId); err != nil {
		err = errors.New("ошибка удаления теста из раздела")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// AddLabToSection @Summary add lab to section
// @Security ApiKeyAuthLecturer
// @Tags section
// @Description add lab to section
// @Id add-lab-to-section
// @Accept json
// @Produce json
// @Param input body model.LabToSectionInput true "add lab to section input"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline/section/laboratory-work [post]
func (h *LecturerSectionHandler) AddLabToSection(c *gin.Context) {
	var sectionAndLab model.LabToSectionInput
	if err := c.BindJSON(&sectionAndLab); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Service.AddLabToSection(sectionAndLab.ExternalLabId, sectionAndLab.SectionId, sectionAndLab.DefaultMark)
	if err != nil {
		err = errors.New("ошибка добавления лабораторной работы к секции")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteLabFromSection @Summary delete lab from section
// @Security ApiKeyAuthLecturer
// @Tags section
// @Description delete lab from section
// @Id delete-lab-from-section
// @Accept json
// @Produce json
// @Param laboratory_id query string true "laboratory_id"
// @Param section_id query string true "section_id"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline/section/laboratory-work [delete]
func (h *LecturerSectionHandler) DeleteLabFromSection(c *gin.Context) {
	lab := c.Query("laboratory_id")
	section := c.Query("section_id")
	laboratoryId, err := strconv.Atoi(lab)
	if err != nil {
		err = errors.New("ошибка получения лабораторной работы")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	sectionId, err := strconv.Atoi(section)
	if err != nil {
		err = errors.New("ошибка получения раздела")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Service.DeleteLabFromSection(laboratoryId, sectionId)
	if err != nil {
		err = errors.New("ошибка удаления лабораторной работы из раздела")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetLabFromSection @Summary GetLabFromSection
// @Security ApiKeyAuthLecturer
// @Tags section
// @Description GetLabFromSection
// @Id getLabFromSection
// @Produce json
// @Param id body int true "section id"
// @Success 200 {object} model.StudentWithGroupWithClosedDateResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/discipline/section/laboratory-work/{id} [get]
func (h *LecturerSectionHandler) GetLabFromSection(c *gin.Context) {
	sectionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка получения раздела")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	labs, err := h.Service.GetLabFromSection(sectionId)
	if err != nil {
		err = errors.New("ошибка получения лабораторных работ")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"labs": labs,
	})
}
