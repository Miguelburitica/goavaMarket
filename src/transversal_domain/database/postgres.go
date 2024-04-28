package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/Miguelburitica/goavaMarket/src/public_domain/models"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type PostgresRepository struct {
	db *sql.DB
}

func (p *PostgresRepository) Disconnect() error {
	return p.db.Close()
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("libsql", url)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db: db}, err
}

func (p *PostgresRepository) GetUser(ctx context.Context, id string) (models.User, error) {
	var user models.User
	err := p.db.QueryRowContext(ctx, "SELECT id, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Email)
	if err != nil {
		log.Printf("Error getting the user: %v", err)
		return user, err
	}

	return user, nil
}

func (p *PostgresRepository) GetUsers(ctx context.Context, props models.GetUsersRequest) ([]models.User, error) {
	var users []models.User
	rows, err := p.db.QueryContext(ctx, "SELECT id, email FROM users")
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Email)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (p *PostgresRepository) CreateUser(ctx context.Context, user models.User) error {
	_, err := p.db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.ID, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresRepository) UpdateUser(ctx context.Context, userInfo models.UserToUpdate) error {
	_, err := p.db.ExecContext(ctx, "UPDATE users SET email = $1, password = $2 WHERE id = $3", userInfo.Email, userInfo.Password, userInfo.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresRepository) DeleteUser(ctx context.Context, id string) error {
	_, err := p.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
