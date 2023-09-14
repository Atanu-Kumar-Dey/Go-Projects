package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func getResponse(client gpt3.Client, ctx context.Context, input string) {
	err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt: []string{
			input,
		},
		MaxTokens:   gpt3.IntPtr(3000),
		Temperature: gpt3.Float32Ptr(0),
	}, func(cr *gpt3.CompletionResponse) {
		fmt.Println(cr.Choices[0].Text)
	},
	)
	if err != nil {
		log.Fatalln(err)
		os.Exit(10)
	}
	fmt.Printf("\n")
}

type NullWriter int

func (NullWriter) Write([]byte) (int, error) {
	return 0, nil
}

func main() {
	log.SetOutput(new(NullWriter))
	godotenv.Load()
	apikey := os.Getenv("API_KEY")
	if apikey == "" {
		log.Fatalln("Missing api key")
	}
	ctx := context.Background()
	client := gpt3.NewClient(apikey)

	rootCmd := &cobra.Command{
		Use:   "chatgpt",
		Short: "Chat with ChatGPT in console",
		Run: func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			quit := false

			for !quit {
				fmt.Println("Say something ('quit' to exit)")
				if !scanner.Scan() {

					break
				}
				input := scanner.Text()
				switch input {
				case "quit":
					quit = true
				default:
					getResponse(client, ctx, input)
				}

			}

		},
	}
	rootCmd.Execute()
}
