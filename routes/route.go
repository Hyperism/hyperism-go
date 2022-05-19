package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/hyperism/hyperism-go/controllers" // replace
    "os"
    "net/http"
    jwtware "github.com/gofiber/jwt/v3"
    _"github.com/hyperism/hyperism-go/actions"
    "github.com/hyperism/hyperism-go/auth"
    "github.com/hyperism/hyperism-go/utix"
)

func Meta(route fiber.Router) {
    api := route.Group("/api")

    secret := os.Getenv("JWT_SECRET_KEY")
	api.Post("/SignUp", auth.SignUp)
	api.Post("/LoginIn", auth.Login)

	meta := api.Group("/meta")
	meta.Use(jwtware.New(jwtware.Config{
		SigningKey:    []byte(secret),
		SigningMethod: "HS256",
		TokenLookup:   "header:Authorization",
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			return c.
				Status(http.StatusUnauthorized).
				JSON(utix.NewJError(e))
		},
	}),
	)

    // from here
    // we need to add bearer token
    meta.Get("/", controllers.GetAllMeta)
    meta.Get("/:owner", controllers.GetMeta)
    meta.Post("/", controllers.AddMeta)
    meta.Put("/:id", controllers.UpdateMeta)
    meta.Delete("/:id", controllers.DeleteMeta)

    meta.Post("/getshader/:id", controllers.GetShader)
}