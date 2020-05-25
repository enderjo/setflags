package main

import (
	"fmt"

	"github.com/PioneerDev/setflags/config"
	"github.com/PioneerDev/setflags/model/database"
	"github.com/PioneerDev/setflags/router"
)

func main() {
	db, err := database.Init()
	if err != nil {
		panic(fmt.Errorf("Fatal to connect database: %v ", err.Error))
	}
	defer db.Close()

	r := router.InitRouter()
	r.Run(config.ServerConfig.Port)
}
