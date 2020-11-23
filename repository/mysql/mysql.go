package mysql

import (
	"database/sql"

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
