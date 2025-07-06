package auth

import (
	"fmt"
	"github.com/joho/godotenv"
	api "github.com/spitfiregg/cerebro/internal/api/gemini"
	"os"
)

func LoadAPiKey() (AppCfg *api.AppConfig) {
	// load the environment variable
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Unable to load environment variable, are you sure you set the .env in the current directory?")
	}
	gemini_key := os.Getenv("GEMINI_API_KEY")
	if gemini_key == "" {
		fmt.Println("GEMINI_API_KEY environment variable not found. Add that to your .env and try again, no reload needed!")
		os.Exit(1)
	}
	AppCfg = api.NewDefaultAppConfig(gemini_key)
	return AppCfg
}
