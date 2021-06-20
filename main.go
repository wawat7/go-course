package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go-course/app/config"
	"go-course/app/containers/student"
	"log"
)

func main() {

	//setup configuration
	configuration := config.New()
	database := config.NewMongoDatabase(configuration)


	//setup Repository
	studentRepository := student.NewStudentRepository(database)

	//setup Service
	studentService := student.NewService(studentRepository)

	//setup Handler
	studentHandler := student.NewHandler(studentService)

	//setup fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	//setup routing
	studentHandler.Route(app)

	log.Fatal(app.Listen(":3000"))

}
