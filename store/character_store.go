package store

import (
	"database/sql"

	"github.com/marcusziade/dragondex/models"
)

type CharacterStore struct {
	db *sql.DB
}

func NewCharacterStore(db *sql.DB) *CharacterStore {
	return &CharacterStore{db: db}
}

func (s *CharacterStore) Create(c *models.Character) error {
	query := `INSERT INTO characters (name, race, description, power_level) VALUES (?, ?, ?, ?)`
	_, err := s.db.Exec(query, c.Name, c.Race, c.Description, c.PowerLevel)
	return err
}

func (s *CharacterStore) Get(id int) (*models.Character, error) {
	query := `SELECT name, race, description, power_level FROM characters WHERE id = ?`
	row := s.db.QueryRow(query, id)

	var c models.Character
	if err := row.Scan(&c.Name, &c.Race, &c.Description, &c.PowerLevel); err != nil {
		return nil, err
	}
	return &c, nil
}

func (s *CharacterStore) Update(c *models.Character) error {
	query := `UPDATE characters SET name = ?, race = ?, description = ?, power_level = ? WHERE id = ?`
	_, err := s.db.Exec(query, c.Name, c.Race, c.Description, c.PowerLevel, c.ID)
	return err
}

func (s *CharacterStore) Delete(id int) error {
	query := `DELETE FROM characters WHERE id = ?`
	_, err := s.db.Exec(query, id)
	return err
}
