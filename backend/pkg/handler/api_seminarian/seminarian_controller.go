package api_seminarian

import (
	"backend/pkg/service"
	"github.com/gin-gonic/gin"
)

type SeminarianPersonalData interface {
	GetPersonalData(c *gin.Context)
	UpdatePersonalData(c *gin.Context)
}

type SeminarianGroup interface {
	GetAllStudentsFromGroup(c *gin.Context)
	GetOwnGroup(c *gin.Context)
}

type SeminarianDiscipline interface {
	GetOwnDiscipline(c *gin.Context)
	GetDisciplineSections(c *gin.Context)
	GetAllInfoAboutDiscipline(c *gin.Context)
}

type SeminarianMarks interface {
	GetTestMarksFromGroup(c *gin.Context)
	GetLaboratoryMarksFromGroup(c *gin.Context)
	GiveExamMark(c *gin.Context)
	GetExamMark(c *gin.Context)
}

type SeminarianAttendance interface {
	GetAllLessons(c *gin.Context)
	AddSeminar(c *gin.Context)
	GetAllSeminars(c *gin.Context)
	ChangeSeminar(c *gin.Context)
	DeleteSeminar(c *gin.Context)
	GetLessonVisitingGroup(c *gin.Context)
	GetSeminarVisitingGroup(c *gin.Context)
	AddLessonVisiting(c *gin.Context)
	AddSeminarVisiting(c *gin.Context)
	ChangeSeminarVisiting(c *gin.Context)
	ChangeLessonVisiting(c *gin.Context)
	GetLessonDate(c *gin.Context)
	GetTableLessons(c *gin.Context)
	GetTableSeminars(c *gin.Context)
}

type SeminarianTestAndLab interface {
	GetAllTestFromSection(c *gin.Context)
	GetAllLabFromSection(c *gin.Context)
	OpenTest(c *gin.Context)
	GetOpenedTestForStudent(c *gin.Context)
	CloseOpenedTestForStudent(c *gin.Context)
	GetReportForTest(c *gin.Context)
	GetUsersWithDoneTests(c *gin.Context)
	GetUsersWithDoneLab(c *gin.Context)
	OpenLab(c *gin.Context)
	CloseOpenedLabForStudent(c *gin.Context)
}

type SeminarianStudyGuide interface {
	GetAllLessonsForDiscipline(c *gin.Context)
	GetGuide(c *gin.Context)
	GetGuideFile(c *gin.Context)
}

type SeminarianReport interface {
	GetReport(c *gin.Context)
}

type SeminarianController struct {
	SeminarianPersonalData
	SeminarianGroup
	SeminarianDiscipline
	SeminarianMarks
	SeminarianAttendance
	SeminarianTestAndLab
	SeminarianStudyGuide
	SeminarianReport
}

func NewSeminarianController(service *service.Service) *SeminarianController {
	return &SeminarianController{
		SeminarianPersonalData: NewSeminarianPersonalDataHandler(service.SeminarianPersonalData),
		SeminarianGroup:        NewSeminarianGroupHandler(service.SeminarianGroup),
		SeminarianDiscipline:   NewSeminarianDisciplineHandler(service.SeminarianDiscipline),
		SeminarianMarks:        NewSeminarianMarksHandler(service.SeminarianMark),
		SeminarianAttendance:   NewSeminarianAttendanceHandler(service.SeminarianAttendance),
		SeminarianTestAndLab:   NewSeminarianTestAndLabHandler(service.SeminarianTestAndLab),
		SeminarianStudyGuide:   NewSeminarianStudyGuideHandler(service.SeminarianStudyGuide),
		SeminarianReport:       NewSeminarianReportHandler(service.SeminarianReport),
	}
}
