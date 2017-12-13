//
// db_test.go
//

package models_test

import (
	"fmt"
	"github.com/xemoe/go-step/example/sqlite-users-model/models"
	"os"
	"testing"
)

var (
	dbname string = "testing-sqlite-users-model.db3"
	MGR    *models.UserManager
	DB     *models.DB
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
