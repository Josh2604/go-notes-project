package mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Josh2604/go-notes-project/core/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NoteModel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Deleted     bool               `bson:"deleted"`
	CreatedDate time.Time          `bson:"created_date"`
	UpdatedDate time.Time          `bson:"updated_date"`
	DeletedDate time.Time          `bson:"deleted_date"`
}

type NotesRepositoryImplementation struct {
	Db *mongo.Collection
}

func (r *NotesRepositoryImplementation) Get(ctx context.Context, noteID string) (*entities.Note, error) {
	note := new(NoteModel)
	mNoteID, _ := primitive.ObjectIDFromHex(noteID)
	err := r.Db.FindOne(ctx, bson.M{
		"_id": mNoteID,
	}).Decode(note)

	if err != nil {
		return nil, err
	}

	return toNoteModel(note), nil
}

func (r *NotesRepositoryImplementation) GetAll(ctx context.Context) (*[]entities.Note, error) {
	notesCursor, err := r.Db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var notes []NoteModel
	if err = notesCursor.All(ctx, &notes); err != nil {
		log.Fatal(err)
	}

	allNotes := []entities.Note{}
	for _, val := range notes {
		auxNote := toNoteModel(&val)
		allNotes = append(allNotes, *auxNote)
	}

	return &allNotes, nil
}

func (r *NotesRepositoryImplementation) Create(ctx context.Context, note *entities.Note) error {
	noteModel := toMongoNote(note)
	res, err := r.Db.InsertOne(ctx, noteModel)
	if err != nil {
		return err
	}

	note.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (r *NotesRepositoryImplementation) Update(ctx context.Context, note *entities.ShortNote) error {
	noteModelToSave := transformInterface(*note)
	mNoteID, err := primitive.ObjectIDFromHex(note.ID)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR", err)
	}

	noteModel := primitive.M{}
	for key, val := range noteModelToSave {
		noteModel[key] = val
	}

	_, err = r.Db.UpdateOne(
		ctx,
		bson.M{"_id": bson.M{"$eq": mNoteID}},
		bson.M{"$set": noteModel},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *NotesRepositoryImplementation) Delete(ctx context.Context, noteID string) error {
	mnoteID, _ := primitive.ObjectIDFromHex(noteID)
	_, err := r.Db.DeleteOne(ctx, bson.M{
		"_id": mnoteID,
	})
	if err != nil {
		return err
	}
	return nil
}

func transformInterface(note entities.ShortNote) map[string]interface{} {
	var inIterface map[string]interface{}
	outInterface := make(map[string]interface{})

	inrec, _ := json.Marshal(&note)
	json.Unmarshal(inrec, &inIterface)

	for field, val := range inIterface {
		if val != nil && field != "_id" {
			outInterface[field] = val
		}
	}

	return outInterface
}

func toMongoNote(n *entities.Note) *NoteModel {
	return &NoteModel{
		Name:        n.Name,
		Description: n.Description,
		Deleted:     n.Deleted,
		CreatedDate: n.DateCreated,
		UpdatedDate: n.DateUpdated,
		DeletedDate: n.DateDeleted,
	}
}

func toNoteModel(n *NoteModel) *entities.Note {
	return &entities.Note{
		ID:          n.ID.Hex(),
		Name:        n.Name,
		Description: n.Description,
		Deleted:     n.Deleted,
		DateCreated: n.CreatedDate,
		DateUpdated: n.UpdatedDate,
		DateDeleted: n.DeletedDate,
	}
}
