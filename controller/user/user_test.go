package user

import (
	"reflect"
	"testing"

	"github.com/evoke365/webserver/controller/internal/responder"
	"github.com/evoke365/webserver/models"
	"github.com/evoke365/webserver/restapi/operations/user"
	"github.com/evoke365/webserver/store/fakedb"
	"github.com/evoke365/webserver/store/data"
)

func TestSignupSuccess(t *testing.T) {
	c := NewController(fakedb.NewFakeDB(), nil, DefaultConfig())

	res := c.Signup(&user.SignupUserParams{
		Body: &models.UserSignupRequests{
			Email:    "test@evoke365.net",
			Password: "mock",
			Timezone: 1,
		},
	})

	exp := responder.DefaultOK()
	if !reflect.DeepEqual(exp, res) {
		t.Fatalf("expected %+v but got %+v", exp, res)
	}
}

func TestFindUserNotFound(t *testing.T) {
	c := NewController(fakedb.NewFakeDB(), nil, DefaultConfig())

	res := c.FindUser(&user.FindUserParams{
		ID: "6d04eea2-3d4d-4593-886f-65bde3de7889",
	})
	exp := responder.DefaultNoContent()
	if !reflect.DeepEqual(exp, res) {
		t.Fatalf("expected %+v but got %+v", exp, res)
	}
}

func TestFindUser(t *testing.T) {
	// setup
	db := fakedb.NewFakeDB()
	mockUser := &data.User{
		Email:    "test@test.com",
		Password: "tobeencrypted",
	}
	id, err := db.InsertUser(mockUser)
	if err != nil {
		t.Fatal(err)
	}

	c := NewController(db, nil, DefaultConfig())

	res := c.FindUser(&user.FindUserParams{
		ID: id,
	})

	exp := responder.DefaultOK()
	if !reflect.DeepEqual(exp, res) {
		t.Fatalf("expected %+v but got %+v", exp, res)
	}
}