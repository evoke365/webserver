package fakedb

import (
	"reflect"
	"testing"
	"time"

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

func TestGetUserNotFound(t *testing.T) {
	ins := NewFakeDB()
	mock := &data.User{}
	err := ins.GetUser("id", mock)
	if !ins.IsErrNotFound(err) {
		t.Fatalf("expected err %v but got %v", ErrNoDocument, err)
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

func TestFindUserByTok(t *testing.T) {
	ins := NewFakeDB()
	mockUser := &data.User{
		Email:    "test@test.com",
		Password: "tobeencrypted",
		Token:    "tok",
	}

	_, err := ins.InsertUser(mockUser)
	if err != nil {
		t.Fatal(err)
	}

	mockUser2 := &data.User{}
	if err := ins.FindUserByTok("tok", mockUser2); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(mockUser, mockUser2) {
		t.Fatalf("results do not match! expected %+v but got %+v", mockUser, mockUser2)
	}
}

func TestFindUserByTokNotFound(t *testing.T) {
	ins := NewFakeDB()
	mockUser2 := &data.User{}
	err := ins.FindUserByTok("tok", mockUser2)
	if !ins.IsErrNotFound(err) {
		t.Fatalf("expected err %v but got %v", ErrNoDocument, err)
	}
}

func TestUpdateActiveCode(t *testing.T) {
	ins := NewFakeDB()
	now := time.Now()
	mockUser := &data.User{
		Email:                "test@test.com",
		Password:             "tobeencrypted",
		Token:                "tok",
		ActivationCode:       "code1",
		ActivationCodeExpiry: now,
	}

	id, err := ins.InsertUser(mockUser)
	if err != nil {
		t.Fatal(err)
	}

	now2 := time.Now().Add(time.Second * 60)
	user, err := ins.UpdateActiveCode(id, "code2", now2)
	if err != nil {
		t.Fatal(err)
	}

	if user.ActivationCode != "code2" {
		t.Fatalf("expected code2 but got %s", user.ActivationCode)
	}

	if !reflect.DeepEqual(user.ActivationCodeExpiry, now2) {
		t.Fatal("exp mismatch")
	}
}

func TestTouchTok(t *testing.T) {
	ins := NewFakeDB()
	now := time.Now().Add(time.Minute * -1)
	mockUser := &data.User{
		Email:       "test@test.com",
		Password:    "tobeencrypted",
		Token:       "tok",
		TokenExpiry: now,
	}

	id, err := ins.InsertUser(mockUser)
	if err != nil {
		t.Fatal(err)
	}

	if err := ins.TouchTok(id); err != nil {
		t.Fatal(err)
	}

	if !ins.isTokenValid(id) {
		t.Fatalf("expected valid token")
	}
}

func TestUpdatePassword(t *testing.T) {
	ins := NewFakeDB()
	mockUser := &data.User{
		Email:    "test@test.com",
		Password: "pass1",
		Token:    "tok",
	}

	id, err := ins.InsertUser(mockUser)
	if err != nil {
		t.Fatal(err)
	}

	if err := ins.UpdatePassword(id, "tok", "pass2"); err != nil {
		t.Fatal(err)
	}

	mockUser2 := &data.User{}
	if err := ins.GetUser(id, mockUser2); err != nil {
		t.Fatal(err)
	}

	if mockUser2.Password != "pass2" {
		t.Fatal("password mismatch")
	}
}
