package repository

type Repository interface {
	Close()
	// FindByID(id string) (*model.UserModel, error)
}
