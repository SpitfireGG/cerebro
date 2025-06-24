package window

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var BaseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type LLMmodel struct {
	table table.Model
}

type ModelSelectedMsg struct {
	ModelName string
}

type Greeter struct {
	Greet string
}

func GreetUser() Greeter {
	return Greeter{Greet: "welcome to the playground user\n"}
}

func NewModel() LLMmodel {
	columns := []table.Column{
		{Title: "Id", Width: 4},
		{Title: "Model", Width: 20},
	}
	rows := []table.Row{
		{"1", "Gemini-flash 2.5"},
		{"2", "Claude-4 sonet"},
		{"3", "Grok-3 smartest"},
		{"4", "Chatgpt-4.0 opus"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return LLMmodel{table: t}
}

func (lm LLMmodel) Init() tea.Cmd {
	return nil
}

func (lm LLMmodel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		lm.table.SetWidth(msg.Width - 4)
		lm.table.SetHeight(msg.Height - 4)

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if lm.table.Focused() {
				lm.table.Blur()
			} else {
				lm.table.Focus()
			}
		case "q", "ctrl+c":
			return lm, tea.Quit
		case "enter":
			// When enter is pressed, send a message with the selected model name.
			// This is the key change to make the selection work.
			selectedModel := lm.table.SelectedRow()[1]
			return lm, func() tea.Msg {
				return ModelSelectedMsg{ModelName: selectedModel}
			}
		}
	}

	lm.table, cmd = lm.table.Update(msg)
	return lm, cmd
}

func (lm LLMmodel) View() string {
	return BaseStyle.Render(lm.table.View()) + "\n"
}
