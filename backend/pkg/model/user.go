package model

import (
	"github.com/spf13/viper"
	"net/url"
	"time"
)

type User struct {
	Id        int    `json:"-" db:"id"`
	Name      string `json:"name" binding:"required"`
	Surname   string `json:"surname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Post      string `json:"post" binding:"required"`
	GroupName string `json:"group_name"`
	Password  string `json:"password" binding:"required"`
}

type UserData struct {
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	GroupName string `json:"group_name"`
}

type NewName struct {
	Name string `json:"name" binding:"required"`
}

type NewSurname struct {
	Surname string `json:"surname" binding:"required"`
}

type Student struct {
	Id      int    `json:"student_id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type StudentResponse struct {
	Students []Student `json:"students"`
}

type Seminarian struct {
	Id      int    `json:"seminarian_id" db:"id"`
	Name    string `json:"name" db:"name"`
	Surname string `json:"surname" db:"surname"`
	Email   string `json:"email" db:"email"`
}

type SeminarianAddToGroupInput struct {
	Id           int `json:"seminarian_id" binding:"required"`
	GroupId      int `json:"group_id" binding:"required"`
	DisciplineId int `json:"discipline_id" binding:"required"`
}

type SeminarianResponse struct {
	Seminarians []Seminarian `json:"seminarians"`
}

type StudentChangeGroup struct {
	UserId  int `json:"user_id" binding:"required"`
	GroupId int `json:"group_id" binding:"required"`
}

type StudentWithGroup struct {
	Id        int    `json:"student_id" db:"id"`
	Name      string `json:"name" db:"name"`
	Surname   string `json:"surname" db:"surname"`
	GroupName string `json:"group_name" db:"group_name"`
}

type StudentWithGroupWithClosedDate struct {
	Id         int    `json:"student_id" db:"id"`
	Name       string `json:"name" db:"name"`
	Surname    string `json:"surname" db:"surname"`
	GroupName  string `json:"group_name" db:"group_name"`
	ClosedDate int    `json:"closed_date" db:"closed_date"`
}

type StudentWithGroupResponse struct {
	Students []StudentWithGroup `json:"students"`
}

type StudentWithGroupWithClosedDateResponse struct {
	Students []StudentWithGroupWithClosedDate `json:"students"`
}

type OpenTestInput struct {
	UserId int   `json:"user_id" binding:"required"`
	TestId int   `json:"test_id" binding:"required"`
	Date   int64 `json:"date"`
}

type OpenLabInput struct {
	UserId int   `json:"user_id" binding:"required"`
	LabId  int   `json:"laboratory_id" binding:"required"`
	Date   int64 `json:"date"`
}

type ChangePasswordInput struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type ForgotPasswordInput struct {
	Email string `json:"email" binding:"required"`
}

type LecturerChangePasswordForOtherInput struct {
	UserId      int    `json:"user_id" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type RestorePasswordInput struct {
	NewPassword string `json:"new_password" binding:"required"`
}

type PasswordResetLink struct {
	Email  string
	Token  string
	Expiry time.Time
}

func (link *PasswordResetLink) ToURL() string {
	u := &url.URL{
		Scheme: "https",
		Host:   viper.GetString("host"),
		Path:   "/reset-password",
		RawQuery: url.Values{
			"token":  {link.Token},
			"expiry": {link.Expiry.Format(time.RFC3339)},
		}.Encode(),
	}
	return u.String()
}
