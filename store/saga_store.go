package store

import (
	"database/sql"

	"github.com/marcusziade/dragondex/models"
)

type SagaStore struct {
	db *sql.DB
}

func NewSagaStore(db *sql.DB) *SagaStore {
	return &SagaStore{db: db}
}

func (s *SagaStore) Create(sa *models.Saga) error {
	query := `INSERT INTO sagas (name, description, start_date, end_date) VALUES (?, ?, ?, ?)`
	_, err := s.db.Exec(query, sa.Name, sa.Description, sa.StartDate, sa.EndDate)
	return err
}

func (s *SagaStore) Get(id int) (*models.Saga, error) {
	query := `SELECT name, description, start_date, end_date FROM sagas WHERE id = ?`
	row := s.db.QueryRow(query, id)

	var sa models.Saga
	if err := row.Scan(&sa.Name, &sa.Description, &sa.StartDate, &sa.EndDate); err != nil {
		return nil, err
	}
	return &sa, nil
}

func (s *SagaStore) Update(sa *models.Saga) error {
	query := `UPDATE sagas SET name = ?, description = ?, start_date = ?, end_date = ? WHERE id = ?`
	_, err := s.db.Exec(query, sa.Name, sa.Description, sa.StartDate, sa.EndDate, sa.ID)
	return err
}

func (s *SagaStore) Delete(id int) error {
	query := `DELETE FROM sagas WHERE id = ?`
	_, err := s.db.Exec(query, id)
	return err
}
