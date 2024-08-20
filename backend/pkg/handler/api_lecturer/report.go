package api_lecturer

import (
	"encoding/csv"
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"

	"backend/pkg/handler/error_response"
	"backend/pkg/service"
)

type LecturerReportHandler struct {
	Service service.LecturerReport
}

func NewLecturerReportHandler(service service.LecturerReport) *LecturerReportHandler {
	return &LecturerReportHandler{
		Service: service,
	}
}

// GetReport @Summary get report
// @Security ApiKeyAuthLecturer
// @Tags report
// @Description get report
// @Id get-report
// @Accept json
// @Produce json
// @Param group_id query string true "group_id"
// @Param discipline_id query string true "discipline_id"
// @Success 200 {file} report
// @Failure 400 {object} error_response.errorWeb
// @Failure 404 {object} error_response.errorWeb
// @Failure 500 {object} error_response.errorWeb
// @Failure default {object} error_response.errorWeb
// @Router /api/lecturer/group/report  [get]
func (h *LecturerReportHandler) GetReport(c *gin.Context) {
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

	f := excelize.NewFile()
	students, err := h.Service.GetAllStudents(groupId)
	if err != nil {
		err = errors.New("ошибка получения отчета")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	cellX := 1
	cellY := 2

	for i, student := range students {
		cell, err := excelize.CoordinatesToCellName(cellX, cellY)
		if err != nil {
			err = errors.New("ошибка получения отчета")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		f.SetCellValue("Sheet1", cell, i+1)
		cellX++
		cell, err = excelize.CoordinatesToCellName(cellX, cellY)
		if err != nil {
			err = errors.New("ошибка получения отчета")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		f.SetCellValue("Sheet1", cell, student.StudentSurname+" "+student.StudentName)
		cellY++
		cellX--
	}

	cellX = 1
	cellY = 1
	cell, err := excelize.CoordinatesToCellName(cellX, cellY)
	if err != nil {
		err = errors.New("ошибка получения отчета")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	f.SetCellValue("Sheet1", cell, "№ ПП")

	cellX = 2
	cellY = 1
	cell, err = excelize.CoordinatesToCellName(cellX, cellY)
	if err != nil {
		err = errors.New("ошибка получения отчета")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	f.SetCellValue("Sheet1", cell, "Фаилиия И.О")

	cellX = 3
	cellY = 1
	cell, err = excelize.CoordinatesToCellName(cellX, cellY)
	if err != nil {
		err = errors.New("ошибка получения отчета")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	f.SetCellValue("Sheet1", cell, "№ зач.книжки")

	cellX = 4
	cellY = 1
	sections, err := h.Service.GetThemesFromDiscipline(disciplineId)
	if err != nil {
		err = errors.New("ошибка получения отчета")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	for _, section := range sections {
		cellY = 1
		cell, err = excelize.CoordinatesToCellName(cellX, cellY)
		if err != nil {
			err = errors.New("ошибка получения отчета")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		f.SetCellValue("Sheet1", cell, section.Name)
		for _, student := range students {
			cellY++
			cell, _ = excelize.CoordinatesToCellName(cellX, cellY)
			mark, err := h.Service.GetMarkFromSection(student.StudentId, section.SectionId)
			if err != nil {
				err = errors.New("ошибка получения отчета")
				error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
			f.SetCellValue("Sheet1", cell, mark)
		}
		cellX++
	}

	cellY = 1
	cell, err = excelize.CoordinatesToCellName(cellX, cellY)
	if err != nil {
		err = errors.New("ошибка получения отчета")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	f.SetCellValue("Sheet1", cell, "Сумма балл. за разделы")
	for _, student := range students {
		cellY++
		cell, err = excelize.CoordinatesToCellName(cellX, cellY)
		if err != nil {
			err = errors.New("ошибка получения отчета")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		sumMark := h.Service.GetSummaryMarkFromSections(student.StudentId, disciplineId)
		f.SetCellValue("Sheet1", cell, sumMark)
	}

	cellY = 1

	cellX++
	cell, err = excelize.CoordinatesToCellName(cellX, cellY)
	if err != nil {
		err = errors.New("ошибка получения отчета")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	f.SetCellValue("Sheet1", cell, "Отметка об аттест.всех разд.(а, н/а)")
	for _, student := range students {
		cellY++
		cell, err = excelize.CoordinatesToCellName(cellX, cellY)
		if err != nil {
			err = errors.New("ошибка получения отчета")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		res := h.Service.GetSectionsResult(student.StudentId, disciplineId)
		f.SetCellValue("Sheet1", cell, res)
	}

	cellY = 1
	cellX++
	cell, err = excelize.CoordinatesToCellName(cellX, cellY)
	if err != nil {
		err = errors.New("ошибка получения отчета")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	f.SetCellValue("Sheet1", cell, "Итог.баллы")
	for _, student := range students {
		cellY++
		cell, err = excelize.CoordinatesToCellName(cellX, cellY)
		if err != nil {
			err = errors.New("ошибка получения отчета")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		finalMark := h.Service.GetFinalGrade(student.StudentId, disciplineId)
		f.SetCellValue("Sheet1", cell, finalMark)
	}

	cellY = 1
	cellX++
	cell, err = excelize.CoordinatesToCellName(cellX, cellY)
	if err != nil {
		err = errors.New("ошибка получения отчета")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	f.SetCellValue("Sheet1", cell, "Отметка (зачтено/не зачтено)")
	for _, student := range students {
		cellY++
		cell, err = excelize.CoordinatesToCellName(cellX, cellY)
		if err != nil {
			err = errors.New("ошибка получения отчета")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		res, _ := h.Service.GetResult(student.StudentId, disciplineId)
		f.SetCellValue("Sheet1", cell, res)
	}

	cellY = 1
	cellX++
	cell, err = excelize.CoordinatesToCellName(cellX, cellY)
	if err != nil {
		err = errors.New("ошибка получения отчета")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	f.SetCellValue("Sheet1", cell, "Код оценки")

	cellX++
	cell, err = excelize.CoordinatesToCellName(cellX, cellY)
	if err != nil {
		err = errors.New("ошибка получения отчета")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	f.SetCellValue("Sheet1", cell, "ECTS")
	for _, student := range students {
		cellY++
		cell, err = excelize.CoordinatesToCellName(cellX, cellY)
		if err != nil {
			err = errors.New("ошибка получения отчета")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		_, ects := h.Service.GetResult(student.StudentId, disciplineId)
		f.SetCellValue("Sheet1", cell, ects)
	}

	cellX++
	cellY = 1
	cell, err = excelize.CoordinatesToCellName(cellX, cellY)
	if err != nil {
		err = errors.New("ошибка получения отчета")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	f.SetCellValue("Sheet1", cell, "Подтв./подпись")

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		err = errors.New("ошибка получения отчета")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Создать новый файл CSV
	file, err := os.Create("src/output.csv")
	if err != nil {
		err = errors.New("ошибка получения отчета")
		error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()

	// Записать данные в файл CSV
	writer := csv.NewWriter(file)
	for _, row := range rows {
		err := writer.Write(row)
		if err != nil {
			err = errors.New("ошибка получения отчета")
			error_response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	writer.Flush()
	//_ = f.SaveAs("src/report.xlsx")
	c.FileAttachment("src/output.csv", "file")
	go func() {
		time.Sleep(5 * time.Second)
		_ = os.Remove("src/output.csv")
	}()
}
