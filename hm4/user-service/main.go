package main

import (
	"user-service/db"
	"user-service/router"
)

func main() {
	db.InitPostgresDB()
	router.InitRouter().Run(":8080")
}
