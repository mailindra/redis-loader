package main

import (
	"context"
	"fmt"
	"os"
	"redis-loader/internal/config"
	"redis-loader/internal/console"
	"redis-loader/internal/repository"
	"redis-loader/internal/service"
)

func main() {
	ctx := context.Background()

	// Initialize Redis connection
	redisClient, err := config.NewRedisClient()
	if err != nil {
		fmt.Printf("Failed to initialize Redis: %v\n", err)
		os.Exit(1)
	}
	defer redisClient.Close()

	// Initialize repository and service
	repo := repository.NewRedisRepository(redisClient)
	loaderService := service.NewLoaderService(repo)

	// Get user input
	console := console.NewConsoleInput()
	count := console.GetNumberInput("Enter the number of key-value pairs to generate")

	// Execute data loading
	if err := loaderService.LoadRandomData(ctx, count); err != nil {
		fmt.Printf("Error loading data: %v\n", err)
		os.Exit(1)
	}
}
