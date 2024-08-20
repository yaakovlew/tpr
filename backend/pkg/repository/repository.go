package repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/api_common_repository"
	"backend/pkg/repository/api_lecturer_repository"
	"backend/pkg/repository/api_seminarian_repository"
	"backend/pkg/repository/api_student_repository"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(email string) (model.User, error)
	GetPost(id int) (string, error)
	AddPost(id int, user model.User, tx *sql.Tx) error
	CheckGroup(groupId string) (int, error)
	GetUserMail(userId int) (string, error)
	ChangePassword(userId int, password string) error
}

type CommonGroup interface {
	GetAllGroups() ([]model.Group, error)
}

type CommonLab interface {
	ChangeLabDate(studentId, laboratoryId int) error
	ChangeLabMark(studentId, laboratoryId, mark int) error
	GiveAccessForLab(userId, labId int) (bool, error)
}

type StudentMarks interface {
	GetAllTestsMarks(id, disciplineId int) ([]model.TestsMarks, error)
	GetAllLaboratoryMarks(id, disciplineId int) ([]model.LaboratoryMarks, error)
	GetExamMark(userId, disciplineId int) (int, error)
}

type StudentPersonalData interface {
	GetPersonalData(id int) (model.User, error)
	UpdateSurname(id int, surname string) error
	UpdateName(id int, surname string) error
}

type StudentDiscipline interface {
	GetAllUserDiscipline(id int) ([]model.Discipline, []model.Discipline, error)
	GetDisciplineSections(disciplineId int) ([]model.Section, []model.Section, error)
}

type StudentAttendance interface {
	GetAllSeminarVisiting(disciplineId, userId int) ([]model.SeminarVisiting, error)
	GetAllLessonVisiting(disciplineId, userId int) ([]model.LessonVisiting, error)
	GetLessonDate(lessonId, groupId int) (model.LessonDate, error)
	GetAllSeminars(userId, disciplineId int) ([]model.SeminarDate, error)
	GetAllLessons(userId, disciplineId int) ([]model.LessonDate, error)
}

type StudentStudyGuide interface {
	GetDigitalDiscipline(disciplineId int) ([]model.DigitalDisciplineWithInfo, []model.DigitalDisciplineWithInfo, error)
	GetFilesIdFromDigital(digitalId int) ([]model.FileId, error)
	GetFile(fileId int) (string, error)
	CheckAccessForDiscipline(studentId, disciplineId int) error
	CheckAccessForLesson(studentId, lessonId int) error
	CheckAccessForFile(studentId, fileId int) error
}

type StudentTestAndLab interface {
	GetAllTestFromSection(userId, sectionId int) ([]model.Test, []model.Test, error)
	GetAllLab(userId, sectionId int) ([]model.LaboratoryWork, []model.LaboratoryWork, error)
	GiveAccessForTest(userId, testId int) (bool, error)
	CheckAccessForSection(studentId, sectionId int) error
	CheckAccessToOpenTest(userId, testId int) bool
	CheckAccessToOpenLab(userId, labId int) (bool, error)
	GetThemesQuestions(testId int) ([]model.QuestionWithTheme, int, error)
	GetQuestionsAndAnswersForTest(questions []model.QuestionWithTheme) ([]model.QuestionsWithAnswers, []model.QuestionsWithAnswers, error)
	GetTestMark(testId int) (int, error)
	GetTestMaxPoint(testId int) (int, error)
	GetQuestionAnswers(questionId int) (model.GetQuestionWithRightAnswer, int, error)
	GetWeightFromQuestionTheme(themeId int) int
	DoneOfTest(userId, testId int) error
	GetPersonalData(id int) (model.User, error)
	GetTest(testId int) (model.Test, model.Test, error)
	GetQuestion(id int) (model.Question, model.Question, error)
	GetAnswers(id int) ([]model.AllAnswersName, error)
	SaveResultOfTest(userId, testId, mark int) error
	GetResultOfTest(userId, testId int) (model.TestResult, error)
	GetResultOfLab(userId, labId int) (model.LabResult, error)
	GetAllOpenedTests(userId int, timeNow int64) ([]model.TestWithClosedDate, []model.TestWithClosedDate, error)
	GetAllDoneTests(userId int, timeNow int64) ([]model.TestWithClosedDate, []model.TestWithClosedDate, error)
	GetAllDoneLabs(userId int, timeNow int64) ([]model.LabWithClosedDate, []model.LabWithClosedDate, error)
	GetAllOpenedLabs(userId int, timeNow int64) ([]model.LabWithClosedDate, []model.LabWithClosedDate, error)
	GetLabLink(labId int) (model.LabCredential, error)
}

