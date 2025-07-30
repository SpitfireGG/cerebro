package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	Black          = "#000000"
	AlmostBlack    = "#1a1a1a"
	EerieBlack     = "#1f1f1f"
	White          = "#ffffff"
	AntiFlashWhite = "#f2f3f4"
	DarkVanilla    = "#d1bea8"
	Grey           = "#808080"
	LightGrey      = "#dddddd"
	DimGrey        = "#777777"
	SonicSilver    = "#757575"
	ImperialRed    = "#ed2939"
	Crimson        = "#dc143c"
	NeonFuchsia    = "#fe4164"
	Gold           = "#ffd700"
	ElectricBlue   = "#7df9ff"
	NeonGreen      = "#39ff14"
	Purple         = "#7D56F4"
	PrussianBlue   = "#003153"

	// extra sauce
	ICS             = "#33b5e5"
	DarkICS         = "#0099CC"
	CherryRed       = "#c40000"
	McLarenEdition  = "#ff9f34"
	Darkorchid      = "#9932CC"
	YellowGold      = "#FFDF00"
	MetallicGold    = "#D4AF37"
	LightSkyBlue    = "#81C4FF"
	YaleBlue        = "#16588E"
	AlizarinCrimson = "#E7222E"

	// status colors
	SuccessGreen = "#00ff7f"
	WarningAmber = "#ffbf00"
	ErrorRed     = "#ff6b6b"
	InfoBlue     = "#4a9eff"
)

// color themes
var (
	TextPrimaryColor   = lipgloss.AdaptiveColor{Light: AlmostBlack, Dark: AntiFlashWhite}
	TextSecondaryColor = lipgloss.AdaptiveColor{Light: SonicSilver, Dark: LightGrey}
	TextMutedColor     = lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: DimGrey}
	TextDisabledColor  = lipgloss.AdaptiveColor{Light: "#555156", Dark: "#505050"}
	TextTitleColor     = lipgloss.AdaptiveColor{Light: Black, Dark: White}

	BorderPrimaryColor   = lipgloss.AdaptiveColor{Light: EerieBlack, Dark: LightGrey}
	BorderSecondaryColor = lipgloss.AdaptiveColor{Light: Grey, Dark: DimGrey}
	BorderAccentColor    = lipgloss.AdaptiveColor{Light: NeonFuchsia, Dark: ElectricBlue}

	BackgroundPrimary   = lipgloss.AdaptiveColor{Light: White, Dark: AlmostBlack}
	BackgroundSecondary = lipgloss.AdaptiveColor{Light: AntiFlashWhite, Dark: EerieBlack}
	BackgroundAccent    = lipgloss.AdaptiveColor{Light: DarkVanilla, Dark: "#2a2a2a"}

	SelectedColor = lipgloss.AdaptiveColor{Light: NeonFuchsia, Dark: ElectricBlue}
	HoverColor    = lipgloss.AdaptiveColor{Light: Purple, Dark: NeonGreen}
)

// Message role styles
var (
	UserStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(White)).
			Background(lipgloss.Color(NeonFuchsia)).
			Padding(0, 0).
			Margin(0, 1, 1, 0).
			Bold(true).
			Align(lipgloss.Left)

	AssistantStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(AlmostBlack)).
			Background(lipgloss.Color(ElectricBlue)).
			Padding(0, 0).
			Margin(0, 0, 1, 1).
			Bold(true).
			Align(lipgloss.Left)

	SystemStyle = lipgloss.NewStyle().
			Background(lipgloss.Color(WarningAmber)).
			Foreground(lipgloss.Color(White)).
			Padding(0, 0).
			Margin(0, 1, 1, 1).
			Bold(true).
			Align(lipgloss.Center)

	SelectedModelStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(AntiFlashWhite)).
				Background(lipgloss.Color("#F35865")).
				Padding(0, 1).
				Margin(2, 1).
				Bold(true).
				Italic(false).
				Align(lipgloss.Left)

	TimeDisplayStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(AntiFlashWhite)).
				Background(lipgloss.Color("#F35865")).
				Bold(true).
				Italic(false).
				Align(lipgloss.Right)
)

// Chat ui styles
var (
	UserMessageStyle = lipgloss.NewStyle().
				Border(RoundedBorder).
				BorderForeground(lipgloss.Color(WarningAmber)).
				Background(BackgroundSecondary).
				Padding(0, 1).
				Margin(0, 1, 0, 1).
		/* 				MaxWidth(60). */
		Align(lipgloss.Left)

	AssistantMessageStyle = lipgloss.NewStyle().
				Border(RoundedBorder).
				BorderForeground(lipgloss.Color(ElectricBlue))

	SystemMessageStyle = lipgloss.NewStyle().
				Border(SharpBorder).
				BorderForeground(lipgloss.Color(WarningAmber)).
				Background(BackgroundAccent).
				Padding(0, 1).
				Margin(0, 1, 0, 1).
				MaxWidth(50).
				Align(lipgloss.Center)
)

// Interactive UI styles
var (
	InputStyle = lipgloss.NewStyle().
			Border(RoundedBorder).
			BorderForeground(lipgloss.Color(PrussianBlue)).
			Padding(0, 1).
			Margin(0, 0).
			Foreground(TextPrimaryColor)

	InputFocusedStyle = lipgloss.NewStyle().
				Foreground(TextPrimaryColor).Bold(true).Italic(true)

	StatusBarStyle = lipgloss.NewStyle().
			Padding(0, 1, 0, 1).
			Italic(true)

	ThinkingStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(AntiFlashWhite)).
			Bold(true).
			Italic(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(White)).
			Background(lipgloss.Color(ErrorRed)).
			Padding(0, 1).
			Margin(0, 1).
			Bold(true)
)

// GetRoleStyle: Helper functions for dynamic styling
func GetRoleStyle(role string) lipgloss.Style {
	switch role {
	case "user":
		return UserStyle
	case "assistant":
		return AssistantStyle
	case "System":
		return SystemStyle
	default:
		return lipgloss.NewStyle()
	}
}

func GetMessageBubbleStyle(role string) lipgloss.Style {
	switch role {
	case "user":
		return UserMessageStyle
	case "assistant":
		return AssistantMessageStyle
	case "System":
		return SystemMessageStyle
	default:
		return lipgloss.NewStyle()
	}
}

func GetButtonStyle(state string) lipgloss.Style {
	switch state {
	case "hover":
		return ButtonHoverStyle
	case "active":
		return ButtonActiveStyle
	default:
		return ButtonStyle
	}
}

// Utility functions for borders and separators
func CreateSeparator(width int, char string, color string) string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(color)).
		Render(lipgloss.JoinHorizontal(lipgloss.Center, char))
}

func CreateTitle(title string, width int) string {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(TextTitleColor).
		Background(BackgroundAccent).
		Padding(0, 1).
		Width(width).
		Align(lipgloss.Center).
		Render(title)
}
