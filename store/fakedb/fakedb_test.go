package fakedb

import (
	"reflect"
	"testing"

	"github.com/evoke365/webserver/store/data"
)

func TestInsertGetUser(t *testing.T) {
	ins := NewFakeDB()
	mockUser := &data.User{
		Email:    "test@test.com",
		Password: "tobeencrypted",
	}

	id, err := ins.InsertUser(mockUser)
	if err != nil {
		t.Fatal(err)
	}

	mockUser2 := &data.User{}
	if err := ins.GetUser(id, mockUser2); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(mockUser, mockUser2) {
		t.Fatalf("results do not match! expected %+v but got %+v", mockUser, mockUser2)
	}
}

func TestActivateUser(t *testing.T) {
	ins := NewFakeDB()
	mockUser := &data.User{
		Email:    "test@test.com",
		Password: "tobeencrypted",
	}
	id, err := ins.InsertUser(mockUser)
	if err != nil {
		t.Fatal(err)
	}

	if err := ins.ActivateUser(id); err != nil {
		t.Fatal(err)
	}

	mockUser2 := &data.User{}
	if err := ins.GetUser(id, mockUser2); err != nil {
		t.Fatal(err)
	}

	if !mockUser2.IsActive {
		t.Fatal("expected user to be activated")
	}
}