type LecturerPersonalData interface {
	GetPersonalData(id int) (model.User, error)
	UpdateSurname(id int, surname string) error
	UpdateName(id int, name string) error
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
	ChangeExamMark(userId, disciplineId, mark int) error
	GetAllMarksForExam(groupId, disciplineId int) ([]model.ExamMark, error)
	CheckExistMark(userId, disciplineId int) error
	MaxExamMark(disciplineId int) int
}

type LecturerAttendance interface {
	GetAllLessons(disciplineId int) ([]model.Lesson, error)
	AddLesson(disciplineId int, name string) error
	DeleteLesson(lessonId int) error
	ChangeLesson(lessonId int, name string) error
	AddSeminar(disciplineId, groupId int, date int, name string) error
	ChangeSeminarDate(seminarId int, date int) error
	GetLessonDate(lessonId, groupId int) (model.LessonDate, error)
	ChangeLessonDate(lessonId, groupId, date int) error
	ChangeLessonDateDescription(lessonId, groupId int, description string) error
	AddLessonDate(lessonId, groupId, date int, description string) error
	DeleteLessonDate(lessonId int, groupId int) error
	GetAllSeminars(disciplineId, groupId int) ([]model.Seminar, error)
	ChangeSeminar(seminarId int, name string) error
	DeleteSeminar(seminarId int) error
	GetLessonVisitingGroup(lessonId, groupId int) ([]model.LessonVisitingStudent, error)
	GetSeminarVisitingGroup(seminarId int) ([]model.SeminarVisitingStudent, error)
	AddLessonVisiting(lessonId, userId int, isAbsent bool) error
	AddSeminarVisiting(seminarId, userId int, isAbsent bool) error
	ChangeSeminarVisiting(seminarId, userId int, isAbsent bool) error
	ChangeLessonVisiting(lessonId, userId int, isAbsent bool) error
	GetTableLessons(disciplineId int) ([]model.LessonDate, error)
	GetTableSeminars(disciplineId int) ([]model.SeminarDate, error)
}

type LecturerTestAndLab interface {
	GetAllTests() ([]model.Test, []model.Test, error)
	GetAllExternalLab() ([]model.CommonLaboratoryWork, []model.CommonLaboratoryWork, error)
	GetAllTestFromSection(sectionId int) ([]model.Test, []model.Test, error)
	GetAllLabFromSection(sectionId int) ([]model.LaboratoryWork, []model.LaboratoryWork, error)
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
	DeleteQuestion(questionId int) error
	ChangeQuestionName(questionId int, name string) error
	ChangeQuestionNameEn(questionId int, name string) error
	AddQuestionForTheme(themeId, questionId int) error
	CreateQuestion(isVariable int, question string, questionEn string) (int, error)
	ChangeThemeWeight(themeId, weight int) error
	AddAnswerForQuestion(questionId int, name string, nameEn string, isRight bool) error
	ChangeAnswerName(answerId int, name string) error
	ChangeAnswerNameEn(answerId int, name string) error
	ChangeAnswerRight(answerId int, isRight bool) error
	GetAnswers(questionId int) ([]model.Answer, []model.Answer, error)
	CreateTestDate(userId, testId, date int) error
	UpdateTestDate(userId, testId, date int) error
	GetAllQuestions() ([]model.QuestionWithAmountAnswers, []model.QuestionWithAmountAnswers, error)
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
	ChangeTestMarkForStudent(studentId, testId, mark int) error
	GetTestMarkForStudent(studentId, testId int) (int, error)
	GetQuestionWithoutEnglishVersion() ([]model.Question, error)
	GetQuestionWithoutTheme() ([]model.Question, error)
	GetQuestionsByName(name string) ([]model.Question, error)
	GetUsersWithOpenedTest(testId int, timeNow int64) ([]model.StudentWithGroupWithClosedDate, error)
	GetUsersWithDoneTests(testId int, timeNow int64) ([]model.StudentWithGroupWithClosedDate, error)
	GetThemesByQuestion(questionId int) ([]model.Theme, error)
	CreateExternalLab(lab model.LaboratoryWorkInputWithoutId) (int, error)
	DeleteExternalLab(labId int) error
	ChangeLabLinc(labId int, linc string) error
	ChangeLabToken(labId int, token string) error
	GetLabToken(labId int) (string, error)
	GetLabInfo(labId int) (int, error)
	GetUsersWithOpenedLab(labId int, timeNow int64) ([]model.StudentWithGroupWithClosedDate, error)
	GetUsersWithDoneLaboratory(labId int, timeNow int64) ([]model.StudentWithGroupWithClosedDate, error)
	OpenLabForStudent(studentId, labId int, date int64) error
	CloseOpenedLabForStudent(studentId, labId int) error
	ChangeLabMarkForStudent(studentId, labId, mark int) error
	GetLabMarkForStudent(studentId, labId int) (int, error)
	GetExternalLabInfo(labId int) (model.LaboratoryWorkResponse, error)
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
	GetAttendanceSeminar(studentId, disciplineId int) int
	GetAttendanceLesson(studentId, disciplineId int) int
	GetThemesFromDiscipline(disciplineId int) ([]model.Section, error)
	GetMarkFromSection(userId, sectionId int) (int, error)
	GetMarkFromExam(userId, disciplineId int) int
	GetSummaryMarkFromSections(userId, disciplineId int) int
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
	GetAllStudentsFromGroup(id int) ([]model.Student, error)
	CheckAccessForGroup(seminarianId, groupId int) error
}

