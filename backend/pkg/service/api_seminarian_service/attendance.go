package api_seminarian_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type SeminarianAttendanceService struct {
	repo repository.SeminarianAttendance
}

func NewSeminarianAttendanceService(repo repository.SeminarianAttendance) *SeminarianAttendanceService {
	return &SeminarianAttendanceService{repo: repo}
}

func (s *SeminarianAttendanceService) AddSeminar(seminarianId, disciplineId, groupId int, name string, date int) error {
	if err := s.repo.CheckAccessForDiscipline(seminarianId, disciplineId); err != nil {
		return err
	}
	if err := s.repo.CheckAccessForGroup(seminarianId, groupId, disciplineId); err != nil {
		return err
	}
	return s.repo.AddSeminar(disciplineId, groupId, name, date)
}

func (s *SeminarianAttendanceService) GetAllSeminars(seminarianId, disciplineId, groupId int) ([]model.Seminar, error) {
	if err := s.repo.CheckAccessForDiscipline(seminarianId, disciplineId); err != nil {
		return nil, err
	}
	if err := s.repo.CheckAccessForGroup(seminarianId, groupId, disciplineId); err != nil {
		return nil, err
	}
	return s.repo.GetAllSeminars(disciplineId, groupId)
}

func (s *SeminarianAttendanceService) ChangeSeminar(seminarianId, seminarId int, name string) error {
	if err := s.repo.CheckAccessForSeminar(seminarianId, seminarId); err != nil {
		return err
	}
	return s.repo.ChangeSeminar(seminarId, name)
}

func (s *SeminarianAttendanceService) DeleteSeminar(seminarianId, seminarId int) error {
	if err := s.repo.CheckAccessForSeminar(seminarianId, seminarId); err != nil {
		return err
	}
	return s.repo.DeleteSeminar(seminarId)
}

func (s *SeminarianAttendanceService) GetLessonVisitingGroup(seminarianId, lessonId, groupId int) ([]model.LessonVisitingStudent, error) {
	if err := s.repo.CheckAccessForLesson(seminarianId, lessonId, groupId); err != nil {
		return nil, err
	}
	return s.repo.GetLessonVisitingGroup(lessonId, groupId)
}

func (s *SeminarianAttendanceService) GetSeminarVisitingGroup(seminarianId, seminarId int) ([]model.SeminarVisitingStudent, error) {
	if err := s.repo.CheckAccessForSeminar(seminarianId, seminarId); err != nil {
		return nil, err
	}
	return s.repo.GetSeminarVisitingGroup(seminarId)
}

func (s *SeminarianAttendanceService) AddLessonVisiting(seminarianId, lessonId, userId int, isAbsent bool) error {
	if err := s.repo.CheckAccessForStudentToLesson(seminarianId, userId, lessonId); err != nil {
		return err
	}
	return s.repo.AddLessonVisiting(lessonId, userId, isAbsent)
}

func (s *SeminarianAttendanceService) AddSeminarVisiting(seminarianId, seminarId, userId int, isAbsent bool) error {
	if err := s.repo.CheckAccessForStudentToSeminar(seminarianId, userId, seminarId); err != nil {
		return err
	}
	return s.repo.AddSeminarVisiting(seminarId, userId, isAbsent)
}

func (s *SeminarianAttendanceService) ChangeSeminarVisiting(seminarianId, seminarId, userId int, isAbsent bool) error {
	if err := s.repo.CheckAccessForStudentToSeminar(seminarianId, userId, seminarId); err != nil {
		return err
	}
	return s.repo.ChangeSeminarVisiting(seminarId, userId, isAbsent)
}

func (s *SeminarianAttendanceService) ChangeLessonVisiting(seminarianId, lessonId, userId int, isAbsent bool) error {
	if err := s.repo.CheckAccessForStudentToLesson(seminarianId, userId, lessonId); err != nil {
		return err
	}
	return s.repo.ChangeLessonVisiting(lessonId, userId, isAbsent)
}

func (s *SeminarianAttendanceService) GetAllLessons(seminarianId, disciplineId int) ([]model.Lesson, error) {
	if err := s.repo.CheckAccessForDiscipline(seminarianId, disciplineId); err != nil {
		return nil, err
	}
	return s.repo.GetAllLessons(disciplineId)
}

func (s *SeminarianAttendanceService) ChangeSeminarDate(seminarianId, seminarId int, date int) error {
	if err := s.repo.CheckAccessForSeminar(seminarianId, seminarId); err != nil {
		return err
	}
	return s.repo.ChangeSeminarDate(seminarId, date)
}

func (s *SeminarianAttendanceService) GetLessonDate(seminarianId, lessonId, groupId int) (model.LessonDate, error) {
	if err := s.repo.CheckAccessForLesson(seminarianId, lessonId, groupId); err != nil {
		return model.LessonDate{}, err
	}
	return s.repo.GetLessonDate(lessonId, groupId)
}

func (s *SeminarianAttendanceService) GetTableLessons(seminarianId, disciplineId int) ([]model.LessonDate, error) {
	if err := s.repo.CheckAccessForDiscipline(seminarianId, disciplineId); err != nil {
		return nil, err
	}
	return s.repo.GetTableLessons(disciplineId)
}

func (s *SeminarianAttendanceService) GetTableLessonsByGroup(seminarianId, disciplineId, groupId int) ([]model.LessonDate, error) {
	if err := s.repo.CheckAccessForDiscipline(seminarianId, disciplineId); err != nil {
		return nil, err
	}
	groups, err := s.repo.GetTableLessons(disciplineId)
	if err != nil {
		return nil, err
	}

	var result []model.LessonDate
	for _, group := range groups {
		if group.GroupId == groupId {
			result = append(result, group)
		}
	}

	return result, nil
}

func (s *SeminarianAttendanceService) GetTableSeminars(seminarianId, disciplineId, groupId int) ([]model.SeminarDate, error) {
	if err := s.repo.CheckAccessForGroup(seminarianId, groupId, disciplineId); err != nil {
		return nil, err
	}
	return s.repo.GetTableSeminars(disciplineId)
}
