package bubble

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spitfiregg/RTUI_chatbot/internal/debug"
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
		if m.isLLMthinking && msg.String() != "ctrl+c" && msg.String() != "q" {
			break
		}
		switch msg.String() {
		case "q", "Q", "quit", "ctrl-c":
			return m, tea.Quit

		case "enter":
			if m.textInput.Value() != "" && !m.isLLMthinking {
				prompt := m.textInput.Value()
				m.isLLMthinking = true

				m.chat.AddUserMessage(prompt)
				m.updateViewportContent()

				m.textInput.SetValue("")

				cmd = m.GenerateReponse(prompt)
				cmds = append(cmds, cmd)
			}
		}
	case LLMreponseMsg:
		m.isLLMthinking = false
		if msg.err != nil {
			m.chat.AddSystemMessage("Error: " + msg.err.Error())
		} else {
			m.chat.AddAssistantMessage(msg.response)
		}
		m.updateViewportContent()
		m.viewPort.GotoBottom()
	}
	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)

	m.viewPort, cmd = m.viewPort.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)

}