type SeminarianDiscipline interface {
	GetOwnDiscipline(userId int) ([]model.Discipline, []model.Discipline, error)
	GetDisciplineSections(disciplineId int) ([]model.Section, []model.Section, error)
	CheckAccessForDiscipline(seminarianId, disciplineId int) error
	GetAllInfoAboutDiscipline(id int) (model.DisciplineInfoDoubleLang, model.DisciplineInfoDoubleLang, error)
}

type SeminarianMark interface {
	GetTestMarksFromGroup(groupId, testId int) ([]model.GroupTestMarks, error)
	GetLaboratoryMarksFromGroup(groupId, laboratoryId int) ([]model.GroupLaboratoryMarks, error)
	CheckAccessForTest(seminarianId, groupId, testId int) error
	CheckAccessForLaboratory(seminarianId, groupId, labId int) error
	CheckAccessForGroup(seminarianId, groupId, disciplineId int) error
	CheckAccessForStudent(seminarianId, userId, disciplineId int) error
	CheckExistMark(userId, disciplineId int) error
	GetAllMarksForExam(groupId, disciplineId int) ([]model.ExamMark, error)
	ChangeExamMark(userId, disciplineId, mark int) error
	GiveExamMark(userId, disciplineId, mark int) error
	MaxExamMark(disciplineId int) int
}

type SeminarianAttendance interface {
	GetAllLessons(disciplineId int) ([]model.Lesson, error)
	AddSeminar(disciplineId, groupId int, name string, date int) error
	GetAllSeminars(disciplineId, groupId int) ([]model.Seminar, error)
	ChangeSeminar(seminarId int, name string) error
	DeleteSeminar(seminarId int) error
	GetLessonVisitingGroup(lessonId, groupId int) ([]model.LessonVisitingStudent, error)
	GetSeminarVisitingGroup(seminarId int) ([]model.SeminarVisitingStudent, error)
	AddLessonVisiting(lessonId, userId int, isAbsent bool) error
	AddSeminarVisiting(seminarId, userId int, isAbsent bool) error
	ChangeSeminarVisiting(seminarId, userId int, isAbsent bool) error
	ChangeLessonVisiting(lessonId, userId int, isAbsent bool) error
	ChangeSeminarDate(seminarId int, date int) error
	GetLessonDate(lessonId, groupId int) (model.LessonDate, error)
	GetTableLessons(disciplineId int) ([]model.LessonDate, error)
	GetTableSeminars(disciplineId int) ([]model.SeminarDate, error)
	CheckAccessForGroup(seminarianId, groupId, disciplineId int) error
	CheckAccessForStudent(seminarianId, userId, disciplineId int) error
	CheckAccessForDiscipline(seminarianId, disciplineId int) error
	CheckAccessForSeminar(seminarianId, seminarId int) error
	CheckAccessForLesson(seminarianId, lessonId, groupId int) error
	CheckAccessForStudentToLesson(seminarianId, userId, lessonId int) error
	CheckAccessForStudentToSeminar(seminarianId, userId, seminarId int) error
}

type SeminarianTestAndLab interface {
	GetAllTestFromSection(sectionId int) ([]model.Test, []model.Test, error)
	GetAllLab(sectionId int) ([]model.LaboratoryWork, []model.LaboratoryWork, error)
	CreateTestDate(userId, testId, date int) error
	UpdateTestDate(userId, testId, date int) error
	CheckExistTestDate(userId, testId int) bool
	CheckAccessForDiscipline(seminarianId, disciplineId int) error
	CheckAccessForSection(seminarianId, sectionId int) error
	CheckAccessToOpenTest(seminarianId, userId, testId int) error
	OpenTestForStudent(studentId, testId int, date int64) error
	GetOpenedTestForStudent(studentId, testId int) (model.OpenedTest, error)
	CloseOpenedTestForStudent(studentId, testId int) error
	GetUsersWithOpenedTest(seminarianId, testId int, timeNow int64) ([]model.StudentWithGroupWithClosedDate, error)
	GetUsersWithDoneTest(seminarianId, testId int, timeNow int64) ([]model.StudentWithGroupWithClosedDate, error)
	GetOpenedLabForStudent(studentId, labId int) (model.OpenedLab, error)
	CloseOpenedLabForStudent(studentId, labId int) error
	GetUsersWithOpenedLab(seminarianId, labId, groupId int, timeNow int64) ([]model.StudentWithGroupWithClosedDate, error)
	GetUsersWithDoneLab(seminarianId, labId, groupId int, timeNow int64) ([]model.StudentWithGroupWithClosedDate, error)
	CheckAccessToOpenLab(seminarianId, userId, labId int) error
	OpenLabForStudent(studentId, labId int, date int64) error
	GetLabToken(labId int) (string, error)
	GetInternalLabInfo(labId int) (int, error)
	GetExternalLabInfo(labId int) (model.LaboratoryWorkResponse, error)
}

