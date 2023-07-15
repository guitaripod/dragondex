package store

import (
	"database/sql"

	"github.com/marcusziade/dragondex/models"
)

type AbilityStore struct {
	db *sql.DB
}

func NewAbilityStore(db *sql.DB) *AbilityStore {
	return &AbilityStore{db: db}
}

func (s *AbilityStore) Create(a *models.Ability) error {
	query := `INSERT INTO abilities (name, description, power_required) VALUES (?, ?, ?)`
	_, err := s.db.Exec(query, a.Name, a.Description, a.PowerRequired)
	return err
}

func (s *AbilityStore) Get(id int) (*models.Ability, error) {
	query := `SELECT name, description, power_required FROM abilities WHERE id = ?`
	row := s.db.QueryRow(query, id)

	var a models.Ability
	if err := row.Scan(&a.Name, &a.Description, &a.PowerRequired); err != nil {
		return nil, err
	}
	return &a, nil
}

func (s *AbilityStore) Update(a *models.Ability) error {
	query := `UPDATE abilities SET name = ?, description = ?, power_required = ? WHERE id = ?`
	_, err := s.db.Exec(query, a.Name, a.Description, a.PowerRequired, a.ID)
	return err
}

func (s *AbilityStore) Delete(id int) error {
	query := `DELETE FROM abilities WHERE id = ?`
	_, err := s.db.Exec(query, id)
	return err
}
