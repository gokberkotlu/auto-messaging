package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gokberkotlu/auto-messaging/entity"
)

type MessageClient struct {
	url         string
	token       string
	xInsAuthKey string
}

func New(token string) *MessageClient {
	return &MessageClient{
		url:         "https://webhook.site/",
		token:       token,
		xInsAuthKey: "INS.me1x9uMcyYGlhKKQVPoc.bO3j9aZwRTOcA2Ywo",
	}
}

func (c *MessageClient) SendMessage(message entity.Message) {
	method := "POST"
	jsonBytes, _ := json.Marshal(message)
	payload := string(jsonBytes)

	httpClient := &http.Client{}
	req, err := http.NewRequest(method, c.url+c.token, strings.NewReader(payload))

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-ins-auth-key", c.xInsAuthKey)

	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
