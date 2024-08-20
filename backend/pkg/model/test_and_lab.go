package model

type Test struct {
	Id              int    `json:"test_id" db:"id"`
	Name            string `json:"name" db:"name"`
	TaskDescription string `json:"task_description" db:"task_description"`
	MinutesDuration int    `json:"minutes_duration" db:"minutes_duration"`
	DefaultMark     int    `json:"default_mark" db:"default_mark"`
}

type TestWithClosedDate struct {
	Id              int    `json:"test_id" db:"id"`
	Name            string `json:"name" db:"name"`
	TaskDescription string `json:"task_description" db:"task_description"`
	MinutesDuration int    `json:"minutes_duration" db:"minutes_duration"`
	DefaultMark     int    `json:"default_mark" db:"default_mark"`
	ClosedDate      int    `json:"closed_date" db:"closed_date"`
}

type LabWithClosedDate struct {
	Id              int    `json:"laboratory_id" db:"id"`
	Name            string `json:"name" db:"name"`
	TaskDescription string `json:"task_description" db:"task_description"`
	DefaultMark     int    `json:"default_mark" db:"default_mark"`
	ClosedDate      int    `json:"closed_date" db:"closed_date"`
}

type TestAdd struct {
	Name              string `json:"name" binding:"required"`
	TaskDescription   string `json:"task_description" binding:"required"`
	NameEn            string `json:"name_en" binding:"required"`
	TaskDescriptionEn string `json:"task_description_en" binding:"required"`
	MinutesDuration   int    `json:"minutes_duration" binding:"required"`
	DefaultMark       int    `json:"default_mark" binding:"required"`
}

type TestResponse struct {
	Ru []Test `json:"ru"`
	En []Test `json:"en"`
}

type TestResponseWithClosedDate struct {
	Ru []TestWithClosedDate `json:"ru"`
	En []TestWithClosedDate `json:"en"`
}

type LabsResponseWithClosedDate struct {
	Ru []LabWithClosedDate `json:"ru"`
	En []LabWithClosedDate `json:"en"`
}

type LaboratoryWork struct {
	Id              int    `json:"laboratory_id" db:"id"`
	Name            string `json:"name" db:"name"`
	TaskDescription string `json:"task_description" db:"task_description"`
	Link            string `json:"link" db:"link"`
	DefaultMark     int    `json:"default_mark" db:"default_mark"`
}

type LaboratoryWorkWithExternal struct {
	Id              int    `json:"laboratory_id" db:"laboratory_id"`
	ExternalLabId   int    `json:"external_laboratory_id" db:"external_laboratory_id"`
	Name            string `json:"name" db:"name"`
	TaskDescription string `json:"task_description" db:"task_description"`
	Link            string `json:"link" db:"link"`
	DefaultMark     int    `json:"default_mark" db:"default_mark"`
}

type CommonLaboratoryWork struct {
	Id              int    `json:"external_laboratory_id" db:"id"`
	Name            string `json:"name" db:"name"`
	TaskDescription string `json:"task_description" db:"task_description"`
	Linc            string `json:"link" db:"link"`
}

type LaboratoryWorkInput struct {
	Id                int    `json:"laboratory_id" db:"id"`
	Name              string `json:"name" db:"name"`
	NameEn            string `json:"name_en" db:"name_en"`
	TaskDescription   string `json:"task_description" db:"task_description"`
	TaskDescriptionEn string `json:"task_description_en" db:"task_description_en"`
	MinutesDuration   int    `json:"minutes_duration" db:"minutes_duration"`
	Linc              string `json:"linc" db:"linc"`
	Token             string `json:"token" db:"token"`
	DefaultMark       int    `json:"default_mark" db:"default_mark"`
	DayFine           int    `json:"day_fine" db:"day_fine"`
}

