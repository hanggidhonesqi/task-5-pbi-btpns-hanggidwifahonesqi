package main

import (
	"github.com/hanggi/task-5-vix-btpns-hanggidwifahonesqi/database"
	"github.com/hanggi/task-5-vix-btpns-hanggidwifahonesqi/router"
)

func main() {
	database.InitDB()
	database.MigrateDB()
	r := router.RouteInit()
	r.Run(":8080")
}
