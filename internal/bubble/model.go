package bubble

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spitfiregg/RTUI_chatbot/internal/bubble/chat"
	"os"
)

// TODO: make struct member for more models ( add more models )
// TODO: make user be able to recurse through the chat history

// Viewport dimension
const (
	vpHeight = 40
	vpWidth  = 120
)

// define the main program state
type UI struct {
	textInput textinput.Model
	viewPort  viewport.Model
	height    int
	width     int
}

type App struct {
	ui   *UI
	chat *chat.Session
}

type LLMreponseMsg struct {
	response string
	err      error
}

type DebugModel struct {
	Dump *os.File
}

type Model struct {
	isLLMthinking bool
	api_key       string

	//embed the defined structs into the main Model

	App
	UI
	LLMreponseMsg
	DebugModel
}

func TextInputHandler() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "Talk to Gemini"
	ti.Focus()
	ti.Cursor.Blink = true
	ti.CharLimit = 512
	ti.Width = 80
	return ti

}

func InitialModel(apiKey string) Model {

	// jump straight to prompting the model
	vp := viewport.New(vpWidth, vpHeight)
	vp.SetContent("welcome to the Playground...")

	return Model{
		isLLMthinking: false, // initially set the model thinking to be false
		api_key:       apiKey,

		UI: UI{
			textInput: TextInputHandler(),
			viewPort:  vp,
		},
		App: App{
			ui:   &UI{},
			chat: chat.NewSession(), // default to new session
		},
		LLMreponseMsg: LLMreponseMsg{
			response: "This is an initial test reponse",
			err:      nil,
		},
		DebugModel: DebugModel{},
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink)
}
