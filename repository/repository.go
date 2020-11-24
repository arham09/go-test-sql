package repository

import "github.com/arham09/go-test-sql/model"

type Repository interface {
	FindByID(id string) (*model.UserModel, error)
	Find() ([]*model.UserModel, error)
	Create(user *model.UserModel) error
	Update(user *model.UserModel) error
	Delete(id string) error
}
