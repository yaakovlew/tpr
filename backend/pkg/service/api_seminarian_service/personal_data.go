package api_seminarian_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type SeminarianPersonalDataService struct {
	repo repository.SeminarianPersonalData
}

func NewSeminarianPersonalDataService(repo repository.SeminarianPersonalData) *SeminarianPersonalDataService {
	return &SeminarianPersonalDataService{repo: repo}
}

func (s *SeminarianPersonalDataService) GetPersonalData(id int) (model.User, error) {
	return s.repo.GetPersonalData(id)
}

func (s *SeminarianPersonalDataService) UpdateName(id int, name string) error {
	return s.repo.UpdateName(id, name)
}

func (s *SeminarianPersonalDataService) UpdateSurname(id int, surname string) error {
	return s.repo.UpdateSurname(id, surname)
}
