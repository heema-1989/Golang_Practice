package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heema-1989/Go-Fiber/User"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Welcome")
}
func Routers(app *fiber.App) {
	app.Get("/User", User.GetUsers)
	app.Get("/User/:id", User.GetUser)
	app.Post("/User", User.SaveUser)
	app.Put("/User/:id", User.UpdateUser)
	app.Delete("User/:id", User.DeleteUser)
}
func main() {
	User.InitialMigration()
	app := fiber.New()
	app.Get("/", hello)
	Routers(app)
	app.Listen(":3000")
}
