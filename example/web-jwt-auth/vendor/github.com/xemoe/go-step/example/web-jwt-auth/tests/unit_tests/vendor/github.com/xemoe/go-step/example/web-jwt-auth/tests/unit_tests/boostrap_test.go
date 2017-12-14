//
// boostrap_test.go
//

package unit_test

import (
	"fmt"
	"github.com/xemoe/go-step/example/web-jwt-auth/models"
	"os"
	"testing"
)

var (
	MGR             *models.UserManager
	DB              *models.DB
	dbname          string = "testing-web-jwt-auth.db3"
	defaultUsername string = "myname"
	defaultPassword string = "password"
)

func TestMain(m *testing.M) {
	code := m.Run()
	//
	// Remove sqlite test db file after complete test
	//
	ClearDB()
	os.Exit(code)
}

func ClearDB() {
	os.Remove(fmt.Sprintf("%s", dbname))
}

func GetUser() (user *models.User) {
	ClearDB()
	DB = models.NewSqliteDB(dbname)
	MGR, _ = models.NewUserManager(DB)
	username := defaultUsername
	password := defaultPassword
	user = MGR.AddUser(username, password)
	return
}
