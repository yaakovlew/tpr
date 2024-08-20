package service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"backend/pkg/service/api_common_service"
	"backend/pkg/service/api_lecturer_service"
	"backend/pkg/service/api_seminarian_service"
	"backend/pkg/service/api_student_service"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(email, password string) (string, error)
	GenerateTokenByUserId(userId int) (string, error)
	ParseToken(token string) (int, string, error)
	ChangePassword(userId int, changePassword model.ChangePasswordInput) error
	ChangePasswordForStudent(studentId int, newPassword string) error
	ParseTokenForRestorePassword(accessToken string) (int, error)
	GenerateTokenForRestorePassword(email string) (string, error)
	GeneratePasswordResetLink(email string) (*model.PasswordResetLink, error)
	RestorePassword(userId int, password string) error
	GetPost(email string) (string, error)
}

type CommonGroup interface {
	GetAllGroups() ([]model.Group, error)
}

type CommonLab interface {
	ChangeLabDateAndMark(studentId, laboratoryId, percentage int) error
}

type StudentMarks interface {
	GetAllTestsMarks(id, disciplineId int) ([]model.TestsMarks, error)
	GetAllLaboratoryMarks(id, disciplineId int) ([]model.LaboratoryMarks, error)
	GetExamMark(userId, disciplineId int) (int, error)
}

type StudentPersonalData interface {
	GetPersonalData(id int) (model.User, error)
	UpdateName(id int, name string) error
	UpdateSurname(id int, surname string) error
}

type StudentDiscipline interface {
	GetAllUserDiscipline(id int) ([]model.Discipline, []model.Discipline, error)
	GetDisciplineSections(userId, disciplineId int) ([]model.Section, []model.Section, error)
}

type StudentAttendance interface {
	GetAllSeminarVisiting(disciplineId, userId int) ([]model.SeminarVisiting, error)
	GetAllLessonVisiting(disciplineId, userId int) ([]model.LessonVisiting, error)
	GetAllSeminars(userId, disciplineId int) ([]model.SeminarDate, error)
	GetAllLessons(userId, disciplineId int) ([]model.LessonDate, error)
}

type StudentStudyGuide interface {
	GetDigitalDiscipline(studentId, disciplineId int) ([]model.DigitalDisciplineWithInfo, []model.DigitalDisciplineWithInfo, error)
	GetFilesIdFromDigital(studentId, digitalId int) ([]model.FileId, error)
	GetFile(studentId, fileId int) (string, error)
}

type StudentTestAndLab interface {
	GetAllTestFromSection(userId, sectionId int) ([]model.Test, []model.Test, error)
	GetAllLab(userId, sectionId int) ([]model.LaboratoryWork, []model.LaboratoryWork, error)
	GetQuestionsForTest(userId, testId int) ([]model.QuestionsWithAnswers, []model.QuestionsWithAnswers, error)
	CheckAnswers(userId, testId int, answers []model.QuestionAndAnswerResponse) (int, []model.QuestionPercentage, error)
	GetResultOfTest(userId, testId int) (model.TestResult, error)
	GetPathForReportTest(userId, testId int) (string, error)
	GetAllOpenedTests(userId int) ([]model.TestWithClosedDate, []model.TestWithClosedDate, error)
	GetAllDoneTests(userId int) ([]model.TestWithClosedDate, []model.TestWithClosedDate, error)
	GetAllDoneLabs(userId int) ([]model.LabWithClosedDate, []model.LabWithClosedDate, error)
	GetAllOpenedLabs(userId int) ([]model.LabWithClosedDate, []model.LabWithClosedDate, error)
	GetResultOfLab(userId, labId int) (model.LabResult, error)
	GetLinkForLab(userId, labId int) (model.LabCredential, error)
}

type LecturerPersonalData interface {
	GetPersonalData(id int) (model.User, error)
	UpdateName(id int, name string) error
	UpdateSurname(id int, surname string) error
}

