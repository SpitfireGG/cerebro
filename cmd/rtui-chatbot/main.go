package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"bufio"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

// TODO: add feat:  make this shit a TUI program for better interactions ( set buffers for inputs)
// TODO: add feat:  initiate multiple reponses with concurrency
// TODO: add feat: syntax/keywords highlighting for better readability ( parsing tokens )

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("could not find .env")
	}

	gemini_key := os.Getenv("GEMINI_API_KEY")
	if gemini_key == "" {
		fmt.Println("key was not set or found")
		os.Exit(1)
	}

}
