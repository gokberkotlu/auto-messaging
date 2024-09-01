package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gokberkotlu/auto-messaging/dto"
	"github.com/gokberkotlu/auto-messaging/service"
)

type IMessageClient interface {
	SendNextTwoUnsentMessages()
	sendMessageByExternalAPI(messageDTO dto.MessageDTO) error
}

type MessageClient struct {
	url            string
	token          string
	xInsAuthKey    string
	messageService service.IMessageService
}

func New() IMessageClient {
	return &MessageClient{
		url:            "https://webhook.site/",
		token:          os.Getenv("CLIENT_TOKEN"),
		xInsAuthKey:    os.Getenv("CLIENT_X_INS_AUTH_KEY"),
		messageService: service.NewMessageService(),
	}
}

func (c *MessageClient) sendMessageByExternalAPI(messageDTO dto.MessageDTO) error {
	method := "POST"
	jsonBytes, _ := json.Marshal(messageDTO)
	payload := string(jsonBytes)

	httpClient := &http.Client{}
	req, err := http.NewRequest(method, c.url+c.token, strings.NewReader(payload))

	if err != nil {
		fmt.Println(err)
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-ins-auth-key", c.xInsAuthKey)

	res, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request to external api: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response from external api: %w", err)
	}
	fmt.Println(string(body))
	return nil
}

func (c *MessageClient) SendNextTwoUnsentMessages() {
	messages := c.messageService.GetNextTwoUnsentMessages()

	if messages != nil && len(messages) > 0 {
		for _, message := range messages {
			fmt.Println(message.ID)
			err := c.sendMessageByExternalAPI(dto.ToMessageDTO(message))
			if err != nil {
				fmt.Println(err)
				return
			}

			c.messageService.UpdateMessageStatusAsSent(message)
		}
	}
}
