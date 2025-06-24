package bubble

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spitfiregg/RTUI_chatbot/internal/debug"
	"github.com/spitfiregg/RTUI_chatbot/window"
	"time"
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
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		}

	case window.ModelSelectedMsg:
		m.selectedLLM = msg.ModelName
		m.currentState = LLMwindow // transition to the chat window
		m.chat.AddSystemMessage("Model selected: " + msg.ModelName)
		m.updateViewportContent()
		m.viewPort.GotoBottom()
		return m, textinput.Blink

	case LLMreponseMsg:
		if m.currentState == LLMwindow {
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
		Transition(time.Second * 3)

	case MainWindow:
		var newTableModel tea.Model
		newTableModel, cmd = m.LLMSelectorWindow.Update(msg)
		m.LLMSelectorWindow = newTableModel.(window.LLMmodel)
		cmds = append(cmds, cmd)

	case LLMwindow:
		// If the model is thinking, don't process any input
		if m.isLLMthinking {
			return m, nil
		}

		// Handle 'enter' to send the prompt
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "enter" {
			if m.textInput.Value() != "" {
				prompt := m.textInput.Value()
				m.isLLMthinking = true
				m.chat.AddUserMessage(prompt)
				m.updateViewportContent()
				m.textInput.SetValue("")

				cmd = m.GenerateReponse(prompt)
				cmds = append(cmds, cmd)
			}
		} else {
			// updathe the text input with the user's typing
			m.textInput, cmd = m.textInput.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	// always update the viewport for scrolling
	m.viewPort, cmd = m.viewPort.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
