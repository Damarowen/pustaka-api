package main

import (
	"pustaka-api/config"
	"pustaka-api/routes"
)



func main() {

	db, _ := config.ConnectDatabase()

	r := Routes.SetupRouter(db)

	r.Run("127.0.0.1:9090")

}
