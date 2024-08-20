package api_seminarian_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type SeminarianReportService struct {
	repo repository.SeminarianReport
}

func NewSeminarianReportService(repo repository.SeminarianReport) *SeminarianReportService {
	return &SeminarianReportService{repo: repo}
}

func (s *SeminarianReportService) CheckAccessForGroup(seminarianId, groupId, disciplineId int) error {
	return s.repo.CheckAccessForGroup(seminarianId, groupId, disciplineId)
}

func (s *SeminarianReportService) GetAllStudents(groupId int) ([]model.StudentReport, error) {
	return s.repo.GetAllStudents(groupId)
}

func (s *SeminarianReportService) GetThemesFromDiscipline(disciplineId int) ([]model.Section, error) {
	return s.repo.GetThemesFromDiscipline(disciplineId)
}

func (s *SeminarianReportService) GetMarkFromSection(userId, sectionId int) (int, error) {
	return s.repo.GetMarkFromSection(userId, sectionId)
}

func (s *SeminarianReportService) GetSummaryMarkFromSections(userId, disciplineId int) int {
	return s.repo.GetSummaryMarkFromSections(userId, disciplineId)
}

func (s *SeminarianReportService) GetMarkFromExam(userId, disciplineId int) int {
	return s.repo.GetMarkFromExam(userId, disciplineId)
}

func (s *SeminarianReportService) GetFinalGrade(userId, disciplineId int) int {
	mark := s.GetSummaryMarkFromSections(userId, disciplineId) + s.GetMarkFromExam(userId, disciplineId)
	return mark
}

func (s *SeminarianReportService) GetResult(studentId, disciplineId int) (string, string) {
	return s.repo.GetResult(studentId, disciplineId)
}

func (s *SeminarianReportService) GetSectionsResult(studentId, disciplineId int) string {
	return s.repo.GetSectionsResult(studentId, disciplineId)
}
