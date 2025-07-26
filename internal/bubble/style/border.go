package style

import "github.com/charmbracelet/lipgloss"

const (
	HLine          = "─"
	VLine          = "│"
	DottedVLine    = "╎"
	LeftUp         = "┌"
	RightUp        = "┐"
	LeftDown       = "└"
	RightDown      = "┘"
	RoundLeftUp    = "╭"
	RoundRightUp   = "╮"
	RoundLeftDown  = "╰"
	RoundRightDown = "╯"
	TitleLeftDown  = "┘"
	TitleRightDown = "└"
	TitleLeft      = "┐"
	TitleRight     = "┌"
	DivRight       = "┤"
	DivLeft        = "├"
	DivUp          = "┬"
	DivDown        = "┴"
	Cross          = "┼"
	Up             = "↑"
	Down           = "↓"
	Left           = "←"
	Right          = "→"
	Enter          = "↵"
	Bullet         = "•"
	Arrow          = "→"
)

var (
	RoundedBorder = lipgloss.Border{
		Top:         HLine,
		Bottom:      HLine,
		Left:        VLine,
		Right:       VLine,
		TopLeft:     RoundLeftUp,
		TopRight:    RoundRightUp,
		BottomLeft:  RoundLeftDown,
		BottomRight: RoundRightDown,
	}

	SharpBorder = lipgloss.Border{
		Top:         HLine,
		Bottom:      HLine,
		Left:        VLine,
		Right:       VLine,
		TopLeft:     LeftUp,
		TopRight:    RightUp,
		BottomLeft:  LeftDown,
		BottomRight: RightDown,
	}

	ThickBorder = lipgloss.Border{
		Top:         "━",
		Bottom:      "━",
		Left:        "┃",
		Right:       "┃",
		TopLeft:     "┏",
		TopRight:    "┓",
		BottomLeft:  "┗",
		BottomRight: "┛",
	}
)
