package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func MainPage(c *fiber.Ctx) error {
	return c.SendString("Hi I am the go fiber an inspired Express")
}
func TestRouter(c *fiber.Ctx) error {
	return c.SendString("FOr the test router")
}

func TestSleep(c *fiber.Ctx) error {
	cha := make(chan time.Time)
	start := time.Now()
	go func(cha chan time.Time) {
		time.Sleep(5 * time.Second)
		cha <- time.Now()
	}(cha)
	t := <-cha
	elapsed := t.Sub(start).Milliseconds()
	str := strconv.FormatInt(elapsed, 10)
	return c.SendString("For the test router " + str)
}
