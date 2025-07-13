package store

import (
	"testing"
	"time"

	"github.com/marcusziade/dragondex/models"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestEpisodeStoreCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(`INSERT INTO episodes`).WithArgs("The Arrival of Raditz", time.Date(1989, time.April, 26, 0, 0, 0, 0, time.UTC), "Raditz arrives on Earth.").WillReturnResult(sqlmock.NewResult(1, 1))

	store := NewEpisodeStore(db)

	err = store.Create(&models.Episode{
		Title:       "The Arrival of Raditz",
		AirDate:     time.Date(1989, time.April, 26, 0, 0, 0, 0, time.UTC),
		Description: "Raditz arrives on Earth.",
	})
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEpisodeStoreGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(`SELECT`).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"title", "air_date", "description"}).AddRow("The Arrival of Raditz", time.Date(1989, time.April, 26, 0, 0, 0, 0, time.UTC), "Raditz arrives on Earth."))

	store := NewEpisodeStore(db)

	episode, err := store.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, &models.Episode{
		Title:       "The Arrival of Raditz",
		AirDate:     time.Date(1989, time.April, 26, 0, 0, 0, 0, time.UTC),
		Description: "Raditz arrives on Earth.",
	}, episode)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEpisodeStoreUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(`UPDATE episodes`).WithArgs("The Arrival of Raditz", time.Date(1989, time.April, 26, 0, 0, 0, 0, time.UTC), "Raditz arrives on Earth.", 1).WillReturnResult(sqlmock.NewResult(1, 1))

	store := NewEpisodeStore(db)

	err = store.Update(&models.Episode{
		ID:          1,
		Title:       "The Arrival of Raditz",
		AirDate:     time.Date(1989, time.April, 26, 0, 0, 0, 0, time.UTC),
		Description: "Raditz arrives on Earth.",
	})
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEpisodeStoreDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(`DELETE FROM episodes`).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	store := NewEpisodeStore(db)

	err = store.Delete(1)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
