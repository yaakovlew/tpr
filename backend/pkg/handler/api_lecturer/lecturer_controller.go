package api_lecturer

import (
	"backend/pkg/service"
	"github.com/gin-gonic/gin"
)

type LecturerPersonalData interface {
	UpdatePersonalData(c *gin.Context)
	GetPersonalData(c *gin.Context)
}

type LecturerGroup interface {
	AddGroup(c *gin.Context)
	DeleteGroup(c *gin.Context)
	GetAllStudentsFromGroup(c *gin.Context)
	GetAllGroups(c *gin.Context)
	GetGroupsDisciplines(c *gin.Context)
	ChangeGroupName(c *gin.Context)
	AddGroupInArchive(c *gin.Context)
}

type LecturerDiscipline interface {
	GetAllGroupForDiscipline(c *gin.Context)
	AddDiscipline(c *gin.Context)
	GetAllDisciplines(c *gin.Context)
	ChangeLessonMarks(c *gin.Context)
	ChangeSeminarMarks(c *gin.Context)
	GetAllInfoAboutDiscipline(c *gin.Context)
	DeleteDiscipline(c *gin.Context)
	AddGroupToDiscipline(c *gin.Context)
	DeleteGroupFromDiscipline(c *gin.Context)
	ChangeDiscipline(c *gin.Context)
	ChangeExamMark(c *gin.Context)
	ArchiveGroupToDiscipline(c *gin.Context)
}

type LecturerSeminarian interface {
	GetSeminarianFromGroupAndDiscipline(c *gin.Context)
	AddSeminarianToGroup(c *gin.Context)
	GetAllSeminarians(c *gin.Context)
	DeleteSeminarianFromGroupAndDiscipline(c *gin.Context)
}

