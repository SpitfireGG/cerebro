package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spitfiregg/cerebro/internal/debug"
	"github.com/spitfiregg/cerebro/internal/ui/states"
	"github.com/spitfiregg/cerebro/internal/ui/styles"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	Dmodel := debug.Debug{
		DumpFile: m.Dump,
	}
	Dmodel.WriteLog(msg)
	var cmds []tea.Cmd
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.viewPort.Width = msg.Width
		m.viewPort.Height = msg.Height
		m.textArea.SetWidth(msg.Width - 4)
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case window.ModelSelectedMsg:
		m.selectedLLM = msg.ModelName
		m.currentState = MainWindow
		m.chat.AddSystemMessage(styles.SelectedModelStyle.Render(m.selectedLLM))
		m.updateViewportContent()
		m.viewPort.GotoBottom()
		return m, textinput.Blink
	case LLMreponseMsg:
		if m.currentState == MainWindow {
			m.isLLMthinking = false
			if msg.err != nil {
				m.chat.AddSystemMessage("Error: " + msg.err.Error())
			} else {
				m.chat.AddAssistantMessage(msg.response)
			}
			m.updateViewportContent()
			m.viewPort.GotoBottom()
		}
		return m, nil
	}
	switch m.currentState {
	case GreetWindow:
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "enter" {
			m.currentState = ModelSelection
		}
	case ModelSelection:
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "esc" {
			m.currentState = GreetWindow
		}
		var newTableModel tea.Model
		newTableModel, cmd = m.LLMSelectorWindow.Update(msg)
		if lm, ok := newTableModel.(window.LLMmodel); ok {
			m.LLMSelectorWindow = lm
		}
		cmds = append(cmds, cmd)
	case MainWindow:
		var prompt string
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "esc" {
			m.currentState = ModelSelection
		}
		if m.isLLMthinking {
			m.spinner, cmd = m.spinner.Update(msg)
			cmds = append(cmds, cmd)
		} else {
			if key, ok := msg.(tea.KeyMsg); ok {
				// Use proper conditional logic with else if
				if key.String() == "tab" {
					m.textArea.InsertString("\n")
					m.textArea, cmd = m.textArea.Update(msg)
					cmds = append(cmds, cmd)
				} else if key.String() == "enter" && m.textArea.Value() != "" {
					prompt = m.textArea.Value()
					m.isLLMthinking = true
					m.chat.AddUserMessage(prompt)
					m.updateViewportContent()
					m.textArea.Reset()
					cmd = tea.Batch(
						m.StartStream(prompt),
						m.spinner.Tick,
					)
					cmds = append(cmds, cmd)
				} else {
					// Only update textarea if we didn't handle the key above
					m.textArea, cmd = m.textArea.Update(msg)
					cmds = append(cmds, cmd)
				}
			}
		}
		m.viewPort, cmd = m.viewPort.Update(msg)
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}
