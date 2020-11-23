package repository

import "github.com/arham09/go-test-sql/model"

type Repository interface {
	Close()
	FindByID(id string) (*model.UserModel, error)
}
