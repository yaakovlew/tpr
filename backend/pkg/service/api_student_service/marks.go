package api_student_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type StudentMarksService struct {
	repo repository.StudentMarks
}

func NewStudentMarksService(repo repository.StudentMarks) *StudentMarksService {
	return &StudentMarksService{repo: repo}
}

func (s *StudentMarksService) GetAllTestsMarks(id, disciplineId int) ([]model.TestsMarks, error) {
	return s.repo.GetAllTestsMarks(id, disciplineId)
}

func (s *StudentMarksService) GetAllLaboratoryMarks(id, disciplineId int) ([]model.LaboratoryMarks, error) {
	return s.repo.GetAllLaboratoryMarks(id, disciplineId)
}

func (s *StudentMarksService) GetExamMark(userId, disciplineId int) (int, error) {
	return s.repo.GetExamMark(userId, disciplineId)
}
