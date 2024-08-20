package api_lecturer_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type LecturerSeminarianPostgres struct {
	db *sqlx.DB
}

func NewLecturerSeminarianPostgres(db *sqlx.DB) *LecturerSeminarianPostgres {
	return &LecturerSeminarianPostgres{db: db}
}

func (r *LecturerSeminarianPostgres) GetSeminarianFromGroupsAndDiscipline(groupId, disciplineId int) ([]model.Seminarian, error) {
	var seminarian []model.Seminarian
	query := fmt.Sprintf("SELECT id, name, surname, email FROM %s INNER JOIN %s ON %s.user_id = %s.id WHERE group_id = $1 AND discipline_id = $2 ORDER BY %s.user_id",
		table_name.SeminariansGroupsTable, table_name.UsersTable, table_name.SeminariansGroupsTable, table_name.UsersTable, table_name.SeminariansGroupsTable)
	err := r.db.Select(&seminarian, query, groupId, disciplineId)
	if err != nil {
		return []model.Seminarian{}, err
	}
	return seminarian, nil
}

func (r *LecturerSeminarianPostgres) GetAllSeminarians() ([]model.Seminarian, error) {
	var seminarians []model.Seminarian
	query := fmt.Sprintf("SELECT id, name, surname, email FROM %s INNER JOIN %s ON %s.id = %s.user_id ORDER BY %s.id",
		table_name.UsersTable, table_name.SeminariansTable, table_name.UsersTable, table_name.SeminariansTable, table_name.UsersTable)
	err := r.db.Select(&seminarians, query)
	if err != nil {
		return []model.Seminarian{}, err
	}
	return seminarians, nil
}

func (r *LecturerSeminarianPostgres) AddSeminarian(seminarianId, groupId, disciplineId int) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, group_id, discipline_id) VALUES ($1, $2, $3)",
		table_name.SeminariansGroupsTable)
	_, err := r.db.Exec(query, seminarianId, groupId, disciplineId)
	return err
}

func (r *LecturerSeminarianPostgres) DeleteSeminarianFromGroupAndDiscipline(seminarianId, groupId, disciplineId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE ((user_id = $1 AND group_id = $2) AND discipline_id = $3)",
		table_name.SeminariansGroupsTable)
	_, err := r.db.Exec(query, seminarianId, groupId, disciplineId)
	return err
}
