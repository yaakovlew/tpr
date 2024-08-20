package api_lecturer_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type LecturerSectionService struct {
	repo repository.LecturerSection
}

func NewLecturerSectionService(repo repository.LecturerSection) *LecturerSectionService {
	return &LecturerSectionService{repo: repo}
}

func (s *LecturerSectionService) AddSection(name string, nameEn string, disciplineId int) error {
	return s.repo.AddSection(name, nameEn, disciplineId)
}

func (s *LecturerSectionService) GetDisciplineSections(disciplineId int) ([]model.Section, []model.Section, error) {
	return s.repo.GetDisciplineSections(disciplineId)
}

func (s *LecturerSectionService) DeleteSection(sectionId int) error {
	return s.repo.DeleteSection(sectionId)
}

func (s *LecturerSectionService) ChangeSectionName(sectionId int, name string) error {
	return s.repo.ChangeSectionName(sectionId, name)
}

func (s *LecturerSectionService) ChangeSectionNameEn(sectionId int, name string) error {
	return s.repo.ChangeSectionNameEn(sectionId, name)
}

func (s *LecturerSectionService) AddTestToSection(sectionId, testId int) error {
	return s.repo.AddTestToSection(sectionId, testId)
}

func (s *LecturerSectionService) DeleteTestFromSection(sectionId, testId int) error {
	return s.repo.DeleteTestFromSection(sectionId, testId)
}

func (s *LecturerSectionService) AddLabToSection(labId, sectionId, defaultMark int) error {
	return s.repo.AddLabToSection(labId, sectionId, defaultMark)
}

func (s *LecturerSectionService) DeleteLabFromSection(labId, sectionId int) error {
	return s.repo.DeleteLabFromSection(labId, sectionId)
}

func (s *LecturerSectionService) GetLabFromSection(sectionId int) ([]model.LaboratoryWorkWithExternal, error) {
	return s.repo.GetLabFromSection(sectionId)
}
