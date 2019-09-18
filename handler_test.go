package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

// TODO: implement unit test
func TestHealth(t *testing.T) {
	c := Config{1, 2}
	testHandler := NewHandler(c, &noopModel{}, &noopCallback{})

	router := httprouter.New()
	router.GET("/health", testHandler.Health)

	req, _ := http.NewRequest("GET", "/health", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatal("incorrect status")
	}
}

type noopModel struct {
	Model
}

type noopCallback struct {
}

func (cb *noopCallback) OnSignup(user *User) error {
	return nil
}

func (cb *noopCallback) OnVerify(to string, data interface{}) error {
	return nil
}
