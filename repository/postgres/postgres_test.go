package postgres

import (
	"database/sql"
	"log"
	"testing"

	"github.com/arham09/go-test-sql/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var u = &model.UserModel{
	ID:    uuid.New().String(),
	Name:  "Arham",
	Email: "test@gmail.com",
	Phone: "08123456789",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestFindByID(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		db.Close()
	}()

	query := "SELECT id, name, email, phone FROM users WHERE id = \\$1"

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone"}).
		AddRow(u.ID, u.Name, u.Email, u.Phone)

	mock.ExpectQuery(query).WithArgs(u.ID).WillReturnRows(rows)

	user, err := repo.FindByID(u.ID)
	assert.NotNil(t, user)
	assert.NoError(t, err)
}
