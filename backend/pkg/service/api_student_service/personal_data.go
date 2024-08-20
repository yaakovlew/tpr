package api_student_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type StudentPersonalDataServices struct {
	repo repository.StudentPersonalData
}

func NewStudentPersonalDataServices(repo repository.StudentPersonalData) *StudentPersonalDataServices {
	return &StudentPersonalDataServices{repo: repo}
}

func (s *StudentPersonalDataServices) GetPersonalData(id int) (model.User, error) {
	return s.repo.GetPersonalData(id)
}

func (s *StudentPersonalDataServices) UpdateName(id int, name string) error {
	return s.repo.UpdateName(id, name)
}

func (s *StudentPersonalDataServices) UpdateSurname(id int, surname string) error {
	return s.repo.UpdateSurname(id, surname)
}
