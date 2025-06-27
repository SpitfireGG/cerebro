package api

import (
	"fmt"
	"github.com/spitfiregg/garlic/internal/config"
	"google.golang.org/genai"
)

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
