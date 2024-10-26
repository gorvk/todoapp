package initializers

import (
	"fmt"

	"github.com/joho/godotenv"
)

// loading environment variables
func LoadEnv() {
	fmt.Println("Loading Environment Variables...")
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf("Failed to load Environment Variables!")
		panic(err)
	}
	fmt.Println("Environment Variables Loaded!")
}
