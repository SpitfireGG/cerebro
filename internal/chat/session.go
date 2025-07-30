package chat

import (
	"fmt"
	"time"
)

type Role string

const (
	RoleUser      Role = "User"
	RoleAssistant Role = "Assistant"
	RoleSystem    Role = "System"
)

type Message struct {
	Role      Role      `json:"role"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

type Session struct {
	messages  []Message
	id        string
	startTime time.Time
}

func (s *Session) GetFormattedTimeString(now time.Time) string {
	return now.Format("2006-01-02 15:04:05")
}

// New method: Get session duration
func (s *Session) GetSessionDuration() string {
	elapsed := time.Since(s.startTime)
	hours := int(elapsed.Hours())
	minutes := int(elapsed.Minutes()) % 60
	seconds := int(elapsed.Seconds()) % 60

	if hours > 0 {
		return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	}
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

func NewSession() *Session {
	return &Session{
		messages:  make([]Message, 0),
		id:        generateSessionID(),
		startTime: time.Now(), // Initialize start time
	}
}

func generateSessionID() string {
	return time.Now().Format("20060102-150405-000")
}

func (s *Session) addMessage(role Role, content string) *Message {
	if len(content) == 0 {
		return nil
	}

	msg := &Message{
		Role:      role,
		Content:   content,
		Timestamp: time.Now(), // Set timestamp when message is created
	}

	s.messages = append(s.messages, *msg)
	return msg
}

// ... rest of the session methods remain the same ...

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
	if len(s.messages) == 0 {
		return []Message{}
	}

	history := make([]Message, len(s.messages))
	copy(history, s.messages)
	return history
}

func (s *Session) GetLastMessage() *Message {
	if len(s.messages) == 0 {
		return nil
	}

	lastMsg := s.messages[len(s.messages)-1]
	return &lastMsg
}

func (s *Session) Clear() {
	s.messages = s.messages[:0]
}

func (s *Session) IsEmpty() bool {
	return len(s.messages) == 0
}

func (s *Session) GetID() string {
	return s.id
}

func (s *Session) GetMessageCount() int {
	return len(s.messages)
}
