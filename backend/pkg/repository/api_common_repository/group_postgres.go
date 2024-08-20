package api_common_repository

import (
	"backend/pkg/model"
	"backend/pkg/repository/table_name"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CommonGroupPostgres struct {
	db *sqlx.DB
}

func NewCommonGroupPostgres(db *sqlx.DB) *CommonGroupPostgres {
	return &CommonGroupPostgres{db: db}
}

func (r *CommonGroupPostgres) GetAllGroups() ([]model.Group, error) {
	var groups []model.Group
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id", table_name.GroupsTable)
	err := r.db.Select(&groups, query)
	if err != nil {
		return []model.Group{}, err
	}
	return groups, nil
}
