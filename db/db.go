package db

import (
	"os"
	"path/filepath"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", filepath.Join(os.TempDir(), "gorm.db"))
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	db.AutoMigrate(&Permission{}, &Role{}, &RolePermission{}, &UserRole{}, &User{})
}

func DB() *gorm.DB {
	return db
}
