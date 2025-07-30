package debug

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/davecgh/go-spew/spew"
	"log"
	"os"
)

type Debug struct {
	DumpFile *os.File
}

func (dbg *Debug) EnterDebug(dumpFile string, filename string) {

	if _, ok := os.LookupEnv("DEBUG"); ok {
		var err error

		dbg.DumpFile, err = os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("ran into some error when opening the file: %v", err)
			os.Exit(1)
		}
		fmt.Printf("Debug has been logged into file %s\n", filename)
	} else {
		dbg.DumpFile = nil
	}
}

func (dbg *Debug) WriteLog(msg tea.Msg) {
	if dbg.DumpFile != nil {
		spew.Fdump(dbg.DumpFile, msg)
	}
}

func (dm *Debug) CloseDebug() {
	if dm.DumpFile != nil {
		err := dm.DumpFile.Close()
		if err != nil {
			log.Printf("Error closing debug dump file: %v", err)
		}
		dm.DumpFile = nil
	}
}
