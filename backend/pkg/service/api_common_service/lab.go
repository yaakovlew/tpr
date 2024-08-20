package api_common_service

import (
	"backend/pkg/repository"
)

type CommonLabService struct {
	repo repository.CommonLab
}

func NewCommonLabService(repository repository.CommonLab) *CommonLabService {
	return &CommonLabService{repo: repository}
}

func (s *CommonLabService) ChangeLabDateAndMark(studentId, laboratoryId, percentage int) error {
	if _, err := s.repo.GiveAccessForLab(studentId, laboratoryId); err != nil {
		return err
	}

	if err := s.repo.ChangeLabDate(studentId, laboratoryId); err != nil {
		return err
	}
	if err := s.repo.ChangeLabMark(studentId, laboratoryId, percentage); err != nil {
		return err
	}

	return nil
}
