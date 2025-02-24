package main

import (
	"log"
	"sinappsebackend/app"
	"sinappsebackend/routes"
)

func main() {
	err := app.LoadConfig()
	if err != nil {
		log.Fatalln("A fatal error occurred when trying to load config:", err.Error())
		return
	}
	log.Println("Successfully loaded environment variables!")

	err = app.InitDatabase()
	if err != nil {
		log.Fatalln("A fatal error occurred when trying to initiate database:", err.Error())
		return
	}
	log.Println("Successfully connected to the database!")

	err = app.InitHttpServer()
	if err != nil {
		log.Fatalln("A fatal error occurred when trying to init the http server:", err.Error())
		return
	}

	routes.RegisterRoutes(app.Server)
	log.Println("Registered routes successfully!")
	
	log.Println("Running http server...")
	err = app.RunHttpServer()
	if err != nil {
		log.Fatalln("A fatal error occurred when trying to run the http server:", err.Error())
		return
	}
}