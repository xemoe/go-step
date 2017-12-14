//
// auth_test.go
//

package unit_test

import (
	"github.com/xemoe/go-step/example/web-jwt-auth/auth"
	"reflect"
	"testing"
)

//
// @testobject auth.GetToken(user *models.User)
//
func TestGetToken(t *testing.T) {

	user := GetUser()
	token := auth.GetToken(user)

	if reflect.DeepEqual(token, nil) {
		t.Error("Failed to get token")
	}
}

//
// @testobject auth.GetJSONToken(user *models.User)
//
func TestGetJSONToken(t *testing.T) {

	user := GetUser()
	jstoken := auth.GetJSONToken(user)

	if reflect.DeepEqual(jstoken, nil) {
		t.Error("Failed to get json token")
	}
}
