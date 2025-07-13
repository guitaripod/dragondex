package store

import (
	"database/sql"

	"github.com/marcusziade/dragondex/models"
)

type EpisodeStore struct {
	db *sql.DB
}

func NewEpisodeStore(db *sql.DB) *EpisodeStore {
	return &EpisodeStore{db: db}
}

func (s *EpisodeStore) Create(e *models.Episode) error {
	query := `INSERT INTO episodes (title, air_date, description) VALUES (?, ?, ?)`
	_, err := s.db.Exec(query, e.Title, e.AirDate, e.Description)
	return err
}

func (s *EpisodeStore) Get(id int) (*models.Episode, error) {
	query := `SELECT title, air_date, description FROM episodes WHERE id = ?`
	row := s.db.QueryRow(query, id)

	var e models.Episode
	if err := row.Scan(&e.Title, &e.AirDate, &e.Description); err != nil {
		return nil, err
	}
	return &e, nil
}

func (s *EpisodeStore) Update(e *models.Episode) error {
	query := `UPDATE episodes SET title = ?, air_date = ?, description = ? WHERE id = ?`
	_, err := s.db.Exec(query, e.Title, e.AirDate, e.Description, e.ID)
	return err
}

func (s *EpisodeStore) Delete(id int) error {
	query := `DELETE FROM episodes WHERE id = ?`
	_, err := s.db.Exec(query, id)
	return err
}
