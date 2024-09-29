package http

import (
	"fiber-server-1/internal/adapter/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Router struct {
	*fiber.App
}

func CreateRouter(

	config *config.HTTP,
	userHandler UserHandler,

) *Router {

	/* Init Fiber App */
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Server-v1",
	})

	/* CORS */
	// app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.AllowedOrigins,
	}))

	/* App logger */
	app.Use(logger.New())
	// router.Use(gin.Logger(), gin.Recovery(), cors.New(ginConfig))

	/* Custom validators */

	/* Define routes */
	user := app.Group("/users")
	{
		user.Post("/", userHandler.Register)
		user.Get("/:id", userHandler.GetUserInfo)
		user.Get("/:id/friends", userHandler.GetUserFriends)
		user.Patch("/:id/:friendId", userHandler.AddRemoveFriend)
	}

	return &Router{
		app,
	}

}

/* Serve */
func (r *Router) Serve(addr string) error {
	return r.Listen(addr)
}
