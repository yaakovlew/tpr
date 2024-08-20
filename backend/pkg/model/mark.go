package model

type TestsMarks struct {
	TestName string `json:"test_name" db:"name"`
	Mark     int    `json:"mark" db:"mark"`
}

type LaboratoryMarks struct {
	LaboratoryName string `json:"laboratory_name" db:"name"`
	Mark           int    `json:"mark" db:"mark"`
}

type InputLessonMark struct {
	DisciplineId int `json:"discipline_id" binding:"required"`
	Mark         int `json:"lesson_mark"`
}

type InputSeminarMark struct {
	DisciplineId int `json:"discipline_id" binding:"required"`
	Mark         int `json:"seminar_mark"`
}

type InputExamMark struct {
	DisciplineId int `json:"discipline_id" binding:"required"`
	Mark         int `json:"exam_mark"`
}

type GetAllLaboratoryMarksResponse struct {
	LaboratoryWorks []LaboratoryMarks `json:"laboratory_works"`
}

type GetAllTestsMarksResponse struct {
	Tests []TestsMarks `json:"test"`
}

type ChangeTestMarkInput struct {
	UserId int `json:"user_id" binding:"required"`
	TestId int `json:"test_id" binding:"required"`
	Mark   int `json:"mark" binding:"required"`
}

type ChangeLaboratoryMarkInput struct {
	UserId       int `json:"user_id" binding:"required"`
	LaboratoryId int `json:"laboratory_id" binding:"required"`
	Mark         int `json:"mark" binding:"required"`
}

type GroupTestMarks struct {
	UserId      int    `json:"user_id" db:"user_id"`
	UserName    string `json:"user_name" db:"user_name"`
	UserSurname string `json:"user_surname" db:"user_surname"`
	Mark        int    `json:"mark" db:"mark"`
}

type GroupLaboratoryMarks struct {
	UserId      int    `json:"user_id" db:"user_id"`
	UserName    string `json:"user_name" db:"user_name"`
	UserSurname string `json:"user_surname" db:"user_surname"`
	Mark        int    `json:"mark" db:"mark"`
}

type GetTestMarksInput struct {
	GroupId int `json:"group_id" binding:"required"`
	TestId  int `json:"test_id" binding:"required"`
}

type GetLaboratoryMarksInput struct {
	GroupId      int `json:"group_id" binding:"required"`
	LaboratoryId int `json:"laboratory_id" binding:"required"`
}

type GroupTestMarksResponse struct {
	Marks []GroupTestMarks `json:"marks"`
}

type GroupLaboratoryMarksResponse struct {
	Marks []GroupLaboratoryMarks `json:"marks"`
}

type ExamMark struct {
	UserId      int    `json:"user_id" db:"user_id"`
	UserName    string `json:"user_name" db:"user_name"`
	UserSurname string `json:"user_surname" db:"user_surname"`
	Mark        int    `json:"mark" db:"mark"`
}

type ExamMarkResponse struct {
	Marks []ExamMark `json:"marks"`
}

type ExamMarkInput struct {
	UserId       int `json:"user_id" binding:"required"`
	DisciplineId int `json:"discipline_id" binding:"required"`
	Mark         int `json:"mark" binding:"required"`
}

type MarkResponse struct {
	Mark int `json:"mark"`
}

type TestResult struct {
	UserId int `json:"user_id" db:"user_id"`
	TestId int `json:"test_id" db:"test_id"`
	Mark   int `json:"mark" db:"mark"`
}

type LabResult struct {
	UserId int `json:"user_id" db:"user_id"`
	LabId  int `json:"laboratory_id" db:"laboratory_id"`
	Mark   int `json:"mark" db:"mark"`
}

type LabCredential struct {
	Link  string `json:"link" db:"link"`
	Token string `json:"token" db:"token"`
}