type LaboratoryWorkInputWithoutId struct {
	Name              string `json:"name" db:"name"`
	NameEn            string `json:"name_en" db:"name_en"`
	TaskDescription   string `json:"task_description" db:"task_description"`
	TaskDescriptionEn string `json:"task_description_en" db:"task_description_en"`
	Linc              string `json:"linc" db:"linc"`
	Token             string `json:"token" db:"token"`
}

type LaboratoryWorkResponse struct {
	Id                int    `json:"laboratory_id" db:"id"`
	Name              string `json:"name" db:"name"`
	NameEn            string `json:"name_en" db:"name_en"`
	TaskDescription   string `json:"task_description" db:"task_description"`
	TaskDescriptionEn string `json:"task_description_en" db:"task_description_en"`
	Link              string `json:"link" db:"link"`
}

type QuestionPercentage struct {
	QuestionId int     `json:"question_id"`
	Percentage float64 `json:"percentage"`
}

type LabsResponse struct {
	Ru []LaboratoryWork `json:"ru"`
	En []LaboratoryWork `json:"en"`
}

type ChangeLabNameInput struct {
	LaboratoryId int    `json:"laboratory_id" binding:"required"`
	Name         string `json:"name" binding:"required"`
}

type ChangeTestNameInput struct {
	TestId int    `json:"test_id" binding:"required"`
	Name   string `json:"name" binding:"required"`
}

type ChangeTestDescriptionInput struct {
	TestId      int    `json:"test_id" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type ChangeLabDescriptionInput struct {
	LaboratoryId int    `json:"laboratory_id" binding:"required"`
	Description  string `json:"description" binding:"required"`
}

type ChangeLabDefaultMarkInput struct {
	LaboratoryId int `json:"laboratory_id" binding:"required"`
	Mark         int `json:"default_mark" binding:"required"`
}

type ChangeTestDurationInput struct {
	TestId   int `json:"test_id" binding:"required"`
	Duration int `json:"minutes_duration" binding:"required"`
}

type ChangeLabDurationInput struct {
	LaboratoryId int `json:"laboratory_id" binding:"required"`
	Duration     int `json:"minutes_duration" binding:"required"`
}

type ChangeLabFineInput struct {
	LaboratoryId int `json:"laboratory_id" binding:"required"`
	Fine         int `json:"day_fine" binding:"required"`
}

type ChangeLabLincInput struct {
	LaboratoryId int    `json:"laboratory_id" binding:"required"`
	Linc         string `json:"linc" binding:"required"`
}

type ChangeLabTokenInput struct {
	LaboratoryId int    `json:"laboratory_id" binding:"required"`
	Token        string `json:"token" binding:"required"`
}

type Theme struct {
	ThemeId int    `json:"theme_id" db:"id"`
	Name    string `json:"name" db:"name"`
	Weight  int    `json:"weight" db:"weight"`
}

type ThemeOutput struct {
	ThemeId int    `json:"theme_id" db:"id"`
	Name    string `json:"name" db:"name"`
	Weight  int    `json:"weight" db:"weight"`
	Count   int    `json:"count" db:"count"`
}

type Answer struct {
	AnswerId int    `json:"answer_id" db:"id"`
	Name     string `json:"answer" db:"name"`
	IsRight  *bool  `json:"is_right" db:"is_right"`
}

type Question struct {
	QuestionId int    `json:"question_id" db:"id"`
	Name       string `json:"question" db:"name"`
	IsVariable int    `json:"is_variable" db:"is_variable"`
}

type QuestionWithAmountAnswers struct {
	QuestionId    int    `json:"question_id" db:"id"`
	Name          string `json:"question" db:"name"`
	IsVariable    int    `json:"is_variable" db:"is_variable"`
	AmountAnswers int    `json:"amount_answers" db:"amount_answers"`
}

type TestId struct {
	Id int `json:"test_id" binding:"required"`
}

type AddedThemeForTestInput struct {
	TestId  int `json:"test_id" binding:"required"`
	ThemeId int `json:"theme_id" binding:"required"`
	Count   int `json:"count" binding:"required"`
}

type ChangeThemeNameInput struct {
	Id   int    `json:"theme_id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type ThemesResponse struct {
	Themes []Theme `json:"themes"`
}

