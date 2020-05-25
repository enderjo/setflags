package database

import (
	"github.com/PioneerDev/setflags/config"
	"github.com/PioneerDev/setflags/model"
	"github.com/jinzhu/gorm"
)

var (
	//ServerDb 数据库连接
	ServerDb *gorm.DB
)

//Init init
func Init() (*gorm.DB, error) {
	db := config.ServerConfig.Db
	var err error
	switch db.DbType {
	case "mysql":
		ServerDb, err = mysql(db)
	case "postgresql":
		ServerDb, err = postgresql(db)
	default:
		ServerDb, err = postgresql(db)
	}
	if err == nil {
		autoMigrate()
	}
	return ServerDb, err
}

func autoMigrate() {
	ServerDb.AutoMigrate(model.Flag{},
		model.Evidence{}, model.Asset{}, model.User{}, model.Witness{})

}
