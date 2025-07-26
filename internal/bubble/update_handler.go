package bubble

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	api "github.com/spitfiregg/cerebro/internal/api/gemini"
	"github.com/spitfiregg/cerebro/internal/bubble/chat"
	"github.com/spitfiregg/cerebro/internal/bubble/style"
)

// reposible for generating the reponse
func (m *Model) GenerateReponse(prompt string) tea.Cmd {
	return func() tea.Msg {
		resp, err := api.GenerateContent(m.api_key, prompt)
		return LLMreponseMsg{
			response: resp,
			err:      err,
		}
	}
}

// updateViewportContent method is responsible for updating the contents being displayed in
// the viewport of the TUI interface
func (m *Model) updateViewportContent() {

	// create a new string buffer to hold contents/strings
	var content strings.Builder

	history := m.chat.GetHistory()

	// if lengt of history is `0`, we create a new session for the user
	// works everytime

	m.chat.Clear()
	if len(history) == 0 {

		welcomeMsg := style.CreateTitle("Welcome to Cerebro!", m.width-4)
		instructions := lipgloss.NewStyle().
			Foreground(style.TextSecondaryColor).
			Margin(2, 0).
			Align(lipgloss.Center).
			Render("Start typing to chat with Gemini AI\nPress 'h' for help, 'q' to quit")
		content.WriteString(welcomeMsg + "\n" + instructions)

	} else {
		for i, msg := range history {
			// Create styled message bubble based on role
			var styledMessage string
			roleLabel := style.GetRoleStyle(string(msg.Role)).Render(strings.Title(string(msg.Role)))
			messageBubble := style.GetMessageBubbleStyle(string(msg.Role))

			switch msg.Role {
			case chat.RoleUser:
				styledMessage = lipgloss.JoinVertical(
					lipgloss.Left,
					roleLabel, messageBubble.Render(msg.Content),
				)
			case chat.RoleAssistant:
				// Add typing indicator if this is the last message and still generating
				content := msg.Content
				if i == len(history)-1 && m.isLLMthinking && content == "" {
					content = fmt.Sprintf("Thinking %s", m.SpinnerModel.spinner.View())
				}
				styledMessage = lipgloss.JoinVertical(
					lipgloss.Left,
					roleLabel,
					messageBubble.Render(content),
				)
			case chat.RoleSystem:
				styledMessage = lipgloss.JoinVertical(
					lipgloss.Center,
					roleLabel,
					messageBubble.Render(msg.Content),
				)
			}

			content.WriteString(styledMessage)
			if i < len(history)-1 {
				content.WriteString("\n\n")
			}
		}
	}

	m.viewPort.SetContent(content.String())
}
