package api_lecturer_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type LecturerSeminarianService struct {
	repo repository.LecturerSeminarian
}

func NewLecturerSeminarianService(repo repository.LecturerSeminarian) *LecturerSeminarianService {
	return &LecturerSeminarianService{repo: repo}
}

func (s *LecturerSeminarianService) GetSeminarianFromGroupsAndDiscipline(groupId, disciplineId int) ([]model.Seminarian, error) {
	return s.repo.GetSeminarianFromGroupsAndDiscipline(groupId, disciplineId)
}

func (s *LecturerSeminarianService) GetAllSeminarians() ([]model.Seminarian, error) {
	return s.repo.GetAllSeminarians()
}

func (s *LecturerSeminarianService) AddSeminarian(seminarianId, groupId, disciplineId int) error {
	return s.repo.AddSeminarian(seminarianId, groupId, disciplineId)
}

func (s *LecturerSeminarianService) DeleteSeminarianFromGroupAndDiscipline(seminarianId, groupId, disciplineId int) error {
	return s.repo.DeleteSeminarianFromGroupAndDiscipline(seminarianId, groupId, disciplineId)
}
