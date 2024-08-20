package api_lecturer_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type LecturerAttendanceService struct {
	repo repository.LecturerAttendance
}

func NewLecturerAttendanceService(repo repository.LecturerAttendance) *LecturerAttendanceService {
	return &LecturerAttendanceService{repo: repo}
}

func (s *LecturerAttendanceService) GetAllLessons(disciplineId int) ([]model.Lesson, error) {
	return s.repo.GetAllLessons(disciplineId)
}

func (s *LecturerAttendanceService) AddLesson(disciplineId int, name string) error {
	return s.repo.AddLesson(disciplineId, name)
}

func (s *LecturerAttendanceService) DeleteLesson(lessonId int) error {
	return s.repo.DeleteLesson(lessonId)
}

func (s *LecturerAttendanceService) ChangeLesson(lessonId int, name string) error {
	return s.repo.ChangeLesson(lessonId, name)
}

func (s *LecturerAttendanceService) AddSeminar(disciplineId, groupId, date int, name string) error {
	return s.repo.AddSeminar(disciplineId, groupId, date, name)
}

func (s *LecturerAttendanceService) GetAllSeminars(disciplineId, groupId int) ([]model.Seminar, error) {
	return s.repo.GetAllSeminars(disciplineId, groupId)
}

func (s *LecturerAttendanceService) ChangeSeminar(seminarId int, name string) error {
	return s.repo.ChangeSeminar(seminarId, name)
}

func (s *LecturerAttendanceService) DeleteSeminar(seminarId int) error {
	return s.repo.DeleteSeminar(seminarId)
}

func (s *LecturerAttendanceService) GetLessonVisitingGroup(lessonId, groupId int) ([]model.LessonVisitingStudent, error) {
	return s.repo.GetLessonVisitingGroup(lessonId, groupId)
}

func (s *LecturerAttendanceService) GetSeminarVisitingGroup(seminarId int) ([]model.SeminarVisitingStudent, error) {
	return s.repo.GetSeminarVisitingGroup(seminarId)
}

func (s *LecturerAttendanceService) AddLessonVisiting(lessonId, userId int, isAbsent bool) error {
	return s.repo.AddLessonVisiting(lessonId, userId, isAbsent)
}

func (s *LecturerAttendanceService) AddSeminarVisiting(seminarId, userId int, isAbsent bool) error {
	return s.repo.AddSeminarVisiting(seminarId, userId, isAbsent)
}

func (s *LecturerAttendanceService) ChangeSeminarVisiting(seminarId, userId int, isAbsent bool) error {
	return s.repo.ChangeSeminarVisiting(seminarId, userId, isAbsent)
}

func (s *LecturerAttendanceService) ChangeLessonVisiting(lessonId, userId int, isAbsent bool) error {
	return s.repo.ChangeLessonVisiting(lessonId, userId, isAbsent)
}

func (s *LecturerAttendanceService) ChangeSeminarDate(seminarId int, date int) error {
	return s.repo.ChangeSeminarDate(seminarId, date)
}

func (s *LecturerAttendanceService) GetLessonDate(lessonId, groupId int) (model.LessonDate, error) {
	return s.repo.GetLessonDate(lessonId, groupId)
}

func (s *LecturerAttendanceService) ChangeLessonDate(lessonId, groupId, date int) error {
	return s.repo.ChangeLessonDate(lessonId, groupId, date)
}

func (s *LecturerAttendanceService) ChangeLessonDateDescription(lessonId, groupId int, description string) error {
	return s.repo.ChangeLessonDateDescription(lessonId, groupId, description)
}

func (s *LecturerAttendanceService) AddLessonDate(lessonId, groupId, date int, description string) error {
	lessonDate, _ := s.repo.GetLessonDate(lessonId, groupId)
	if lessonDate.Date != 0 {
		return s.repo.ChangeLessonDate(lessonId, groupId, date)
	}
	if lessonDate.Description != "" {
		return s.repo.ChangeLessonDateDescription(lessonId, groupId, description)
	}
	return s.repo.AddLessonDate(lessonId, groupId, date, description)
}

func (s *LecturerAttendanceService) DeleteLessonDate(lessonId int, groupId int) error {
	return s.repo.DeleteLessonDate(lessonId, groupId)
}

func (s *LecturerAttendanceService) GetTableLessons(disciplineId int) ([]model.LessonDate, error) {
	return s.repo.GetTableLessons(disciplineId)
}
func (s *LecturerAttendanceService) GetTableSeminars(disciplineId int) ([]model.SeminarDate, error) {
	return s.repo.GetTableSeminars(disciplineId)
}
