package main

import (
	"github.com/rootemanuel/api-fiber-crud/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var sper = service.PersonService{}

const (
	ServicePort = ":8080"
)

func main() {

	app := InitApi()
	app.Listen(ServicePort)
}

func InitApi() *fiber.App {

	app := fiber.New()
	app.Use(cors.New())

	//URL`s GROUP API
	api := app.Group("/api")
	api.Get("/ping", sper.Ping)
	api.Get("/test", sper.TestCache)

	//URL`s GROUP API - V1
	v1 := api.Group("/v1")
	v1.Get("/person", sper.GetPersons)
	v1.Get("/person/:cpf", sper.GetPerson)
	v1.Post("/person", sper.CreatePerson)
	v1.Delete("/person/:cpf", sper.DeletePerson)
	v1.Put("/person/:cpf", sper.UpdatePerson)

	return app
}
