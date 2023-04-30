package agents

import (
	"GoGPT/promts"
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

func RunAgent(prompt promts.PromptTemplate, input []string) {
	chatMsg := prompt.Format(input)
	resp, err := predict(chatMsg)
	if err {
		println(err)
		return
	}

	println(resp)
}

func predict(chatMsg string) (string, bool) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo0301,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: chatMsg,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "No response from gpt", false
	}

	return resp.Choices[0].Message.Content, false
}
