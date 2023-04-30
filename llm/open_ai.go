package llm

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

type OpenAi struct {
	apiKey string
}

func NewOpenAi() (*OpenAi, error) {
	var apiKey = os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		err := errors.New("missing OpenAI API key")
		return &OpenAi{}, err
	}
	return &OpenAi{apiKey: apiKey}, nil
}

func (oa OpenAi) Predict(msg string) (Response, bool) {
	client := openai.NewClient(oa.apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo0301,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: msg,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return Response{}, false
	}

	return Response{[]string{resp.Choices[0].Message.Content}}, false
}
