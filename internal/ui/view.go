package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/spitfiregg/cerebro/internal/ui/states"
	"github.com/spitfiregg/cerebro/internal/ui/styles"
)

func (m Model) View() string {
	switch m.currentState {

	case GreetWindow:
		return window.RenderLogo(m.width, m.height)

	case ModelSelection:
		tableView := m.LLMSelectorWindow.View()
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, tableView)

	case MainWindow:

		placeholder := "Start a new conversation with Gemini..."
		m.textArea.Placeholder = placeholder

		var status = lipgloss.NewStyle().Bold(true).Italic(true).Foreground(lipgloss.Color(styles.AntiFlashWhite)).Render("Gemini")
		/* 			lipgloss.NewStyle().Bold(true).Underline(true).Background(lipgloss.Color(style.ImperialRed)).Foreground(lipgloss.Color(style.AntiFlashWhite)).Padding(0, 1, 0, 1).Render("Gemini") */

		var promptStat string

		if m.SpinnerModel.err != nil {
			fmt.Println(m.SpinnerModel.err.Error())
		}
		if m.isLLMthinking {
			promptStat = "Thinking"
			status = fmt.Sprintf("%s%s", styles.ThinkingStyle.Render("Thinking"), m.spinner.View())
			placeholder = "Wait for the reponse to end..."
			m.textArea.Placeholder = placeholder
		} else {
			promptStat = "Ready"
		}

		/* 		statusHeight := lipgloss.Height(status) */
		inputAreaHeight := 1 + lipgloss.Height(m.textArea.View()) // 1 for status line, plus text input height

		m.viewPort.Height = m.height - inputAreaHeight - 2

		statusBar := styles.StatusBarStyle.Render(fmt.Sprintf("• Model: %s • Status: %s • [Tab] Newline • [Enter] Send • q/esc: Quit", m.selectedLLM, promptStat))

		promptBoxContent := status + "\n" + m.textArea.View()
		/* PromptBoxHeight := styles.InputStyle.Height(lipgloss.Height(promptBoxContent))
		// init
		promptBox := window.PromptBox{
			Width: m.width - 2,
			Height: PromptBoxHeight,
		} */

		promptBox := styles.InputStyle.Width(m.width - 2).Height(lipgloss.Height(promptBoxContent)).Render(promptBoxContent)

		ren := lipgloss.JoinVertical(lipgloss.Center, m.viewPort.View(), promptBox, statusBar)
		return ren

	case SettingsWindow:
		return "Settings (Not Implemented)"
	}
	return ""
}
