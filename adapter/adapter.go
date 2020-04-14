package adapter

import (
	"github.com/CyrivlClth/rbac/db"
	"github.com/CyrivlClth/rbac/rbac"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func AccountAdapter(username string) *rbac.Account {
	conn := db.DB()
	user := &db.User{}
	if err := conn.Where(&db.User{Username: username}).First(user).Error; err != nil {
		panic(err)
	}

	return &rbac.Account{}
}
