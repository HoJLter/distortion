package dto

import "server/internal/repository"

type GetUserResponse struct {
	UserUID    string  `json:"user_uid"`
	Username   string  `json:"username"`
	AvatarPath *string `json:"avatar_path"`
}

func RecordToUserResponse(record repository.UserRecord) GetUserResponse {
	return GetUserResponse{
		UserUID:    record.UserUID,
		Username:   record.Username,
		AvatarPath: record.AvatarPath,
	}
}

type CreateUserRequest struct {
	Username   string  `json:"username"`
	AvatarPath *string `json:"avatar_path"`
}
