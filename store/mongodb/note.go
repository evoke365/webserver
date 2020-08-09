package mongodb

import (
	"context"
	"time"

	"github.com/evoke365/webserver/store/data"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func (m *MongoDB) GetNoteById(ctx context.Context, objID primitive.ObjectID) (*data.Note, error) {
	note := &data.Note{}
	if err := m.note.FindOne(ctx, bson.M{"_id": objID}).Decode(&note); err != nil {
		return nil, err
	}

	return note, nil
}

func (m *MongoDB) GetNotes(ctx context.Context, userID string) ([]*data.Note, error) {
	var result []*data.Note

	cur, err := m.note.Find(ctx, bson.M{"userId": userID})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		// To decode into a struct, use cursor.Decode()
		note := &data.Note{}
		if err := cur.Decode(note); err != nil {
			return nil, err
		}
		result = append(result, note)
	}
	return result, nil
}

// InsertNote implements query logic that inserts a note.
func (m *MongoDB) InsertNote(ctx context.Context, notes ...interface{}) error {
	if len(notes) == 0 {
		return nil
	}
	if len(notes) == 1 {
		if _, err := m.note.InsertOne(ctx, notes[0]); err != nil {
			return err
		}
		return nil
	}
	if _, err := m.note.InsertMany(ctx, notes); err != nil {
		return err
	}
	return nil
}

// DeleteNote implements query logic that deletes a note.
func (m *MongoDB) DeleteNote(ctx context.Context, userID, noteID string) error {
	objID, err := primitive.ObjectIDFromHex(noteID)
	if err != nil {
		return err
	}
	if _, err := m.note.DeleteOne(ctx, bson.M{"_id": objID, "userId": userID}); err != nil {
		return err
	}
	return nil
}

// UpdateNote implements query logic that updates a note.
func (m *MongoDB) UpdateNote(ctx context.Context, userID, noteID string, newData map[string]interface{}) (*data.Note, error) {
	objID, err := primitive.ObjectIDFromHex(noteID)
	if err != nil {
		return nil, err
	}
	note := &data.Note{}
	opts := options.FindOneAndUpdate()
	opts.SetReturnDocument(options.After)
	if err := m.note.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objID, "userId": userID},
		bson.M{"$set": bson.M(newData)},
		opts,
	).Decode(note); err != nil {
		return nil, err
	}
	return note, nil
}

func (m *MongoDB) GetUpdatedNotes(ctx context.Context, fromDate time.Time, toDate time.Time) ([]*data.Note, error) {
	var notes []*data.Note
	cur, err := m.note.Find(
		ctx,
		bson.M{"modified": bson.M{"$gte": fromDate, "$lt": toDate}},
	)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		// To decode into a struct, use cursor.Decode()
		note := &data.Note{}
		if err := cur.Decode(note); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}
