package api

import (
	"bufio"
	"context"
	"fmt"
	"os"
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
		fmt.Println("creating a gemini client failed")
	}
	fmt.Println("new gemini client initialized")
	fmt.Println()

	// make a reader to read inputs from the stdin
	reader := bufio.NewReader(os.Stdin)

	var botReponse strings.Builder

	for {
		fmt.Print("User: ")
		userInput, _ := reader.ReadString('\n')

		if strings.ToLower(userInput) == "q" || strings.ToLower(userInput) == "quit" {
			fmt.Println("exiting...")
			break
		}
		if userInput == "" {
			fmt.Println("please try entering something")
			continue
		}
		res, err := client.Models.GenerateContent(
			ctx,
			"gemini-2.0-flash",
			genai.Text(fmt.Sprint(userInput)),
			nil,
		)
		if err != nil {
			fmt.Println("an error occured when repsonding, please try again!")
			continue
		}

		// candiates are the differenct responses the LLM redponds with
		// res.PromptFeedback is recieved when any violation prompt is sent to the LLM is found, eg: pornographic or hacking questions or something
		if len(res.Candidates) == 0 {
			fmt.Println("no response candiate found, issue with the model or something.... blah blah blah")
			if res.PromptFeedback != nil && len(res.PromptFeedback.BlockReason) > 0 {
				fmt.Printf("due to voilation the bot was blocked\n\n err: %v", res.PromptFeedback.BlockReason)
			}
			continue
		}

		// create a string slice to hold the reponse
		for _ = range res.Candidates[0].Content.Parts {
			botReponse.WriteString(string(res.Text()))
		}

	}
	return botReponse.String(), nil
}
