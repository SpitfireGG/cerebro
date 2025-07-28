package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spitfiregg/cerebro/cmd/auth"
	"github.com/spitfiregg/cerebro/internal/ui"
	"github.com/spitfiregg/cerebro/internal/debug"
)

// TODO: add feat:  initiate multiple reponses with concurrency

func main() {

	model := ui.InitialModel(auth.LoadAPiKey())

	var dbgModel debug.Debug
	dbgModel.EnterDebug("", "debug.log")

	model.Dump = dbgModel.DumpFile

	program := tea.NewProgram(&model, tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := program.Run(); err != nil {
		log.Fatalf("something went wrong: %v", err)
		os.Exit(1)
	}
	dbgModel.CloseDebug()

	// ui tea debugging
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

}
