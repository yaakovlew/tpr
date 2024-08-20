package model

type StudyGuideHeader struct {
	Id          int     `json:"study_guide_id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
}

type AddStudyGuideHeaderInput struct {
	Name          string `json:"name" binding:"required"`
	NameEn        string `json:"name_en" binding:"required"`
	Description   string `json:"description"`
	DescriptionEn string `json:"description_en"`
}

type DeleteStudyGuideHeaderInput struct {
	Id int `json:"digital_guide_id" binding:"required"`
}

type ChangeNameStudyGuideHeaderInput struct {
	Id   int    `json:"digital_guide_id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type ChangeDescriptionStudyGuideHeaderInput struct {
	Id          int    `json:"digital_guide_id" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type GetStudyGuideHeaderInput struct {
	DisciplineId int `json:"discipline_id" binding:"required"`
}

type StudyGuideHeaderResponse struct {
	Ru []StudyGuideHeader `json:"ru"`
	En []StudyGuideHeader `json:"en"`
}

type GuideIdInput struct {
	StudyGuideId int `json:"study_guide_id" binding:"required"`
}

type DeleteGuideInput struct {
	StudyGuideId int    `json:"study_guide_id" binding:"required"`
	FileName     string `json:"file_name" binding:"required"`
}

type PathToFile struct {
	Path string `json:"path" binding:"required"`
}

type FilesPathResponse struct {
	Files []string `json:"files"`
}

type DigitalDiscipline struct {
	DigitalMaterialId int `json:"digital_material_id" db:"digital_material_id" binding:"required"`
	DisciplineId      int `json:"discipline_id" db:"discipline_id" binding:"required"`
}

type DigitalDisciplineWithInfo struct {
	Name              string `json:"name" db:"name"`
	Description       string `json:"description" db:"description"`
	DigitalMaterialId int    `json:"digital_material_id" db:"digital_material_id"`
}

type DigitalDisciplinesResponse struct {
	DigitalGuides []DigitalDiscipline `json:"digital_guides"`
}

type DigitalDisciplinesInfoResponse struct {
	Ru []DigitalDisciplineWithInfo `json:"ru"`
	En []DigitalDisciplineWithInfo `json:"en"`
}

type FileId struct {
	Id   int    `json:"file_id" db:"id" binding:"required"`
	Name string `json:"file_name" db:"name" binding:"required"`
}

type FilesResponse struct {
	Files []FileId `json:"files"`
}
