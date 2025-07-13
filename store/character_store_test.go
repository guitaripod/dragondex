package store

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/marcusziade/dragondex/models"
)

func TestCreateCharacter(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(`INSERT INTO characters`).
		WithArgs("Goku", "Saiyan", "Strong fighter", 9001).
		WillReturnResult(sqlmock.NewResult(1, 1))

	store := NewCharacterStore(db)

	err = store.Create(&models.Character{
		Name:        "Goku",
		Race:        "Saiyan",
		Description: "Strong fighter",
		PowerLevel:  9001,
	})
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetCharacter(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name", "race", "description", "power_level"}).
		AddRow("Goku", "Saiyan", "Strong fighter", 9001)

	mock.ExpectQuery(`SELECT name, race, description, power_level FROM characters WHERE id = ?`).
		WithArgs(1).
		WillReturnRows(rows)

	store := NewCharacterStore(db)

	character, err := store.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, "Goku", character.Name)
	assert.Equal(t, "Saiyan", character.Race)
	assert.Equal(t, "Strong fighter", character.Description)
	assert.Equal(t, 9001, character.PowerLevel)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateCharacter(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(`UPDATE characters`).
		WithArgs("Vegeta", "Saiyan", "Prince of all Saiyans", 8500, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	store := NewCharacterStore(db)

	err = store.Update(&models.Character{
		ID:          1,
		Name:        "Vegeta",
		Race:        "Saiyan",
		Description: "Prince of all Saiyans",
		PowerLevel:  8500,
	})
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteCharacter(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(`DELETE FROM characters WHERE id = ?`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	store := NewCharacterStore(db)

	err = store.Delete(1)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
