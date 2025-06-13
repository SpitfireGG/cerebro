package bubble

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spitfiregg/RTUI_chatbot/internal/api"
	"github.com/spitfiregg/RTUI_chatbot/internal/bubble/chat"
	"strings"
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

func (m *Model) updateViewportContent() {

	var content strings.Builder

	history := m.chat.GetHistory()
	if len(history) == 0 {
		content.WriteString("Welcome! Start typing to chat with Gemini...\n")
	} else {
		for i, msg := range history {
			switch msg.Role {
			case chat.RoleUser:
				content.WriteString(fmt.Sprintf("👤 You: %s\n", msg.Content))
			case chat.RoleAssistant:
				content.WriteString(fmt.Sprintf("🤖 Gemini: %s\n", msg.Content))
			case chat.RoleSystem:
				content.WriteString(fmt.Sprintf("⚠️  System: %s\n", msg.Content))
			}

			// Add spacing between messages except for the last one
			if i < len(history)-1 {
				content.WriteString("\n")
			}
		}
	}

	m.viewPort.SetContent(content.String())
}
