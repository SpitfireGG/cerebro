package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spitfiregg/cerebro/cmd/auth"
	"github.com/spitfiregg/cerebro/internal/bubble"
	"github.com/spitfiregg/cerebro/internal/debug"
	"log"
	"os"
)

// TODO: add feat:  initiate multiple reponses with concurrency

func main() {

	model := bubble.InitialModel(auth.LoadAPiKey())

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
