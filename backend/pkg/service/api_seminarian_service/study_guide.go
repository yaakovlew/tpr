package api_seminarian_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type SeminarianStudyGuideService struct {
	repo repository.SeminarianStudyGuide
}

func NewSeminarianStudyGuideService(repo repository.SeminarianStudyGuide) *SeminarianStudyGuideService {
	return &SeminarianStudyGuideService{repo: repo}
}

func (s *SeminarianStudyGuideService) GetDigitalDiscipline(seminarianId, disciplineId int) ([]model.DigitalDisciplineWithInfo, []model.DigitalDisciplineWithInfo, error) {
	if err := s.repo.CheckAccessForDiscipline(seminarianId, disciplineId); err != nil {
		return nil, nil, err
	}
	return s.repo.GetDigitalDiscipline(disciplineId)
}

func (s *SeminarianStudyGuideService) GetFilesIdFromDigital(seminarianId, digitalId int) ([]model.FileId, error) {
	if err := s.repo.CheckAccessForLesson(seminarianId, digitalId); err != nil {
		return nil, err
	}
	return s.repo.GetFilesIdFromDigital(digitalId)
}

func (s *SeminarianStudyGuideService) GetFile(seminarianId, fileId int) (string, error) {
	if err := s.repo.CheckAccessForFile(seminarianId, fileId); err != nil {
		return "", err
	}
	return s.repo.GetFile(fileId)
}
