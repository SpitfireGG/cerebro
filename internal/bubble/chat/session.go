package chat

import "time"

type Role string

const (
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
	RoleSystem    Role = "system"
)

type Message struct {
	Role      Role
	Content   string
	Timestamp time.Time
}

type Session struct {
	messages []Message
}

func NewSession() *Session {
	return &Session{
		messages: make([]Message, 0),
	}
}
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

func (s *Session) GetHistory() []Message {
	return s.messages
}

func (s *Session) Clear() {
	s.messages = s.messages[:0]
}
