package api_seminarian_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"errors"
)

type SeminarianMarksService struct {
	repo repository.SeminarianMark
}

func NewSeminarianMarksService(repo repository.SeminarianMark) *SeminarianMarksService {
	return &SeminarianMarksService{repo: repo}
}

func (s *SeminarianMarksService) GetTestMarksFromGroup(seminarianId, groupId, testId int) ([]model.GroupTestMarks, error) {
	if err := s.repo.CheckAccessForTest(seminarianId, groupId, testId); err != nil {
		return nil, err
	}
	return s.repo.GetTestMarksFromGroup(groupId, testId)
}

func (s *SeminarianMarksService) GetLaboratoryMarksFromGroup(seminarianId, groupId, laboratoryId int) ([]model.GroupLaboratoryMarks, error) {
	if err := s.repo.CheckAccessForLaboratory(seminarianId, groupId, laboratoryId); err != nil {
		return nil, err
	}
	return s.repo.GetLaboratoryMarksFromGroup(groupId, laboratoryId)
}

func (s *SeminarianMarksService) GiveExamMark(seminarianId, userId, disciplineId, mark int) error {
	maxMark := s.repo.MaxExamMark(disciplineId)
	if maxMark < mark {
		err := errors.New("big mark")
		return err
	}
	if err := s.repo.CheckAccessForStudent(seminarianId, userId, disciplineId); err != nil {
		return err
	}
	if s.repo.CheckExistMark(userId, disciplineId) == nil {
		return s.repo.ChangeExamMark(userId, disciplineId, mark)
	}
	return s.repo.GiveExamMark(userId, disciplineId, mark)
}

func (s *SeminarianMarksService) GetAllMarksForExam(seminarianId, groupId, disciplineId int) ([]model.ExamMark, error) {
	if err := s.repo.CheckAccessForGroup(seminarianId, groupId, disciplineId); err != nil {
		return nil, err
	}
	return s.repo.GetAllMarksForExam(groupId, disciplineId)
}
