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
	textIP   textinput.Model
	viewPort viewport.Model
	height   int
	width    int
}

type App struct {
	ui      *UI
	chat    *chat.Session
	api_key string
}

type LLMreponseMsg struct {
	response string
	err      error
}

type DebugModel struct {
	Dump *os.File
}

type Model struct {
	LLMreponse    string
	Userprompt    string
	isLLMthinking bool

	//embed the defined structs into the main Model

	App
	UI
	LLMreponseMsg
	DebugModel
}

func InitialModel(apiKey string) Model {

	t1 := textinput.New()
	t1.Placeholder = "Talk to Gemini"
	t1.Focus()
	t1.Cursor.Blink = true
	t1.CharLimit = 512
	t1.Width = 80

	// jump straight to prompting the model

	return Model{
		LLMreponse:    "",    // nil initially
		Userprompt:    "",    // nil initially
		isLLMthinking: false, // initially set the model thinking to be false

		UI: UI{
			textIP:   t1,
			viewPort: viewport.New(vpWidth, vpHeight),
		},
		App: App{
			ui:      &UI{},
			chat:    &chat.Session{},
			api_key: apiKey,
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
