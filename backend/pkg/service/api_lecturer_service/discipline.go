package api_lecturer_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type LecturerDisciplineService struct {
	repo repository.LecturerDiscipline
}

func NewLecturerDisciplineService(repo repository.LecturerDiscipline) *LecturerDisciplineService {
	return &LecturerDisciplineService{repo: repo}
}

func (s *LecturerDisciplineService) GetGroupForDiscipline(id int) ([]model.Group, error) {
	return s.repo.GetGroupForDiscipline(id)
}

func (s *LecturerDisciplineService) AddDiscipline(discipline model.AddNewDiscipline) (int, error) {
	return s.repo.AddDiscipline(discipline)
}

func (s *LecturerDisciplineService) ChangeLessonMarks(id, lessonMarks int) error {
	return s.repo.ChangeLessonMarks(id, lessonMarks)
}
func (s *LecturerDisciplineService) ChangeSeminarMarks(id, seminarMarks int) error {
	return s.repo.ChangeSeminarMarks(id, seminarMarks)
}
func (s *LecturerDisciplineService) GetAllDisciplines() ([]model.Discipline, []model.Discipline, error) {
	return s.repo.GetAllDisciplines()
}

func (s *LecturerDisciplineService) GetAllInfoAboutDiscipline(id int) (model.DisciplineInfo, error) {
	return s.repo.GetAllInfoAboutDiscipline(id)
}

func (s *LecturerDisciplineService) DeleteDiscipline(id int) error {
	return s.repo.DeleteDiscipline(id)
}

func (s *LecturerDisciplineService) AddGroupToDiscipline(groupId, disciplineId int) error {
	return s.repo.AddGroupToDiscipline(groupId, disciplineId)
}

func (s *LecturerDisciplineService) DeleteGroupFromDiscipline(groupId, disciplineId int) error {
	return s.repo.DeleteGroupFromDiscipline(groupId, disciplineId)
}

func (s *LecturerDisciplineService) ChangeDiscipline(disciplineId int, name string) error {
	return s.repo.ChangeDiscipline(disciplineId, name)
}

func (s *LecturerDisciplineService) ChangeExamMark(id, examMark int) error {
	return s.repo.ChangeExamMark(id, examMark)
}

func (s *LecturerDisciplineService) ChangeDisciplineEn(disciplineId int, name string) error {
	return s.repo.ChangeDisciplineEn(disciplineId, name)
}

func (s *LecturerDisciplineService) GetGroupsAvailableToAddForDiscipline(id int) ([]model.Group, error) {
	return s.repo.GetGroupsAvailableToAddForDiscipline(id)
}

func (s *LecturerDisciplineService) ArchiveGroupToDiscipline(groupId, disciplineId int) error {
	return s.repo.ArchiveGroupToDiscipline(groupId, disciplineId)
}
