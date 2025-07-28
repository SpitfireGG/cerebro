package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spitfiregg/cerebro/internal/debug"
	"github.com/spitfiregg/cerebro/internal/ui/states"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	Dmodel := debug.Debug{
		DumpFile: m.DebugModel.Dump,
	}
	Dmodel.WriteLog(msg)

	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.viewPort.Width = msg.Width
		m.viewPort.Height = msg.Height - 6
		m.textInput.Width = msg.Width - 4

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

	case window.ModelSelectedMsg:

		m.selectedLLM = msg.ModelName
		m.currentState = MainWindow // transition to the chat window
		m.chat.AddSystemMessage("Model selected: " + msg.ModelName)
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
		m.LLMSelectorWindow = newTableModel.(window.LLMmodel)
		cmds = append(cmds, cmd)

	case MainWindow:

		// revert back to LLM selection window  on `esc`
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "esc" {
			m.currentState = ModelSelection
		}

		// Only update spinner when thinking
		if m.isLLMthinking {
			m.spinner, cmd = m.spinner.Update(msg)
			cmds = append(cmds, cmd)
		} else {
			// Handle user input only when not thinking
			if key, ok := msg.(tea.KeyMsg); ok && key.String() == "enter" {
				if m.textInput.Value() != "" {
					prompt := m.textInput.Value()
					m.isLLMthinking = true
					m.chat.AddUserMessage(prompt)
					m.updateViewportContent()
					m.textInput.SetValue("")

					// start the spinner when we begin thinking
					cmd = tea.Batch(
						m.GenerateReponse(prompt),
						m.spinner.Tick,
					)
					cmds = append(cmds, cmd)
				}
			} else {
				m.textInput, cmd = m.textInput.Update(msg)
				cmds = append(cmds, cmd)
			}
		}

		m.viewPort, cmd = m.viewPort.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
