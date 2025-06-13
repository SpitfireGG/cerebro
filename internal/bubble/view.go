package bubble

import (
	"fmt"
	"strings"
)

func (m Model) View() string {

	status := "Gemini ready for prompt..."
	if m.isLLMthinking {
		status = "Thinking..."
	}
	return fmt.Sprintf("%s\n%s\n%s\n%s",
		m.viewPort.View(),
		strings.Repeat("─", m.width),
		status,
		m.textInput.View(),
	)
}
