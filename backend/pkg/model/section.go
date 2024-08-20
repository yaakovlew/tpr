package model

type Section struct {
	SectionId int    `json:"section_id" db:"id"`
	Name      string `json:"name" db:"name"`
}

type AddSectionInput struct {
	Name         string `json:"name" binding:"required"`
	NameEn       string `json:"name_en" binding:"required"`
	DisciplineId int    `json:"discipline_id" binding:"required"`
}

type DeleteSectionInput struct {
	Id int `json:"section_id" binding:"required"`
}

type ChangeSectionNameInput struct {
	Id   int    `json:"section_id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type GetDisciplineSectionsInput struct {
	DisciplineId int `json:"discipline_id" binding:"required"`
}

type SectionsResponse struct {
	Ru []Section `json:"ru"`
	En []Section `json:"en"`
}

type TestToSectionInput struct {
	SectionId int `json:"section_id" binding:"required"`
	TestId    int `json:"test_id" binding:"required"`
}

type LabToSectionInput struct {
	SectionId     int `json:"section_id" binding:"required"`
	ExternalLabId int `json:"external_lab_id" binding:"required"`
	DefaultMark   int `json:"default_mark" binding:"required"`
}
