package styles

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type pos int

// Box struct dictates the position to set the border and dimensions to put the box on
type Box struct {
	// We would like to add a title to in the border itself and the library's border is crap for that
	X, Y          pos
	Width, Height int
}

// Title struct defines the properties of the title
type Title struct {
	Text     string
	Color    lipgloss.ANSIColor
	Position int
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

// ANSI escape codes for terminal control
const (
	Reset      = "\033[0m"
	Clear      = "\033[2J"
	Home       = "\033[H"
	HideCursor = "\033[?25l"
	ShowCursor = "\033[?25h"

	ColorReset = "\033[0m"
	ColorBold  = "\033[1m"
	ColorDim   = "\033[2m"
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

var (
	prussianBlue = "#003153"
	cherryRed    = "#c40000"
)

// moveCursor places the cursor to specific locations in the screeen
// ANSI code: ESC[{line};{column}H
func moveCursor(x, y int) {
	fmt.Printf("\033[%d;%dH", y, x)
}

// DrawTop function draws the border with title in it
func DrawTop(x, y int, width int, title string) {

	style := lipgloss.NewStyle().Foreground(lipgloss.Color(prussianBlue))

	moveCursor(x, y)
	style.Render(RoundLeftUp)

	pos := 1
	moveCursor(x+pos, y)
	style.Render(HLine)

	titleLen := len(title)
	fmt.Print(ColorReset + cherryRed + "" + title + "" + ColorReset + cherryRed)

	moveCursor(x+pos+titleLen, y)

	currentPos := x + pos + titleLen

	for currentPos < width-1 {
		style.Render(HLine)
		currentPos++

		if currentPos < width-1 {
			fmt.Print(VLine)
		}
	}

}
