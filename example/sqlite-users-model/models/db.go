//
// db.go
//

package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"path/filepath"
)

//
// DB Abstraction
//
type DB struct {
	*gorm.DB
}

const (
	DBFILE_EXT = "db3"
)

//
// New Connection
//
func NewSqliteDB(dbname string) *DB {

	dbfile := ""
	ext := filepath.Ext(dbname)

	//
	// Manipulate sqlite3 db filename with extension
	//
	if ext != DBFILE_EXT {
		dbfile = fmt.Sprintf("%s.%s", dbname, DBFILE_EXT)
	} else {
		dbfile = fmt.Sprintf("%s", dbname)
	}

	db, err := gorm.Open("sqlite3", dbfile)
	if err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	return &DB{db}
}
