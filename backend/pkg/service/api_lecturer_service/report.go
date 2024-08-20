package api_lecturer_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type LecturerReportService struct {
	repo repository.LecturerReport
}

func NewLecturerReportService(repo repository.LecturerReport) *LecturerReportService {
	return &LecturerReportService{repo: repo}
}

func (s *LecturerReportService) GetAllStudents(groupId int) ([]model.StudentReport, error) {
	return s.repo.GetAllStudents(groupId)
}

func (s *LecturerReportService) GetThemesFromDiscipline(disciplineId int) ([]model.Section, error) {
	return s.repo.GetThemesFromDiscipline(disciplineId)
}

func (s *LecturerReportService) GetMarkFromSection(userId, sectionId int) (int, error) {
	return s.repo.GetMarkFromSection(userId, sectionId)
}

func (s *LecturerReportService) GetSummaryMarkFromSections(userId, disciplineId int) int {
	return s.repo.GetSummaryMarkFromSections(userId, disciplineId)
}

func (s *LecturerReportService) GetMarkFromExam(userId, disciplineId int) int {
	return s.repo.GetMarkFromExam(userId, disciplineId)
}

func (s *LecturerReportService) GetFinalGrade(userId, disciplineId int) int {
	mark := s.GetSummaryMarkFromSections(userId, disciplineId) + s.GetMarkFromExam(userId, disciplineId)
	return mark
}

func (s *LecturerReportService) GetResult(studentId, disciplineId int) (string, string) {
	return s.repo.GetResult(studentId, disciplineId)
}

func (s *LecturerReportService) GetSectionsResult(studentId, disciplineId int) string {
	return s.repo.GetSectionsResult(studentId, disciplineId)
}
