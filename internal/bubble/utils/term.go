package utils

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type TermSize struct {
	Width int
	Float int
}

// GetTermSize() outputs the current length and width of the terminal
func GetTermSize() (width, height int) {
	fd := int(os.Stdout.Fd())

	if !term.IsTerminal(fd) {
		fmt.Println("Not running in a terminal.")
		return 80, 24 //fallback to default terminal size
	}
	width, height, err := term.GetSize(fd)
	if err != nil {
		fmt.Printf("Error getting terminal size: %v\n", err)
		return 80, 24 // fallback to default size
	}
	return width, height
}

// PlaceContentCenter() function helps place contents to the center of the terminal
// Align.Center() dont work for contents inside borders
func PlaceContentCenter(content string) {
	termWidth, _ := GetTermSize()

	contentWidth := lipgloss.Width(content)

	remainedSpace := termWidth - contentWidth

	if remainedSpace < 0 {
		remainedSpace = 0
	}
	margin := remainedSpace / 2
}
