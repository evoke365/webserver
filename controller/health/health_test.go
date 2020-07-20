package health

import (
	"reflect"
	"testing"
)

func TestHealthz(t *testing.T) {
	testController := NewController()
	responder := testController.Healthz()
	expected := &HealthzOK{}
	if !reflect.DeepEqual(responder, expected) {
		t.Fatalf("results do not match! expected %+v but got %+v", expected, responder)
	}
}
