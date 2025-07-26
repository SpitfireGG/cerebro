package config

// define the configuration for models
type AppConfig struct {
	GeminiDefault GeminiConfig
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
func NewDefaultAppConfig() *AppConfig {
	return &AppConfig{
		GeminiDefault: GeminiConfig{
			Model:            "gemini-2.5-flash",
			Temperature:      0.9,
			TopP:             0.5,
			TopK:             20.0,
			CandidateCount:   1,
			Seed:             5,
			StopSequences:    []string{"Stop!"},
			PresencePenalty:  0.0,
			FrequencyPenalty: 0.0,
			IncludeThoughts:  false, // default to false
			ThinkingBudget:   0,     // 0 for disabling and 1 for enabling
			SafetySettings: []SafetySettingConfig{
				{Category: "HarmCategoryDangerousContent", Threshold: "BlockLowAndAbove"},
				// define sensible default safety settings
			},
			ResponseMimeType:  "", // no default MIME type
			SystemInstruction: "",
		},
		/* OpenAIDefault: OpenAIConfig{},
		GrokDefault:   GrokConfig{},
		ClaudeDefault: ClaudeConfig{}, */
	}
}
