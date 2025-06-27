package api

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/spitfiregg/garlic/internal/config"
	"google.golang.org/genai"
)

type ClientWrapper struct {
	base_client   *genai.Client
	defaultConfig *config.GeminiConfig
}

// initialize a new Gemini client
func NewGeminiClient(appCfg *config.AppConfig, apiKey string, prompt string) (*ClientWrapper, error) {

	clientWr := &ClientWrapper{}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return clientWr, fmt.Errorf("creating a gemini client failed: %w", err)
	}
	defer client.ClientConfig().HTTPClient.CloseIdleConnections()

	return &ClientWrapper{
		base_client:   client,
		defaultConfig: &appCfg.GeminiDefault,
	}, nil

}

func (cw *ClientWrapper) GenerateContent(ctx context.Context, userConfig *config.GeminiConfig, parts ...genai.Part) (*genai.GenerateContentResponse, error) {

	// create a working copy of the defaultConfig that the user overrides
	workingConfig := *cw.defaultConfig

	if userConfig != nil {

		if userConfig.Model != "" {
			workingConfig.Model = userConfig.Model
		}
		if userConfig.Temperature != 0 {
			workingConfig.Temperature = userConfig.Temperature
		}
		if userConfig.TopP != 0 {
			workingConfig.TopP = userConfig.TopP
		}
		if userConfig.TopK != 0 {
			workingConfig.TopK = userConfig.TopK
		}
		if userConfig.CandidateCount != 0 {
			workingConfig.CandidateCount = userConfig.CandidateCount
		}
		if userConfig.Seed != 0 {
			workingConfig.Seed = userConfig.Seed
		}
		if len(userConfig.StopSequences) > 0 {
			workingConfig.StopSequences = userConfig.StopSequences
		}

		workingConfig.PresencePenalty = userConfig.PresencePenalty
		workingConfig.FrequencyPenalty = userConfig.FrequencyPenalty

		workingConfig.IncludeThoughts = userConfig.IncludeThoughts
		workingConfig.ThinkingBudget = userConfig.ThinkingBudget

		if len(userConfig.SafetySettings) > 0 {
			workingConfig.SafetySettings = userConfig.SafetySettings
		}
		if workingConfig.ResponseMimeType != "" {
			workingConfig.ResponseMimeType = userConfig.ResponseMimeType
		}
		if userConfig.SystemInstruction != "" {
			workingConfig.SystemInstruction = userConfig.SystemInstruction
		}
	}

	res, err := genai.Client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return clientWr, fmt.Errorf("an error occurred when responding: %w", err)
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

	r, _ := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(120),
	)

	renderedResponse, err := r.Render(botResponse.String())

	return renderedResponse, nil
}
