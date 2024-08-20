package api_student_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type StudentDisciplineService struct {
	repo repository.StudentDiscipline
}

func NewStudentDisciplinesService(repo repository.StudentDiscipline) *StudentDisciplineService {
	return &StudentDisciplineService{repo: repo}
}

func (s *StudentDisciplineService) GetAllUserDiscipline(id int) ([]model.Discipline, []model.Discipline, error) {
	return s.repo.GetAllUserDiscipline(id)
}

func (s *StudentDisciplineService) GetDisciplineSections(userId, disciplineId int) ([]model.Section, []model.Section, error) {
	return s.repo.GetDisciplineSections(disciplineId)
}
