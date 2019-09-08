package auth

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	Collection *mongo.Collection
}

func NewMongo(client *mongo.Client, dbName, collectionName string) (*Mongo, error) {
	return &Mongo{
		Collection: client.Database(dbName).Collection(collectionName),
	}, nil
}

func (m *Mongo) GetUser(id string, user *User) error {
	return m.Collection.FindOne(context.Background(), bson.M{"email": id}).Decode(&user)
}

func (m *Mongo) InsertUser(user *User) (string, error) {
	now := time.Now()
	tok := newToken()
	res, err := m.Collection.InsertOne(context.Background(),
		bson.M{
			"_id":                    primitive.NewObjectID(),
			"email":                  user.Email,
			"password":               user.Password,
			"token":                  tok,
			"token_expiry":           now.Add(time.Second * defaultTokenExpirySec),
			"timezone":               user.Timezone,
			"activation_code":        user.ActivationCode,
			"activation_code_expiry": user.ActivationCodeExpiry,
			"is_active":              user.IsActive,
			"created":                now,
			"modified":               now,
		})
	if err != nil {
		return "", err
	}
	return res.InsertedID.(string), nil
}

func (m *Mongo) ActivateUser(id string) error {
	if _, err := m.Collection.UpdateOne(context.Background(), bson.M{"email": id}, bson.M{"$set": bson.M{"is_active": true}}); err != nil {
		return err
	}
	return nil
}

func (m *Mongo) VerifyUser(id, code string, user *User) error {
	return m.Collection.FindOne(
		context.Background(),
		bson.M{
			"email":           id,
			"activation_code": code,
			"activation_code_expiry": bson.M{
				"$gte": time.Now(),
			},
		},
	).Decode(&user)
}

func (m *Mongo) UpdateActiveCode(id, code string, exp time.Time) error {
	if _, err := m.Collection.UpdateOne(
		context.Background(),
		bson.M{"email": id},
		bson.M{"$set": bson.M{
			"activation_code":        code,
			"activation_code_expiry": exp,
		},
		}); err != nil {
		return err
	}
	return nil
}

func (m *Mongo) FindUserByTok(tok string, user *User) error {
	return m.Collection.FindOne(context.Background(), bson.M{"token": tok}).Decode(&user)
}

func (m *Mongo) TouchTok(id string) error {
	if _, err := m.Collection.UpdateOne(
		context.Background(),
		bson.M{"email": id},
		bson.M{
			"$set": bson.M{
				"token_expiry": time.Now().Add(time.Second * defaultTokenExpirySec),
			},
		},
	); err != nil {
		return err
	}
	return nil
}

func (m *Mongo) UpdatePassword(id, tok, pwd string) error {
	if _, err := m.Collection.UpdateOne(
		context.Background(),
		bson.M{"email": id, "token": tok},
		bson.M{"$set": bson.M{"password": pwd}},
	); err != nil {
		return err
	}
	return nil
}

func (m *Mongo) IsErrNotFound(err error) bool {
	return err == mongo.ErrNilDocument
}
