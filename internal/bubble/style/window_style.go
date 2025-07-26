package style

import "github.com/charmbracelet/lipgloss"

var (
	MainWindowStyle = lipgloss.NewStyle().
			Border(ThickBorder).
			BorderForeground(BorderPrimaryColor).
			Background(BackgroundPrimary).
			Padding(1).
			Margin(1)

	ChatWindowStyle = lipgloss.NewStyle().
			Border(RoundedBorder).
			BorderForeground(BorderAccentColor).
			Background(BackgroundPrimary).
			Padding(1).
			Height(30)

	SidebarStyle = lipgloss.NewStyle().
			Border(SharpBorder).
			BorderForeground(BorderSecondaryColor).
			Background(BackgroundSecondary).
			Padding(1).
			Width(25)
)
