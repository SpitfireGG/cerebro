package debug

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/davecgh/go-spew/spew"
	"github.com/spitfiregg/RTUI_chatbot/internal/bubble"
)


type DebugModel struct {
	bubble.Model
	dumpFile *os.File
}

func (dbg *DebugModel) EnterDebug(filename string) {

	if _, ok := os.LookupEnv("DEBUG"); ok {
		var err error

		dbg.dumpFile, err = os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("ran into some error when opening the file: %v", err)
			os.Exit(1)
		}
		fmt.Printf("Debug has been logged into file %s\n", filename)
	}else{
		dbg.dumpFile = nil
	}
}



func (dbg *DebugModel) WriteLog(msg tea.Msg) {
	if dbg.dumpFile != nil {
		spew.Dump(dbg.dumpFile, msg)
	}
}

func (dm *DebugModel) CloseDebug() {
	if dm.dumpFile != nil {
		err := dm.dumpFile.Close()
		if err != nil {
			log.Printf("Error closing debug dump file: %v", err)
		}
		dm.dumpFile = nil 
	}
}



