package api_lecturer_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type LecturerPersonalDataServices struct {
	repo repository.LecturerPersonalData
}

func NewLecturerPersonalDataServices(repo repository.LecturerPersonalData) *LecturerPersonalDataServices {
	return &LecturerPersonalDataServices{repo: repo}
}

func (s *LecturerPersonalDataServices) GetPersonalData(id int) (model.User, error) {
	return s.repo.GetPersonalData(id)
}

func (s *LecturerPersonalDataServices) UpdateName(id int, name string) error {
	return s.repo.UpdateName(id, name)
}

func (s *LecturerPersonalDataServices) UpdateSurname(id int, surname string) error {
	return s.repo.UpdateSurname(id, surname)
}
