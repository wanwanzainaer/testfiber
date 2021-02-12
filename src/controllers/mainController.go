package controllers

import (
	"fmt"
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
	fmt.Println("I am in sleep")
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

// func TestSleep(c *fiber.Ctx) error {
// 	fmt.Println("I am in sleep")
// 	start := time.Now()
// 	time.Sleep(5 * time.Second)
// 	t := time.Now()
// 	elapsed := t.Sub(start).Milliseconds()
// 	str := strconv.FormatInt(elapsed, 10)
// 	return c.SendString("For the test router " + str)
// }

///////////////////////////////
// func TestSleep(c *fiber.Ctx) error {
// 	cha := make(chan time.Time)
// 	fmt.Println("I am in sleep")
// 	start := time.Now()
// 	go stopTime(cha)
// 	t := <-cha
// 	elapsed := t.Sub(start).Milliseconds()
// 	str := strconv.FormatInt(elapsed, 10)
// 	return c.SendString("For the test router " + str)
// }
// func stopTime(cha chan time.Time) {
// 	time.Sleep(5 * time.Second)
// 	cha <- time.Now()
// }
