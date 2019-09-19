package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/julienschmidt/httprouter"
)

// TODO: implement unit test
func TestHealth(t *testing.T) {
	c := Config{1, 2}
	testHandler := NewHandler(c, &model200{}, &noopCallback{})

	router := httprouter.New()
	router.GET("/health", testHandler.Health)

	req, _ := http.NewRequest("GET", "/health", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatal("non 200 status")
	}
}

func TestUser200(t *testing.T) {
	c := Config{1, 2}
	testHandler := NewHandler(c, &model200{}, &noopCallback{})

	router := httprouter.New()
	router.GET("/user/find/:id", testHandler.User)

	req, _ := http.NewRequest("GET", "/user/find/123", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatal("non 200 status")
	}
}

func TestForget200(t *testing.T) {
	c := Config{1, 2}
	testHandler := NewHandler(c, &model200{}, &noopCallback{})

	router := httprouter.New()
	router.PUT("/user/forget/:id", testHandler.Forget)

	req, _ := http.NewRequest("PUT", "/user/forget/123", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatal("non 200 status")
	}
}

func TestSetPassword200(t *testing.T) {
	c := Config{1, 2}
	testHandler := NewHandler(c, &model200{}, &noopCallback{})

	router := httprouter.New()
	router.PUT("/user/forget", testHandler.SetPassword)

	obj := struct {
		Email    string `json:"email"`
		Token    string `json:"token"`
		Password string `json:"password"`
	}{
		Email:    "test@test.com",
		Token:    "test-token",
		Password: "12345",
	}
	body, err := json.Marshal(obj)
	if err != nil {
		t.Fatal(err)
	}

	req, _ := http.NewRequest("PUT", "/user/forget", bytes.NewReader(body))
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatal("non 200 status")
	}
}

func TestUser500(t *testing.T) {
	c := Config{1, 2}
	testHandler := NewHandler(c, &model500{}, &noopCallback{})

	router := httprouter.New()
	router.GET("/user/find/:id", testHandler.User)

	req, _ := http.NewRequest("GET", "/user/find/123", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Fatal("non 500 status")
	}
}

type model200 struct {
	Model
}

func (m *model200) GetUser(id string, user *User) error {
	user = &User{
		Email:    "test@test.com",
		Password: "test-password",
	}
	return nil
}

func (m *model200) UpdateActiveCode(id, code string, exp time.Time) (*User, error) {
	user := &User{
		Email:          "test@test.com",
		Password:       "test-password",
		ActivationCode: "123",
	}
	return user, nil
}

func (m *model200) UpdatePassword(id, tok, pwd string) error {
	return nil
}

type model500 struct {
	Model
}

var errNotFound = errors.New("not found")
var errMock = errors.New("mock error")

func (m *model500) GetUser(id string, user *User) error {
	return errMock
}

func (m model500) IsErrNotFound(err error) bool {
	return err == errNotFound
}

type noopCallback struct {
}

func (cb *noopCallback) OnSignup(user *User) error {
	return nil
}

func (cb *noopCallback) OnVerify(to string, data interface{}) error {
	return nil
}
