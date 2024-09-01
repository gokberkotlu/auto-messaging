package dto

import "github.com/gokberkotlu/auto-messaging/entity"

type MessageDTO struct {
	To      string `json:"to"`
	Content string `json:"content"`
}

func ToMessageDTO(message entity.Message) MessageDTO {
	return MessageDTO{
		To:      message.To,
		Content: message.Content,
	}
}
