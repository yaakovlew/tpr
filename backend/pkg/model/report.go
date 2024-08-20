package model

type MarksForTest struct {
	Name string
}

type TestsReport struct {
	TestId int `json:"test_id" db:"test_id"`
}

type LabsReport struct {
	LabId int `json:"laboratory_id" db:"laboratory_id"`
}

type StudentReport struct {
	StudentId      int    `json:"student_id" db:"student_id"`
	StudentName    string `json:"student_name" db:"student_name"`
	StudentSurname string `json:"student_surname" db:"student_surname"`
}

type GroupReport struct {
	GroupId      int   `json:"group_id" binding:"required"`
	DisciplineId int   `json:"discipline_id" binding:"required"`
	IsExam       *bool `json:"is_exam" binding:"required"`
}

type MarkReport struct {
	Mark int `json:"mark" db:"mark"`
}

type CountSeminar struct {
	Amount int `json:"amount" db:"amount"`
}

type CountSeminarVisitingStudent struct {
	Amount int `json:"amount" db:"amount"`
}