type LecturerGroup interface {
	AddGroup(name string) error
	DeleteGroup(id int) error
	GetAllStudentsFromGroup(id int) ([]model.Student, error)
	GetAllGroups() ([]model.Group, error)
	GetGroupsDisciplines(groupId int) ([]model.Discipline, []model.Discipline, error)
	ChangeName(groupId int, name string) error
	AddGroupInArchive(groupId int) error
	DeleteGroupFromArchive(groupId int) error
}

type LecturerDiscipline interface {
	GetGroupForDiscipline(id int) ([]model.Group, error)
	AddDiscipline(discipline model.AddNewDiscipline) (int, error)
	ChangeLessonMarks(id, lessonMarks int) error
	ChangeSeminarMarks(id, seminarMarks int) error
	GetAllDisciplines() ([]model.Discipline, []model.Discipline, error)
	GetAllInfoAboutDiscipline(id int) (model.DisciplineInfo, error)
	DeleteDiscipline(id int) error
	AddGroupToDiscipline(groupId, disciplineId int) error
	DeleteGroupFromDiscipline(groupId, disciplineId int) error
	ChangeDiscipline(disciplineId int, name string) error
	ChangeExamMark(id, examMark int) error
	ChangeDisciplineEn(disciplineId int, name string) error
	GetGroupsAvailableToAddForDiscipline(id int) ([]model.Group, error)
	ArchiveGroupToDiscipline(groupId, disciplineId int) error
}

type LecturerSeminarian interface {
	GetSeminarianFromGroupsAndDiscipline(groupId, disciplineId int) ([]model.Seminarian, error)
	GetAllSeminarians() ([]model.Seminarian, error)
	AddSeminarian(seminarianId, groupId, disciplineId int) error
	DeleteSeminarianFromGroupAndDiscipline(seminarianId, groupId, disciplineId int) error
}

type LecturerStudents interface {
	ChangeGroupForStudent(userId, groupId int) error
	GetAllStudents() ([]model.StudentWithGroup, error)
	DeleteUser(userId int) error
}

type LecturerMarks interface {
	ChangeTestMark(userId, testId, mark int) error
	ChangeLaboratoryMark(userId, laboratoryId, mark int) error
	GetTestMarksFromGroup(groupId, testId int) ([]model.GroupTestMarks, error)
	GetLaboratoryMarksFromGroup(groupId, laboratoryId int) ([]model.GroupLaboratoryMarks, error)
	GiveExamMark(userId, disciplineId, mark int) error
	GetAllMarksForExam(groupId, disciplineId int) ([]model.ExamMark, error)
}

type LecturerAttendance interface {
	GetAllLessons(disciplineId int) ([]model.Lesson, error)
	AddLesson(disciplineId int, name string) error
	DeleteLesson(lessonId int) error
	ChangeLesson(lessonId int, name string) error
	GetAllSeminars(disciplineId, groupId int) ([]model.Seminar, error)
	ChangeSeminar(seminarId int, name string) error
	DeleteSeminar(seminarId int) error
	GetLessonVisitingGroup(lessonId, groupId int) ([]model.LessonVisitingStudent, error)
	GetSeminarVisitingGroup(seminarId int) ([]model.SeminarVisitingStudent, error)
	AddLessonVisiting(lessonId, userId int, isAbsent bool) error
	AddSeminarVisiting(seminarId, userId int, isAbsent bool) error
	ChangeSeminarVisiting(seminarId, userId int, isAbsent bool) error
	ChangeLessonVisiting(lessonId, userId int, isAbsent bool) error
	AddSeminar(disciplineId, groupId int, date int, name string) error
	ChangeSeminarDate(seminarId int, date int) error
	GetLessonDate(lessonId, groupId int) (model.LessonDate, error)
	ChangeLessonDate(lessonId, groupId, date int) error
	ChangeLessonDateDescription(lessonId, groupId int, description string) error
	AddLessonDate(lessonId, groupId, date int, description string) error
	DeleteLessonDate(lessonId int, groupId int) error
	GetTableLessons(disciplineId int) ([]model.LessonDate, error)
	GetTableSeminars(disciplineId int) ([]model.SeminarDate, error)
}

