package dto

import "encoding/json"

type WSPostMessageDTO struct {
	ChannelUID string `json:"channel_uid"`
	UserUID    string `json:"user_uid"`
	Content    string `json:"content"`
}

type WSMessageDTO struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}
