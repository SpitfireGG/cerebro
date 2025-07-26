package bubble

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spitfiregg/cerebro/internal/bubble/style"
)

func (m Model) View() string {
	switch m.currentState {
	case GreetWindow:

		logo := `
    ██████╗███████╗██████╗ ███████╗██████╗ ██████╗  ██████╗
   ██╔════╝██╔════╝██╔══██╗██╔════╝██╔══██╗██╔══██╗██╔═══██╗
   ██║     █████╗  ██████╔╝█████╗  ██████╔╝██████╔╝██║   ██║
   ██║     ██╔══╝  ██╔══██╗██╔══╝  ██╔══██╗██╔══██╗██║   ██║
   ╚██████╗███████╗██║  ██║███████╗██████╔╝██║  ██║╚██████╔╝
    ╚═════╝╚══════╝╚═╝  ╚═╝╚══════╝╚═════╝ ╚═╝  ╚═╝ ╚═════╝
	`
		logoStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#B22222")).Bold(true)
		return logoStyle.Render(lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, logo))

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
			strings.Repeat(style.HLine, m.width),
			status,
			m.textInput.View(),
		)
	case SettingsWindow:
		return "Settings (Not Implemented)"
	}
	return ""
}
