package api_lecturer_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type LecturerStudyGuideService struct {
	repo repository.LecturerStudyGuide
}

func NewLecturerStudyGuideService(repo repository.LecturerStudyGuide) *LecturerStudyGuideService {
	return &LecturerStudyGuideService{repo: repo}
}

func (s *LecturerStudyGuideService) AddStudyGuideHeader(name, nameEn, description, descriptionEn string) (int, error) {
	return s.repo.AddStudyGuideHeader(name, nameEn, description, descriptionEn)
}

func (s *LecturerStudyGuideService) GetStudyGuideHeader() ([]model.StudyGuideHeader, []model.StudyGuideHeader, error) {
	return s.repo.GetStudyGuideHeader()
}

func (s *LecturerStudyGuideService) ChangeNameDigitalGuideHeaderEn(digitalGuideId int, name string) error {
	return s.repo.ChangeNameDigitalGuideHeaderEn(digitalGuideId, name)
}

func (s *LecturerStudyGuideService) ChangeDescriptionDigitalGuideHeaderEn(digitalGuideId int, description string) error {
	return s.repo.ChangeDescriptionDigitalGuideHeaderEn(digitalGuideId, description)
}

func (s *LecturerStudyGuideService) DeleteStudyGuideHeader(digitalGuideId int) (int, error) {
	return s.repo.DeleteStudyGuideHeader(digitalGuideId)
}

func (s *LecturerStudyGuideService) ChangeNameDigitalGuideHeader(digitalGuideId int, name string) error {
	return s.repo.ChangeNameDigitalGuideHeader(digitalGuideId, name)
}

func (s *LecturerStudyGuideService) ChangeDescriptionDigitalGuideHeader(digitalGuideId int, description string) error {
	return s.repo.ChangeDescriptionDigitalGuideHeader(digitalGuideId, description)
}

func (s *LecturerStudyGuideService) GetDigitalDiscipline(disciplineId int) ([]model.DigitalDiscipline, error) {
	return s.repo.GetDigitalDiscipline(disciplineId)
}

func (s *LecturerStudyGuideService) AddDigitalDiscipline(digitalMaterialId, disciplineId int) error {
	return s.repo.AddDigitalDiscipline(digitalMaterialId, disciplineId)
}

func (s *LecturerStudyGuideService) DeleteDigitalDiscipline(digitalMaterialId, disciplineId int) error {
	return s.repo.DeleteDigitalDiscipline(digitalMaterialId, disciplineId)
}

func (s *LecturerStudyGuideService) GetFilesIdFromDigital(digitalId int) ([]model.FileId, error) {
	return s.repo.GetFilesIdFromDigital(digitalId)
}

func (s *LecturerStudyGuideService) AddFileToDigital(path string, digitalId int) error {
	return s.repo.AddFileToDigital(path, digitalId)
}

func (s *LecturerStudyGuideService) DeleteFileFromDigital(fileId int) (string, error) {
	return s.repo.DeleteFileFromDigital(fileId)
}

func (s *LecturerStudyGuideService) GetFile(fileId int) (string, error) {
	return s.repo.GetFile(fileId)
}