type LecturerTestAndLab interface {
	GetAllTests() ([]model.Test, []model.Test, error)
	GetAllTestFromSection(sectionId int) ([]model.Test, []model.Test, error)
	GetAllLabFromSection(sectionId int) ([]model.LaboratoryWork, []model.LaboratoryWork, error)
	GetAllExternalLab() ([]model.CommonLaboratoryWork, []model.CommonLaboratoryWork, error)
	ChangeTestDuration(testId, minutes int) error
	ChangeLabDefaultMark(labId, mark int) error
	ChangeLabTaskDescription(labId int, description string) error
	ChangeLabName(labId int, name string) error
	ChangeTestTaskDescription(testId int, description string) error
	ChangeTestName(testId int, name string) error
	ChangeLabTaskDescriptionEn(labId int, description string) error
	ChangeLabNameEn(labId int, name string) error
	ChangeTestTaskDescriptionEn(testId int, description string) error
	ChangeTestNameEn(testId int, name string) error
	CreateTest(test model.TestAdd) error
	DeleteTest(testId int) error
	DeleteAnswer(answerId int) error
	ChangeThemeName(themeId int, name string) error
	DeleteTheme(themeId int) error
	GetQuestions(themeId int) ([]model.Question, []model.Question, error)
	ChangeQuestionNameEn(questionId int, name string) error
	DeleteQuestion(questionId int) error
	ChangeThemeWeight(themeId, weight int) error
	AddAnswerForQuestion(questionId int, name string, nameEn string, isRight bool) error
	ChangeAnswerName(answerId int, name string) error
	ChangeAnswerNameEn(answerId int, name string) error
	ChangeAnswerRight(answerId int, isRight bool) error
	GetAnswers(questionId int) ([]model.Answer, []model.Answer, error)
	ChangeQuestionName(questionId int, name string) error
	CreateQuestion(isVariable int, question string, questionEn string) (int, error)
	AddQuestionForTheme(themeId int, questionId int) error
	GetAllQuestions() ([]model.QuestionWithAmountAnswers, []model.QuestionWithAmountAnswers, error)
	GetThemesByQuestion(questionId int) ([]model.Theme, error)
	GetAllExistThemes() ([]model.Theme, error)
	AddThemeForTest(testId, themeId, count int) error
	CreateTheme(name string, weight int) (int, error)
	GetThemeIdByName(name string) (int, error)
	DeleteQuestionFromTheme(themeId, questionId int) error
	ChangeThemeTestCount(testId, themeId, count int) error
	DeleteThemeFromTest(testId, themeId int) error
	GetAllThemes(testId int) ([]model.ThemeOutput, error)
	OpenTestForStudent(studentId, testId int, date int64) error
	GetOpenedTestForStudent(studentId, testId int) (model.OpenedTest, error)
	CloseOpenedTestForStudent(studentId, testId int) error
	GetPathForReportTest(userId, testId int) (string, error)
	ChangeTestMarkForStudent(studentId, testId, mark int) error
	GetTestMarkForStudent(studentId, testId int) (int, error)
	GetEqualMarkForExport(questionId int) (float64, error)
	GetQuestionWithoutEnglishVersion() ([]model.Question, error)
	GetQuestionWithoutTheme() ([]model.Question, error)
	GetQuestionsByName(name string) ([]model.Question, error)
	GetUsersWithOpenedTest(testId int) ([]model.StudentWithGroupWithClosedDate, error)
	GetUsersWithDoneTests(testId int) ([]model.StudentWithGroupWithClosedDate, error)
	CreateExternalLab(lab model.LaboratoryWorkInputWithoutId) (int, error)
	DeleteExternalLab(labId int) error
	ChangeLabLinc(labId int, linc string) error
	ChangeLabToken(labId int, token string) error
	GetExternalLabInfo(labId int) (model.LaboratoryWorkResponse, error)
	GetUsersWithOpenedLab(labId int) ([]model.StudentWithGroupWithClosedDate, error)
	GetUsersWithDoneLaboratory(labId int) ([]model.StudentWithGroupWithClosedDate, error)
	OpenLabForStudent(studentId, labId int, date int64) error
	CloseOpenedLabForStudent(studentId, labId int) error
	ChangeLabMarkForStudent(studentId, labId, mark int) error
	GetLabMarkForStudent(studentId, labId int) (int, error)
}

