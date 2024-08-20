package handler

import (
	_ "backend/docs"
	"backend/pkg/handler/api_common"
	"backend/pkg/handler/api_lecturer"
	"backend/pkg/handler/api_seminarian"
	"backend/pkg/handler/api_student"
	"backend/pkg/handler/middleware"
	"backend/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

type Controller struct {
	StudentApi    *api_student.StudentController
	LecturerApi   *api_lecturer.LecturerController
	SeminarianApi *api_seminarian.SeminarianController
	CommonApi     *api_common.CommonController
	MiddleWare    *middleware.MiddleWareController
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		StudentApi:    api_student.NewStudentController(service),
		CommonApi:     api_common.NewCommonController(service),
		MiddleWare:    middleware.NewMiddleWareController(service),
		LecturerApi:   api_lecturer.NewLecturerController(service),
		SeminarianApi: api_seminarian.NewSeminarianController(service),
	}
}

func (h *Controller) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://soft-computing-mephi.ru.na4u.ru", "http://soft-computing-mephi.ru"},
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	lab := router.Group("/lab", h.MiddleWare.CheckHeaderForWebhook)
	{
		lab.POST("", h.CommonApi.WebhookForLab)

		security := lab.Group("", h.MiddleWare.UserIdentityStudent)
		{
			security.GET("", h.CommonApi.WebhookForGetUser)
		}
	}

	router.GET("/groups", h.CommonApi.GetAllGroup)
	router.POST("/forget-password", h.CommonApi.RestorePasswordLync)

	changePass := router.Group("/restore-password", h.MiddleWare.UserIdentity)
	{
		changePass.POST("", h.CommonApi.RestorePassword)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.CommonApi.SignUp)
		auth.POST("/sign-in", h.CommonApi.SignIn)
	}

	api := router.Group("/api")
	{
		student := api.Group("/student", h.MiddleWare.UserIdentityStudent)
		{
			marks := student.Group("/marks")
			{
				marks.GET("/test/:id", h.StudentApi.GetAllTestMarks)
				marks.GET("/laboratory-work/:id", h.StudentApi.GetAllLaboratoryMarks)
				marks.GET("/exam/:id", h.StudentApi.GetExamMark)
			}

			discipline := student.Group("/disciplines")
			{
				discipline.GET("", h.StudentApi.GetUserDiscipline)
				discipline.GET("/section/:id", h.StudentApi.GetDisciplineSections)
				discipline.GET("/test/:id", h.StudentApi.GetAllTestFromSection)
				discipline.GET("/laboratory/:id", h.StudentApi.GetAllLabFromSection)
			}

			personalData := student.Group("/personal-data")
			{
				personalData.GET("", h.StudentApi.GetPersonalData)
				personalData.PUT("", h.StudentApi.UpdatePersonalData)
				personalData.PUT("/change-password", h.CommonApi.ChangePassword)
			}

			attendance := student.Group("/attendance")
			{
				attendance.GET("/seminars/:id", h.StudentApi.GetAllSeminarVisiting)
				attendance.GET("/lessons/:id", h.StudentApi.GetAllLessonVisiting)
				attendance.GET("/seminars/date/:id", h.StudentApi.GetSeminars)
				attendance.GET("/lessons/date/:id", h.StudentApi.GetLessons)
			}

			materials := student.Group("/material")
			{
				materials.GET("/:id", h.StudentApi.GetAllLessonsForDiscipline)
				materials.GET("/files/:id", h.StudentApi.GetGuide)
				materials.GET("/download/:id", h.StudentApi.GetGuideFile)
			}

			test := student.Group("/test")
			{
				test.GET("", h.StudentApi.GetAllDoneTests)
				test.GET("/mark/:id", h.StudentApi.GetTestMark)
				test.GET("/report/:id", h.StudentApi.GetReportForTest)
				test.GET("/:id", h.StudentApi.GetQuestionsForTest)
				test.POST("/:id", h.StudentApi.PassTest)
			}

			labs := student.Group("/laboratory-work")
			{
				labs.GET("/:id", h.StudentApi.OpenLab)
				labs.GET("", h.StudentApi.GetAllDoneLabs)
				labs.GET("/mark/:id", h.StudentApi.GetLabMark)
			}
		}

		seminarian := api.Group("/seminarian", h.MiddleWare.UserIdentitySeminarian)
		{

			materials := seminarian.Group("/material")
			{
				materials.GET("/:id", h.SeminarianApi.GetAllLessonsForDiscipline)
				materials.GET("/files/:id", h.SeminarianApi.GetGuide)
				materials.GET("/download/:id", h.SeminarianApi.GetGuideFile)
			}

			personalData := seminarian.Group("personal-data")
			{
				personalData.GET("", h.SeminarianApi.GetPersonalData)
				personalData.PUT("", h.SeminarianApi.UpdatePersonalData)
				personalData.PUT("/change-password", h.CommonApi.ChangePassword)
			}

			discipline := seminarian.Group("/discipline")
			{
				discipline.GET("", h.SeminarianApi.GetOwnDiscipline)
				discipline.GET("/:id", h.SeminarianApi.GetAllInfoAboutDiscipline)
				discipline.GET("/section/:id", h.SeminarianApi.GetDisciplineSections)
				discipline.GET("/test/:id", h.SeminarianApi.GetAllTestFromSection)
				discipline.GET("/laboratory-work/:id", h.SeminarianApi.GetAllLabFromSection)

				group := discipline.Group("/group")
				{
					group.GET("/:id", h.SeminarianApi.GetOwnGroup)
					group.GET("/students/:id", h.SeminarianApi.GetAllStudentsFromGroup)
				}
			}

			tests := seminarian.Group("/test")
			{
				tests.GET("/students", h.SeminarianApi.GetUsersWithDoneTests)

				testActivity := tests.Group("/activity")
				{
					testActivity.GET("", h.SeminarianApi.GetOpenedTestForStudent)
					testActivity.GET("/report", h.SeminarianApi.GetReportForTest)
					testActivity.POST("", h.SeminarianApi.OpenTest)
					testActivity.DELETE("", h.SeminarianApi.CloseOpenedTestForStudent)
				}
			}

			labs := seminarian.Group("/laboratory-work")
			{
				labs.GET("/students", h.SeminarianApi.GetUsersWithDoneLab)

				activity := labs.Group("/activity")
				{
					activity.POST("", h.SeminarianApi.OpenLab)
					activity.DELETE("", h.LecturerApi.CloseOpenedLabForStudent)
				}
			}

			mark := seminarian.Group("/mark")
			{
				mark.GET("/test", h.SeminarianApi.GetTestMarksFromGroup)
				mark.GET("/laboratory", h.SeminarianApi.GetLaboratoryMarksFromGroup)
				mark.GET("/exam", h.SeminarianApi.GetExamMark)
				mark.POST("/exam", h.SeminarianApi.GiveExamMark)
			}

			attendance := seminarian.Group("/attendance")
			{

				lesson := attendance.Group("/lesson")
				{
					lesson.GET("/:id", h.SeminarianApi.GetAllLessons)
					lesson.GET("/date", h.SeminarianApi.GetLessonDate)
					lesson.GET("/table", h.SeminarianApi.GetTableLessons)

					lessonVisiting := lesson.Group("/visiting")
					{
						lessonVisiting.POST("", h.SeminarianApi.AddLessonVisiting)
						lessonVisiting.GET("", h.SeminarianApi.GetLessonVisitingGroup)
						lessonVisiting.PUT("", h.SeminarianApi.ChangeLessonVisiting)
					}
				}

				seminar := attendance.Group("/seminar")
				{
					seminar.GET("", h.SeminarianApi.GetAllSeminars)
					seminar.GET("/table", h.SeminarianApi.GetTableSeminars)
					seminar.POST("", h.SeminarianApi.AddSeminar)
					seminar.PUT("", h.SeminarianApi.ChangeSeminar)
					seminar.DELETE("/:id", h.SeminarianApi.DeleteSeminar)

					seminarVisiting := seminar.Group("/visiting")
					{
						seminarVisiting.POST("", h.SeminarianApi.AddSeminarVisiting)
						seminarVisiting.GET("/:id", h.SeminarianApi.GetSeminarVisitingGroup)
						seminarVisiting.PUT("", h.SeminarianApi.ChangeSeminarVisiting)
					}
				}
			}
			group := seminarian.Group("/group")
			{
				group.GET("/report", h.SeminarianApi.GetReport)
			}

		}

		lecturer := api.Group("/lecturer", h.MiddleWare.UserIdentityLecturer)
		{

			material := lecturer.Group("/material")
			{
				material.GET("", h.LecturerApi.GetStudyGuideHeader)
				material.POST("", h.LecturerApi.AddStudyGuideHeader)
				material.DELETE("/:id", h.LecturerApi.DeleteStudyGuideHeader)
				material.PUT("", h.LecturerApi.ChangeDigitalGuideHeader)
				material.POST("/upload/:id", h.LecturerApi.UploadFile)
				material.GET("/:id", h.LecturerApi.GetGuide)
				material.GET("/download/:id", h.LecturerApi.GetGuideFile)
				material.DELETE("/file/:id", h.LecturerApi.DeleteGuide)

				digitalDiscipline := material.Group("/digital-lesson")
				{
					digitalDiscipline.GET("/:id", h.LecturerApi.GetDigitalDiscipline)
					digitalDiscipline.DELETE("", h.LecturerApi.DeleteDigitalDiscipline)
					digitalDiscipline.POST("", h.LecturerApi.AddDigitalDiscipline)
				}
			}

			attendance := lecturer.Group("/attendance")
			{

				lesson := attendance.Group("/lesson")
				{
					lesson.GET("/:id", h.LecturerApi.GetAllLessons)
					lesson.GET("", h.LecturerApi.GetTableLessons)
					lesson.POST("", h.LecturerApi.AddLesson)
					lesson.PUT("", h.LecturerApi.ChangeLesson)
					lesson.DELETE("/:id", h.LecturerApi.DeleteLesson)

					date := lesson.Group("/date")
					{
						date.PUT("", h.LecturerApi.ChangeLessonDate)
						date.POST("", h.LecturerApi.AddLessonDate)
						date.DELETE("", h.LecturerApi.DeleteLessonDate)
					}

					lessonVisiting := lesson.Group("/visiting")
					{
						lessonVisiting.POST("", h.LecturerApi.AddLessonVisiting)
						lessonVisiting.GET("", h.LecturerApi.GetLessonVisitingGroup)
						lessonVisiting.PUT("", h.LecturerApi.ChangeLessonVisiting)
					}
				}

				seminar := attendance.Group("/seminar")
				{
					seminar.GET("", h.LecturerApi.GetAllSeminars)
					seminar.POST("", h.LecturerApi.AddSeminar)
					seminar.PUT("", h.LecturerApi.ChangeSeminar)
					seminar.DELETE("/:id", h.LecturerApi.DeleteSeminar)
					seminar.GET("/table/:id", h.LecturerApi.GetTableSeminars)

					seminarVisiting := seminar.Group("/visiting")
					{
						seminarVisiting.POST("", h.LecturerApi.AddSeminarVisiting)
						seminarVisiting.GET("/:id", h.LecturerApi.GetSeminarVisitingGroup)
						seminarVisiting.PUT("", h.LecturerApi.ChangeSeminarVisiting)
					}
				}
			}

			mark := lecturer.Group("/mark")
			{
				mark.PUT("/laboratory", h.LecturerApi.ChangeLaboratoryMark)
				mark.PUT("/test", h.LecturerApi.ChangeTestMark)
				mark.GET("/test", h.LecturerApi.GetTestMarksFromGroup)
				mark.GET("/laboratory", h.LecturerApi.GetLaboratoryMarksFromGroup)
				mark.GET("/exam", h.LecturerApi.GetExamMark)
				mark.POST("/exam", h.LecturerApi.GiveExamMark)
			}

			personalData := lecturer.Group("/personal-data")
			{
				personalData.GET("", h.LecturerApi.GetPersonalData)
				personalData.PUT("", h.LecturerApi.UpdatePersonalData)
				personalData.PUT("/change-password", h.CommonApi.ChangePassword)
			}

			group := lecturer.Group("/group")
			{
				group.GET("", h.LecturerApi.GetAllGroups)
				group.GET("/students/:id", h.LecturerApi.GetAllStudentsFromGroup)
				group.GET("/seminarian", h.LecturerApi.GetSeminarianFromGroupAndDiscipline)
				group.GET("/discipline/:id", h.LecturerApi.GetGroupsDisciplines)
				group.POST("", h.LecturerApi.AddGroup)
				group.DELETE("/:id", h.LecturerApi.DeleteGroup)
				group.PUT("", h.LecturerApi.ChangeGroupName)
				group.GET("/report", h.LecturerApi.GetReport)
				group.POST("/archive/:id", h.LecturerApi.AddGroupInArchive)
			}

			students := lecturer.Group("/student")
			{
				students.PUT("", h.LecturerApi.ChangeGroupForStudent)
				students.GET("", h.LecturerApi.GetAllStudents)
				students.DELETE("/:id", h.LecturerApi.DeleteUser)
				students.POST("/open-test", h.LecturerApi.OpenTest)
				students.PUT("/change-password", h.CommonApi.ChangePasswordForStudentAndSeminarianFromLecturer)
			}

			seminarians := lecturer.Group("/seminarian")
			{
				seminarians.GET("", h.LecturerApi.GetAllSeminarians)
				seminarians.POST("", h.LecturerApi.AddSeminarianToGroup)
				seminarians.DELETE("", h.LecturerApi.DeleteSeminarianFromGroupAndDiscipline)
				seminarians.PUT("/change-password", h.CommonApi.ChangePasswordForStudentAndSeminarianFromLecturer)
			}

			discipline := lecturer.Group("/discipline")
			{
				groups := discipline.Group("/group")
				{
					groups.GET("", h.LecturerApi.GetAllGroupForDiscipline)
					groups.POST("", h.LecturerApi.AddGroupToDiscipline)
					groups.DELETE("", h.LecturerApi.DeleteGroupFromDiscipline)
				}

				discipline.POST("", h.LecturerApi.AddDiscipline)
				discipline.GET("", h.LecturerApi.GetAllDisciplines)
				discipline.DELETE("/:id", h.LecturerApi.DeleteDiscipline)
				discipline.GET("/:id", h.LecturerApi.GetAllInfoAboutDiscipline)
				discipline.PUT("", h.LecturerApi.ChangeDiscipline)
				discipline.PUT("/archive", h.LecturerApi.ArchiveGroupToDiscipline)

				marks := discipline.Group("/mark")
				{
					marks.PUT("/seminar", h.LecturerApi.ChangeSeminarMarks)
					marks.PUT("/lesson", h.LecturerApi.ChangeLessonMarks)
					marks.PUT("/exam", h.LecturerApi.ChangeExamMark)
				}

				section := discipline.Group("/section")
				{
					section.GET("/:id", h.LecturerApi.GetDisciplineSections)
					section.POST("", h.LecturerApi.AddSection)
					section.DELETE("/:id", h.LecturerApi.DeleteSection)
					section.PUT("", h.LecturerApi.ChangeSectionName)

					test := section.Group("/test")
					{
						test.GET("/:id", h.LecturerApi.GetAllTestFromSection)
						test.POST("", h.LecturerApi.AddTestToSection)
						test.DELETE("", h.LecturerApi.DeleteTestFromSection)
					}

					laboratory := section.Group("/laboratory-work")
					{
						laboratory.POST("", h.LecturerApi.AddLabToSection)
						laboratory.DELETE("", h.LecturerApi.DeleteLabFromSection)
						laboratory.GET("/:id", h.LecturerApi.GetLabFromSection)
					}
				}

				test := lecturer.Group("/test")
				{
					test.POST("", h.LecturerApi.CreateTest)
					test.DELETE("/:id", h.LecturerApi.DeleteTest)
					test.GET("", h.LecturerApi.GetAllTests)
					test.PUT("", h.LecturerApi.ChangeTest)
					test.GET("/mark", h.LecturerApi.GetMarkTestForStudent)
					test.PUT("/mark", h.LecturerApi.ChangeMarkTestForStudent)
					test.GET("/students", h.LecturerApi.GetUsersWithDoneTests)

					testActivity := test.Group("/activity")
					{
						testActivity.GET("", h.LecturerApi.GetOpenedTestForStudent)
						testActivity.GET("/report", h.LecturerApi.GetReportForTest)
						testActivity.POST("", h.LecturerApi.OpenTest)
						testActivity.DELETE("", h.LecturerApi.CloseOpenedTestForStudent)
					}

					theme := test.Group("/theme")
					{
						theme.GET("", h.LecturerApi.GetAllExistThemes)
						theme.POST("", h.LecturerApi.CreateTheme)
						theme.POST("/add", h.LecturerApi.AddThemeForTest)
						theme.PUT("", h.LecturerApi.ChangeTheme)
						theme.DELETE("/:id", h.LecturerApi.DeleteTheme)
						theme.GET("/:id", h.LecturerApi.GetAllThemes)
						theme.PUT("/count", h.LecturerApi.ChangeThemeTestCount)
						theme.DELETE("", h.LecturerApi.DeleteThemeFromTest)
						theme.GET("/export", h.LecturerApi.ExportTheme)

						question := theme.Group("/question")
						{
							question.GET("", h.LecturerApi.GetAllQuestions)
							question.GET("/by/:id", h.LecturerApi.GetAllThemesByQuestion)
							question.GET("/:id", h.LecturerApi.GetQuestions)
							question.GET("/name", h.LecturerApi.GetQuestionsByName)
							question.POST("", h.LecturerApi.AddQuestionForTheme)
							question.DELETE("", h.LecturerApi.DeleteQuestionFromTheme)
							question.POST("/create", h.LecturerApi.CreateQuestion)
							question.DELETE("/:id", h.LecturerApi.DeleteQuestion)
							question.PUT("", h.LecturerApi.ChangeQuestion)
							question.POST("/import", h.LecturerApi.ImportQuestions)
							question.GET("/without-english", h.LecturerApi.GetQuestionWithoutEnglishVersion)
							question.GET("/without-theme", h.LecturerApi.GetQuestionWithoutTheme)

							answer := question.Group("/answer")
							{
								answer.POST("", h.LecturerApi.AddAnswerForQuestion)
								answer.PUT("", h.LecturerApi.ChangeAnswer)
								answer.GET("/:id", h.LecturerApi.GetAnswers)
								answer.DELETE("/:id", h.LecturerApi.DeleteAnswer)
							}
						}
					}
				}

				lab := lecturer.Group("/laboratory-work")
				{
					lab.GET("/students", h.LecturerApi.GetUsersWithDoneLabs)

					external := lab.Group("/external")
					{
						external.POST("", h.LecturerApi.CreateExternalLab)
						external.GET("", h.LecturerApi.GetAllExternalLab)
						external.GET("/:id", h.LecturerApi.GetExternalLabInfo)
						external.DELETE("/:id", h.LecturerApi.DeleteExternalLab)
						external.PUT("", h.LecturerApi.ChangeExternalLab)
					}

					activity := lab.Group("/activity")
					{
						activity.POST("", h.LecturerApi.OpenLab)
						activity.DELETE("", h.LecturerApi.CloseOpenedLabForStudent)
					}

					marksLabs := lab.Group("/mark")
					{
						marksLabs.GET("", h.LecturerApi.GetMarkLabForStudent)
						marksLabs.PUT("", h.LecturerApi.ChangeMarkLabForStudent)
					}
				}

			}
		}
	}
	return router
}
