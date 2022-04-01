package mongo

import (
	"context"

	"github.com/Josh2604/go-notes-project/core/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository -
type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.User) error
	GetUser(ctx context.Context, username, password string) (*entities.User, error)
}

type UserModel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}

type UserRepositoryImplementation struct {
	DB *mongo.Collection
}

func (r UserRepositoryImplementation) CreateUser(ctx context.Context, user *entities.User) error {
	model := toMongoUser(user)
	res, err := r.DB.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	user.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (r UserRepositoryImplementation) GetUser(ctx context.Context, username, password string) (*entities.User, error) {
	user := new(UserModel)
	err := r.DB.FindOne(ctx, bson.M{
		"username": username,
		"password": password,
	}).Decode(user)

	if err != nil {
		return nil, err
	}

	return toModel(user), nil
}

func toMongoUser(u *entities.User) *UserModel {
	return &UserModel{
		Username: u.Username,
		Password: u.Password,
	}
}

func toModel(u *UserModel) *entities.User {
	return &entities.User{
		ID:       u.ID.Hex(),
		Username: u.Username,
		Password: u.Password,
	}
}