type LecturerStudents interface {
	ChangeGroupForStudent(c *gin.Context)
	GetAllStudents(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type LecturerMarks interface {
	ChangeTestMark(c *gin.Context)
	ChangeLaboratoryMark(c *gin.Context)
	GetTestMarksFromGroup(c *gin.Context)
	GetLaboratoryMarksFromGroup(c *gin.Context)
	GiveExamMark(c *gin.Context)
	GetExamMark(c *gin.Context)
}

type LecturerAttendance interface {
	GetAllLessons(c *gin.Context)
	AddLesson(c *gin.Context)
	DeleteLesson(c *gin.Context)
	ChangeLesson(c *gin.Context)
	GetAllSeminars(c *gin.Context)
	ChangeSeminar(c *gin.Context)
	DeleteSeminar(c *gin.Context)
	AddSeminar(c *gin.Context)
	GetLessonVisitingGroup(c *gin.Context)
	GetSeminarVisitingGroup(c *gin.Context)
	AddLessonVisiting(c *gin.Context)
	AddSeminarVisiting(c *gin.Context)
	ChangeSeminarVisiting(c *gin.Context)
	ChangeLessonVisiting(c *gin.Context)
	GetLessonDate(c *gin.Context)
	ChangeLessonDate(c *gin.Context)
	AddLessonDate(c *gin.Context)
	DeleteLessonDate(c *gin.Context)
	GetTableLessons(c *gin.Context)
	GetTableSeminars(c *gin.Context)
}

type LecturerTestAndLab interface {
	GetAllTestFromSection(c *gin.Context)
	ChangeExternalLab(c *gin.Context)
	ChangeTest(c *gin.Context)
	CreateTest(c *gin.Context)
	DeleteTest(c *gin.Context)
	AddThemeForTest(c *gin.Context)
	DeleteAnswer(c *gin.Context)
	ChangeTheme(c *gin.Context)
	DeleteTheme(c *gin.Context)
	GetAllThemes(c *gin.Context)
	AddQuestionForTheme(c *gin.Context)
	GetQuestions(c *gin.Context)
	DeleteQuestion(c *gin.Context)
	ChangeQuestion(c *gin.Context)
	AddAnswerForQuestion(c *gin.Context)
	ChangeAnswer(c *gin.Context)
	GetAnswers(c *gin.Context)
	GetAllTests(c *gin.Context)
	GetAllExternalLab(c *gin.Context)
	CreateQuestion(c *gin.Context)
	GetAllQuestions(c *gin.Context)
	GetAllExistThemes(c *gin.Context)
	CreateTheme(c *gin.Context)
	DeleteQuestionFromTheme(c *gin.Context)
	DeleteThemeFromTest(c *gin.Context)
	ChangeThemeTestCount(c *gin.Context)
	OpenTest(c *gin.Context)
	GetOpenedTestForStudent(c *gin.Context)
	CloseOpenedTestForStudent(c *gin.Context)
	GetReportForTest(c *gin.Context)
	GetMarkTestForStudent(c *gin.Context)
	ChangeMarkTestForStudent(c *gin.Context)
	ExportTheme(c *gin.Context)
	ImportQuestions(c *gin.Context)
	GetQuestionWithoutEnglishVersion(c *gin.Context)
	GetQuestionWithoutTheme(c *gin.Context)
	GetQuestionsByName(c *gin.Context)
	GetUsersWithDoneTests(c *gin.Context)
	GetAllThemesByQuestion(c *gin.Context)
	ChangeLabLinc(c *gin.Context)
	ChangeLabToken(c *gin.Context)
	DeleteExternalLab(c *gin.Context)
	GetExternalLabInfo(c *gin.Context)
	CreateExternalLab(c *gin.Context)
	GetUsersWithDoneLabs(c *gin.Context)
	OpenLab(c *gin.Context)
	CloseOpenedLabForStudent(c *gin.Context)
	GetMarkLabForStudent(c *gin.Context)
	ChangeMarkLabForStudent(c *gin.Context)
}

type LecturerStudyGuide interface {
	AddStudyGuideHeader(c *gin.Context)
	GetStudyGuideHeader(c *gin.Context)
	DeleteStudyGuideHeader(c *gin.Context)
	ChangeDigitalGuideHeader(c *gin.Context)
	UploadFile(c *gin.Context)
	GetGuideFile(c *gin.Context)
	DeleteGuide(c *gin.Context)
	GetGuide(c *gin.Context)
	GetDigitalDiscipline(c *gin.Context)
	AddDigitalDiscipline(c *gin.Context)
	DeleteDigitalDiscipline(c *gin.Context)
}

type LecturerSection interface {
	AddSection(c *gin.Context)
	GetDisciplineSections(c *gin.Context)
	DeleteSection(c *gin.Context)
	ChangeSectionName(c *gin.Context)
	AddTestToSection(c *gin.Context)
	DeleteTestFromSection(c *gin.Context)
	AddLabToSection(c *gin.Context)
	DeleteLabFromSection(c *gin.Context)
	GetLabFromSection(c *gin.Context)
}

type LecturerReport interface {
	GetReport(c *gin.Context)
}

type LecturerController struct {
	LecturerPersonalData
	LecturerGroup
	LecturerDiscipline
	LecturerSeminarian
	LecturerStudents
	LecturerMarks
	LecturerAttendance
	LecturerTestAndLab
	LecturerSection
	LecturerStudyGuide
	LecturerReport
}

func NewLecturerController(service *service.Service) *LecturerController {
	return &LecturerController{
		LecturerPersonalData: NewLecturerPersonalDataHandler(service.LecturerPersonalData),
		LecturerGroup:        NewLecturerGroupHandler(service.LecturerGroup),
		LecturerDiscipline:   NewLecturerDisciplineHandler(service.LecturerDiscipline),
		LecturerSeminarian:   NewLecturerSeminarianHandler(service.LecturerSeminarian),
		LecturerStudents:     NewLecturerStudentsHandler(service.LecturerStudents),
		LecturerMarks:        NewLecturerMarksHandler(service.LecturerMarks),
		LecturerAttendance:   NewLecturerAttendanceHandler(service.LecturerAttendance),
		LecturerTestAndLab:   NewLecturerTestAndLabHandler(service.LecturerTestAndLab),
		LecturerSection:      NewLecturerSectionHandler(service.LecturerSection),
		LecturerStudyGuide:   NewLecturerStudyGuideHandler(service.LecturerStudyGuide),
		LecturerReport:       NewLecturerReportHandler(service.LecturerReport),
	}
}
