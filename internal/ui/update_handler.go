package ui

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	api "github.com/spitfiregg/cerebro/internal/api/gemini"
	"github.com/spitfiregg/cerebro/internal/chat"
	"github.com/spitfiregg/cerebro/internal/ui/styles"
)

func (m *Model) StartStream(prompt string) tea.Cmd {
	return tea.Sequence(
		func() tea.Msg {
			chunkChan, err := api.GenerateContentStream(m.api_key, prompt)
			if err != nil {
				return api.StreamChunk{IsError: true, Error: err}
			}
			m.streamChan = chunkChan
			return api.StreamStartMsg{}
		},
		tea.Tick(time.Millisecond*100, func(time.Time) tea.Msg {
			return api.PollStreamMsg{}
		}),
	)
}

// updateViewportContent method is responsible for updating the contents being displayed in
// the viewport of the TUI interface
func (m *Model) updateViewportContent() {

	// create a new string buffer to hold contents/strings
	var content strings.Builder

	history := m.chat.GetHistory()

	// if length of history is `0`, we create a new session for the user
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
			/* var title string

			switch msg.Role {
			case chat.RoleUser:
				title = "User"
			case chat.RoleAssistant:
				title = "Assistant"
			case chat.RoleSystem:
				title = "System"
			} */

			msgContent := lipgloss.NewStyle().Foreground(lipgloss.Color(styles.AntiFlashWhite)).
				Background(lipgloss.Color("#F35865")).
				Bold(true).
				Italic(false).
				Align(lipgloss.Center)

				/* 			titleLine := CreateStyledTitleLine(title, m.width) */

			switch msg.Role {
			case chat.RoleUser:

				styledMessage = lipgloss.JoinVertical(
					lipgloss.Center,
					msgContent.Render(msg.Content),
				)

			case chat.RoleAssistant:
				// Add typing indicator if this is the last message and still generating
				content := msg.Content
				if i == len(history)-1 && m.isLLMthinking && content == "" {
					content = fmt.Sprintf("Thinking %s", m.spinner.View())
				}
				styledMessage = lipgloss.JoinVertical(
					lipgloss.Left,
					msgContent.Render(content))

			case chat.RoleSystem:
				styledMessage = lipgloss.JoinVertical(
					lipgloss.Center,
					msgContent.Render(msg.Content))

			}

			content.WriteString(styledMessage)
			if i < len(history)-1 {
				content.WriteString("\n\n")
			}
		}
	}

	m.viewPort.SetContent(content.String())
}

/*
const (
	TitleConnectorLeft  = "┤"
	TitleConnectorRight = "├"

	TitleSimpleLeft  = "─"
	TitleSimpleRight = "─"

	TitleFancyLeft  = "┨"
	TitleFancyRight = "┠"
)

func setPos(x, y int) {
	fmt.Printf("\033[%d;%dH", y, x)

}

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
} */
