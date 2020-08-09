package mongodb

import (
	"context"
	"time"

	"github.com/evoke365/webserver/store/data"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func (m *MongoDB) GetUser(id string, user *data.User) error {
	return m.user.FindOne(context.Background(), bson.M{"email": id}).Decode(&user)
}

func (m *MongoDB) InsertUser(user *data.User) (string, error) {
	now := time.Now()
	tok := newToken()
	id := primitive.NewObjectID()
	if _, err := m.user.InsertOne(context.Background(),
		bson.M{
			"_id":                    id,
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
		}); err != nil {
		return "", err
	}
	return id.String(), nil
}

func (m *MongoDB) ActivateUser(id string) error {
	if _, err := m.user.UpdateOne(context.Background(), bson.M{"email": id}, bson.M{"$set": bson.M{"is_active": true}}); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) VerifyUser(id, code string, user *data.User) error {
	return m.user.FindOne(
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

func (m *MongoDB) UpdateActiveCode(id, code string, exp time.Time) (*data.User, error) {
	user := &data.User{}
	opts := options.FindOneAndUpdate()
	opts.SetReturnDocument(options.After)
	if err := m.user.FindOneAndUpdate(
		context.Background(),
		bson.M{"email": id},
		bson.M{"$set": bson.M{
			"activation_code":        code,
			"activation_code_expiry": exp,
		}},
		opts,
	).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

// FindUserByTok implements query logic that retrieves user info by authentciation token.
func (m *MongoDB) FindUserByTok(tok string, user *data.User) error {
	return m.user.FindOne(context.Background(), bson.M{"token": tok}).Decode(&user)
}

// TouchTok implements query logic that extends a user's authentciation token expiry.
func (m *MongoDB) TouchTok(id string) error {
	if _, err := m.user.UpdateOne(
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

func (m *MongoDB) UpdatePassword(id, tok, pwd string) error {
	if _, err := m.user.UpdateOne(
		context.Background(),
		bson.M{"email": id, "token": tok},
		bson.M{"$set": bson.M{"password": pwd}},
	); err != nil {
		return err
	}
	return nil
}
