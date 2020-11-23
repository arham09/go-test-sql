package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/arham09/go-test-sql/model"
	repo "github.com/arham09/go-test-sql/repository"

	_ "github.com/go-sql-driver/mysql"
)

type repository struct {
	db *sql.DB
}

func NewRepository(dialect string, dsn string, idleConn int, maxConn int) (repo.Repository, error) {
	db, err := sql.Open(dialect, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(idleConn)
	db.SetMaxOpenConns(maxConn)

	return &repository{db}, nil
}

func (r *repository) Close() {
	r.db.Close()
}

func (r *repository) FindByID(id string) (*model.UserModel, error) {
	user := new(model.UserModel)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err := r.db.QueryRowContext(ctx, "SELECT id, name, email, phone FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Phone)
	if err != nil {
		return nil, err
	}
	return user, nil
}
