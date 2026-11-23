package handler

import (
	"encoding/json"
	"server/internal/repository"
)

type MessageAuthorDTO struct {
	UID        string `json:"uid"`
	Username   string `json:"username"`
	AvatarPath string `json:"avatar_path"`
}
type MessageDTO struct {
	MessageID  int              `json:"message_id"`
	ChannelUID string           `json:"channel_uid"`
	Author     MessageAuthorDTO `json:"author"`
	Content    string           `json:"content"`
}

func RecordToMessageDTO(record repository.MessageRecord) MessageDTO {
	return MessageDTO{
		MessageID: record.MessageID,
		Author: MessageAuthorDTO{
			UID:        record.Author.UID,
			Username:   record.Author.Username,
			AvatarPath: record.Author.AvatarPath,
		},
		Content: record.Content,
	}
}

type CreateMessageRequest struct {
	UserUID string `json:"user_uid" validate:"required"`
	Content string `json:"content" validate:"required"`
}
type CreateMessageResponse struct {
	MessageID int `json:"message_id"`
}

type WSPostMessageDTO struct {
	ChannelUID string `json:"channel_uid"`
	UserUID    string `json:"user_uid"`
	Content    string `json:"content"`
}

type WSMessageDTO struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type ChannelType string

const (
	TextChannelType  ChannelType = "text"
	VoiceChannelType ChannelType = "voice"
)

type CreateChannelRequest struct {
	ChannelUID string      `json:"channel_uid"`
	Type       ChannelType `json:"type"`
}
