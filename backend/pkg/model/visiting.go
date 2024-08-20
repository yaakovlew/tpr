package model

type SeminarVisiting struct {
	SeminarID int    `json:"id" db:"id"`
	Name      string `json:"name"`
	ISAbsent  bool   `json:"is_absent" db:"is_absent"`
}

type LessonVisiting struct {
	LessonID int    `json:"id" db:"id"`
	Name     string `json:"name"`
	ISAbsent bool   `json:"is_absent" db:"is_absent"`
}

type SeminarsVisitingResponse struct {
	Seminars []SeminarVisiting `json:"seminars"`
}

type LessonVisitingResponse struct {
	Lessons []LessonVisiting `json:"lessons"`
}

type SeminarVisitingStudent struct {
	Id       int    `json:"student_id" db:"student_id"`
	Name     string `json:"student_name" db:"student_name"`
	Surname  string `json:"student_surname" db:"student_surname"`
	ISAbsent bool   `json:"is_absent" db:"is_absent"`
}

type SeminarVisitingStudentResponse struct {
	Students []SeminarVisitingStudent `json:"students"`
}

type LessonVisitingStudent struct {
	Id       int    `json:"student_id" db:"student_id"`
	Name     string `json:"student_name" db:"student_name"`
	Surname  string `json:"student_surname" db:"student_surname"`
	ISAbsent bool   `json:"is_absent" db:"is_absent"`
}

type LessonVisitingStudentResponse struct {
	Students []LessonVisitingStudent `json:"students"`
}

type LessonVisitingInput struct {
	LessonId int `json:"lesson_id" binding:"required"`
	GroupId  int `json:"group_id" binding:"required"`
}

type SeminarVisitingInput struct {
	SeminarId int `json:"seminar_id" binding:"required"`
}

type AddSeminarVisitingInput struct {
	SeminarId int   `json:"seminar_id" binding:"required"`
	UserId    int   `json:"user_id" binding:"required"`
	IsAbsent  *bool `json:"is_absent" binding:"required"`
}

type AddLessonVisitingInput struct {
	LessonId int   `json:"lesson_id" binding:"required"`
	UserId   int   `json:"user_id" binding:"required"`
	IsAbsent *bool `json:"is_absent" binding:"required"`
}

type LessonDate struct {
	LessonId    int    `json:"lesson_id" db:"lesson_id"`
	Date        int    `json:"date" db:"date"`
	Description string `json:"description" db:"description"`
	LessonName  string `json:"lesson_name" db:"name"`
	GroupId     int    `json:"group_id" db:"group_id"`
}

type LessonDateToChange struct {
	LessonId    int    `json:"lesson_id" binding:"required"`
	Date        int    `json:"date"`
	Description string `json:"description"`
	GroupId     int    `json:"group_id" binding:"required"`
}

type SeminarDateInput struct {
	SeminarId int `json:"seminar_id" binding:"required"`
	Date      int `json:"date" binding:"required"`
}

type LessonDateInput struct {
	LessonId int `json:"lesson_id" binding:"required"`
	GroupId  int `json:"group_id" binding:"required"`
}

type LessonDateToResponse struct {
	LessonId   int    `json:"lesson_id" db:"lesson_id"`
	LessonName string `json:"lesson_name" db:"name"`
	Date       int    `json:"date" db:"date"`
}

type SeminarDate struct {
	SeminarId  int    `json:"seminar_id" db:"id"`
	LessonName string `json:"seminar_name" db:"name"`
	Date       int    `json:"date" db:"date"`
}

type LessonsTableResponse struct {
	Lessons []LessonDate `json:"lessons"`
}

type SeminarsTableResponse struct {
	Seminars []SeminarDate `json:"seminars"`
}
