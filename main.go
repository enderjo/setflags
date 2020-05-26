package main

import (
	"fmt"

	"github.com/PioneerDev/setflags/config"
	"github.com/PioneerDev/setflags/model/database"
	"github.com/PioneerDev/setflags/router"
)

// @title SetFlags API
// @version 0.0.1
// @description SetFlags API
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	db, err := database.Init()
	if err != nil {
		panic(fmt.Errorf("Fatal to connect database: %v ", err.Error))
	}
	defer db.Close()

	r := router.InitRouter()
	r.Run(config.ServerConfig.Port)
}
