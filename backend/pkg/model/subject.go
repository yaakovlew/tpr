package model

type Lesson struct {
	Id   int    `json:"lesson_id" db:"id"`
	Name string `json:"name" db:"name"`
}

type LessonsResponse struct {
	Lessons []Lesson `json:"lessons"`
}

type AddLessonInput struct {
	DisciplineId int    `json:"discipline_id" binding:"required"`
	Name         string `json:"name" binding:"required"`
}

type DeleteLessonInput struct {
	Id int `json:"lesson_id" binding:"required"`
}

type ChangeLessonInput struct {
	Id   int    `json:"lesson_id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type Seminar struct {
	Id   int    `json:"seminar_id" db:"id"`
	Name string `json:"name" db:"name"`
	Date int    `json:"date" db:"date"`
}

type SeminarResponse struct {
	Seminars []Seminar `json:"seminars"`
}

type GetAllSeminarsInput struct {
	DisciplineId int `json:"discipline_id" binding:"required"`
	GroupId      int `json:"group_id" binding:"required"`
}

type ChangeSeminarInput struct {
	SeminarId int    `json:"seminar_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
}

type DeleteSeminarInput struct {
	SeminarId int `json:"seminar_id" binding:"required"`
}

type AddSeminarInput struct {
	DisciplineId int    `json:"discipline_id" binding:"required"`
	GroupId      int    `json:"group_id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Date         int    `json:"date" binding:"required"`
}

type AllInfoSeminar struct {
	Id           int    `json:"seminar_id" db:"id"`
	Name         string `json:"name" db:"name"`
	DisciplineId int    `json:"discipline_id" db:"discipline_id"`
	GroupId      int    `json:"group_id" db:"group_id"`
	Date         int    `json:"date" db:"date"`
}

type AllInfoLesson struct {
	Id           int    `json:"lesson_id" db:"id"`
	Name         string `json:"name" db:"name"`
	DisciplineId int    `json:"discipline_id" db:"discipline_id"`
}
