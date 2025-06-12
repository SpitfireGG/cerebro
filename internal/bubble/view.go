package bubble

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var console strings.Builder

func (m Model) View() string {

	var status string
	if m.isLLMthinking {
		status = ResponseStyle.Render("model is thinking....")
	} else {
		status = ResponseStyle.Render("ready for the prompt...")
	}
	return lipgloss.JoinVertical(
		lipgloss.Left,
		ResponseStyle.Render("Gemini-2.0 flash"),
		"",
		m.viewPort.View(),
		"",
		status,
		m.textIP.View(),
		ResponseStyle.Render("Commands: 'enter' to send • 'ctrl+l' to clear • 'q' to quit"),
	)
}
