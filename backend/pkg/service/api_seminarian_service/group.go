package api_seminarian_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type SeminarianGroupService struct {
	repo repository.SeminarianGroup
}

func NewSeminarianGroupService(repo repository.SeminarianGroup) *SeminarianGroupService {
	return &SeminarianGroupService{repo: repo}
}

func (s *SeminarianGroupService) GetAllStudentsFromGroup(seminarianId, groupId int) ([]model.Student, error) {
	if err := s.repo.CheckAccessForGroup(seminarianId, groupId); err != nil {
		return nil, err
	}
	return s.repo.GetAllStudentsFromGroup(groupId)
}

func (s *SeminarianGroupService) GetOwnGroup(userId int, disciplineId int) ([]model.Group, error) {
	return s.repo.GetOwnGroup(userId, disciplineId)
}
