package main

import (
	"blogapi/cmd/api"
	"log"
)

func main() {
	server := api.NewAPIServer(":8080",nil)
	error:=server.Run()
	if error!=nil{
		log.Fatal(error)
	}
}