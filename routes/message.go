package routes

import (
	"awesomeProject/controllers"
	"github.com/gofiber/fiber/v2"
)

func MessagesRoute(route fiber.Router)  {
	route.Post("/create", controllers.AddNewMessage)
}