package main

import (
	"context"
	"fiber-server-1/internal/adapter/config"
	"fiber-server-1/internal/adapter/database"
	"fiber-server-1/internal/adapter/http"
	"fiber-server-1/internal/core/repository"
	"fiber-server-1/internal/core/service"
	"fmt"
	"os"
)

func main() {

	// Load env variables
	config, err := config.New()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Init database
	ctx := context.Background()
	db, err := database.ConnectDB(&ctx, config.DB)

	if err != nil {
		fmt.Println("Error connecting to database: %s\n", err.Error())
		os.Exit(1)
	}

	/* Dependenct Injection: Connect adapters and cores of hexagonal arhcitecture */

	// User
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := http.NewUserHandler(userService)

	// Post
	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepository)
	postHandler := http.NewPostHandler(postService)

	router := http.CreateRouter(config.HTTP, *userHandler, *postHandler)

	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)
	fmt.Printf("Starting the HTTP server at %s\n", listenAddr)
	err = router.Serve(listenAddr)

	if err != nil {
		fmt.Printf("Error starting the HTTP server %s\n", err.Error())
		os.Exit(1)
	}

}
