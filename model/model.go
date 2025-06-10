package model

import (
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg error

type Model struct {
	textInput  textinput.Model
	prompt     string
	response   string
	err        errMsg
	statusMsg  string
	StatusCode uint8
}

func InitialModel(m Model) Model {

	t1 := textinput.New()
	t1.Placeholder = "enter the prompt..."
	t1.Focus()
	t1.CharLimit = 256
	t1.Width = 1024

	var prompt string = ""

	return Model{
		textInput:  t1,
		prompt:     prompt,
		response:   "",
		err:        nil,
		statusMsg:  "",
		StatusCode: 0,
	}

}

func (m Model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.textInput.Focus())
}
