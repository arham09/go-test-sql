package mysql

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/arham09/go-test-sql/model"
	repo "github.com/arham09/go-test-sql/repository"
)

type repository struct {
	db *mongo.Collection
}

func NewRepository(dsn string, database string) (repo.Repository, error) {
	clientOptions := options.Client().ApplyURI(dsn)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	db := client.Database(database).Collection("users")

	return &repository{db}, nil
}

func (r *repository) FindByID(id string) (*model.UserModel, error) {
	user := new(model.UserModel)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err := r.db.FindOne(ctx, bson.D{{"id", id}}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) Find() ([]*model.UserModel, error) {
	users := make([]*model.UserModel, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	rows, err := r.db.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	for rows.Next(ctx) {
		user := new(model.UserModel)
		err := rows.Decode(&user)

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

	_, err := r.db.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

// Update attaches the user repository and update data based on id
func (r *repository) Update(user *model.UserModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"id": user.ID}

	update := bson.M{
		"$set": bson.M{"name": user.Name, "email": user.Email, "phone": user.Phone},
	}

	result := r.db.FindOneAndUpdate(ctx, filter, update)

	if result.Err() != nil {
		return result.Err()
	}

	return nil
}

// Delete attaches the user repository and delete data based on id
func (r *repository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Delete().SetCollation(&options.Collation{})

	_, err := r.db.DeleteOne(ctx, bson.D{{"id", id}}, opts)

	if err != nil {
		return err
	}

	return nil
}
