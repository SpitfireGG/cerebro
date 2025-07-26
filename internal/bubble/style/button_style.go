package style

import "github.com/charmbracelet/lipgloss"

// Button and selection styles
var (
	ButtonStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(White)).
			Background(lipgloss.Color(Purple)).
			Padding(0, 2).
			Margin(0, 1).
			Bold(true).
			Align(lipgloss.Center)

	ButtonHoverStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(EerieBlack)).
				Background(lipgloss.Color(ElectricBlue)).
				Padding(0, 2).
				Margin(0, 1).
				Bold(true).
				Align(lipgloss.Center)

	ButtonActiveStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(White)).
				Background(lipgloss.Color(NeonFuchsia)).
				Padding(0, 2).
				Margin(0, 1).
				Bold(true).
				Align(lipgloss.Center)

	MenuItemStyle = lipgloss.NewStyle().
			Foreground(TextPrimaryColor).
			Background(BackgroundPrimary).
			Padding(0, 1).
			Margin(0)

	MenuItemSelectedStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(White)).
				Background(lipgloss.Color(SonicSilver)).
				Padding(0, 1).
				Margin(0).
				Bold(true)
)
