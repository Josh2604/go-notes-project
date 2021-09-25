package postgres

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/Josh2604/go-notes-project/core/entities"
	"gorm.io/gorm"
)

type PostgresImplementation struct {
	DB *gorm.DB
}

func (p *PostgresImplementation) Get(ctx context.Context, noteID string) (*entities.Note, error) {
	var note entities.Note

	noteId, err := strconv.Atoi(noteID)
	if err != nil {
		return nil, err
	}

	if result := p.DB.First(&note, []int{noteId}); result.Error != nil {
		return nil, result.Error
	}

	return &note, nil
}

func (p *PostgresImplementation) GetAll(ctx context.Context) (*[]entities.Note, error) {
	var notes []entities.Note

	if result := p.DB.Where("deleted = false").Find(&notes); result.Error != nil {
		return nil, result.Error
	}
	return &notes, nil
}

func (p *PostgresImplementation) Create(ctx context.Context, note *entities.Note) error {
	if result := p.DB.Omit("id", "date_updated", "date_deleted", "deleted").Create(&note); result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *PostgresImplementation) Update(ctx context.Context, note *entities.ShortNote) error {
	noteModelToSave := transformInterface(*note)
	noteModelToSave["date_updated"] = time.Now()

	if result := p.DB.Model(entities.Note{ID: note.ID}).Updates(&noteModelToSave); result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *PostgresImplementation) Delete(ctx context.Context, noteID string) error {
	if result := p.DB.Exec("UPDATE notes SET date_deleted = NOW(), deleted= ? WHERE id = ?", true, noteID); result.Error != nil {
		return result.Error
	}

	return nil
}

func transformInterface(note entities.ShortNote) map[string]interface{} {
	var inInterface map[string]interface{}
	outInterface := make(map[string]interface{})

	inrec, _ := json.Marshal(&note)
	json.Unmarshal(inrec, &inInterface)

	for field, val := range inInterface {
		if val != nil && field != "_id" {
			outInterface[field] = val
		}
	}
	return outInterface
}
