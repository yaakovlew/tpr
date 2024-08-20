package model

type Group struct {
	Id        int    `json:"group_id" db:"id"`
	Name      string `json:"name" db:"name"`
	IsArchive bool   `json:"is_archive" db:"is_archive"`
}

type GroupResponse struct {
	Groups []Group `json:"groups"`
}

type GroupAndDisciplineInput struct {
	GroupId      int `json:"group_id" binding:"required"`
	DisciplineId int `json:"discipline_id" binding:"required"`
}

type GroupChangeNameInput struct {
	GroupId int    `json:"group_id" binding:"required"`
	Name    string `json:"name" binding:"required"`
}

type GroupId struct {
	Id int `json:"group_id" db:"group_id"`
}
