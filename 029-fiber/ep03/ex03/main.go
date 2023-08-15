package main

import (
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
)

const grId = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var idx = 0

func getHandlerId() string {
	c := grId[idx%26]
	idx++
	return fmt.Sprintf("grID-%v-%c", idx, c) // eg: grID-1-A, grID-2-B, grID-3-C, ...
}

func getHandler(c *fiber.Ctx) error {
	ccId := getHandlerId()
	n := c.Params("name")

	doStuff(ccId, n)

	slog.Info("request received", "name", n)
	return nil
}

func doStuff(ccId string, n string) {
	go func() {
		slog.Info("starting handler", "ccId", ccId, "name", n)
		t := time.After(10 * time.Second)
		for {
			select {
			case <-t:
				slog.Info("handler done", "ccId", ccId, "name", n)
				return
			default:
				slog.Info("still running", "ccId", ccId, "name", n)
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

func main() {
	appConfig := fiber.Config{
		AppName:           "My Awesome App v0.0.0-beta1",
		EnablePrintRoutes: true,
		ServerHeader:      "Awesome App 1",
		Immutable:         true,
	}

	app := fiber.New(appConfig)
	app.Get("/:name", getHandler).Name("Get default")

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
