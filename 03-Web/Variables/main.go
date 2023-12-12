package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	/*
		check why godotenv.Load() is not overriding the env variables
	*/
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	// config
	cfgAddr := os.Getenv("ADDR")
	cfgTitle:= os.Getenv("TITLE")

	fmt.Println("ADDR:", cfgAddr)
	fmt.Println("TITLE:", cfgTitle)

}