package main

import (
	"app/internal/application"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// env
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
		return
	}

	// app
	// - config
	cfg := &application.ConfigDefault{
		ServerHost:     os.Getenv("SERVER_HOST"),
		ServerPort:     os.Getenv("SERVER_PORT"),
		FilePathMovies: os.Getenv("PATH_FILE_MOVIES"),	
	}
	// - application
	app := application.NewDefault(cfg)
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}