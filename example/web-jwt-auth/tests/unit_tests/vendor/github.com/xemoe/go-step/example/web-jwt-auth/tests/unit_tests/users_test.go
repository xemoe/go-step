//
// users_test.go
//

package unit_test

import (
	"github.com/xemoe/go-step/example/web-jwt-auth/models"
	"reflect"
	"testing"
)

//
// @testobject models.NewUserManager(DB *models.DB)
//
func TestNewUserManager(t *testing.T) {
	ClearDB()
	DB = models.NewSqliteDB(dbname)
	MGR, _ = models.NewUserManager(DB)

	if MGR == nil {
		t.Error("Failed to create user manager")
	}
}

//
// @testobject models.AddUser(username string, password string)
//
func TestAddUser(t *testing.T) {
	ClearDB()
	DB = models.NewSqliteDB(dbname)
	MGR, _ = models.NewUserManager(DB)

	username := "myname"
	password := "password"

	user := MGR.AddUser(username, password)

	if MGR == nil {
		t.Error("Failed to create user manager")
	}
	if user.Username != username {
		t.Error("Returned user from Adduser() should equals input")
	}
	if user.Password == password {
		t.Error("Returned password from Adduser() should be hash string")
	}
}

func TestHasUser(t *testing.T) {
	ClearDB()
	DB = models.NewSqliteDB(dbname)
	MGR, _ = models.NewUserManager(DB)

	username := "myname"
	password := "password"

	hasUserResult1 := MGR.HasUser(username)

	if hasUserResult1 != false {
		t.Error("Has empty user should return false")
	}

	MGR.AddUser(username, password)
	hasUserResult2 := MGR.HasUser(username)

	if hasUserResult2 != true {
		t.Error("Has created user should return true")
	}
}

func TestFindUser(t *testing.T) {
	ClearDB()
	DB = models.NewSqliteDB(dbname)
	MGR, _ = models.NewUserManager(DB)

	username := "myname"
	password := "password"

	findUserResult1 := MGR.FindUser(username)
	emptyUser := &models.User{}

	if !reflect.DeepEqual(findUserResult1, emptyUser) {
		t.Error("Find empty user should return empty")
	}

	MGR.AddUser(username, password)
	findUserResult2 := MGR.FindUser(username)

	if reflect.DeepEqual(findUserResult2, emptyUser) {
		t.Error("Find created user should return user")
	}
}

func TestFindUserByUUID(t *testing.T) {
	ClearDB()
	DB = models.NewSqliteDB(dbname)
	MGR, _ = models.NewUserManager(DB)

	username := "myname"
	password := "password"

	user := MGR.AddUser(username, password)
	uuid := user.UUID
	findUserResult := MGR.FindUserByUUID(uuid)

	if user.Username != findUserResult.Username {
		t.Error("Returned user from FindUserByUUID() should equals added user")
	}
}

func TestHashPassword(t *testing.T) {
	ClearDB()
	DB = models.NewSqliteDB(dbname)
	MGR, _ = models.NewUserManager(DB)

	username := "myname"
	password := "password"

	hashPassword := MGR.HashPassword(username, password)

	if hashPassword == "" {
		t.Error("Hash password should return string")
	}
}

func TestCheckPassword(t *testing.T) {
	ClearDB()
	DB = models.NewSqliteDB(dbname)
	MGR, _ = models.NewUserManager(DB)

	username := "myname"
	password := "password"

	hashPassword := MGR.HashPassword(username, password)
	checkResult := MGR.CheckPassword(hashPassword, password)

	if checkResult != true {
		t.Error("Check password should return true")
	}
}
