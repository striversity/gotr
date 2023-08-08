package main

import (
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

func getHandler(c *fiber.Ctx) error {
	log.Infof("request received")
	return nil
}

func main() {
	appConfig := fiber.Config{
		AppName: "My Awesome App v0.0.0-beta1",
		EnablePrintRoutes: true,
		ServerHeader: "Awesome App 1",
	}

	app := fiber.New(appConfig)
	app.Get("/", getHandler).Name("Get default")

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
