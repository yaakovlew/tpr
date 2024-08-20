package api_lecturer_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type LecturerStudentService struct {
	repo repository.LecturerStudents
}

func NewLecturerStudentService(repo repository.LecturerStudents) *LecturerStudentService {
	return &LecturerStudentService{repo: repo}
}

func (s *LecturerStudentService) ChangeGroupForStudent(userId, groupId int) error {
	return s.repo.ChangeGroupForStudent(userId, groupId)
}

func (s *LecturerStudentService) GetAllStudents() ([]model.StudentWithGroup, error) {
	return s.repo.GetAllStudents()
}

func (s *LecturerStudentService) DeleteUser(userId int) error {
	return s.repo.DeleteUser(userId)
}
