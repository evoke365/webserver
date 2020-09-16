package bus

import (
	"bytes"
	"testing"
)

func TestEncodeDecodePayload(t *testing.T) {
	data := []byte("123")

	s := encodePayload(data)
	if len(s) == 0 {
		t.Fatal("expected encoded string")
	}

	b, err := decodePayload(s)
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Compare(data, b) != 0 {
		t.Fatalf("expected %+v but got %+v", data, b)
	}
}
