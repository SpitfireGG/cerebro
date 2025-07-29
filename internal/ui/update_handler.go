package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	api "github.com/spitfiregg/cerebro/internal/api/gemini"
	"github.com/spitfiregg/cerebro/internal/chat"
	"github.com/spitfiregg/cerebro/internal/ui/styles"
)

// reposible for generating the reponse
func (m *Model) GenerateReponse(prompt string) tea.Cmd {
	return func() tea.Msg {
		resp, err := api.GenerateContent(m.api_key, prompt)
		return LLMreponseMsg{
			response: resp,
			err:      err,
		}
	}
}

const (
	// Better choices for title connectors
	TitleConnectorLeft  = "┤" // ┤
	TitleConnectorRight = "├" // ├

	// Alternative minimal approach
	TitleSimpleLeft  = "─" // Just use horizontal line
	TitleSimpleRight = "─" // Just use horizontal line

	// Decorative approach
	TitleFancyLeft  = "┨" // ┨
	TitleFancyRight = "┠" // ┠
)

// CreateCenteredTitleLine creates a horizontal line with centered title
func CreateCenteredTitleLine(title string, totalWidth int) string {
	if totalWidth < len(title)+2 {
		// Not enough space, return truncated title
		if totalWidth <= 0 {
			return ""
		}
		if totalWidth <= len(title) {
			return title[:totalWidth]
		}
		return title
	}

	// Calculate spacing
	titleLen := len(title)
	connectorChars := 2 // left + right connectors
	availableSpace := totalWidth - titleLen - connectorChars

	leftPadding := availableSpace / 2
	rightPadding := availableSpace - leftPadding // Handles odd numbers

	// Build the line
	return strings.Repeat(styles.HLine, leftPadding) +
		TitleConnectorLeft +
		title +
		TitleConnectorRight +
		strings.Repeat(styles.HLine, rightPadding)
}

// CreateSimpleTitleLine creates a minimal title line without connectors
func CreateSimpleTitleLine(title string, totalWidth int) string {
	if totalWidth < len(title) {
		if totalWidth <= 0 {
			return ""
		}
		return title[:totalWidth]
	}

	titleLen := len(title)
	availableSpace := totalWidth - titleLen

	leftPadding := availableSpace / 2
	rightPadding := availableSpace - leftPadding

	return strings.Repeat(styles.HLine, leftPadding) +
		title +
		strings.Repeat(styles.HLine, rightPadding)
}

// CreateStyledTitleLine creates a title line with lipgloss styling
func CreateStyledTitleLine(title string, totalWidth int) string {
	borderStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#003153"))
	titleStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#c40000")).Bold(true)

	if totalWidth < len(title)+2 {
		return titleStyle.Render(title)
	}

	titleLen := len(title)
	connectorChars := 2
	availableSpace := totalWidth - titleLen - connectorChars

	leftPadding := availableSpace / 2
	rightPadding := availableSpace - leftPadding

	leftLines := borderStyle.Render(strings.Repeat(styles.HLine, leftPadding))
	rightLines := borderStyle.Render(strings.Repeat(styles.HLine, rightPadding))
	leftConnector := borderStyle.Render(TitleConnectorLeft)
	rightConnector := borderStyle.Render(TitleConnectorRight)
	styledTitle := titleStyle.Render(title)

	return leftLines + leftConnector + styledTitle + rightConnector + rightLines
}

// CreateTitleLineVariants shows different visual styles
func CreateTitleLineVariants(title string, totalWidth int) map[string]string {
	variants := make(map[string]string)

	// Basic centered
	variants["basic"] = CreateSimpleTitleLine(title, totalWidth)

	// With connectors
	variants["connected"] = CreateCenteredTitleLine(title, totalWidth)

	// Minimal brackets
	if totalWidth >= len(title)+4 {
		titleLen := len(title)
		availableSpace := totalWidth - titleLen - 4 // for [ ] brackets
		leftPadding := availableSpace / 2
		rightPadding := availableSpace - leftPadding

		variants["brackets"] = strings.Repeat(styles.HLine, leftPadding) +
			"[ " + title + " ]" +
			strings.Repeat(styles.HLine, rightPadding)
	}

	// Double line style
	if totalWidth >= len(title)+2 {
		titleLen := len(title)
		availableSpace := totalWidth - titleLen - 2
		leftPadding := availableSpace / 2
		rightPadding := availableSpace - leftPadding

		variants["double"] = strings.Repeat("═", leftPadding) +
			"╡" + title + "╞" +
			strings.Repeat("═", rightPadding)
	}

	return variants
}

// DebugTitleLine helps debug width calculations
func DebugTitleLine(title string, totalWidth int) string {
	result := CreateCenteredTitleLine(title, totalWidth)
	actualWidth := len(result)

	debug := lipgloss.NewStyle().
		Foreground(lipgloss.Color("red")).
		Render(fmt.Sprintf("\nDEBUG: Expected width: %d, Actual width: %d, Title: '%s'",
			totalWidth, actualWidth, title))

	return result + debug
}

// updateViewportContent method is responsible for updating the contents being displayed in
// the viewport of the TUI interface
func (m *Model) updateViewportContent() {

	// create a new string buffer to hold contents/strings
	var content strings.Builder

	history := m.chat.GetHistory()

	// if lengt of history is `0`, we create a new session for the user
	// works everytime

	m.chat.Clear()
	if len(history) == 0 {

		welcomeMsg := styles.CreateTitle("Welcome to Cerebro!", m.width-4)
		instructions := lipgloss.NewStyle().
			Foreground(styles.TextSecondaryColor).
			Margin(2, 0).
			Align(lipgloss.Center).
			Render("Start typing to chat with Gemini AI\nPress 'h' for help, 'q' to quit")
		content.WriteString(welcomeMsg + "\n" + instructions)

	} else {
		for i, msg := range history {

			// Create styled message ui based on role
			var styledMessage string
			var title string

			switch msg.Role {
			case chat.RoleUser:
				title = "User"
			case chat.RoleAssistant:
				title = "Assistant"
			case chat.RoleSystem:
				title = "System"
			}

			/* 			titleLen := len(titleHardcoded) */

			// strings.Title is deprecated
			/* 			roleLabel := styles.GetRoleStyle(string(msg.Role)).Render(strings.Title(string(msg.Role))) */
			messageBubble := styles.GetMessageBubbleStyle(string(msg.Role))

			titleLine := CreateStyledTitleLine(title, m.width)

			switch msg.Role {
			case chat.RoleUser:

				styledMessage = lipgloss.JoinVertical(
					lipgloss.Center,
					titleLine,
					messageBubble.Render(msg.Content),
				)
			case chat.RoleAssistant:
				// Add typing indicator if this is the last message and still generating
				content := msg.Content
				if i == len(history)-1 && m.isLLMthinking && content == "" {
					content = fmt.Sprintf("Thinking %s", m.spinner.View())
				}
				styledMessage = lipgloss.JoinVertical(
					lipgloss.Left,
					titleLine,
					messageBubble.Render(content),
				)
				styledMessage = lipgloss.JoinVertical(
					lipgloss.Center,
					lipgloss.NewStyle().Render(titleLine),
					messageBubble.Render(msg.Content),
				)
			case chat.RoleSystem:
				styledMessage = lipgloss.JoinVertical(
					lipgloss.Center,
					titleLine,
					messageBubble.Render(msg.Content),
				)
			}

			content.WriteString(styledMessage)
			if i < len(history)-1 {
				content.WriteString("\n\n")
			}
		}
	}

	m.viewPort.SetContent(content.String())
}
