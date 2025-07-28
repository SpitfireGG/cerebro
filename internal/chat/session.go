package chat

import "time"

type Role string

const (
	// represents the human interacting with the llm
	RoleUser Role = "User"

	// represents the role of the model itself
	RoleAssistant Role = "Assistant"

	// the System role is used to provide setup information or context that informs the behavior of the model
	RoleSystem Role = "System"
)

type Message struct {
	// role determines whether it is User, System or Assistant
	Role Role

	// actual content
	Content string

	// timestamp of the current conversation
	Timestamp time.Time
}

type Session struct {
	messages []Message
}

// NewSession creates an empty session when the user shifts to a new chat
// previous chats are saved for future references and a new one is created
func NewSession() *Session {
	return &Session{
		messages: make([]Message, 0),
	}
}

// addMessage method adds a new message along with following properties:
// - role
// - content
// timestamp
func (s *Session) addMessage(role Role, content string) {
	msg := Message{
		Role:      role,
		Content:   content,
		Timestamp: time.Now(),
	}
	s.messages = append(s.messages, msg)
}

func (s *Session) AddUserMessage(content string) {
	s.addMessage(RoleUser, content)
}

func (s *Session) AddAssistantMessage(content string) {
	s.addMessage(RoleAssistant, content)
}

func (s *Session) AddSystemMessage(content string) {
	s.addMessage(RoleSystem, content)
}

// GetHistory method returns a slice of message that were saved proviously
func (s *Session) GetHistory() []Message {
	return s.messages
}

// Clear method deletes the marked history from the history surf
func (s *Session) Clear() {
	s.messages = s.messages[:0]
}
