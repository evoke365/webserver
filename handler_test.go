package auth

import (
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
	router.GET("/user/forget/:id", testHandler.Forget)

	req, _ := http.NewRequest("GET", "/user/forget/123", nil)
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
