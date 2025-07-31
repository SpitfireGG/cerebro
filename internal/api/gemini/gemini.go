package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/spitfiregg/cerebro/internal/config"
	"google.golang.org/genai"
)

type StreamChunk struct {
	Content string
	IsError bool
	Error   error
	IsEnd   bool
}

type LogTransport struct {
	RoundTripper http.RoundTripper
}

func (lt *LogTransport) RoundTrip(req *http.Request) (*http.Response, error) {

	Df, err := os.Create("gemini_req_dump.log")
	if err != nil {
		log.Printf("Error creating gemini_req_dump.log: %v", err)
		return lt.RoundTripper.RoundTrip(req)
	}

	requestDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Printf("Error dump request: %v", err)
	}

	httpRequestDumpFile, err := os.OpenFile("http_req_dump_gemini.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Printf("Error http dump request failed: %v", err)
	} else {
		httpRequestDumpFile.WriteString(fmt.Sprintf("--- %s Request to %s ---\n\n", req.Method, req.URL.String()))
		httpRequestDumpFile.Write(requestDump)
		httpRequestDumpFile.WriteString("\n--- End Request ---\n\n")
	}

	if req.Body != nil {
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			log.Printf("Error reading request body for logging: %v", err)
		} else {
			_, writeErr := Df.Write(bodyBytes)
			if writeErr != nil {
				log.Printf("Error writing to gemini_req_dump: %v", writeErr)
			}

			_, writeErr = Df.WriteString("\n")
			if writeErr != nil {
				log.Printf("Error writing newline to gemini_req_dump.log: %v", writeErr)
			}
		}
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}
	return lt.RoundTripper.RoundTrip(req)
}

// initialize a new Gemini client
func NewGeminiClient(apiKey string) (*genai.Client, error) {
	ctx := context.Background()

	httpClient := &http.Client{
		Transport: &LogTransport{RoundTripper: http.DefaultTransport},
		Timeout:   0,
	}

	if apiKey == "" {
		return nil, fmt.Errorf("API key was not found, please check again if the API Key exists")
	}
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:     apiKey,
		Backend:    genai.BackendGeminiAPI,
		HTTPClient: httpClient,
	})
	if err != nil {
		return nil, fmt.Errorf("Could not initialize a Gemini client")
	}
	return client, nil
}

func GenerateContentStream(api, prompt string) (<-chan StreamChunk, error) {

	ctx := context.Background()
	client, _ := NewGeminiClient(api)
	model := NewDefaultAppConfig(api).GeminiDefault.GeminiConfig.Model

	cfg := GenerateContentConfigFromGeminiConfig(&config.NewDefaultAppConfig().GeminiDefault)

	// create a buffered channel for non-blocking IO
	chunkChan := make(chan StreamChunk, 10)

	go func() {

		defer close(chunkChan)
		defer client.ClientConfig().HTTPClient.CloseIdleConnections()

		response := client.Models.GenerateContentStream(ctx, model, genai.Text(prompt), cfg)

		r, err := glamour.NewTermRenderer(

			glamour.WithStandardStyle("dark"),
			glamour.WithWordWrap(120),
		)
		if err != nil {
			chunkChan <- StreamChunk{
				IsError: true,
				Error:   fmt.Errorf("Markdown render failed to init: %v", err),
			}
		}

		for chunk, err := range response {
			if err != nil {
				chunkChan <- StreamChunk{
					IsError: true,
					Error:   fmt.Errorf("Markdown render failed to init: %v", err),
				}
			} else {
				part := chunk.Candidates[0].Content.Parts[0]
				rendredChunk, _ := r.Render(part.Text)
				chunkChan <- StreamChunk{
					IsError: false,
					Content: rendredChunk,
				}
			}

		}
		chunkChan <- StreamChunk{IsEnd: true}

	}()
	return chunkChan, nil

	/* if err != nil {
		return "", fmt.Errorf("error calling GenerateContent: %w", err)
	} */

	/* if len(response.Candidates) == 0 || response.Candidates == nil {
		// candiates are the differenct responses the LLM redponds with
		// response.PromptFeedback is recieved when any violation prompt is sent to the LLM is found, eg: pornographic or hacking questions or something
		if response.PromptFeedback != nil && len(response.PromptFeedback.BlockReason) > 0 {
			return "", fmt.Errorf("no response candidate found, bot was blocked due to violation: %v", response.PromptFeedback.BlockReason)
		}
		return "", fmt.Errorf("no response candidate found, issue with the model or something")
	} */

	/* 	var ( */
	/* 		resp_rank   int = 0 */
	/* 		botResponse strings.Builder */
	/* 		parts       []*genai.Part */
	/* 	) */
	/* 	parts = response.Candidates[resp_rank].Content.Parts */

	/* if len(parts) == 0 {
		return "empty content", nil

	} else {
		for _, part := range parts {
			if part.Thought {
				botResponse.WriteString(part.Text)
			} else {
				botResponse.WriteString(part.Text)
			}
		}
	} */

	// markdown response, seems pretty easy for small rendering , i should be making my own i guessss.................
	/* renderedResponse, err := r.Render(botResponse.String())
	if err != nil {
		fmt.Printf("got an error rendering the reponse: %v", err)
	}
	return renderedResponse, nil */
}

func GenerateContent(api, prompt string) (string, error) {

	chunkChan, err := GenerateContentStream(api, prompt)
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	var botResponse strings.Builder

	for chunk := range chunkChan {
		if chunk.IsError {
			return "", chunk.Error
		}
		if chunk.IsEnd {
			break
		}

		botResponse.WriteString(chunk.Content)
	}
	return botResponse.String(), nil

}
