package database

import (
	"fmt"

	"github.com/PioneerDev/setflags/config"
	"github.com/jinzhu/gorm"

	//postgres driver init
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func postgresql(db config.DB) (*gorm.DB, error) {
	serverDb, err := gorm.Open(db.DbType, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", db.Host, db.Port, db.Username, db.DbName, db.Password))
	return serverDb, err
}
