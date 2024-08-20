package api_student_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type StudentStudyGuideService struct {
	repo repository.StudentStudyGuide
}

func NewStudentStudyGuideService(repo repository.StudentStudyGuide) *StudentStudyGuideService {
	return &StudentStudyGuideService{repo: repo}
}

func (s *StudentStudyGuideService) GetDigitalDiscipline(studentId, disciplineId int) ([]model.DigitalDisciplineWithInfo, []model.DigitalDisciplineWithInfo, error) {
	if err := s.repo.CheckAccessForDiscipline(studentId, disciplineId); err != nil {
		return nil, nil, err
	}
	return s.repo.GetDigitalDiscipline(disciplineId)
}

func (s *StudentStudyGuideService) GetFilesIdFromDigital(studentId, digitalId int) ([]model.FileId, error) {
	if err := s.repo.CheckAccessForLesson(studentId, digitalId); err != nil {
		return nil, err
	}
	return s.repo.GetFilesIdFromDigital(digitalId)
}

func (s *StudentStudyGuideService) GetFile(studentId, fileId int) (string, error) {
	if err := s.repo.CheckAccessForFile(studentId, fileId); err != nil {
		return "", err
	}
	return s.repo.GetFile(fileId)
}
