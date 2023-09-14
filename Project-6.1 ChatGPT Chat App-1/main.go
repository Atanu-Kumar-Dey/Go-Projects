package main

import (
	"context"
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load()
	apikey := os.Getenv("API_KEY")
	if apikey == "" {
		log.Fatalln("Missing api key")
	}
	ctx := context.Background()
	client := gpt3.NewClient(apikey)

	response, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:    []string{"The first thing you should know about GoLang is"},
		MaxTokens: gpt3.IntPtr(30),
		Stop:      []string{"."},
		Echo:      true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(response.Choices[0].Text)
}
