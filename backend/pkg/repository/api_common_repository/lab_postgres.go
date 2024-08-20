package api_common_repository

import (
	"backend/pkg/repository/table_name"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type CommonLabPostgres struct {
	db *sqlx.DB
}

func NewCommonLabPostgres(db *sqlx.DB) *CommonLabPostgres {
	return &CommonLabPostgres{db: db}
}

func (r *CommonLabPostgres) ChangeLabDate(studentId, laboratoryId int) error {
	var isDone bool
	query := fmt.Sprintf("SELECT is_done FROM %s WHERE user_id = $1 AND laboratory_id = $2", table_name.LaboratoryDateTable)
	if err := r.db.QueryRow(query, studentId, laboratoryId).Scan(&isDone); err != nil {
		query := fmt.Sprintf("INSERT INTO %s (user_id, laboratory_id, closed_date, is_done) VALUES ($1, $2, $3, true)", table_name.LaboratoryDateTable)
		if _, err := r.db.Exec(query, studentId, laboratoryId, time.Now().Unix()); err != nil {
			return err
		}
		return nil
	}

	if isDone {
		return fmt.Errorf("already done")
	}
	query = fmt.Sprintf("UPDATE %s SET is_done = true, closed_date = $1  WHERE user_id = $2 AND laboratory_id = $3", table_name.LaboratoryDateTable)
	if _, err := r.db.Exec(query, time.Now().Unix(), studentId, laboratoryId); err != nil {
		return err
	}

	return nil
}

func (r *CommonLabPostgres) ChangeLabMark(studentId, laboratoryId, percentage int) error {
	var defaultMark int
	query := fmt.Sprintf("SELECT default_mark FROM %s WHERE id = $1", table_name.LaboratoryTable)
	if err := r.db.Get(&defaultMark, query, laboratoryId); err != nil {
		return err
	}

	var markLast int
	query = fmt.Sprintf("SELECT mark FROM %s WHERE user_id = $1 AND laboratory_id = $2", table_name.LaboratoryMarkTable)
	if err := r.db.QueryRow(query, studentId, laboratoryId).Scan(&markLast); err != nil {
		query := fmt.Sprintf("INSERT INTO %s (user_id, laboratory_id, mark) VALUES ($1, $2, $3)", table_name.LaboratoryMarkTable)
		if _, err := r.db.Exec(query, studentId, laboratoryId, percentage*defaultMark/100); err != nil {
			return err
		}
		return nil
	}

	query = fmt.Sprintf("UPDATE %s SET mark = $1 WHERE user_id = $2 AND laboratory_id = $3", table_name.LaboratoryMarkTable)
	if _, err := r.db.Exec(query, percentage*defaultMark/100, studentId, laboratoryId); err != nil {
		return err
	}

	return nil
}

func (r *CommonLabPostgres) GiveAccessForLab(userId, labId int) (bool, error) {
	var isDone bool
	currentTime := time.Now().Unix()
	query := fmt.Sprintf("SELECT is_done FROM %s WHERE user_id = $1 AND laboratory_id = $2 AND closed_date + 6000 > $3", table_name.LaboratoryDateTable)
	errAccess := fmt.Errorf("access to lab denied")
	row := r.db.QueryRow(query, userId, labId, currentTime)
	if err := row.Scan(&isDone); err != nil {
		return false, errAccess
	}
	if isDone {
		return false, errAccess
	} else {
		return true, nil
	}
}
