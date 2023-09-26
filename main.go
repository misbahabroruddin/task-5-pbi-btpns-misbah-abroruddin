package main

import (
	db "github.com/misbahabroruddin/task-5-pbi-btpns-misbah-abroruddin/database"

	"github.com/misbahabroruddin/task-5-pbi-btpns-misbah-abroruddin/controllers"
	"github.com/misbahabroruddin/task-5-pbi-btpns-misbah-abroruddin/routes"
)

func main()  {
	// ping := controllers.Ping
	db.Init()
	r := routes.Routes()

	r.GET("/ping", controllers.Ping)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}