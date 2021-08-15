package user

import (
	"reflect"
	"testing"

	"github.com/evoke365/webserver/controller/internal/responder"
	"github.com/evoke365/webserver/models"
	"github.com/evoke365/webserver/restapi/operations/user"
	"github.com/evoke365/webserver/store/fakedb"
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
