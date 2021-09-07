package postgres

import (
	"context"
	"database/sql"

	"github.com/Josh2604/go-notes-project/core/entities"
)

type PostgresImplementation struct {
	DB *sql.DB
}

func (p *PostgresImplementation) Get(ctx context.Context, noteID string) (*entities.Note, error) {
	return &entities.Note{}, nil
}

func (p *PostgresImplementation) GetAll(ctx context.Context) (*[]entities.Note, error) {
	return nil, nil
}
func (p *PostgresImplementation) Create(ctx context.Context, note *entities.Note) error {
	return nil
}
func (p *PostgresImplementation) Update(ctx context.Context, note *entities.ShortNote) error {
	return nil
}
func (p *PostgresImplementation) Delete(ctx context.Context, noteID string) error {
	return nil
}
