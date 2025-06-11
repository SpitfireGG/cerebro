package bubble

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	/* 	var cmds []tea.Cmd */

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:

	case tea.KeyMsg:
		switch msg.String() {

		case "q", "quit", "ctrl-c":
			return m, tea.Quit

		case "enter":
			if m.isLLMthinking {
				break
			}
		}
	}
	m.textIP, cmd = m.textIP.Update(msg)
	return m, cmd
}
