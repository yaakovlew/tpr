package api_student

import (
	"backend/pkg/service"
	"github.com/gin-gonic/gin"
)

type StudentMarks interface {
	GetAllLaboratoryMarks(c *gin.Context)
	GetAllTestMarks(c *gin.Context)
	GetExamMark(c *gin.Context)
}

type StudentPersonalData interface {
	UpdatePersonalData(c *gin.Context)
	GetPersonalData(c *gin.Context)
}

type StudentDiscipline interface {
	GetUserDiscipline(c *gin.Context)
	GetDisciplineSections(c *gin.Context)
}

type StudentAttendance interface {
	GetAllSeminarVisiting(c *gin.Context)
	GetAllLessonVisiting(c *gin.Context)
	GetLessons(c *gin.Context)
	GetSeminars(c *gin.Context)
}

type StudentStudyGuide interface {
	GetAllLessonsForDiscipline(c *gin.Context)
	GetGuide(c *gin.Context)
	GetGuideFile(c *gin.Context)
}

type StudentTestAndLab interface {
	GetAllTestFromSection(c *gin.Context)
	GetAllLabFromSection(c *gin.Context)
	GetQuestionsForTest(c *gin.Context)
	PassTest(c *gin.Context)
	GetTestMark(c *gin.Context)
	GetReportForTest(c *gin.Context)
	GetAllDoneTests(c *gin.Context)
	GetAllDoneLabs(c *gin.Context)
	GetLabMark(c *gin.Context)
	OpenLab(c *gin.Context)
}

type StudentController struct {
	StudentDiscipline
	StudentAttendance
	StudentMarks
	StudentPersonalData
	StudentStudyGuide
	StudentTestAndLab
}

func NewStudentController(service *service.Service) *StudentController {
	return &StudentController{
		StudentDiscipline:   NewStudentDisciplineHandler(service.StudentDiscipline),
		StudentAttendance:   NewStudentAttendanceHandler(service.StudentAttendance),
		StudentMarks:        NewStudentMarksHandler(service.StudentMarks),
		StudentPersonalData: NewStudentPersonalDataHandler(service.StudentPersonalData),
		StudentStudyGuide:   NewStudentStudyGuideHandler(service.StudentStudyGuide),
		StudentTestAndLab:   NewStudentTestAndLabHandler(service.StudentTestAndLab),
	}
}
