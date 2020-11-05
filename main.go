package main

import (
	"log"

	"github.com/nguyenbt456/demo-clean-architecture/app"
	"github.com/nguyenbt456/demo-clean-architecture/database"
)

func main() {
	database.ConnectDB()

	router := app.InitRouter()
	if err := router.Run(":9000"); err != nil {
		log.Fatalln("Error:", err)
	}
}
