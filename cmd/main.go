package main

import (
	"blogapi/cmd/api"
	"blogapi/config"
	"log"

	"gorm.io/gorm"
)

func main() {
	var db *gorm.DB = config.ConnectWithDb()
	defer config.CloseDbConnection(db)
	port := config.GetEnvOrDefault("APP_PORT", "8080")
	server := api.NewAPIServer(port, db)
	error := server.Run()
	if error != nil {
		log.Fatal(error)
	}
}
