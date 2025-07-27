package bubble

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spitfiregg/cerebro/internal/bubble/style"
	"github.com/spitfiregg/cerebro/internal/bubble/window"
)

func (m Model) View() string {
	switch m.currentState {

	case GreetWindow:
		return window.RenderLogo(m.width, m.height)

	case ModelSelection:
		tableView := m.LLMSelectorWindow.View()
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, tableView)

	case MainWindow:

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
