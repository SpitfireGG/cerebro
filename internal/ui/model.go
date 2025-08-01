package ui

import (
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	api "github.com/spitfiregg/cerebro/internal/api/gemini"
	"github.com/spitfiregg/cerebro/internal/chat"
	"github.com/spitfiregg/cerebro/internal/ui/states"
)

// TODO: make struct member for more models ( add more models )
// TODO: make user be able to recurse through the chat history

// Viewport dimension
const (
	vpHeight = 40
	vpWidth  = 120
)

type State int

const (
	GreetWindow State = iota
	ModelSelection
	MainWindow
	SettingsWindow
)

type Key struct {
	KeyHandlerChan chan (tea.KeyMsg)
	KeyPressRecv   string
}

// define the main program state
type UI struct {
	textArea textarea.Model
	viewPort viewport.Model
	height   int
	width    int
}

type SpinnerModel struct {
	spinner  spinner.Model
	quitting bool
	err      error
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

	currentResponse strings.Builder
	isStreaming     bool

	// recieves a stream chunk from the StreamChunk channnel
	streamChan <-chan api.StreamChunk

	// for window selection
	currentState      State
	LLMSelectorWindow window.LLMmodel
	selectedLLM       string

	Timestamp time.Time

	//embed the defined structs into the main Model
	App
	UI
	Key
	SpinnerModel
	LLMreponseMsg
	DebugModel
}

type TransitionToMain struct{}

func TextInputHandler() textarea.Model {
	textarea := textarea.New()
	textarea.Placeholder = ""
	textarea.Focus()
	textarea.Cursor.Blink = true
	textarea.CharLimit = 512
	textarea.Prompt = "┃ "
	textarea.SetWidth(30)
	textarea.SetHeight(3)
	textarea.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#FE7743"))
	textarea.Cursor.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FEFFFF"))

	textarea.ShowLineNumbers = false

	return textarea
}

func InitialModel(config *api.AppConfig) Model {

	// jump straight to prompting the model
	vp := viewport.New(vpWidth, vpHeight)

	// spinner
	s := spinner.New()
	s.Spinner = spinner.Meter
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#1658EE"))

	tnow := time.Now()
	return Model{

		Timestamp: tnow,

		currentState:      GreetWindow,
		LLMSelectorWindow: window.NewModel(),
		isLLMthinking:     false, // initially set the model thinking to be false
		api_key:           config.GeminiDefault.ApiKey,

		// the whole ui of the program
		UI: UI{
			textArea: TextInputHandler(),
			viewPort: vp,
		},

		// the app itself
		App: App{
			ui:   &UI{},
			chat: chat.NewSession(), // default to new session
		},

		// response from the LLM
		LLMreponseMsg: LLMreponseMsg{
			response: "This is an initial test reponse",
			err:      nil,
		},

		// debugging
		DebugModel:   DebugModel{},
		SpinnerModel: SpinnerModel{spinner: s}}

}

func Transition(d time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(d)
		return TransitionToMain{}
	}
}

func (m Model) Init() tea.Cmd {
	return m.spinner.Tick
}
