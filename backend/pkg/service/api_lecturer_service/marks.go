package api_lecturer_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"errors"
)

type LecturerMarksService struct {
	repo repository.LecturerMarks
}

func NewLecturerMarksService(repo repository.LecturerMarks) *LecturerMarksService {
	return &LecturerMarksService{repo: repo}
}

func (s *LecturerMarksService) ChangeTestMark(userId, testId, mark int) error {
	return s.repo.ChangeTestMark(userId, testId, mark)
}

func (s *LecturerMarksService) ChangeLaboratoryMark(userId, laboratoryId, mark int) error {
	return s.repo.ChangeLaboratoryMark(userId, laboratoryId, mark)
}

func (s *LecturerMarksService) GetTestMarksFromGroup(groupId, testId int) ([]model.GroupTestMarks, error) {
	return s.repo.GetTestMarksFromGroup(groupId, testId)
}

func (s *LecturerMarksService) GetLaboratoryMarksFromGroup(groupId, laboratoryId int) ([]model.GroupLaboratoryMarks, error) {
	return s.repo.GetLaboratoryMarksFromGroup(groupId, laboratoryId)
}

func (s *LecturerMarksService) GiveExamMark(userId, disciplineId, mark int) error {
	maxMark := s.repo.MaxExamMark(disciplineId)
	if maxMark < mark {
		err := errors.New("big mark")
		return err
	}
	if s.repo.CheckExistMark(userId, disciplineId) == nil {
		return s.repo.ChangeExamMark(userId, disciplineId, mark)
	}
	return s.repo.GiveExamMark(userId, disciplineId, mark)
}

func (s *LecturerMarksService) GetAllMarksForExam(groupId, disciplineId int) ([]model.ExamMark, error) {
	return s.repo.GetAllMarksForExam(groupId, disciplineId)
}
