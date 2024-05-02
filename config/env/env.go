package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	LogLevel string
	Port     string
	Host     string
	Token    string
)

func LoadEnvs() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("runnning the application without a .env file")
	}

	Port = MustGet("PORT")
	Host = MustGet("HOST")
	LogLevel = MustGet("LOG_LEVEL")
	Token = MustGet("TOKEN")
}

// MustGet environment variable or panic if empty
func MustGet(varName string) string {
	value := os.Getenv(varName)
	if value == "" {
		panic(fmt.Sprintf("%s not set", varName))
	}

	return value
}
