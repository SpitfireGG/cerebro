package bubble

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

func (m Model) View() string {
	switch m.currentState {
	case GreetWindow:

		greetMsg := "Welcome to the Playground!"
		greetStr := lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, greetMsg)
		return greetStr

	case MainWindow:
		tableView := m.LLMSelectorWindow.View()

		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, tableView)

	case LLMwindow:

		var status string = "Prompt Gemini"

		if m.SpinnerModel.err != nil {
			fmt.Println(m.SpinnerModel.err.Error())
		}
		if m.isLLMthinking {
			status = fmt.Sprintf("Thinking%s", m.SpinnerModel.spinner.View())
		}
		return fmt.Sprintf("%s\n%s\n%s\n%s",
			m.viewPort.View(),
			strings.Repeat("─", m.width),
			status,
			m.textInput.View(),
		)
	case SettingsWindow:
		return "Settings (Not Implemented)"
	}
	return ""
}