type LecturerSection interface {
	AddSection(name string, nameEn string, disciplineId int) error
	GetDisciplineSections(disciplineId int) ([]model.Section, []model.Section, error)
	DeleteSection(sectionId int) error
	ChangeSectionName(sectionId int, name string) error
	AddTestToSection(sectionId, testId int) error
	DeleteTestFromSection(sectionId, testId int) error
	AddLabToSection(labId, sectionId, defaultMark int) error
	DeleteLabFromSection(labId, sectionId int) error
	ChangeSectionNameEn(sectionId int, name string) error
	GetLabFromSection(sectionId int) ([]model.LaboratoryWorkWithExternal, error)
}

type LecturerStudyGuide interface {
	AddStudyGuideHeader(name, nameEn, description, descriptionEn string) (int, error)
	GetStudyGuideHeader() ([]model.StudyGuideHeader, []model.StudyGuideHeader, error)
	DeleteStudyGuideHeader(digitalGuideId int) (int, error)
	ChangeNameDigitalGuideHeader(digitalGuideId int, name string) error
	ChangeDescriptionDigitalGuideHeader(digitalGuideId int, description string) error
	GetDigitalDiscipline(disciplineId int) ([]model.DigitalDiscipline, error)
	AddDigitalDiscipline(digitalMaterialId, disciplineId int) error
	DeleteDigitalDiscipline(digitalMaterialId, disciplineId int) error
	GetFilesIdFromDigital(digitalId int) ([]model.FileId, error)
	AddFileToDigital(path string, digitalId int) error
	DeleteFileFromDigital(fileId int) (string, error)
	GetFile(fileId int) (string, error)
	ChangeNameDigitalGuideHeaderEn(digitalGuideId int, name string) error
	ChangeDescriptionDigitalGuideHeaderEn(digitalGuideId int, description string) error
}

type LecturerReport interface {
	GetAllStudents(groupId int) ([]model.StudentReport, error)
	GetThemesFromDiscipline(disciplineId int) ([]model.Section, error)
	GetMarkFromSection(userId, sectionId int) (int, error)
	GetSummaryMarkFromSections(userId, disciplineId int) int
	GetMarkFromExam(userId, disciplineId int) int
	GetFinalGrade(userId, disciplineId int) int
	GetResult(studentId, disciplineId int) (string, string)
	GetSectionsResult(studentId, disciplineId int) string
}

type SeminarianPersonalData interface {
	GetPersonalData(id int) (model.User, error)
	UpdateSurname(id int, surname string) error
	UpdateName(id int, name string) error
}

type SeminarianGroup interface {
	GetOwnGroup(userId int, disciplineId int) ([]model.Group, error)
	GetAllStudentsFromGroup(seminarianId, groupId int) ([]model.Student, error)
}

type SeminarianDiscipline interface {
	GetOwnDiscipline(userId int) ([]model.Discipline, []model.Discipline, error)
	GetDisciplineSections(seminarianId, disciplineId int) ([]model.Section, []model.Section, error)
	GetAllInfoAboutDiscipline(seminarianId, disciplineId int) (model.DisciplineInfoDoubleLang, model.DisciplineInfoDoubleLang, error)
}

