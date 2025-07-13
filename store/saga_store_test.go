package store

import (
	"testing"
	"time"

	"github.com/marcusziade/dragondex/models"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestSagaStoreCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(`INSERT INTO sagas`).WithArgs("Saiyan Saga", "The Saiyans arrive on Earth.", time.Date(1989, time.April, 26, 0, 0, 0, 0, time.UTC), time.Date(1990, time.March, 7, 0, 0, 0, 0, time.UTC)).WillReturnResult(sqlmock.NewResult(1, 1))

	store := NewSagaStore(db)

	err = store.Create(&models.Saga{
		Name:        "Saiyan Saga",
		Description: "The Saiyans arrive on Earth.",
		StartDate:   time.Date(1989, time.April, 26, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(1990, time.March, 7, 0, 0, 0, 0, time.UTC),
	})
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSagaStoreGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(`SELECT`).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"name", "description", "start_date", "end_date"}).AddRow("Saiyan Saga", "The Saiyans arrive on Earth.", time.Date(1989, time.April, 26, 0, 0, 0, 0, time.UTC), time.Date(1990, time.March, 7, 0, 0, 0, 0, time.UTC)))

	store := NewSagaStore(db)

	saga, err := store.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, &models.Saga{
		Name:        "Saiyan Saga",
		Description: "The Saiyans arrive on Earth.",
		StartDate:   time.Date(1989, time.April, 26, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(1990, time.March, 7, 0, 0, 0, 0, time.UTC),
	}, saga)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSagaStoreUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(`UPDATE sagas`).WithArgs("Saiyan Saga", "The Saiyans arrive on Earth.", time.Date(1989, time.April, 26, 0, 0, 0, 0, time.UTC), time.Date(1990, time.March, 7, 0, 0, 0, 0, time.UTC), 1).WillReturnResult(sqlmock.NewResult(1, 1))

	store := NewSagaStore(db)

	err = store.Update(&models.Saga{
		ID:          1,
		Name:        "Saiyan Saga",
		Description: "The Saiyans arrive on Earth.",
		StartDate:   time.Date(1989, time.April, 26, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(1990, time.March, 7, 0, 0, 0, 0, time.UTC),
	})
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSagaStoreDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(`DELETE FROM sagas`).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	store := NewSagaStore(db)

	err = store.Delete(1)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
