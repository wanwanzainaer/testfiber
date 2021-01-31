package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func MainPage(c *fiber.Ctx) error{
	return c.SendString("Hi I am the go fiber an inspired Express")
}
func TestRouter(c *fiber.Ctx) error{
	return c.SendString("FOr the test router");
}