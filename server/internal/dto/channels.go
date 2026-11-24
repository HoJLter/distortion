package dto

import "server/internal/repository"

type CreateChannelRequest struct {
	ChannelUID string                 `json:"channel_uid"`
	Type       repository.ChannelType `json:"type"`
}

type GetChannelResponse struct {
	ChannelUID string                 `json:"channel_uid"`
	Type       repository.ChannelType `json:"type"`
}

func RecordToChannelDTO(channel repository.ChannelRecord) GetChannelResponse {
	return GetChannelResponse{
		ChannelUID: channel.UID,
		Type:       channel.Type,
	}
}
