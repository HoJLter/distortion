package dto

import "server/internal/repository"

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
