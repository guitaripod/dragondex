package store

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/marcusziade/dragondex/models"
	"github.com/stretchr/testify/assert"
)

func TestAbilityStoreCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(`INSERT INTO abilities`).WithArgs("Kamehameha", "Energy wave", 5000).WillReturnResult(sqlmock.NewResult(1, 1))

	store := NewAbilityStore(db)

	err = store.Create(&models.Ability{
		Name:          "Kamehameha",
		Description:   "Energy wave",
		PowerRequired: 5000,
	})
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAbilityStoreGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(`SELECT`).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"name", "description", "power_required"}).AddRow("Kamehameha", "Energy wave", 5000))

	store := NewAbilityStore(db)

	ability, err := store.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, &models.Ability{
		Name:          "Kamehameha",
		Description:   "Energy wave",
		PowerRequired: 5000,
	}, ability)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAbilityStoreUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(`UPDATE abilities`).WithArgs("Kamehameha", "Energy wave", 6000, 1).WillReturnResult(sqlmock.NewResult(1, 1))

	store := NewAbilityStore(db)

	err = store.Update(&models.Ability{
		ID:            1,
		Name:          "Kamehameha",
		Description:   "Energy wave",
		PowerRequired: 6000,
	})
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAbilityStoreDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(`DELETE FROM abilities`).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	store := NewAbilityStore(db)

	err = store.Delete(1)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
