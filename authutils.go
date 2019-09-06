package auth

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

// TODO: implement config driven CORS.
func intercept(w http.ResponseWriter, req *http.Request) {
	log.Printf("Incoming request: %v", req)
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func respond200(w http.ResponseWriter, res interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		return err
	}

	return nil
}

func respond500(w http.ResponseWriter, err error) {
	log.Println(err.Error())
	w.WriteHeader(http.StatusInternalServerError)
}

func respond204(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func respond400(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}

func respond401(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
}

func respond404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encode(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}

	return string(b)
}

func newToken() string {
	return uuid.New().String()
}
