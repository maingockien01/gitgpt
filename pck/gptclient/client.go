package gptclient

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

// Client is a client for the OpenAI API.
type Client struct {
	// APIKey is the API key to use for requests.
	APIKey string

	// HTTPClient is the HTTP client to use for requests.
	Client *openai.Client
}

// NewClient returns a new OpenAI API client.
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
		Client: openai.NewClient(apiKey),
	}
}

// GetCompletion returns a completion for the given prompt.
func (c *Client) GetCompletion(ctx context.Context, prompt string) (string, error) {
	// Create request
	req := &openai.ChatCompletionRequest{
		MaxTokens: 100,
		Temperature: 0.5,
		TopP: 1,
		N: 1,
		Stop: []string{"\n"},
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage {
			{
				Role: openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},

	}

	// Send request
	resp, err := c.Client.CreateChatCompletion(ctx, *req)
	if err != nil {
		return "", err
	}

	// Return completion
	return resp.Choices[0].Message.Content, nil
}

