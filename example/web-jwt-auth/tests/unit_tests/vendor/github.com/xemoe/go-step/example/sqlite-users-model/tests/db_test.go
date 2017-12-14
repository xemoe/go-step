//
// db_test.go
//

package models_test

import (
	"github.com/xemoe/go-step/example/sqlite-users-model/models"
	"testing"
)

func TestNewSqliteDB_ShouldReturnExpectedValue(t *testing.T) {
	DB = models.NewSqliteDB(dbname)
	if DB == nil {
		t.Error("Failed to new sqlite3 db")
	}
}

//
// Copy from https://github.com/jinzhu/gorm/blob/master/main_test.go
//
func TestSetAndGet(t *testing.T) {
	DB = models.NewSqliteDB(dbname)
	if value, ok := DB.Set("hello", "world").Get("hello"); !ok {
		t.Errorf("Should be able to get setting after set")
	} else {
		if value.(string) != "world" {
			t.Errorf("Setted value should not be changed")
		}
	}

	if _, ok := DB.Get("non_existing"); ok {
		t.Errorf("Get non existing key should return error")
	}
}
