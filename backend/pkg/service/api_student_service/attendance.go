package api_student_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type StudentAttendanceService struct {
	repo repository.StudentAttendance
}

func NewStudentAttendanceService(repo repository.StudentAttendance) *StudentAttendanceService {
	return &StudentAttendanceService{repo: repo}
}

func (s *StudentAttendanceService) GetAllSeminarVisiting(disciplineId, userId int) ([]model.SeminarVisiting, error) {
	return s.repo.GetAllSeminarVisiting(disciplineId, userId)
}

func (s *StudentAttendanceService) GetAllLessonVisiting(disciplineId, userId int) ([]model.LessonVisiting, error) {
	return s.repo.GetAllLessonVisiting(disciplineId, userId)
}

func (s *StudentAttendanceService) GetAllSeminars(userId, disciplineId int) ([]model.SeminarDate, error) {
	return s.repo.GetAllSeminars(userId, disciplineId)
}

func (s *StudentAttendanceService) GetAllLessons(userId, disciplineId int) ([]model.LessonDate, error) {
	return s.repo.GetAllLessons(userId, disciplineId)
}