type SeminarianMark interface {
	GetTestMarksFromGroup(seminarianId, groupId, testId int) ([]model.GroupTestMarks, error)
	GetLaboratoryMarksFromGroup(seminarianId, groupId, laboratoryId int) ([]model.GroupLaboratoryMarks, error)
	GiveExamMark(seminarianId, userId, disciplineId, mark int) error
	GetAllMarksForExam(seminarianId, groupId, disciplineId int) ([]model.ExamMark, error)
}

type SeminarianAttendance interface {
	GetAllLessons(seminarianId, disciplineId int) ([]model.Lesson, error)
	AddSeminar(seminarianId, disciplineId, groupId int, name string, date int) error
	GetAllSeminars(seminarianId, disciplineId, groupId int) ([]model.Seminar, error)
	ChangeSeminar(seminarianId, seminarId int, name string) error
	DeleteSeminar(seminarianId, seminarId int) error
	GetLessonVisitingGroup(seminarianId, lessonId, groupId int) ([]model.LessonVisitingStudent, error)
	GetSeminarVisitingGroup(seminarianId, seminarId int) ([]model.SeminarVisitingStudent, error)
	AddLessonVisiting(seminarianId, lessonId, userId int, isAbsent bool) error
	AddSeminarVisiting(seminarianId, seminarId, userId int, isAbsent bool) error
	ChangeSeminarVisiting(seminarianId, seminarId, userId int, isAbsent bool) error
	ChangeLessonVisiting(seminarianId, lessonId, userId int, isAbsent bool) error
	ChangeSeminarDate(seminarianId, seminarId int, date int) error
	GetLessonDate(seminarianId, lessonId, groupId int) (model.LessonDate, error)
	GetTableLessons(seminarianId, disciplineId int) ([]model.LessonDate, error)
	GetTableLessonsByGroup(seminarianId, disciplineId, groupId int) ([]model.LessonDate, error)
	GetTableSeminars(seminarianId, disciplineId, groupId int) ([]model.SeminarDate, error)
}

type SeminarianTestAndLab interface {
	GetAllTestFromSection(seminarianId, sectionId int) ([]model.Test, []model.Test, error)
	GetAllLab(seminarianId, sectionId int) ([]model.LaboratoryWork, []model.LaboratoryWork, error)
	OpenTestForStudent(seminarianId, studentId, testId int, date int64) error
	GetOpenedTestForStudent(seminarianId, studentId, testId int) (model.OpenedTest, error)
	CloseOpenedTestForStudent(seminarianId, studentId, testId int) error
	GetPathForReportTest(seminarianId, userId, testId int) (string, error)
	GetUsersWithOpenedTest(seminarianId, testId int) ([]model.StudentWithGroupWithClosedDate, error)
	GetUsersWithDoneTest(seminarianId, testId int) ([]model.StudentWithGroupWithClosedDate, error)
	GetOpenedLabForStudent(seminarianId, studentId, labId int) (model.OpenedLab, error)
	CloseOpenedLabForStudent(seminarianId, studentId, labId int) error
	GetUsersWithOpenedLab(seminarianId, labId, groupId int) ([]model.StudentWithGroupWithClosedDate, error)
	GetUsersWithDoneLab(seminarianId, labId, groupId int) ([]model.StudentWithGroupWithClosedDate, error)
	OpenLabForStudent(seminarianId, studentId, labId int, date int64) error
}

type SeminarianStudyGuide interface {
	GetDigitalDiscipline(seminarianId, disciplineId int) ([]model.DigitalDisciplineWithInfo, []model.DigitalDisciplineWithInfo, error)
	GetFilesIdFromDigital(seminarianId, digitalId int) ([]model.FileId, error)
	GetFile(seminarianId, fileId int) (string, error)
}

