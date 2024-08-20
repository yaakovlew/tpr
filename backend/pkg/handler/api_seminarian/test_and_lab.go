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

type SeminarianTestAndLabHandler struct {
	Service service.SeminarianTestAndLab
}

func NewSeminarianTestAndLabHandler(service service.SeminarianTestAndLab) *SeminarianTestAndLabHandler {
	return &SeminarianTestAndLabHandler{Service: service}
}

// GetAllTestFromSection @Summary get all test from section
// @Security ApiKeyAuthSeminarian
// @Tags test and lab
// @Description get all test from section
// @Id get-all-test-from-section-seminarian
// @Produce json
// @Param id body int true "get all test from section id"
// @Success 200 {object} model.TestResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/discipline/test/{id} [get]
func (h *SeminarianTestAndLabHandler) GetAllTestFromSection(c *gin.Context) {
	sectionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	tests, testsEn, err := h.Service.GetAllTestFromSection(seminarianId, sectionId)
	if err != nil {
		err = errors.New("ошибка полчения тестов")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.TestResponse{Ru: tests, En: testsEn})
}

// GetAllLabFromSection @Summary get all labs from section
// @Security ApiKeyAuthSeminarian
// @Tags test and lab
// @Description get all labs from section
// @Id get-all-labs-from-section-seminarian
// @Produce json
// @Param id body int true "get all labs from section id"
// @Success 200 {object} model.LabsResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/discipline/laboratory-work/{id} [get]
func (h *SeminarianTestAndLabHandler) GetAllLabFromSection(c *gin.Context) {
	sectionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	labs, labsEn, err := h.Service.GetAllLab(seminarianId, sectionId)
	if err != nil {
		err = errors.New("ошибка полчения лабораторных работ")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.LabsResponse{Ru: labs, En: labsEn})
}

