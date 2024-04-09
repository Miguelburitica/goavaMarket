package adapters

import (
	"context"

	"github.com/Miguelburitica/goavaMarket/src/public_domain/models"
)

type RepositoryUserActions interface {
	GetUser() (models.User, error)
	CreateUser() error
	UpdateUser() error
	DeleteUser() error
}

type Repository interface {
	Connect() error
	Disconnect() error
	RepositoryUserActions
}

var implementation Repository

func SetRepository(repo Repository) {
	implementation = repo
}

func GetUser(ctx context.Context, id string) (models.User, error) {
	return implementation.GetUser()
}

func CreateUser() error {
	return implementation.CreateUser()
}

func UpdateUser() error {
	return implementation.UpdateUser()
}