type SeminarianReport interface {
	GetAllStudents(groupId int) ([]model.StudentReport, error)
	GetThemesFromDiscipline(disciplineId int) ([]model.Section, error)
	GetMarkFromSection(userId, sectionId int) (int, error)
	GetSummaryMarkFromSections(userId, disciplineId int) int
	GetMarkFromExam(userId, disciplineId int) int
	GetFinalGrade(userId, disciplineId int) int
	GetResult(studentId, disciplineId int) (string, string)
	GetSectionsResult(studentId, disciplineId int) string
	CheckAccessForGroup(seminarianId, groupId, disciplineId int) error
}

type Service struct {
	Authorization
	CommonGroup
	CommonLab
	StudentMarks
	StudentPersonalData
	StudentDiscipline
	StudentAttendance
	StudentStudyGuide
	StudentTestAndLab
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
	SeminarianPersonalData
	SeminarianGroup
	SeminarianDiscipline
	SeminarianMark
	SeminarianAttendance
	SeminarianTestAndLab
	SeminarianStudyGuide
	SeminarianReport
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization:          api_common_service.NewAuthService(repo.Authorization),
		CommonGroup:            api_common_service.NewCommonGroupService(repo.CommonGroup),
		CommonLab:              api_common_service.NewCommonLabService(repo.CommonLab),
		StudentMarks:           api_student_service.NewStudentMarksService(repo.StudentMarks),
		StudentPersonalData:    api_student_service.NewStudentPersonalDataServices(repo.StudentPersonalData),
		StudentDiscipline:      api_student_service.NewStudentDisciplinesService(repo.StudentDiscipline),
		StudentAttendance:      api_student_service.NewStudentAttendanceService(repo.StudentAttendance),
		StudentStudyGuide:      api_student_service.NewStudentStudyGuideService(repo.StudentStudyGuide),
		StudentTestAndLab:      api_student_service.NewStudentTestAndLabService(repo.StudentTestAndLab),
		LecturerPersonalData:   api_lecturer_service.NewLecturerPersonalDataServices(repo.LecturerPersonalData),
		LecturerGroup:          api_lecturer_service.NewLecturerGroupService(repo.LecturerGroup),
		LecturerDiscipline:     api_lecturer_service.NewLecturerDisciplineService(repo.LecturerDiscipline),
		LecturerSeminarian:     api_lecturer_service.NewLecturerSeminarianService(repo.LecturerSeminarian),
		LecturerStudents:       api_lecturer_service.NewLecturerStudentService(repo.LecturerStudents),
		LecturerMarks:          api_lecturer_service.NewLecturerMarksService(repo.LecturerMarks),
		LecturerAttendance:     api_lecturer_service.NewLecturerAttendanceService(repo.LecturerAttendance),
		LecturerTestAndLab:     api_lecturer_service.NewLecturerTestAndLabService(repo.LecturerTestAndLab),
		LecturerSection:        api_lecturer_service.NewLecturerSectionService(repo.LecturerSection),
		LecturerStudyGuide:     api_lecturer_service.NewLecturerStudyGuideService(repo.LecturerStudyGuide),
		LecturerReport:         api_lecturer_service.NewLecturerReportService(repo.LecturerReport),
		SeminarianPersonalData: api_seminarian_service.NewSeminarianPersonalDataService(repo.SeminarianPersonalData),
		SeminarianGroup:        api_seminarian_service.NewSeminarianGroupService(repo.SeminarianGroup),
		SeminarianDiscipline:   api_seminarian_service.NewSeminarianDisciplineService(repo.SeminarianDiscipline),
		SeminarianMark:         api_seminarian_service.NewSeminarianMarksService(repo.SeminarianMark),
		SeminarianAttendance:   api_seminarian_service.NewSeminarianAttendanceService(repo.SeminarianAttendance),
		SeminarianTestAndLab:   api_seminarian_service.NewSeminarianTestAndLabService(repo.SeminarianTestAndLab),
		SeminarianStudyGuide:   api_seminarian_service.NewSeminarianStudyGuideService(repo.SeminarianStudyGuide),
		SeminarianReport:       api_seminarian_service.NewSeminarianReportService(repo.SeminarianReport),
	}
}
