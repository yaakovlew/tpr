package api_lecturer_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type LecturerGroupService struct {
	repo repository.LecturerGroup
}

func NewLecturerGroupService(repo repository.LecturerGroup) *LecturerGroupService {
	return &LecturerGroupService{repo: repo}
}

func (s *LecturerGroupService) AddGroup(name string) error {
	return s.repo.AddGroup(name)
}

func (s *LecturerGroupService) DeleteGroup(id int) error {
	return s.repo.DeleteGroup(id)
}

func (s *LecturerGroupService) GetAllStudentsFromGroup(id int) ([]model.Student, error) {
	return s.repo.GetAllStudentsFromGroup(id)
}

func (s *LecturerGroupService) GetAllGroups() ([]model.Group, error) {
	return s.repo.GetAllGroups()
}

func (s *LecturerGroupService) GetGroupsDisciplines(groupId int) ([]model.Discipline, []model.Discipline, error) {
	return s.repo.GetGroupsDisciplines(groupId)
}

func (s *LecturerGroupService) ChangeName(groupId int, name string) error {
	return s.repo.ChangeName(groupId, name)
}

func (s *LecturerGroupService) AddGroupInArchive(groupId int) error {
	return s.repo.AddGroupInArchive(groupId)
}

func (s *LecturerGroupService) DeleteGroupFromArchive(groupId int) error {
	return s.repo.DeleteGroupFromArchive(groupId)
}
