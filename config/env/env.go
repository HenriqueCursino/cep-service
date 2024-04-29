package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	LogLevel string
	Port     string
)

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("runnning the application without a .env file")
	}

	LogLevel = os.Getenv("LOG_LEVEL")
	Port = os.Getenv("PORT")
}
