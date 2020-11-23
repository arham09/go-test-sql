package repository

import "github.com/arham09/go-test-sql/model"

type Repository interface {
	FindByID(id string) (*model.UserModel, error)
}
