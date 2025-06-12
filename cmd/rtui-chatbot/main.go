package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"github.com/spitfiregg/RTUI_chatbot/internal/bubble"
	"github.com/spitfiregg/RTUI_chatbot/internal/debug"
)

// TODO: add feat:  make this shit a TUI program for better interactions ( set buffers for inputs)
// TODO: add feat:  initiate multiple reponses with concurrency
// TODO: add feat: syntax/keywords highlighting for better readability ( parsing tokens )

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

	model := bubble.InitialModel(gemini_key)

	var dbgModel debug.Debug
	dbgModel.EnterDebug("", "debug.log")

	model.Dump = dbgModel.DumpFile

	program := tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := program.Run(); err != nil {
		log.Fatalf("something went wrong: %v", err)
		os.Exit(1)
	}
	dbgModel.CloseDebug()
}