type ThemesResponseOutput struct {
	Themes []ThemeOutput `json:"themes"`
}

type AddQuestion struct {
	ThemeId    int `json:"theme_id" binding:"required"`
	QuestionId int `json:"question_id" binding:"required"`
}

type QuestionsResponse struct {
	Ru []Question `json:"ru"`
	En []Question `json:"en"`
}

type QuestionsResponseWithAnswersAmount struct {
	Ru []QuestionWithAmountAnswers `json:"ru"`
	En []QuestionWithAmountAnswers `json:"en"`
}

type Questions struct {
	Questions []Question `json:"questions"`
}

type ChangeQuestionNameInput struct {
	Id   int    `json:"question_id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type ChangeThemeWeightInput struct {
	Id     int `json:"theme_id" binding:"required"`
	Weight int `json:"weight" binding:"required"`
}

type AddAnswerInput struct {
	Id      int    `json:"question_id" binding:"required"`
	Name    string `json:"name" binding:"required"`
	NameEn  string `json:"name_en" binding:"required"`
	IsRight *bool  `json:"is_right" binding:"required"`
}

type ChangeAnswerNameInput struct {
	Id   int    `json:"answer_id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type ChangeAnswerIsRightInput struct {
	Id      int   `json:"answer_id" binding:"required"`
	IsRight *bool `json:"is_right" binding:"required"`
}

type AnswersResponse struct {
	Ru []Answer `json:"ru"`
	En []Answer `json:"en"`
}

type AnswerDuringTest struct {
	AnswerId   int    `json:"answer_id" db:"id"`
	Name       string `json:"name" db:"name"`
	QuestionId int    `json:"question_id" db:"question_id"`
}

type QuestionsWithAnswers struct {
	Question         Question           `json:"questions"`
	ThemeId          int                `json:"theme_id"`
	Answers          []AnswerDuringTest `json:"answers"`
	AnswersToCompare []AnswerDuringTest `json:"comparing_answers"`
}

type QuestionsWithAnswResp struct {
	QuestionId int    `json:"question_id" db:"question_id"`
	Name       string `json:"question" db:"name"`
	IsVariable int    `json:"is_variable" db:"is_variable"`
	AnswerId   int    `json:"answer_id" db:"id"`
	AnsName    string `json:"name" db:"ans_name"`
	ThemeId    int    `json:"theme_id"`
}

type QuestionInput struct {
	Name       string `json:"question" binding:"required"`
	NameEn     string `json:"question_en"`
	IsVariable int    `json:"is_variable"`
}

type ThemeInput struct {
	Name   string `json:"name" binding:"required"`
	Weight int    `json:"weight" binding:"required"`
}

type TestThemeInput struct {
	TestId  int `json:"test_id" binding:"required"`
	ThemeId int `json:"theme_id" binding:"required"`
}

type UserTestAccess struct {
	UserId     int   `json:"user_id" db:"user_id"`
	TestId     int   `json:"test_id" db:"test_id"`
	ClosedDate int64 `json:"closed_date" db:"closed_date"`
	IsDone     bool  `json:"is_done" db:"is_done"`
}

type UserLabAccess struct {
	UserId     int   `json:"user_id" db:"user_id"`
	LabId      int   `json:"laboratory_id" db:"laboratory_id"`
	ClosedDate int64 `json:"closed_date" db:"closed_date"`
	IsDone     bool  `json:"is_done" db:"is_done"`
}

type CheckTimeForTest struct {
	TestId          int `json:"test_id" db:"id"`
	MinutesDuration int `json:"minutes_duration" db:"minutes_duration"`
}

