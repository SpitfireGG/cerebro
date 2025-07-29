package styles

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type pos int

// Box struct dictates the position to set the border and dimensions to put the box on
type Box struct {
	X, Y          pos
	Width, Height int
}

// Title struct defines the properties of the title
type Title struct {
	Text string
}

// define ASCII chars for custom border
const (
	HLine       = "─"
	VLine       = "│"
	DottedVLine = "╎"

	LeftUp    = "┌"
	RightUp   = "┐"
	LeftDown  = "└"
	RightDown = "┘"

	RoundLeftUp    = "╭"
	RoundRightUp   = "╮"
	RoundLeftDown  = "╰"
	RoundRightDown = "╯"

	TitleLeftDown  = "┘"
	TitleRightDown = "└"
	TitleLeft      = "┐"
	TitleRight     = "┌"

	DivRight = "┤"
	DivLeft  = "├"
	DivUp    = "┬"
	DivDown  = "┴"

	Cross  = "┼"
	Up     = "↑"
	Down   = "↓"
	Left   = "←"
	Right  = "→"
	Enter  = "↵"
	Bullet = "•"
	Arrow  = "→"
)

var title = Title{
	Text: "title",
}

var TitleHardcoded = fmt.Sprintf(RoundRightUp + HLine + DivRight + title.Text + DivLeft)

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

var (
	prussianBlue = "#003153"
	cherryRed    = "#c40000"
)

func drawTitleArea() {}
