package bubble

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spitfiregg/RTUI_chatbot/internal/debug"
)

// BUG: input is not being updated and key mappings wont work

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	Dmodel := debug.Debug{
		DumpFile: m.DebugModel.Dump,
	}

	Dmodel.WriteLog(msg)

	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.handleResize(msg)

	case tea.KeyMsg:
		_, cmd = m.handleKeyPress(msg)
		cmds = append(cmds, cmd)

	case LLMreponseMsg:
		cmd = m.handleLLMResponse(msg)
		cmds = append(cmds, cmd)
	}
	// Update UI components
	var inputCmd tea.Cmd
	m.ui.textIP, inputCmd = m.ui.textIP.Update(msg)
	cmds = append(cmds, inputCmd)

	return m, tea.Batch(cmds...)

}
