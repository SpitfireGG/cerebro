package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	api "github.com/spitfiregg/garlic/internal/api/gemini"
	"github.com/spitfiregg/garlic/internal/bubble"
	"github.com/spitfiregg/garlic/internal/debug"
)

// TODO: add feat:  initiate multiple reponses with concurrency

func main() {

	// load the environment variable
	err := godotenv.Load()
	if err != nil {
		fmt.Println("could not find .env")
	}
	gemini_key := os.Getenv("GEMINI_API_KEY")
	if gemini_key == "" {
		fmt.Println("key was not set or found")
		os.Exit(1)
	}

	appCfg := api.NewDefaultAppConfig(gemini_key)
	model := bubble.InitialModel(appCfg)

	var dbgModel debug.Debug
	dbgModel.EnterDebug("", "debug.log")

	model.Dump = dbgModel.DumpFile

	program := tea.NewProgram(&model, tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := program.Run(); err != nil {
		log.Fatalf("something went wrong: %v", err)
		os.Exit(1)
	}
	dbgModel.CloseDebug()
}
