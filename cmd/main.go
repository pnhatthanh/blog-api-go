package main

import (
	"blogapi/cmd/api"
	"blogapi/config"
	"log"

	"gorm.io/gorm"
)

func main() {
	var db *gorm.DB=config.ConnectWithDb()
	defer config.CloseDbConnection(db)
	server := api.NewAPIServer(":8080",db)
	error:=server.Run()
	if error!=nil{
		log.Fatal(error)
	}
}