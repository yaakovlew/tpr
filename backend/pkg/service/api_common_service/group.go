package api_common_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type CommonGroupService struct {
	repo repository.CommonGroup
}

func NewCommonGroupService(repository repository.CommonGroup) *CommonGroupService {
	return &CommonGroupService{repo: repository}
}

func (s *CommonGroupService) GetAllGroups() ([]model.Group, error) {
	return s.repo.GetAllGroups()
}
