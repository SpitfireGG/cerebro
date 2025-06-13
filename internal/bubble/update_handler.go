package bubble

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spitfiregg/RTUI_chatbot/internal/api"
)

// handles the window reszing stuffs
func (m *Model) handleResize(msg tea.WindowSizeMsg) {

	m.ui.width = msg.Width
	m.ui.height = msg.Height
	m.ui.viewPort.Height = msg.Height - 4
	m.ui.viewPort.Width = msg.Width

}

// handles the keypresses
func (m *Model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {

	switch msg.String() {
	case "q", "ctrl-c", "quit":
		return m, tea.Quit

	case "enter":
		if m.isLLMthinking {
			return m, nil
		}

		prompt := m.ui.textIP.Value()
		if prompt == "" {

			return m, nil
		}

		m.App.chat.AddUserMessage(prompt)
		m.ui.textIP.SetValue("")
		m.isLLMthinking = true

		return m, m.GenerateReponse(prompt)

	}
	return m, nil
}

func (m *Model) handleLLMResponse(msg LLMreponseMsg) tea.Cmd {

	m.isLLMthinking = false

	if msg.err != nil {
		m.chat.AddSystemMessage("Error: " + msg.err.Error())
	} else {
		m.chat.AddAssistantMessage(msg.response)
	}

	return nil
}

func (a *App) GenerateReponse(prompt string) tea.Cmd {
	return func() tea.Msg {
		resp, err := api.GenerateContent(a.api_key, prompt)
		return LLMreponseMsg{
			response: resp,
			err:      err,
		}
	}
}
