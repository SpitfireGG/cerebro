package api

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/genai"
)

func GenerateContent(apiKey, prompt string) (string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return "", fmt.Errorf("creating a gemini client failed: %w", err)
	}
	defer client.ClientConfig().HTTPClient.CloseIdleConnections() // Critical: Clean up resources

	res, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("an error occurred when responding: %w", err)
	}

	if len(res.Candidates) == 0 {
		// candiates are the differenct responses the LLM redponds with
		// res.PromptFeedback is recieved when any violation prompt is sent to the LLM is found, eg: pornographic or hacking questions or something
		if res.PromptFeedback != nil && len(res.PromptFeedback.BlockReason) > 0 {
			return "", fmt.Errorf("no response candidate found, bot was blocked due to violation: %v", res.PromptFeedback.BlockReason)
		}
		return "", fmt.Errorf("no response candidate found, issue with the model or something")
	}

	var botResponse strings.Builder
	for _, part := range res.Candidates[0].Content.Parts {
		botResponse.WriteString(part.Text)
	}

	return botResponse.String(), nil
}
