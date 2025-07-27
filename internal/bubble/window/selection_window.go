package window

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type LLMmodel struct {
	table table.Model
	title string
}

type ModelSelectedMsg struct{ ModelName string }

type ModelSelector interface {
	NewModel() LLMmodel
}

// NewModel returns a LLM model upon selection that will be used as the AI agent
func NewModel() LLMmodel {
	columns := []table.Column{
		{Title: "ID", Width: 6},
		{Title: "Model Name", Width: 25},
		{Title: "Provider", Width: 15},
		{Title: "Status", Width: 10},
	}

	rows := []table.Row{
		{"1", "Gemini Flash 2.5", "Google", "✓ Active"},
		{"2", "Claude 4 Sonnet", "Anthropic", "✓ Active"},
		{"3", "Grok 3 Smartest", "xAI", "✓ Active"},
		{"4", "GPT-4 Opus", "OpenAI", "⚠ Limited"},
		{"5", "Llama 3.1 405B", "Meta", "✓ Active"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(8),
	)

	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(borderColor).
		BorderBottom(true).
		Background(fireBrick).
		Foreground(accentBlue).
		Bold(true).
		Padding(0, 1).
		Align(lipgloss.Center)

	s.Selected = s.Selected.
		Foreground(primaryBg).
		Background(fireBrick).
		Bold(true)

	s.Cell = s.Cell.
		Foreground(textPrimary).
		Padding(0, 1)

	t.SetStyles(s)

	return LLMmodel{
		table: t,
		title: "Model Selection Window",
	}
}
func (lm LLMmodel) Init() tea.Cmd {
	return nil
}

func (lm LLMmodel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:

		availableWidth := msg.Width - 8
		availableHeight := msg.Height - 8

		lm.table.SetWidth(availableWidth)
		lm.table.SetHeight(availableHeight)

	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q", "ctrl+c":
			return lm, tea.Quit

		case "enter", " ":

			selectedRow := lm.table.SelectedRow()
			if len(selectedRow) >= 2 {
				selectedModel := selectedRow[1]
				return lm, func() tea.Msg {
					return ModelSelectedMsg{ModelName: selectedModel}
				}
			}

		case "tab":

			if lm.table.Focused() {
				lm.table.Blur()
			} else {
				lm.table.Focus()
			}
		}
	}

	lm.table, cmd = lm.table.Update(msg)
	return lm, cmd
}

func (lm LLMmodel) View() string {

	title := TitleStyle.Render(lm.title)
	tableView := BaseStyle.Render(lm.table.View())
	footer := FooterStyle.Render(
		"↑↓: navigate • enterspace: select • tab: focus • q/esc: quit",
	)

	return lipgloss.JoinVertical(
		lipgloss.Center,
		title,
		tableView,
		footer,
	)
}
