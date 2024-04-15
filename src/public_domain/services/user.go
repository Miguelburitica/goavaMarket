package services

import (
	"context"

	"github.com/Miguelburitica/goavaMarket/src/public_domain/models"
	"github.com/Miguelburitica/goavaMarket/src/public_domain/repository"
)

func GetUser(ctx context.Context, id string) (models.User, error) {
	return repository.GetUser(ctx, id)
}

func GetUsers(ctx context.Context, props models.GetUsersRequest) ([]models.User, error) {
	return repository.GetUsers(ctx, props)
}