type CheckTimeForLab struct {
	LabId           int `json:"laboratory_id" db:"id"`
	MinutesDuration int `json:"minutes_duration" db:"minutes_duration"`
}

type ThemeForChoose struct {
	ThemeId int `json:"theme_id" db:"theme_id"`
	Count   int `json:"count" db:"count"`
}

type QuestionsForTestResponse struct {
	Ru []QuestionsWithAnswers `json:"ru"`
	En []QuestionsWithAnswers `json:"en"`
}

type ThemePoint struct {
	Count  int `json:"count" db:"count"`
	Weight int `json:"weight" db:"weight"`
}

type QuestionInputToCheck struct {
	Id         int    `json:"question_id" binding:"required"`
	Name       string `json:"question" binding:"required"`
	NameEn     string `json:"question_en" binding:"required"`
	IsVariable int    `json:"is_variable" binding:"required"`
}

type QuestionAndAnswerResponse struct {
	QuestionId int      `json:"question_id" binding:"required"`
	ThemeId    int      `json:"theme_id" binding:"required"`
	Answer     []string `json:"answer" binding:"required"`
}

type QuestionAndAnswersParser struct {
	Answers []QuestionAndAnswerResponse `json:"answers"`
}

type TestMarkResponse struct {
	Mark               int                  `json:"mark"`
	QuestionPercentage []QuestionPercentage `json:"question_percentage"`
}

type GetQuestionWithRightAnswer struct {
	Answers []AllAnswersName `json:"answers" db:"name"`
}

type AllAnswersName struct {
	Name   string `json:"name" db:"name"`
	NameEn string `json:"name_en" db:"name_en"`
}

type QuestionWithTheme struct {
	QuestionId int `json:"question_id" db:"question_id"`
	ThemeId    int `json:"theme_id" db:"theme_id"`
}

type OpenedTest struct {
	UserId     int   `json:"user_id" db:"user_id"`
	TestId     int   `json:"test_id" db:"test_id"`
	ClosedDate int64 `json:"closed_date" db:"closed_date"`
	IsDone     bool  `json:"is_done" db:"is_done"`
}

type OpenedLab struct {
	UserId     int   `json:"user_id" db:"user_id"`
	LabId      int   `json:"laboratory_id" db:"laboratory_id"`
	ClosedDate int64 `json:"closed_date" db:"closed_date"`
	IsDone     bool  `json:"is_done" db:"is_done"`
}

type GetOpenedTestInput struct {
	UserId int `json:"user_id" binding:"required"`
	TestId int `json:"test_id" binding:"required"`
}

type CloseTestInput struct {
	UserId int `json:"user_id" db:"user_id"`
	TestId int `json:"test_id" db:"test_id"`
}

type CloseLabInput struct {
	UserId int `json:"user_id" db:"user_id"`
	LabId  int `json:"laboratory_id" db:"laboratory_id"`
}

type GetTestMark struct {
	UserId int `json:"user_id" binding:"required"`
	TestId int `json:"test_id" binding:"required"`
}

type ChangeTestMark struct {
	UserId int `json:"user_id" binding:"required"`
	TestId int `json:"test_id" binding:"required"`
	Mark   int `json:"mark" binding:"required"`
}

type LabPercentage struct {
	UserId     int `json:"user_id" binding:"required"`
	LabId      int `json:"laboratory_id" binding:"required"`
	Percentage int `json:"percentage"`
}

type ChangeLabMark struct {
	UserId int `json:"user_id" binding:"required"`
	LabId  int `json:"laboratory_id" binding:"required"`
	Mark   int `json:"mark" binding:"required"`
}

type ThemeInputForMultiLanguage struct {
	ThemeId         int   `json:"theme_id" binding:"required"`
	IsMultiLanguage *bool `json:"is_multi_language" binding:"required"`
}

type QuestionResponse struct {
	Questions []Question `json:"questions"`
}
