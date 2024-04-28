package repository

import (
	"context"

	"github.com/Miguelburitica/goavaMarket/src/public_domain/models"
)

type RepositoryUserActions interface {
	GetUser(ctx context.Context, id string) (models.User, error)
	GetUsers(ctx context.Context, props models.GetUsersRequest) ([]models.User, error)
	// CreateUser(ctx context.Context, user models.User) error
	// UpdateUser(ctx context.Context, userInfo models.UserToUpdate) error
	// DeleteUser() error
}

type Repository interface {
	Disconnect() error
	RepositoryUserActions
}

var implementation Repository

func SetRepository(repo Repository) {
	implementation = repo
}

func GetUser(ctx context.Context, id string) (models.User, error) {
	return implementation.GetUser(ctx, id)
}

func GetUsers(ctx context.Context, props models.GetUsersRequest) ([]models.User, error) {
	return implementation.GetUsers(ctx, props)
}

// func CreateUser(ctx context.Context, user models.User) error {
// 	return implementation.CreateUser(ctx, user)
// }

// func UpdateUser(ctx context.Context, user models.UserToUpdate) error {
// 	return implementation.UpdateUser(ctx, user)
// }
