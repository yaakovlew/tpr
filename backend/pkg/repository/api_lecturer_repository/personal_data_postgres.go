package api_lecturer_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type LecturerPersonalDataPostgres struct {
	db *sqlx.DB
}

func NewLecturerPersonalDataPostgres(db *sqlx.DB) *LecturerPersonalDataPostgres {
	return &LecturerPersonalDataPostgres{db: db}
}

func (r *LecturerPersonalDataPostgres) GetPersonalData(id int) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", table_name.UsersTable)
	err := r.db.Get(&user, query, id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *LecturerPersonalDataPostgres) UpdateName(id int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1 WHERE id = $2", table_name.UsersTable)
	_, err := r.db.Exec(query, name, id)
	return err
}

func (r *LecturerPersonalDataPostgres) UpdateSurname(id int, surname string) error {
	query := fmt.Sprintf("UPDATE %s SET surname=$1 WHERE id = $2", table_name.UsersTable)
	_, err := r.db.Exec(query, surname, id)
	return err
}
