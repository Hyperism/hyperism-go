package main

import (
	"log"
	"os"
    _"fmt"
    _"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hyperism/hyperism-go/config"
	"github.com/hyperism/hyperism-go/routes"
	"github.com/joho/godotenv"

    // shell "github.com/ipfs/go-ipfs-api"
)

func setupRoutes(app *fiber.App) {
    app.Get("/", func(c *fiber.Ctx) error {
        return c.Status(fiber.StatusOK).JSON(fiber.Map{
            "success":     true,
            "message":     "test ?",
            "github_repo": "https://github.com/MikeFMeyer/catchphrase-go-mongodb-rest-api",
        })
    })

    // api := app.Group("/api")
}

func main() {
    if os.Getenv("APP_ENV") != "production" {
        err := godotenv.Load()
        if err != nil {
            log.Fatal("Error loading .env file")
        }
    }

    app := fiber.New()

    app.Use(cors.New())
    app.Use(logger.New())

    config.ConnectDB()

    setupRoutes(app)

    routes.Meta(app)
    port := "3000"
    err := app.Listen(":" + port)

    
    if err != nil {
        log.Fatal("Error app failed to start")
        panic(err)
    }

}