// OpenTest @Summary open test for user
// @Security ApiKeyAuthSeminarian
// @Tags test and lab
// @Description open test for user
// @Id open-test-for-user-seminarian
// @Accept json
// @Produce json
// @Param input body model.OpenTestInput true "open test for user"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/test/activity/open [post]
func (h *SeminarianTestAndLabHandler) OpenTest(c *gin.Context) {
	var test model.OpenTestInput
	if err := c.BindJSON(&test); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	err = h.Service.OpenTestForStudent(seminarianId, test.UserId, test.TestId, test.Date)
	if err != nil {
		err = errors.New("ошибка открытия теста")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetOpenedTestForStudent @Summary get opened test for student
// @Security ApiKeyAuthSeminarian
// @Tags test and lab
// @Description get opened test for user
// @Id get-opened-test-for-user-seminarian
// @Accept json
// @Produce json
// @Param user_id query string true "user_id"
// @Param test_id query string true "test_id"
// @Success 200 {object} model.OpenedTest
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/test/activity/get [get]
func (h *SeminarianTestAndLabHandler) GetOpenedTestForStudent(c *gin.Context) {
	user := c.Query("user_id")
	test := c.Query("test_id")
	userId, err := strconv.Atoi(user)
	if err != nil {
		err = errors.New("ошибка получения пользователя")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	testId, err := strconv.Atoi(test)
	if err != nil {
		err = errors.New("ошибка получения теста")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	testTake, err := h.Service.GetOpenedTestForStudent(seminarianId, userId, testId)
	if err != nil {
		err = errors.New("ошибка получения теста")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, testTake)
}

// CloseOpenedTestForStudent @Summary close opened test for student
// @Security ApiKeyAuthSeminarian
// @Tags test and lab
// @Description close opened test for student
// @Id close-opened-test-for-student-seminarian
// @Accept json
// @Produce json
// @Param input body model.OpenTestInput true "open test for user"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/test/activity/close [post]
func (h *SeminarianTestAndLabHandler) CloseOpenedTestForStudent(c *gin.Context) {
	var test model.CloseTestInput
	if err := c.BindJSON(&test); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	err = h.Service.CloseOpenedTestForStudent(seminarianId, test.UserId, test.TestId)
	if err != nil {
		err = errors.New("ошибка закрытия теста")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetReportForTest @Summary get report for test
// @Security ApiKeyAuthSeminarian
// @Tags test and lab
// @Description get report for test
// @Id get-report-for-test-seminarian
// @Produce json
// @Param user_id query string true "user_id"
// @Param test_id query string true "test_id"
// @Success 200 {file} file
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/test/activity/report [get]
func (h *SeminarianTestAndLabHandler) GetReportForTest(c *gin.Context) {
	user := c.Query("user_id")
	test := c.Query("test_id")
	userId, err := strconv.Atoi(user)
	if err != nil {
		err = errors.New("ошибка получения пользователя")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	testId, err := strconv.Atoi(test)
	if err != nil {
		err = errors.New("ошибка получения теста")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}

	path, err := h.Service.GetPathForReportTest(seminarianId, userId, testId)
	if err != nil {
		err = errors.New("ошибка сервера: возможно отсутствуют данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.FileAttachment(path, "file")
}

// GetUsersWithDoneTests @Summary get all students with done test
// @Security ApiKeyAuthSeminarian
// @Tags test and lab
// @Description get all students with done test
// @Id get-all-students-with-done-test-seminarian
// @Produce json
// @Param test_id query string true "test_id"
// @Param is_done query string false "is_done"
// @Success 200 {object} model.StudentWithGroupWithClosedDateResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/test/students/done [get]
func (h *SeminarianTestAndLabHandler) GetUsersWithDoneTests(c *gin.Context) {
	testId, err := strconv.Atoi(c.Query("test_id"))
	if err != nil {
		err = errors.New("ошибка получения теста")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	isDone := c.Query("is_done")
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	if isDone == "1" {
		users, err := h.Service.GetUsersWithDoneTest(seminarianId, testId)
		if err != nil {
			err = errors.New("ошибка получения студентов")
			error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, model.StudentWithGroupWithClosedDateResponse{Students: users})
		return
	} else {
		users, err := h.Service.GetUsersWithOpenedTest(seminarianId, testId)
		if err != nil {
			err = errors.New("ошибка получения студентов")
			error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, model.StudentWithGroupWithClosedDateResponse{Students: users})
		return
	}
}

// GetUsersWithDoneLab @Summary get all students with done lab
// @Security ApiKeyAuthSeminarian
// @Tags test and lab
// @Description get all students with done lab
// @Id get-all-students-with-done-lab-seminarian
// @Produce json
// @Param laboratory_id query string true "laboratory_id"
// @Param group_id query string true "group_id"
// @Param is_done query string false "is_done"
// @Success 200 {object} model.StudentWithGroupWithClosedDateResponse
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/laboratory-work/students/done/{id} [get]
func (h *SeminarianTestAndLabHandler) GetUsersWithDoneLab(c *gin.Context) {
	labId, err := strconv.Atoi(c.Query("laboratory_id"))
	if err != nil {
		err = errors.New("ошибка получения лабораторной работы")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	groupId, err := strconv.Atoi(c.Query("group_id"))
	if err != nil {
		err = errors.New("ошибка получения группы")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	isDone := c.Query("is_done")
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	if isDone == "1" {
		users, err := h.Service.GetUsersWithDoneLab(seminarianId, labId, groupId)
		if err != nil {
			err = errors.New("ошибка получения студентов")
			error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, model.StudentWithGroupWithClosedDateResponse{Students: users})
		return
	} else {
		users, err := h.Service.GetUsersWithOpenedLab(seminarianId, labId, groupId)
		if err != nil {
			err = errors.New("ошибка получения студентов")
			error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, model.StudentWithGroupWithClosedDateResponse{Students: users})
		return
	}
}

// OpenLab @Summary open lab for user
// @Security ApiKeyAuthSeminarian
// @Tags test and lab
// @Description open lab for user
// @Id open-lab-for-user-seminarian
// @Accept json
// @Produce json
// @Param input body model.OpenLabInput true "open lab for user"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/laboratory-work/activity [post]
func (h *SeminarianTestAndLabHandler) OpenLab(c *gin.Context) {
	var lab model.OpenLabInput
	if err := c.BindJSON(&lab); err != nil {
		err = errors.New("ошибка запроса: неверные данные")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	err = h.Service.OpenLabForStudent(seminarianId, lab.UserId, lab.LabId, lab.Date)
	if err != nil {
		err = errors.New("ошибка открытия лабораторной работы")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// CloseOpenedLabForStudent @Summary close opened lab for student
// @Security ApiKeyAuthSeminarian
// @Tags test and lab
// @Description close opened lab for student
// @Id close-opened-lab-for-student-seminarian
// @Accept json
// @Produce json
// @Param theme_id query string true "theme_id"
// @Param test_id query string true "test_id"
// @Success 200
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/seminarian/laboratory-work/activity [delete]
func (h *SeminarianTestAndLabHandler) CloseOpenedLabForStudent(c *gin.Context) {
	user := c.Query("user_id")
	lab := c.Query("laboratory_id")
	labId, err := strconv.Atoi(lab)
	if err != nil {
		err = errors.New("ошибка получения лабораторной работы")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := strconv.Atoi(user)
	if err != nil {
		err = errors.New("ошибка получения пользователя")
		error_response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	seminarianId, err := middleware.GetUserId(c)
	if err != nil {
		return
	}
	err = h.Service.CloseOpenedLabForStudent(seminarianId, userId, labId)
	if err != nil {
		err = errors.New("ошибка закрытия лабораторной работы")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
