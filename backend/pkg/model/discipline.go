package model

type Discipline struct {
	Id   int    `json:"discipline_id" db:"id"`
	Name string `json:"name" db:"name"`
}

type DisciplineId struct {
	Id int `json:"discipline_id" binding:"required"`
}

type AddNewDiscipline struct {
	Name                string `json:"discipline_name" binding:"required"`
	NameEn              string `json:"discipline_name_en" binding:"required"`
	SeminarVisitingMark int    `json:"seminar_visiting_mark"`
	LessonVisitingMark  int    `json:"lesson_visiting_mark"`
	ExamMark            int    `json:"exam_mark"`
}

type DisciplineInfo struct {
	Id                  int    `json:"id" db:"id"`
	Name                string `json:"discipline_name" db:"name"`
	NameEn              string `json:"discipline_name_en" db:"name_en"`
	SeminarVisitingMark *int   `json:"seminar_visiting_mark" db:"seminar_visiting_mark"`
	LessonVisitingMark  *int   `json:"lesson_visiting_mark" db:"lesson_visiting_mark"`
	ExamMark            int    `json:"exam_mark" db:"exam_mark"`
}

type DisciplineInfoDoubleLang struct {
	Id                  int    `json:"id" db:"id"`
	Name                string `json:"discipline_name" db:"name"`
	SeminarVisitingMark *int   `json:"seminar_visiting_mark" db:"seminar_visiting_mark"`
	LessonVisitingMark  *int   `json:"lesson_visiting_mark" db:"lesson_visiting_mark"`
	ExamMark            int    `json:"exam_mark" db:"exam_mark"`
}

type InputGroupToDiscipline struct {
	GroupId      int `json:"group_id" binding:"required"`
	DisciplineId int `json:"discipline_id" binding:"required"`
}

type DisciplineForGroup struct {
	Id   int    `json:"discipline_id" db:"id"`
	Name string `json:"name" db:"name"`
}

type DisciplinesResponseRu struct {
}

type DisciplinesResponse struct {
	Ru []Discipline `json:"ru"`
	En []Discipline `json:"en"`
}

type DisciplineInput struct {
	Id int `json:"discipline_id" binding:"required"`
}

type DisciplineNewNameInput struct {
	Id   int    `json:"discipline_id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
