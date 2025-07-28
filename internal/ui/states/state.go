package window

import (
	"github.com/charmbracelet/lipgloss"
)

// nolint:unused

// state.go defines the structs, interfaces and various stylings for the TUI program

// define colors for the TUI window
var (
	primaryBg     = lipgloss.Color("#1e1e2e")
	surfaceBg     = lipgloss.Color("#313244")
	overlayBg     = lipgloss.Color("#6c7086")
	textPrimary   = lipgloss.Color("#cdd6f4")
	textSecondary = lipgloss.Color("#a6adc8")
	accentBlue    = lipgloss.Color("#89b4fa")
	accentPurple  = lipgloss.Color("#cba6f7")
	borderColor   = lipgloss.Color("#45475a")
	fireBrick     = lipgloss.Color("#B22222")
)

// define styles for the TUI window
var (
	BaseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(borderColor).
			Padding(1, 2).
			Margin(1, 0)

	TitleStyle = lipgloss.NewStyle().
			Foreground(accentPurple).
			Bold(true).
			Align(lipgloss.Center).
			Margin(0, 0, 1, 0)

	FooterStyle = lipgloss.NewStyle().
			Foreground(textSecondary).
			Italic(true).
			Align(lipgloss.Center).
			Margin(1, 0, 0, 0)
)
