package api

import (
	"fmt"

	"github.com/spitfiregg/garlic/internal/config"
	"google.golang.org/genai"
)

// define the configuration for models
type AppConfig struct {
	GeminiDefault struct {
		GeminiConfig
		ApiKey string
	}
	/* OpenAIDefault OpenAIConfig
	GrokDefault   GrokConfig
	ClaudeDefault ClaudeConfig */
}

// GeminiConfig holds some parameters for the model to access
type GeminiConfig struct {
	Model             string
	Temperature       float32
	TopP              float32
	TopK              float32
	CandidateCount    int32
	Seed              int32
	StopSequences     []string
	PresencePenalty   float32
	FrequencyPenalty  float32
	IncludeThoughts   bool
	ThinkingBudget    int32
	SafetySettings    []SafetySettingConfig
	ResponseMimeType  string
	SystemInstruction string
}

// TODO: make up configuration for each of the models
/* type OpenAIConfig struct{}
type GrokConfig struct{}
type ClaudeConfig struct{} */

// define the safety settings
type SafetySettingConfig struct {
	Category  string // "HarmCategoryHateSpeech"
	Threshold string // "BlockMediumAndAbove"
}

// return the app config with defined parameters
func NewDefaultAppConfig(apiKey string) *AppConfig {
	return &AppConfig{
		GeminiDefault: struct {
			GeminiConfig
			ApiKey string
		}{
			GeminiConfig: GeminiConfig{
				Model:            "gemini-2.5-flash",
				Temperature:      0.9,
				TopP:             0.5,
				TopK:             20.0,
				CandidateCount:   1,
				Seed:             5,
				StopSequences:    []string{"Stop!"},
				PresencePenalty:  0.0,
				FrequencyPenalty: 0.0,
				IncludeThoughts:  true, // default to false
				ThinkingBudget:   0,    // 0 for disabling and 1 for enabling
				SafetySettings: []SafetySettingConfig{
					{Category: "HarmCategoryDangerousContent", Threshold: "BlockLowAndAbove"},
					// define sensible default safety settings
				},
				ResponseMimeType:  "", // no default MIME type
				SystemInstruction: "",
			},
			ApiKey: apiKey,
		},
		/* OpenAIDefault: OpenAIConfig{},
		GrokDefault:   GrokConfig{},
		ClaudeDefault: ClaudeConfig{}, */
	}
}

// (Helper functions parseHarmCategory and parseHarmBlockThreshold remain the same)
func parseHarmCategory(s string) (genai.HarmCategory, error) {
	return genai.HarmCategoryUnspecified, nil
}

func parseHarmBlockThreshold(s string) (genai.HarmBlockThreshold, error) {
	return genai.HarmBlockThresholdUnspecified, nil
}

// parses the provided safety settings to genai safety settings
func ToGenaiSafetySettings(cfg []config.SafetySettingConfig) []*genai.SafetySetting {

	if len(cfg) == 0 {
		return nil
	}

	// make a slice of safetySetting to hold safety contents
	settings := make([]*genai.SafetySetting, len(cfg))
	for i, s := range cfg {
		category, err := parseHarmCategory(s.Category)
		if err != nil {
			fmt.Printf("Warning: Unknown harm category '%s', skipping. Error: %v\n", s.Category, err)
			continue
		}
		threshold, err := parseHarmBlockThreshold(s.Threshold)
		if err != nil {
			fmt.Printf("Warning: Unknown harm block threshold '%s' for category '%s', skipping. Error: %v\n", s.Threshold, s.Category, err)
			continue
		}
		settings[i] = &genai.SafetySetting{
			Category:  category,
			Threshold: threshold,
		}
	}
	return settings
}

// GenerateContentConfigFromGeminiConfig creates a genai.GenerateContentConfig from a GeminiConfig.
func GenerateContentConfigFromGeminiConfig(cfg *config.GeminiConfig) *genai.GenerateContentConfig {
	if cfg == nil {
		return nil
	}

	thinkingBudget := cfg.ThinkingBudget
	includeThoughts := cfg.IncludeThoughts

	genConfig := &genai.GenerateContentConfig{
		Temperature:      &cfg.Temperature,
		TopP:             &cfg.TopP,
		TopK:             &cfg.TopK,
		CandidateCount:   cfg.CandidateCount,
		Seed:             &cfg.Seed,
		StopSequences:    cfg.StopSequences,
		PresencePenalty:  &cfg.PresencePenalty,
		FrequencyPenalty: &cfg.FrequencyPenalty,
		ResponseMIMEType: cfg.ResponseMimeType,
	}

	if includeThoughts || thinkingBudget > 0 {
		genConfig.ThinkingConfig = &genai.ThinkingConfig{
			IncludeThoughts: includeThoughts,
			ThinkingBudget:  &thinkingBudget,
		}
	}
	return genConfig
}
