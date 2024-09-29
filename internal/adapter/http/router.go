package http

import (
	"fiber-server-1/internal/adapter/config"
)

type Router struct {
}

func CreateRouter(
	config *config.HTTP,
	userHandler UserHandler,
) *Router {

	// if config.Env == "production" {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	/* CORS */
	// ginConfig := cors.DefaultConfig()
	// allowedOrigins := config.AllowedOrigins
	// originsList := strings.Split(allowedOrigins, ",")
	// ginConfig.AllowOrigins = originsList

	// router := gin.New()
	// router.Use(gin.Logger(), gin.Recovery(), cors.New(ginConfig))

	/* Custom validators */

	/* Define routes */
	user := router.Group("/users")
	{
		user.POST("/", userHandler.Register)
		user.GET("/:id", userHandler.GetUserInfo)
		user.GET("/:id/friends", userHandler.GetUserFriends)
		user.PATCH("/:id/:friendId", userHandler.AddRemoveFriend)
	}

	return &Router{
		router,
	}

}

/* Serve */
func (r *Router) Serve(addr string) error {
	return r.Run(addr)
}
