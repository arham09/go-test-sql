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

func (r *repository) Find() ([]*model.UserModel, error) {
	users := make([]*model.UserModel, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, "SELECT id, name, email, phone FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := new(model.UserModel)
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Phone,
		)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Create attaches the user repository and creating the data
func (r *repository) Create(user *model.UserModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "INSERT INTO users (id, name, email, phone) VALUES (?, ?, ?, ?)"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.ID, user.Name, user.Email, user.Phone)
	return err
}

// Update attaches the user repository and update data based on id
func (r *repository) Update(user *model.UserModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "UPDATE users SET name = ?, email = ?, phone = ? WHERE id = ?"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.Name, user.Email, user.Phone, user.ID)
	return err
}

// Delete attaches the user repository and delete data based on id
func (r *repository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "DELETE FROM users WHERE id = ?"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	return err
}
