package main

import (
	"github.com/rootemanuel/api-fiber-crud/service"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

var person = service.PersonService{}

const (
	ServicePort = 8080
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

	//URL`s GROUP API - V1
	v1 := api.Group("/v1")
	v1.Get("/person", person.GetPersons)
	v1.Get("/person/:cpf", person.GetPerson)
	v1.Post("/person", person.CreatePerson)
	v1.Delete("/person/:cpf", person.DeletePerson)
	v1.Put("/person/:cpf", person.UpdatePerson)

	return app
}
