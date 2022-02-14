package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rickschubert/scenemover/handlers/getscenes"
	"github.com/rickschubert/scenemover/handlers/transitionscene"
	"github.com/rickschubert/scenemover/recompile"
)

func launchServer() {
	fmt.Println("Launching server on port 8080")
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/scenes", getscenes.Handler)
	app.Post("/scenes/transition", transitionscene.Handler)
	log.Fatal(app.Listen(":8080"))
}

func main() {
	done := make(chan bool)
	go recompile.LaunchWatcher()
	go launchServer()
	<-done
}
