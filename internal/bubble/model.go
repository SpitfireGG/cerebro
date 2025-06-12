package bubble

import (
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spitfiregg/RTUI_chatbot/internal/api"
)

// TODO: make struct member for more models ( add more models )
// TODO: make user be able to recurse through the chat history

// Viewport dimension
const (
	vpHeight = 80
	vpWidth  = 20
)

type LLMreponseMsg struct {
	response string
	err      error
}

type Model struct {
	textIP         textinput.Model
	viewPort       viewport.Model
	LLMreponse     string
	Userprompt     string
	isLLMthinking  bool
	responseHeight uint8
	responseWidth  uint8
	promptHeight   uint8
	promptWidth    uint8
	api_key        string
	Dump           *os.File
}

func InitialModel(apiKey string) Model {

	t1 := textinput.New()
	t1.Placeholder = "Talk to Gemini"
	t1.Focus()
	t1.Cursor.Blink = true
	t1.CharLimit = 512
	t1.Width = 80

	vp := viewport.New(vpWidth, vpHeight)

	// jump straight to prompting the model

	return Model{
		textIP:         t1,
		viewPort:       vp,
		LLMreponse:     "",    // nil initially
		Userprompt:     "",    // nil initially
		isLLMthinking:  false, // initially set the model thinking to be false
		responseHeight: 0,
		responseWidth:  0,
		promptHeight:   0,
		promptWidth:    0,
		Dump:           nil,
	}
}

func (m Model) Init() tea.Cmd {
	return (textinput.Blink)
}

func (m Model) GenerateReponse(prompt string) tea.Cmd {
	return func() tea.Msg {
		resp, err := api.GenerateContent(m.api_key, prompt)
		return LLMreponseMsg{
			response: resp,
			err:      err,
		}
	}
}