type SeminarianStudyGuide interface {
	GetDigitalDiscipline(disciplineId int) ([]model.DigitalDisciplineWithInfo, []model.DigitalDisciplineWithInfo, error)
	GetFilesIdFromDigital(digitalId int) ([]model.FileId, error)
	GetFile(fileId int) (string, error)
	CheckAccessForDiscipline(seminarianId, disciplineId int) error
	CheckAccessForLesson(seminarianId, lessonId int) error
	CheckAccessForFile(seminarianId, fileId int) error
}

type SeminarianReport interface {
	GetAllStudents(groupId int) ([]model.StudentReport, error)
	GetAttendanceSeminar(studentId, disciplineId int) int
	GetAttendanceLesson(studentId, disciplineId int) int
	GetThemesFromDiscipline(disciplineId int) ([]model.Section, error)
	GetMarkFromSection(userId, sectionId int) (int, error)
	GetMarkFromExam(userId, disciplineId int) int
	GetSummaryMarkFromSections(userId, disciplineId int) int
	GetResult(studentId, disciplineId int) (string, string)
	GetSectionsResult(studentId, disciplineId int) string
	CheckAccessForGroup(seminarianId, groupId, disciplineId int) error
}

type Repository struct {
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

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:          api_common_repository.NewAuthPostgres(db),
		CommonGroup:            api_common_repository.NewCommonGroupPostgres(db),
		CommonLab:              api_common_repository.NewCommonLabPostgres(db),
		StudentMarks:           api_student_repository.NewStudentMarksPostgres(db),
		StudentPersonalData:    api_student_repository.NewStudentPersonalDataPostgres(db),
		StudentDiscipline:      api_student_repository.NewStudentDisciplinePostgres(db),
		StudentAttendance:      api_student_repository.NewStudentAttendancePostgres(db),
		StudentStudyGuide:      api_student_repository.NewStudentStudyGuidePostgres(db),
		StudentTestAndLab:      api_student_repository.NewStudentTestAndLabPostgres(db),
		LecturerPersonalData:   api_lecturer_repository.NewLecturerPersonalDataPostgres(db),
		LecturerGroup:          api_lecturer_repository.NewLecturerGroupPostgres(db),
		LecturerDiscipline:     api_lecturer_repository.NewLecturerDisciplinePostgres(db),
		LecturerSeminarian:     api_lecturer_repository.NewLecturerSeminarianPostgres(db),
		LecturerStudents:       api_lecturer_repository.NewLecturerStudentPostgres(db),
		LecturerMarks:          api_lecturer_repository.NewLecturerMarksPostgres(db),
		LecturerAttendance:     api_lecturer_repository.NewLecturerAttendancePostgres(db),
		LecturerTestAndLab:     api_lecturer_repository.NewLecturerTestAndLabPostgres(db),
		LecturerSection:        api_lecturer_repository.NewLecturerSectionPostgres(db),
		LecturerStudyGuide:     api_lecturer_repository.NewLecturerStudyGuidePostgres(db),
		LecturerReport:         api_lecturer_repository.NewLecturerReport(db),
		SeminarianPersonalData: api_seminarian_repository.NewSeminarianPersonalDataPostgres(db),
		SeminarianGroup:        api_seminarian_repository.NewSeminarianGroupPostgres(db),
		SeminarianDiscipline:   api_seminarian_repository.NewSeminarianDisciplinePostgres(db),
		SeminarianMark:         api_seminarian_repository.NewSeminarianMarksPostgres(db),
		SeminarianAttendance:   api_seminarian_repository.NewSeminarianAttendancePostgres(db),
		SeminarianTestAndLab:   api_seminarian_repository.NewSeminarianTestAndLabPostgres(db),
		SeminarianStudyGuide:   api_seminarian_repository.NewSeminarianStudyGuidePostgres(db),
		SeminarianReport:       api_seminarian_repository.NewSeminarianReport(db),
	}
}
