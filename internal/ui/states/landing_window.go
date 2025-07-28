package window

import "github.com/charmbracelet/lipgloss"

func RenderLogo(width, height int) string {
	logo := `
    ██████╗███████╗██████╗ ███████╗██████╗ ██████╗  ██████╗
   ██╔════╝██╔════╝██╔══██╗██╔════╝██╔══██╗██╔══██╗██╔═══██╗
   ██║     █████╗  ██████╔╝█████╗  ██████╔╝██████╔╝██║   ██║
   ██║     ██╔══╝  ██╔══██╗██╔══╝  ██╔══██╗██╔══██╗██║   ██║
   ╚██████╗███████╗██║  ██║███████╗██████╔╝██║  ██║╚██████╔╝
    ╚═════╝╚══════╝╚═╝  ╚═╝╚══════╝╚═════╝ ╚═╝  ╚═╝ ╚═════╝
	`

	logoStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#B22222")).Bold(true)
	title := `Smart TUI for seamless LLM interactions`
	logo_ret := lipgloss.JoinVertical(lipgloss.Center, logo, title)

	return logoStyle.Render(lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, logo_ret))

}
