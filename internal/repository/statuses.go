package repository

import (
	"database/sql"
	"example.com/m/v2/internal/domain"
)

//go:generate mockgen -source=statuses.go -destination=mocks/statuses.go

type StatusesStorage interface {
	GetStatus(int) (domain.Status, error)
	GetStatuses() ([]domain.Status, error)
}

type StatusRepository struct {
	db *sql.DB
}

func NewStatusRepository(db *sql.DB) *StatusRepository {
	return &StatusRepository{db}
}

func (res *StatusRepository) GetStatus(statusID int) (domain.Status, error) {
	row := res.db.QueryRow("SELECT * FROM statuses WHERE id = ?", statusID)

	status := domain.Status{}

	err := row.Scan(&status.ID, &status.Name)
	if err != nil {
		return domain.Status{}, err
	}

	return status, nil
}

func (res *StatusRepository) GetStatuses() ([]domain.Status, error) {
	rows, err := res.db.Query("select * from statuses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []domain.Status

	for rows.Next() {
		p := domain.Status{}
		err = rows.Scan(&p.ID, &p.Name)
		if err != nil {
			return nil, err
		}
		statuses = append(statuses, p)
	}
	return statuses, nil
}
