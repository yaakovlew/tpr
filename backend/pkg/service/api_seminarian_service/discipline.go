package api_seminarian_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type SeminarianDisciplineService struct {
	repo repository.SeminarianDiscipline
}

func NewSeminarianDisciplineService(repo repository.SeminarianDiscipline) *SeminarianDisciplineService {
	return &SeminarianDisciplineService{repo: repo}
}

func (s *SeminarianDisciplineService) GetOwnDiscipline(userId int) ([]model.Discipline, []model.Discipline, error) {
	return s.repo.GetOwnDiscipline(userId)
}

func (s *SeminarianDisciplineService) GetDisciplineSections(seminarianId, disciplineId int) ([]model.Section, []model.Section, error) {
	if err := s.repo.CheckAccessForDiscipline(seminarianId, disciplineId); err != nil {
		return nil, nil, err
	}
	return s.repo.GetDisciplineSections(disciplineId)
}

func (s *SeminarianDisciplineService) GetAllInfoAboutDiscipline(seminarianId, disciplineId int) (model.DisciplineInfoDoubleLang, model.DisciplineInfoDoubleLang, error) {
	if err := s.repo.CheckAccessForDiscipline(seminarianId, disciplineId); err != nil {
		return model.DisciplineInfoDoubleLang{}, model.DisciplineInfoDoubleLang{}, err
	}

	return s.repo.GetAllInfoAboutDiscipline(disciplineId)
}
