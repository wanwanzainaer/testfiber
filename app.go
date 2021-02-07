package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wanwanzainaer/testfiber/src/controllers"
)

func main() {
	app := fiber.New()
	app.Get("/", controllers.MainPage)
	app.Get("/test", controllers.TestRouter)
	app.Get("/ttt", controllers.TestSleep)
	app.Listen(":4000")
}
