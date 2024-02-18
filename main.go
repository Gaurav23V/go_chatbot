package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	// Access your API key
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY environment variable is not set")
	}

	// Create a new client for the Gemini API with the API key as an option
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Initialize the Generative Model
	model := client.GenerativeModel("gemini-pro")

	// Start the chat session
	cs := model.StartChat()

	// Prompt the user for input until "quit" is typed
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Input: ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		// Exit if "quit" is typed
		if userInput == "quit" {
			break
		}

		// Send the user input to the chat session and receive the response
		resp, err := cs.SendMessage(ctx, genai.Text(userInput))
		if err != nil {
			log.Fatal(err)
		}

		// Print the response to the terminal
		fmt.Println("Response:", resp.Candidates[0].Content)
	}
}
