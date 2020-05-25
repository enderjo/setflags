package database

import (
	"fmt"

	"github.com/PioneerDev/setflags/config"
	"github.com/jinzhu/gorm"

	//mysql driver init
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func mysql(db config.DB) (*gorm.DB, error) {
	connStr := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=true", db.Username, db.Password, db.Host, db.Port, db.DbName)
	serverDb, err := gorm.Open(db.DbType, connStr)
	//TODO: param setting

	return serverDb, err
}
