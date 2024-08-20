package api_seminarian_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type SeminarianTestAndLabService struct {
	repo repository.SeminarianTestAndLab
}

func NewSeminarianTestAndLabService(repo repository.SeminarianTestAndLab) *SeminarianTestAndLabService {
	return &SeminarianTestAndLabService{repo: repo}
}

func (s *SeminarianTestAndLabService) GetAllTestFromSection(seminarianId, sectionId int) ([]model.Test, []model.Test, error) {
	if err := s.repo.CheckAccessForSection(seminarianId, sectionId); err != nil {
		return nil, nil, err
	}
	return s.repo.GetAllTestFromSection(sectionId)
}

func (s *SeminarianTestAndLabService) GetAllLab(seminarianId, sectionId int) ([]model.LaboratoryWork, []model.LaboratoryWork, error) {
	if err := s.repo.CheckAccessForSection(seminarianId, sectionId); err != nil {
		return nil, nil, err
	}
	return s.repo.GetAllLab(sectionId)
}

func (s *SeminarianTestAndLabService) OpenTestForStudent(seminarianId, studentId, testId int, date int64) error {
	if err := s.repo.CheckAccessToOpenTest(seminarianId, studentId, testId); err != nil {
		return err
	}

	return s.repo.OpenTestForStudent(studentId, testId, date)
}

func (s *SeminarianTestAndLabService) GetOpenedTestForStudent(seminarianId, studentId, testId int) (model.OpenedTest, error) {
	if err := s.repo.CheckAccessToOpenTest(seminarianId, studentId, testId); err != nil {
		return model.OpenedTest{}, err
	}
	return s.repo.GetOpenedTestForStudent(studentId, testId)
}

func (s *SeminarianTestAndLabService) CloseOpenedTestForStudent(seminarianId, studentId, testId int) error {
	if err := s.repo.CheckAccessToOpenTest(seminarianId, studentId, testId); err != nil {
		return err
	}
	return s.repo.CloseOpenedTestForStudent(studentId, testId)
}

func (s *SeminarianTestAndLabService) GetUsersWithOpenedTest(seminarianId, testId int) ([]model.StudentWithGroupWithClosedDate, error) {
	return s.repo.GetUsersWithOpenedTest(seminarianId, testId, time.Now().Unix())
}

func (s *SeminarianTestAndLabService) GetUsersWithDoneTest(seminarianId, testId int) ([]model.StudentWithGroupWithClosedDate, error) {
	return s.repo.GetUsersWithDoneTest(seminarianId, testId, time.Now().Unix())
}

func (s *SeminarianTestAndLabService) GetOpenedLabForStudent(seminarianId, studentId, labId int) (model.OpenedLab, error) {
	if err := s.repo.CheckAccessToOpenLab(seminarianId, studentId, labId); err != nil {
		return model.OpenedLab{}, err
	}
	return s.repo.GetOpenedLabForStudent(studentId, labId)
}

func (s *SeminarianTestAndLabService) CloseOpenedLabForStudent(seminarianId, studentId, labId int) error {
	if err := s.repo.CheckAccessToOpenLab(seminarianId, studentId, labId); err != nil {
		return err
	}
	externalLabId, err := s.repo.GetInternalLabInfo(labId)
	if err != nil {
		return err
	}

	token, err := s.repo.GetLabToken(externalLabId)
	if err != nil {
		return err
	}
	lab, err := s.repo.GetExternalLabInfo(externalLabId)
	if err != nil {
		return err
	}
	if err := s.sendRequestToOpenLab(lab, studentId, labId, token, false); err != nil {
		return err
	}

	return s.repo.CloseOpenedLabForStudent(studentId, labId)
}

func (s *SeminarianTestAndLabService) GetUsersWithOpenedLab(seminarianId, labId, groupId int) ([]model.StudentWithGroupWithClosedDate, error) {
	return s.repo.GetUsersWithOpenedLab(seminarianId, labId, groupId, time.Now().Unix())
}

func (s *SeminarianTestAndLabService) GetUsersWithDoneLab(seminarianId, labId, groupId int) ([]model.StudentWithGroupWithClosedDate, error) {
	return s.repo.GetUsersWithDoneLab(seminarianId, labId, groupId, time.Now().Unix())
}

func (s *SeminarianTestAndLabService) GetPathForReportTest(seminarianId, userId, testId int) (string, error) {
	if err := s.repo.CheckAccessToOpenTest(seminarianId, userId, testId); err != nil {
		return "", err
	}

	path := viper.GetString("test") + "/" + strconv.Itoa(userId) + "-" + strconv.Itoa(testId) + ".txt"

	return path, nil
}

func (s *SeminarianTestAndLabService) OpenLabForStudent(seminarianId, studentId, labId int, date int64) error {
	if err := s.repo.CheckAccessToOpenLab(seminarianId, studentId, labId); err != nil {
		return err
	}

	externalLabId, err := s.repo.GetInternalLabInfo(labId)
	if err != nil {
		return err
	}
	token, err := s.repo.GetLabToken(externalLabId)
	if err != nil {
		return err
	}
	lab, err := s.repo.GetExternalLabInfo(externalLabId)
	if err != nil {
		return err
	}
	if err := s.sendRequestToOpenLab(lab, studentId, labId, token, true); err != nil {
		return err
	}

	return s.repo.OpenLabForStudent(studentId, labId, date)
}

func (s *SeminarianTestAndLabService) sendRequestToOpenLab(lab model.LaboratoryWorkResponse, userId, labId int, token string, isOpen bool) error {
	url := fmt.Sprintf("%s/%s?user_id=%d&is_open=%t&lab_id=%d", lab.Link, "open", userId, isOpen, labId)
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return err
	}
	req.Header.Add("lab-token", token)
	req.Header.Add("lecturer-token", os.Getenv("LECTURER_HEADER"))

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if _, err := ioutil.ReadAll(res.Body); err != nil {
		return err
	}

	return nil
}
