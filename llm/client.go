
package llm

import (
    "context"
    "fmt"

    openai "github.com/sashabaranov/go-openai"

    "mini-agent/config"
)

type Client struct {
    client *openai.Client
}

func NewClient() *Client {
    cfg := openai.DefaultConfig(config.APIKey)
    cfg.BaseURL = config.BaseURL

    return &Client{
        client: openai.NewClientWithConfig(cfg),
    }
}

// Chat 发送一条消息，并返回模型的回答。
func (c *Client) Chat(message string) (string, error) {
    resp, err := c.client.CreateChatCompletion(
        context.Background(),
        openai.ChatCompletionRequest{
            Model: config.Model,
            Messages: []openai.ChatCompletionMessage{
                {
                    Role:    openai.ChatMessageRoleUser,
                    Content: message,
                },
            },
        },
    )
    if err != nil {
        return "", fmt.Errorf("调用大模型失败: %w", err)
    }

    if len(resp.Choices) == 0 {
        return "", fmt.Errorf("大模型没有返回回答")
    }

    return resp.Choices[0].Message.Content, nil
}